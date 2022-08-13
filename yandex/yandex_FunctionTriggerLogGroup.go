// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type FunctionTriggerLogGroup struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#batch_cutoff FunctionTrigger#batch_cutoff}.
	BatchCutoff *string `field:"required" json:"batchCutoff" yaml:"batchCutoff"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#log_group_ids FunctionTrigger#log_group_ids}.
	LogGroupIds *[]*string `field:"required" json:"logGroupIds" yaml:"logGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#batch_size FunctionTrigger#batch_size}.
	BatchSize *string `field:"optional" json:"batchSize" yaml:"batchSize"`
}

