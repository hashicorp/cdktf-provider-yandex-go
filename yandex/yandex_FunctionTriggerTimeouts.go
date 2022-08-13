// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type FunctionTriggerTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#create FunctionTrigger#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#delete FunctionTrigger#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#update FunctionTrigger#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

