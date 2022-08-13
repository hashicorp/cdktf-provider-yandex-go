// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type StorageBucketLifecycleRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#enabled StorageBucket#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#abort_incomplete_multipart_upload_days StorageBucket#abort_incomplete_multipart_upload_days}.
	AbortIncompleteMultipartUploadDays *float64 `field:"optional" json:"abortIncompleteMultipartUploadDays" yaml:"abortIncompleteMultipartUploadDays"`
	// expiration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#expiration StorageBucket#expiration}
	Expiration *StorageBucketLifecycleRuleExpiration `field:"optional" json:"expiration" yaml:"expiration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#id StorageBucket#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// noncurrent_version_expiration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#noncurrent_version_expiration StorageBucket#noncurrent_version_expiration}
	NoncurrentVersionExpiration *StorageBucketLifecycleRuleNoncurrentVersionExpiration `field:"optional" json:"noncurrentVersionExpiration" yaml:"noncurrentVersionExpiration"`
	// noncurrent_version_transition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#noncurrent_version_transition StorageBucket#noncurrent_version_transition}
	NoncurrentVersionTransition interface{} `field:"optional" json:"noncurrentVersionTransition" yaml:"noncurrentVersionTransition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#prefix StorageBucket#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// transition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#transition StorageBucket#transition}
	Transition interface{} `field:"optional" json:"transition" yaml:"transition"`
}

