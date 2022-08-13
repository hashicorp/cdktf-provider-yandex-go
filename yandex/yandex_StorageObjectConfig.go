// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageObjectConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_object#bucket StorageObject#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_object#key StorageObject#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_object#access_key StorageObject#access_key}.
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_object#acl StorageObject#acl}.
	Acl *string `field:"optional" json:"acl" yaml:"acl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_object#content StorageObject#content}.
	Content *string `field:"optional" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_object#content_base64 StorageObject#content_base64}.
	ContentBase64 *string `field:"optional" json:"contentBase64" yaml:"contentBase64"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_object#content_type StorageObject#content_type}.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_object#id StorageObject#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_object#secret_key StorageObject#secret_key}.
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_object#source StorageObject#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
}

