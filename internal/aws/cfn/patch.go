package cfn

import (
	"errors"
	"fmt"

	"github.com/aws-cloudformation/rain/internal/aws/lightsail"
)

func patchLightsailInstance(schema *Schema) error {
	blueprintId, found := schema.Properties["BlueprintId"]
	if !found {
		return fmt.Errorf("expected AWS::Lightsail::Instance to have BlueprintId")
	}
	blueprints, err := lightsail.GetBlueprints()
	if err != nil {
		return fmt.Errorf("unable to call aws api to get available lightsail blueprints")
	}
	blueprintId.Enum = blueprints

	bundleId, found := schema.Properties["BundleId"]
	if !found {
		return fmt.Errorf("expected AWS::Lightsail::Instance to have BundleId")
	}
	bundles, err := lightsail.GetBundles()
	if err != nil {
		return fmt.Errorf("unable to call aws api to get available lightsail bundles")
	}
	bundleId.Enum = bundles

	return nil
}

func patchLightsailBucket(schema *Schema) error {
	bundleId, found := schema.Properties["BundleId"]
	if !found {
		return fmt.Errorf("expected AWS::Lightsail::Bucket to have BundleId")
	}
	bundles, err := lightsail.GetBucketBundles()
	if err != nil {
		return fmt.Errorf("unable to call aws api to get available lightsail bucket bundles")
	}
	bundleId.Enum = bundles

	return nil
}

func patchLightsailDistribution(schema *Schema) error {
	bundleId, found := schema.Properties["BundleId"]
	if !found {
		return fmt.Errorf("expected AWS::Lightsail::Distribution to have BundleId")
	}
	bundles, err := lightsail.GetDistributionBundles()
	if err != nil {
		return fmt.Errorf("unable to call aws api to get available lightsail distribution bundles")
	}
	bundleId.Enum = bundles

	return nil
}

func patchLightsailDatabase(schema *Schema) error {
	blueprintId, found := schema.Properties["RelationalDatabaseBlueprintId"]
	if !found {
		return fmt.Errorf("expected AWS::Lightsail::Database to have RelationalDatabaseBlueprintId")
	}
	blueprints, err := lightsail.GetRelationalDatabaseBlueprints()
	if err != nil {
		return fmt.Errorf("unable to call aws api to get available lightsail blueprints")
	}
	blueprintId.Enum = blueprints

	bundleId, found := schema.Properties["RelationalDatabaseBundleId"]
	if !found {
		return fmt.Errorf("expected AWS::Lightsail::Database to have RelationalDatabaseBundleId")
	}
	bundles, err := lightsail.GetRelationalDatabaseBundles()
	if err != nil {
		return fmt.Errorf("unable to call aws api to get available lightsail bundles")
	}
	bundleId.Enum = bundles

	return nil
}

func patchLightsailAlarm(schema *Schema) error {
	// These are documented but not in the schema
	valid := []string{
		"GreaterThanOrEqualToThreshold",
		"GreaterThanThreshold",
		"LessThanThreshold",
		"LessThanOrEqualToThreshold",
	}
	comparisonOperator, found := schema.Properties["ComparisonOperator"]
	if !found {
		return errors.New("expected AWS::Lightsail::Alarm to have ComparisonOperator")
	}
	comparisonOperator.Enum = valid
	return nil
}
