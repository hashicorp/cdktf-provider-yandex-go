// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type YdbDatabaseDedicatedStorageConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/ydb_database_dedicated#group_count YdbDatabaseDedicated#group_count}.
	GroupCount *float64 `field:"required" json:"groupCount" yaml:"groupCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/ydb_database_dedicated#storage_type_id YdbDatabaseDedicated#storage_type_id}.
	StorageTypeId *string `field:"required" json:"storageTypeId" yaml:"storageTypeId"`
}

