// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type FunctionPackage struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function#bucket_name Function#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function#object_name Function#object_name}.
	ObjectName *string `field:"required" json:"objectName" yaml:"objectName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function#sha_256 Function#sha_256}.
	Sha256 *string `field:"optional" json:"sha256" yaml:"sha256"`
}

