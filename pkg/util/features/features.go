package features

import (
	"fmt"
	"maps"

	"github.com/loft-sh/admin-apis/pkg/licenseapi"
	"github.com/stripe/stripe-go/v81"
	stripefeatures "github.com/stripe/stripe-go/v81/entitlements/feature"
	stripeproducts "github.com/stripe/stripe-go/v81/product"
	"github.com/stripe/stripe-go/v81/productfeature"
)

const (
	metadataQueryFmt = "metadata['%s']:'%s'"
)

type SyncedFeature struct {
	name        string
	displayName string
	stripeID    string
}

func EnsureFeatures(syncedFeatures map[string]SyncedFeature, features []*licenseapi.Feature, isLimit bool) error {
	err := EnsureStripeFeatures(syncedFeatures, features, isLimit)
	if err != nil {
		return err
	}

	if err = EnsureFeatureProducts(syncedFeatures); err != nil {
		return err
	}

	if err = EnsureAttachAll(syncedFeatures); err != nil {
		return err
	}
	return nil
}

func EnsureStripeFeatures(syncedFeatures map[string]SyncedFeature, features []*licenseapi.Feature, isLimit bool) error {
	for _, f := range features {
		extraMetadata := map[string]string{}
		if isLimit {
			extraMetadata[licenseapi.MetadataKeyFeatureIsLimit] = licenseapi.MetadataValueTrue
			f.Name = licenseapi.LimitsPrefix + f.Name
		}

		if f.Status == string(licenseapi.FeatureStatusHidden) {
			extraMetadata[licenseapi.MetadataKeyFeatureIsHidden] = licenseapi.MetadataValueTrue
		}

		err := EnsureFeatureExists(syncedFeatures, f.Name, f.DisplayName, extraMetadata)
		if err != nil {
			return err
		}

		if f.Preview {
			extraMetadata[licenseapi.MetadataKeyFeatureIsPreview] = licenseapi.MetadataValueTrue
			err = EnsureFeatureExists(syncedFeatures, f.Name+"-preview", f.DisplayName+" [Preview]", extraMetadata)
			if err != nil {
				return err
			}
		}

		if isLimit {
			extraMetadata[licenseapi.MetadataKeyFeatureLimitType] = licenseapi.MetadataKeyFeatureLimitTypeActive
			err = EnsureFeatureExists(syncedFeatures, f.Name+"-active", f.DisplayName+" [Active]", extraMetadata)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func EnsureFeatureExists(syncedFeatures map[string]SyncedFeature, name, displayName string, extraMetada map[string]string) error {
	id, exists, err := FeatureExists(name)
	if err != nil {
		return err
	}

	if exists {
		syncedFeatures[id] = SyncedFeature{name: name, displayName: displayName, stripeID: id}
		return nil
	}

	params := stripe.EntitlementsFeatureParams{
		Name:      &displayName,
		LookupKey: &name,
	}

	for key, value := range extraMetada {
		params.AddMetadata(key, value)
	}

	feature, err := stripefeatures.New(&params)
	if err != nil {
		return fmt.Errorf("failed to create Stripe feature from feature %s: %v\n", *params.LookupKey, err)
	}
	syncedFeatures[feature.ID] = SyncedFeature{name: *params.LookupKey, displayName: *params.Name, stripeID: feature.ID}
	return nil
}

func FeatureExists(id string) (string, bool, error) {
	list := stripefeatures.List(&stripe.EntitlementsFeatureListParams{
		LookupKey: &id,
	})
	if err := list.Err(); err != nil {
		return "", false, fmt.Errorf("failed to list features while check if feature [%s] exists: %w", id, err)
	}
	if !list.Next() {
		return "", false, nil
	}
	feature, ok := list.Current().(*stripe.EntitlementsFeature)
	if !ok {
		return "", false, fmt.Errorf("failed to ")
	}
	return feature.ID, true, nil
}

func EnsureFeatureProducts(syncedFeatures map[string]SyncedFeature) error {
	for _, feature := range syncedFeatures {
		if err := EnsureFeatureProduct(feature); err != nil {
			return err
		}
	}
	return nil
}

func EnsureFeatureProduct(syncedFeature SyncedFeature) error {
	productSearch := stripeproducts.Search(&stripe.ProductSearchParams{
		SearchParams: stripe.SearchParams{
			Query: fmt.Sprintf(metadataQueryFmt, licenseapi.MetadataKeyProductForFeature, syncedFeature.name),
		},
	})
	if err := productSearch.Err(); err != nil {
		return err
	}
	if productSearch.Next() {
		// a product exists with features name
		return nil
	}

	usdCurrencyCode := "usd"
	unit := int64(2000000) // sample placeholder price
	interval := "year"
	intervalCount := int64(1)
	product, err := stripeproducts.New(&stripe.ProductParams{
		Name: &syncedFeature.displayName,
		DefaultPriceData: &stripe.ProductDefaultPriceDataParams{
			Currency:   &usdCurrencyCode,
			UnitAmount: &unit,
			Recurring: &stripe.ProductDefaultPriceDataRecurringParams{
				Interval:      &interval,
				IntervalCount: &intervalCount,
			},
		},
		Metadata: map[string]string{
			licenseapi.MetadataKeyProductForFeature: syncedFeature.name,
		},
	})
	if err != nil {
		return err
	}

	_, err = productfeature.New(&stripe.ProductFeatureParams{
		Product:            &product.ID,
		EntitlementFeature: &syncedFeature.stripeID,
	})
	if err != nil {
		return err
	}
	return nil
}

func EnsureAttachAll(featureIDs map[string]SyncedFeature) error {
	productSearch := stripeproducts.Search(&stripe.ProductSearchParams{
		SearchParams: stripe.SearchParams{
			Query: fmt.Sprintf(metadataQueryFmt, licenseapi.MetadataKeyAttachAll, licenseapi.MetadataValueTrue),
		},
	})
	if err := productSearch.Err(); err != nil {
		return err
	}
	for productSearch.Next() {
		prod := productSearch.Product()
		featuresToCheck := maps.Clone(featureIDs)

		if err := SearchProductForFeatures(prod.ID, featuresToCheck); err != nil {
			return err
		}
		for featureID := range featuresToCheck {
			_, err := productfeature.New(&stripe.ProductFeatureParams{
				Product:            &prod.ID,
				EntitlementFeature: &featureID,
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func SearchProductForFeatures(productID string, featuresToCheck map[string]SyncedFeature) error {
	productFeaturesList := productfeature.List(&stripe.ProductFeatureListParams{
		Product: &productID,
	})
	for productFeaturesList.Next() {
		if err := productFeaturesList.Err(); err != nil {
			return err
		}
		delete(featuresToCheck, productFeaturesList.ProductFeature().EntitlementFeature.ID)
	}
	return nil
}
