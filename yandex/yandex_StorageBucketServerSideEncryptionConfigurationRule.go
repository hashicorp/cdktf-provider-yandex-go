// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type StorageBucketServerSideEncryptionConfigurationRule struct {
	// apply_server_side_encryption_by_default block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#apply_server_side_encryption_by_default StorageBucket#apply_server_side_encryption_by_default}
	ApplyServerSideEncryptionByDefault *StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault `field:"required" json:"applyServerSideEncryptionByDefault" yaml:"applyServerSideEncryptionByDefault"`
}

