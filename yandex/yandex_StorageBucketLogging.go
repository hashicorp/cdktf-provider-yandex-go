// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type StorageBucketLogging struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#target_bucket StorageBucket#target_bucket}.
	TargetBucket *string `field:"required" json:"targetBucket" yaml:"targetBucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#target_prefix StorageBucket#target_prefix}.
	TargetPrefix *string `field:"optional" json:"targetPrefix" yaml:"targetPrefix"`
}

