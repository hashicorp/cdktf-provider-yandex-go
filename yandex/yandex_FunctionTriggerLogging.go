// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type FunctionTriggerLogging struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#batch_cutoff FunctionTrigger#batch_cutoff}.
	BatchCutoff *string `field:"required" json:"batchCutoff" yaml:"batchCutoff"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#group_id FunctionTrigger#group_id}.
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#levels FunctionTrigger#levels}.
	Levels *[]*string `field:"required" json:"levels" yaml:"levels"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#resource_ids FunctionTrigger#resource_ids}.
	ResourceIds *[]*string `field:"required" json:"resourceIds" yaml:"resourceIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#resource_types FunctionTrigger#resource_types}.
	ResourceTypes *[]*string `field:"required" json:"resourceTypes" yaml:"resourceTypes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#batch_size FunctionTrigger#batch_size}.
	BatchSize *string `field:"optional" json:"batchSize" yaml:"batchSize"`
}

