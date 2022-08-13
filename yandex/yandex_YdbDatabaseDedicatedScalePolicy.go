// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type YdbDatabaseDedicatedScalePolicy struct {
	// fixed_scale block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/ydb_database_dedicated#fixed_scale YdbDatabaseDedicated#fixed_scale}
	FixedScale *YdbDatabaseDedicatedScalePolicyFixedScale `field:"required" json:"fixedScale" yaml:"fixedScale"`
}

