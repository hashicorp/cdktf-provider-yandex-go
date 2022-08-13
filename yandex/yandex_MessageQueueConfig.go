// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MessageQueueConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/message_queue#access_key MessageQueue#access_key}.
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/message_queue#content_based_deduplication MessageQueue#content_based_deduplication}.
	ContentBasedDeduplication interface{} `field:"optional" json:"contentBasedDeduplication" yaml:"contentBasedDeduplication"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/message_queue#delay_seconds MessageQueue#delay_seconds}.
	DelaySeconds *float64 `field:"optional" json:"delaySeconds" yaml:"delaySeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/message_queue#fifo_queue MessageQueue#fifo_queue}.
	FifoQueue interface{} `field:"optional" json:"fifoQueue" yaml:"fifoQueue"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/message_queue#id MessageQueue#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/message_queue#max_message_size MessageQueue#max_message_size}.
	MaxMessageSize *float64 `field:"optional" json:"maxMessageSize" yaml:"maxMessageSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/message_queue#message_retention_seconds MessageQueue#message_retention_seconds}.
	MessageRetentionSeconds *float64 `field:"optional" json:"messageRetentionSeconds" yaml:"messageRetentionSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/message_queue#name MessageQueue#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/message_queue#name_prefix MessageQueue#name_prefix}.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/message_queue#receive_wait_time_seconds MessageQueue#receive_wait_time_seconds}.
	ReceiveWaitTimeSeconds *float64 `field:"optional" json:"receiveWaitTimeSeconds" yaml:"receiveWaitTimeSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/message_queue#redrive_policy MessageQueue#redrive_policy}.
	RedrivePolicy *string `field:"optional" json:"redrivePolicy" yaml:"redrivePolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/message_queue#region_id MessageQueue#region_id}.
	RegionId *string `field:"optional" json:"regionId" yaml:"regionId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/message_queue#secret_key MessageQueue#secret_key}.
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/message_queue#visibility_timeout_seconds MessageQueue#visibility_timeout_seconds}.
	VisibilityTimeoutSeconds *float64 `field:"optional" json:"visibilityTimeoutSeconds" yaml:"visibilityTimeoutSeconds"`
}

