// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageBucketConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#access_key StorageBucket#access_key}.
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#acl StorageBucket#acl}.
	Acl *string `field:"optional" json:"acl" yaml:"acl"`
	// anonymous_access_flags block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#anonymous_access_flags StorageBucket#anonymous_access_flags}
	AnonymousAccessFlags *StorageBucketAnonymousAccessFlags `field:"optional" json:"anonymousAccessFlags" yaml:"anonymousAccessFlags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#bucket StorageBucket#bucket}.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#bucket_prefix StorageBucket#bucket_prefix}.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// cors_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#cors_rule StorageBucket#cors_rule}
	CorsRule interface{} `field:"optional" json:"corsRule" yaml:"corsRule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#default_storage_class StorageBucket#default_storage_class}.
	DefaultStorageClass *string `field:"optional" json:"defaultStorageClass" yaml:"defaultStorageClass"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#folder_id StorageBucket#folder_id}.
	FolderId *string `field:"optional" json:"folderId" yaml:"folderId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#force_destroy StorageBucket#force_destroy}.
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// grant block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#grant StorageBucket#grant}
	Grant interface{} `field:"optional" json:"grant" yaml:"grant"`
	// https block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#https StorageBucket#https}
	Https *StorageBucketHttps `field:"optional" json:"https" yaml:"https"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#id StorageBucket#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// lifecycle_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#lifecycle_rule StorageBucket#lifecycle_rule}
	LifecycleRule interface{} `field:"optional" json:"lifecycleRule" yaml:"lifecycleRule"`
	// logging block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#logging StorageBucket#logging}
	Logging interface{} `field:"optional" json:"logging" yaml:"logging"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#max_size StorageBucket#max_size}.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#policy StorageBucket#policy}.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#secret_key StorageBucket#secret_key}.
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
	// server_side_encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#server_side_encryption_configuration StorageBucket#server_side_encryption_configuration}
	ServerSideEncryptionConfiguration *StorageBucketServerSideEncryptionConfiguration `field:"optional" json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
	// versioning block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#versioning StorageBucket#versioning}
	Versioning *StorageBucketVersioning `field:"optional" json:"versioning" yaml:"versioning"`
	// website block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#website StorageBucket#website}
	Website *StorageBucketWebsite `field:"optional" json:"website" yaml:"website"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#website_domain StorageBucket#website_domain}.
	WebsiteDomain *string `field:"optional" json:"websiteDomain" yaml:"websiteDomain"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#website_endpoint StorageBucket#website_endpoint}.
	WebsiteEndpoint *string `field:"optional" json:"websiteEndpoint" yaml:"websiteEndpoint"`
}

