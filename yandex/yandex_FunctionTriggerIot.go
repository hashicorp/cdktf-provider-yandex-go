// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type FunctionTriggerIot struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#registry_id FunctionTrigger#registry_id}.
	RegistryId *string `field:"required" json:"registryId" yaml:"registryId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#device_id FunctionTrigger#device_id}.
	DeviceId *string `field:"optional" json:"deviceId" yaml:"deviceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#topic FunctionTrigger#topic}.
	Topic *string `field:"optional" json:"topic" yaml:"topic"`
}

