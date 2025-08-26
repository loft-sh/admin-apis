package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/loft-sh/admin-apis/hack/internal/yamlparser"
	"github.com/loft-sh/admin-apis/pkg/licenseapi"
)

const (
	featuresFileTemplate = `package licenseapi

// This code was generated. Change features.yaml to add, remove, or edit features.

// Features
const (
%s
)

func GetFeatures() []FeatureName {
	return []FeatureName{
%s
	}
}
`
	defaultLicenseFileTemplate = `package licenseapi

// This code was generated. Change features.yaml to add, remove, or edit features.

import (
	"cmp"
	"slices"
)

func New() *License {
	limits := make([]*Limit, 0, len(Limits))
	for _, limit := range Limits {
		limits = append(limits, limit)
	}
	slices.SortFunc(limits, func(a, b *Limit) int {
		return cmp.Compare(a.Name, b.Name)
	})

	// Sorting features by module is not requires here. However, to maintain backwards compatibility, the structure of
	// features being contained within a module is still necessary. Therefore, all features are now returned in one module.
	return &License{
		Modules: []*Module{
			{
				DisplayName: "All Features",
				Name:        string(VirtualClusterModule),
				Limits:      limits,
				Features: []*Feature{
%s
				},
			},
		},
	}
}
`

	featuresAllowedBeforeFileTemplate = `package licenseapi

// This code was generated. Change features.yaml to add, remove, or edit features.

import (
	"errors"
	"time"
)

var errNoAllowBefore = errors.New("feature not allowed before license's issued date")

// featureToAllowBefore maps feature names to their corresponding
// RFC3339-formatted allowBefore timestamps. If a license was issued before this
// timestamp, the feature is allowed even if it is not explicitly included in the license.
var featuresToAllowBefore = map[FeatureName]string{
%s
}

// GetFeaturesAllowedBefore returns list of features
// to be allowed before license's issued time
func GetFeaturesAllowedBefore() []FeatureName {
	return []FeatureName{
%s
	}
}

// AllowedBeforeTime returns the parsed allowBefore time for a given feature.
// If the feature does not have an allowBefore date, it returns errNoAllowBefore.
// If the date is present but invalid, it returns the corresponding parsing error.
func AllowedBeforeTime(featureName FeatureName) (*time.Time, error) {
	if date, exists := featuresToAllowBefore[featureName]; exists {
		t, err := time.Parse(time.RFC3339, date)
		if err != nil {
			return nil, err
		}
		return &t, nil

	}
	return nil, errNoAllowBefore
}

// IsAllowBeforeNotDefined determines whether the provided error is
// errNoAllowBefore, indicating that the feature has no allowBefore date.
func IsAllowBeforeNotDefined(err error) bool {
	return errors.Is(err, errNoAllowBefore)
}
`
)

var (
	// the process that maps the hyphenated name to a camel cased name
	// assumes every hyphen delimited string should stay the same but with
	// the leading letter capitalized, e.g. custom-storage-driver, will be
	// CustomStorageDriver. Any exceptions to the aforementioned assumption
	// should be mapped here, e.g. vcluster-auth-sso can become VirtualClusterAuthSSO
	// by adding the mapping `"vcluster": "VirtualCluster"` and `"sso": "SSO"`.
	aliasLookup = map[string]string{
		"authentication": "Auth",
		"vcp":            "VirtualClusterPro",
		"vclusters":      "VirtualCluster",
		"vcluster":       "VirtualCluster",
		"vnode":          "VNode",
		"ui":             "UI",
		"sso":            "SSO",
		"oidc":           "OIDC",
		"ha":             "HighAvailability",
		"coredns":        "CoreDNS",
		"cp":             "ControlPlane",
		"db":             "DB",
	}
	reg = regexp.MustCompile(`^([a-zA-Z]+)|(-[a-zA-Z]+)`)
)

