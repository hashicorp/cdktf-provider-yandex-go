// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type StorageBucketAnonymousAccessFlags struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#list StorageBucket#list}.
	List interface{} `field:"optional" json:"list" yaml:"list"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#read StorageBucket#read}.
	Read interface{} `field:"optional" json:"read" yaml:"read"`
}

