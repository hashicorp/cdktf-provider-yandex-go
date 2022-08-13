// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type YdbDatabaseDedicatedScalePolicyFixedScale struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/ydb_database_dedicated#size YdbDatabaseDedicated#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
}

