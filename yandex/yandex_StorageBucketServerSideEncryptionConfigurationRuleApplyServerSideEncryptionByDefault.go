// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#kms_master_key_id StorageBucket#kms_master_key_id}.
	KmsMasterKeyId *string `field:"required" json:"kmsMasterKeyId" yaml:"kmsMasterKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#sse_algorithm StorageBucket#sse_algorithm}.
	SseAlgorithm *string `field:"required" json:"sseAlgorithm" yaml:"sseAlgorithm"`
}

