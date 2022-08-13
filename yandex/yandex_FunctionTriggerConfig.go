// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FunctionTriggerConfig struct {
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
	// function block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#function FunctionTrigger#function}
	Function *FunctionTriggerFunction `field:"required" json:"function" yaml:"function"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#name FunctionTrigger#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#description FunctionTrigger#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// dlq block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#dlq FunctionTrigger#dlq}
	Dlq *FunctionTriggerDlq `field:"optional" json:"dlq" yaml:"dlq"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#folder_id FunctionTrigger#folder_id}.
	FolderId *string `field:"optional" json:"folderId" yaml:"folderId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#id FunctionTrigger#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// iot block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#iot FunctionTrigger#iot}
	Iot *FunctionTriggerIot `field:"optional" json:"iot" yaml:"iot"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#labels FunctionTrigger#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// logging block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#logging FunctionTrigger#logging}
	Logging *FunctionTriggerLogging `field:"optional" json:"logging" yaml:"logging"`
	// log_group block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#log_group FunctionTrigger#log_group}
	LogGroup *FunctionTriggerLogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// message_queue block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#message_queue FunctionTrigger#message_queue}
	MessageQueue *FunctionTriggerMessageQueue `field:"optional" json:"messageQueue" yaml:"messageQueue"`
	// object_storage block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#object_storage FunctionTrigger#object_storage}
	ObjectStorage *FunctionTriggerObjectStorage `field:"optional" json:"objectStorage" yaml:"objectStorage"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#timeouts FunctionTrigger#timeouts}
	Timeouts *FunctionTriggerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// timer block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#timer FunctionTrigger#timer}
	Timer *FunctionTriggerTimer `field:"optional" json:"timer" yaml:"timer"`
}