func main() {
	yamlContent := struct {
		Features []*licenseapi.Feature `json:"features"`
	}{}

	err := yamlparser.ParseYAML("../../definitions/features.yaml", &yamlContent)
	if err != nil {
		panic(err)
	}

	features := yamlContent.Features

	f, err := os.Create("../../pkg/licenseapi/features.go")
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString(fmt.Sprintf(featuresFileTemplate, generateFeatureConstantsBody(features), generateFeatureSliceBody(features)))
	if err != nil {
		panic(err)
	}

	f, err = os.Create("../../pkg/licenseapi/features_allowed_before.go")
	if err != nil {
		panic(err)
	}

	allowBeforeMap, AllowBeforeList := generateFeatureAllowedBeforeMap(features)
	_, err = f.WriteString(fmt.Sprintf(featuresAllowedBeforeFileTemplate, allowBeforeMap, AllowBeforeList))
	if err != nil {
		panic(err)
	}

	f, err = os.Create("../../pkg/licenseapi/license_new.go")
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString(fmt.Sprintf(defaultLicenseFileTemplate, generatedDefaultLicenseBody(features)))
	if err != nil {
		panic(err)
	}
}

func generateFeatureAllowedBeforeMap(features []*licenseapi.Feature) (string, string) {
	var featureAllowBeforeMap string
	var featureAllowedBeforeList string
	for _, feature := range features {
		if feature.AllowBefore != "" {
			if _, err := time.Parse(time.RFC3339, feature.AllowBefore); err != nil {
				panic(err)
			}
			featureAllowBeforeMap += fmt.Sprintf("\t%s: %q,\n", hyphenatedToCamelCase(replaceAliasWithFull(feature.Name)), feature.AllowBefore)
			featureAllowedBeforeList += fmt.Sprintf(`		%s,`, hyphenatedToCamelCase(replaceAliasWithFull(feature.Name)))
		}
	}
	return strings.TrimSuffix(featureAllowBeforeMap, "\n"), strings.TrimSuffix(featureAllowedBeforeList, "\n")
}

func generateFeatureConstantsBody(features []*licenseapi.Feature) string {
	featureConstants := ""
	for _, feature := range features {
		featureConstants += fmt.Sprintf(`	%s FeatureName = "%s" // %s

`, hyphenatedToCamelCase(replaceAliasWithFull(feature.Name)), feature.Name, feature.DisplayName)
	}
	return strings.TrimSuffix(featureConstants, "\n")
}

func generateFeatureSliceBody(features []*licenseapi.Feature) string {
	featuresList := ""
	for _, feature := range features {
		featuresList += fmt.Sprintf(`		%s,
`, hyphenatedToCamelCase(replaceAliasWithFull(feature.Name)))
	}
	return strings.TrimSuffix(featuresList, "\n")
}

func replaceAliasWithFull(feature string) string {
	for alias, full := range aliasLookup {
		if feature == alias {
			return full
		}
		cutFeature, ok := strings.CutPrefix(feature, alias+"-")
		if ok {
			feature = full + "-" + cutFeature
		}
		cutFeature, ok = strings.CutSuffix(feature, "-"+alias)
		if ok {
			feature = cutFeature + "-" + full
		}
		feature = strings.ReplaceAll(feature, "-"+alias+"-", "-"+full+"-")
	}
	return feature
}

func generatedDefaultLicenseBody(features []*licenseapi.Feature) string {
	moduleFeatures := ""
	for _, feature := range features {
		moduleFeatures += fmt.Sprintf(`					{
						DisplayName: "%s",
						Name:        "%s",
					},
`, feature.DisplayName, feature.Name)
	}
	return strings.TrimSuffix(moduleFeatures, "\n")
}

func hyphenatedToCamelCase(name string) string {
	return reg.ReplaceAllStringFunc(name, func(s string) string {
		return strings.ToUpper(string(strings.TrimPrefix(s, "-")[0])) + strings.TrimPrefix(s, "-")[1:]
	})
}
