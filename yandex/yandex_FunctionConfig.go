// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FunctionConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function#entrypoint Function#entrypoint}.
	Entrypoint *string `field:"required" json:"entrypoint" yaml:"entrypoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function#memory Function#memory}.
	Memory *float64 `field:"required" json:"memory" yaml:"memory"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function#name Function#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function#runtime Function#runtime}.
	Runtime *string `field:"required" json:"runtime" yaml:"runtime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function#user_hash Function#user_hash}.
	UserHash *string `field:"required" json:"userHash" yaml:"userHash"`
	// content block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function#content Function#content}
	Content *FunctionContent `field:"optional" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function#description Function#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function#environment Function#environment}.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function#execution_timeout Function#execution_timeout}.
	ExecutionTimeout *string `field:"optional" json:"executionTimeout" yaml:"executionTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function#folder_id Function#folder_id}.
	FolderId *string `field:"optional" json:"folderId" yaml:"folderId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function#id Function#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function#labels Function#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// package block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function#package Function#package}
	Package *FunctionPackage `field:"optional" json:"package" yaml:"package"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function#service_account_id Function#service_account_id}.
	ServiceAccountId *string `field:"optional" json:"serviceAccountId" yaml:"serviceAccountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function#tags Function#tags}.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function#timeouts Function#timeouts}
	Timeouts *FunctionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

