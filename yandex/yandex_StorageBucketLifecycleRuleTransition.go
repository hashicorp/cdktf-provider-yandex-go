// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type StorageBucketLifecycleRuleTransition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#storage_class StorageBucket#storage_class}.
	StorageClass *string `field:"required" json:"storageClass" yaml:"storageClass"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#date StorageBucket#date}.
	Date *string `field:"optional" json:"date" yaml:"date"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#days StorageBucket#days}.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
}

