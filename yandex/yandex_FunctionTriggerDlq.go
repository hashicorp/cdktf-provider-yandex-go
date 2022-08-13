// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type FunctionTriggerDlq struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#queue_id FunctionTrigger#queue_id}.
	QueueId *string `field:"required" json:"queueId" yaml:"queueId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#service_account_id FunctionTrigger#service_account_id}.
	ServiceAccountId *string `field:"required" json:"serviceAccountId" yaml:"serviceAccountId"`
}

