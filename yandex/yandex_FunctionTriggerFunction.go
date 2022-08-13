// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type FunctionTriggerFunction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#id FunctionTrigger#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#retry_attempts FunctionTrigger#retry_attempts}.
	RetryAttempts *string `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#retry_interval FunctionTrigger#retry_interval}.
	RetryInterval *string `field:"optional" json:"retryInterval" yaml:"retryInterval"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#service_account_id FunctionTrigger#service_account_id}.
	ServiceAccountId *string `field:"optional" json:"serviceAccountId" yaml:"serviceAccountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#tag FunctionTrigger#tag}.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

