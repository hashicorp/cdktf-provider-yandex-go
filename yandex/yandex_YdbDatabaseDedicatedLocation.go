// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type YdbDatabaseDedicatedLocation struct {
	// region block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/ydb_database_dedicated#region YdbDatabaseDedicated#region}
	Region *YdbDatabaseDedicatedLocationRegion `field:"optional" json:"region" yaml:"region"`
}

