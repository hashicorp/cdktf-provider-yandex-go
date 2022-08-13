// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServerlessContainerConfig struct {
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
	// image block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/serverless_container#image ServerlessContainer#image}
	Image *ServerlessContainerImage `field:"required" json:"image" yaml:"image"`
	// Container memory in megabytes, should be aligned to 128.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/serverless_container#memory ServerlessContainer#memory}
	Memory *float64 `field:"required" json:"memory" yaml:"memory"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/serverless_container#name ServerlessContainer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/serverless_container#concurrency ServerlessContainer#concurrency}.
	Concurrency *float64 `field:"optional" json:"concurrency" yaml:"concurrency"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/serverless_container#core_fraction ServerlessContainer#core_fraction}.
	CoreFraction *float64 `field:"optional" json:"coreFraction" yaml:"coreFraction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/serverless_container#cores ServerlessContainer#cores}.
	Cores *float64 `field:"optional" json:"cores" yaml:"cores"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/serverless_container#description ServerlessContainer#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/serverless_container#execution_timeout ServerlessContainer#execution_timeout}.
	ExecutionTimeout *string `field:"optional" json:"executionTimeout" yaml:"executionTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/serverless_container#folder_id ServerlessContainer#folder_id}.
	FolderId *string `field:"optional" json:"folderId" yaml:"folderId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/serverless_container#id ServerlessContainer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/serverless_container#labels ServerlessContainer#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/serverless_container#service_account_id ServerlessContainer#service_account_id}.
	ServiceAccountId *string `field:"optional" json:"serviceAccountId" yaml:"serviceAccountId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/serverless_container#timeouts ServerlessContainer#timeouts}
	Timeouts *ServerlessContainerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

