package main

import (
	"fmt"
	"log"
	"maps"
	"os"

	"github.com/loft-sh/admin-apis/hack/internal/featuresyaml"
	"github.com/loft-sh/admin-apis/pkg/licenseapi"
	"github.com/stripe/stripe-go/v81"
	stripefeatures "github.com/stripe/stripe-go/v81/entitlements/feature"
	stripeproducts "github.com/stripe/stripe-go/v81/product"
	"github.com/stripe/stripe-go/v81/productfeature"
)

const (
	metadataQueryFmt             = "metadata['%s']:'%s'"
	metadataKeyAttachAll         = "attach_all_features"
	metadataKeyProductForFeature = "product_for_feature"
)

type syncedFeature struct {
	name        string
	displayName string
	stripeID    string
}

func main() {
	stripeToken := os.Getenv("STRIPE_API_KEY")
	if stripeToken == "" {
		log.Println("stripe token cannot be empty")
		os.Exit(1)
	}
	stripe.Key = stripeToken

	features, err := featuresyaml.ReadFeaturesYaml("pkg/licenseapi/features.yaml")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	syncedFeatures := syncStripeFeatures(features)

	if err = ensureFeatureProducts(syncedFeatures); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if err = ensureAttachAll(syncedFeatures); err != nil {
		log.Println(err)
		os.Exit(1)
	}

}

func syncStripeFeatures(features []*licenseapi.Feature) map[string]syncedFeature {
	syncedFeatures := make(map[string]syncedFeature, len(features))
	for _, f := range features {
		feature, err := ensureFeatureExists(f.Name, f.DisplayName)
		if err != nil {
			log.Println(err)
			continue
		}
		syncedFeatures[feature.stripeID] = feature

		if !f.Preview {
			continue
		}

		previewFeature, err := ensureFeatureExists(f.Name+"-preview", "Preview: "+f.DisplayName)
		if err != nil {
			log.Println(err)
			continue
		}
		syncedFeatures[previewFeature.stripeID] = previewFeature
	}
	return syncedFeatures
}

func ensureFeatureExists(name, displayName string) (syncedFeature, error) {
	id, exists, err := featureExists(name)
	if err != nil {
		return syncedFeature{}, err
	}

	if exists {
		return syncedFeature{name: name, displayName: displayName, stripeID: id}, nil
	}

	feature, err := stripefeatures.New(&stripe.EntitlementsFeatureParams{
		Name:      &displayName,
		LookupKey: &name,
	})
	if err != nil {
		return syncedFeature{}, fmt.Errorf("failed to create Stripe feature from feature %s: %v\n", name, err)
	}
	return syncedFeature{name: name, displayName: displayName, stripeID: feature.ID}, nil
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
			Query: fmt.Sprintf(metadataQueryFmt, metadataKeyProductForFeature, syncedFeature.name),
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
	unit := int64(2000000) // =20k, this is in cents
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
			metadataKeyProductForFeature: syncedFeature.name,
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
			Query: fmt.Sprintf(metadataQueryFmt, metadataKeyAttachAll, "true"),
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
