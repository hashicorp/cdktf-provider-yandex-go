// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type StorageBucketServerSideEncryptionConfiguration struct {
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#rule StorageBucket#rule}
	Rule *StorageBucketServerSideEncryptionConfigurationRule `field:"required" json:"rule" yaml:"rule"`
}

