// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type StorageBucketLifecycleRuleExpiration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#date StorageBucket#date}.
	Date *string `field:"optional" json:"date" yaml:"date"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#days StorageBucket#days}.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#expired_object_delete_marker StorageBucket#expired_object_delete_marker}.
	ExpiredObjectDeleteMarker interface{} `field:"optional" json:"expiredObjectDeleteMarker" yaml:"expiredObjectDeleteMarker"`
}

