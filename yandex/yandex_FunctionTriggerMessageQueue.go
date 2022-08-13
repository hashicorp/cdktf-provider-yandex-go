// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type FunctionTriggerMessageQueue struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#batch_cutoff FunctionTrigger#batch_cutoff}.
	BatchCutoff *string `field:"required" json:"batchCutoff" yaml:"batchCutoff"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#queue_id FunctionTrigger#queue_id}.
	QueueId *string `field:"required" json:"queueId" yaml:"queueId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#service_account_id FunctionTrigger#service_account_id}.
	ServiceAccountId *string `field:"required" json:"serviceAccountId" yaml:"serviceAccountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#batch_size FunctionTrigger#batch_size}.
	BatchSize *string `field:"optional" json:"batchSize" yaml:"batchSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#visibility_timeout FunctionTrigger#visibility_timeout}.
	VisibilityTimeout *string `field:"optional" json:"visibilityTimeout" yaml:"visibilityTimeout"`
}

