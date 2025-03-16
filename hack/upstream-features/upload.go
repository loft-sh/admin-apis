package main

import (
	"log"
	"os"

	"github.com/loft-sh/admin-apis/hack/internal/yamlparser"
	"github.com/loft-sh/admin-apis/pkg/licenseapi"
	"github.com/loft-sh/admin-apis/pkg/util/features"
	"github.com/stripe/stripe-go/v81/client"
)

func main() {
	stripeKey := os.Getenv("STRIPE_API_KEY")
	if stripeKey == "" {
		log.Fatal("stripe token cannot be empty")
	}
	stripeClient := &client.API{}
	stripeClient.Init(stripeKey, nil)

	syncedFeatures := map[string]features.SyncedFeature{}

	yamlContent := struct {
		Features []*licenseapi.Feature `json:"features"`
		Limits   []*licenseapi.Feature `json:"limits"`
	}{}

	err := yamlparser.ParseYAML("definitions/features.yaml", &yamlContent)
	if err != nil {
		log.Fatal(err)
	}

	err = features.EnsureFeatures(stripeClient, syncedFeatures, yamlContent.Features, false)
	if err != nil {
		log.Fatal(err)
	}

	err = yamlparser.ParseYAML("definitions/limits.yaml", &yamlContent)
	if err != nil {
		log.Fatal(err)
	}

	err = features.EnsureFeatures(stripeClient, syncedFeatures, yamlContent.Limits, true)
	if err != nil {
		log.Fatal(err)
	}
}
