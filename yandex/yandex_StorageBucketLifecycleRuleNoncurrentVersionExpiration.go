// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type StorageBucketLifecycleRuleNoncurrentVersionExpiration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#days StorageBucket#days}.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
}

