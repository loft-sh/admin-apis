package main

import (
	"fmt"
	"log"
	"maps"
	"os"

	"github.com/loft-sh/admin-apis/hack/internal/yamlparser"
	"github.com/loft-sh/admin-apis/pkg/licenseapi"
	"github.com/stripe/stripe-go/v81"
	stripefeatures "github.com/stripe/stripe-go/v81/entitlements/feature"
	stripeproducts "github.com/stripe/stripe-go/v81/product"
	"github.com/stripe/stripe-go/v81/productfeature"
)

const (
	metadataQueryFmt = "metadata['%s']:'%s'"
)

type syncedFeature struct {
	name        string
	displayName string
	stripeID    string
}

func main() {
	stripeToken := os.Getenv("STRIPE_API_KEY")
	if stripeToken == "" {
		log.Fatal("stripe token cannot be empty")
	}
	stripe.Key = stripeToken

	syncedFeatures := map[string]syncedFeature{}

	yamlContent := struct {
		Features []*licenseapi.Feature `json:"features"`
		Limits   []*licenseapi.Feature `json:"limits"`
	}{}

	err := yamlparser.ParseYAML("definitions/features.yaml", &yamlContent)
	if err != nil {
		log.Fatal(err)
	}

	err = createFeatures(syncedFeatures, yamlContent.Features, false)
	if err != nil {
		log.Fatal(err)
	}

	err = yamlparser.ParseYAML("definitions/limits.yaml", &yamlContent)
	if err != nil {
		log.Fatal(err)
	}

	err = createFeatures(syncedFeatures, yamlContent.Limits, true)
	if err != nil {
		log.Fatal(err)
	}
}

func createFeatures(syncedFeatures map[string]syncedFeature, features []*licenseapi.Feature, isLimit bool) error {
	err := ensureStripeFeatures(syncedFeatures, features, isLimit)
	if err != nil {
		return err
	}

	if err = ensureFeatureProducts(syncedFeatures); err != nil {
		return err
	}

	if err = ensureAttachAll(syncedFeatures); err != nil {
		return err
	}
	return nil
}

func ensureStripeFeatures(syncedFeatures map[string]syncedFeature, features []*licenseapi.Feature, isLimit bool) error {
	for _, f := range features {
		extraMetadata := map[string]string{}
		if isLimit {
			extraMetadata[licenseapi.MetadataKeyFeatureIsLimit] = licenseapi.MetadataValueTrue
			f.Name = licenseapi.LimitsPrefix + f.Name
		}

		if f.Status == string(licenseapi.FeatureStatusHidden) {
			extraMetadata[licenseapi.MetadataKeyFeatureIsHidden] = licenseapi.MetadataValueTrue
		}

		err := ensureFeatureExists(syncedFeatures, f.Name, f.DisplayName, extraMetadata)
		if err != nil {
			return err
		}

		if f.Preview {
			extraMetadata[licenseapi.MetadataKeyFeatureIsPreview] = licenseapi.MetadataValueTrue
			err = ensureFeatureExists(syncedFeatures, f.Name+"-preview", f.DisplayName+" [Preview]", extraMetadata)
			if err != nil {
				return err
			}
		}

		if isLimit {
			extraMetadata[licenseapi.MetadataKeyFeatureLimitType] = licenseapi.MetadataKeyFeatureLimitTypeActive
			err = ensureFeatureExists(syncedFeatures, f.Name+"-active", f.DisplayName+" [Active]", extraMetadata)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func ensureFeatureExists(syncedFeatures map[string]syncedFeature, name, displayName string, extraMetada map[string]string) error {
	id, exists, err := featureExists(name)
	if err != nil {
		return err
	}

	if exists {
		syncedFeatures[id] = syncedFeature{name: name, displayName: displayName, stripeID: id}
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
	syncedFeatures[feature.ID] = syncedFeature{name: *params.LookupKey, displayName: *params.Name, stripeID: feature.ID}
	return nil
}

func featureExists(id string) (string, bool, error) {
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

func ensureFeatureProducts(syncedFeatures map[string]syncedFeature) error {
	for _, feature := range syncedFeatures {
		if err := ensureFeatureProduct(feature); err != nil {
			return err
		}
	}
	return nil
}

func ensureFeatureProduct(syncedFeature syncedFeature) error {
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
	unit := int64(2000000) // =20k, this is in cents (sample placeholder price)
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

func ensureAttachAll(featureIDs map[string]syncedFeature) error {
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
		featuresToCheck := maps.Clone[map[string]syncedFeature](featureIDs)

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

func SearchProductForFeatures(productID string, featuresToCheck map[string]syncedFeature) error {
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
