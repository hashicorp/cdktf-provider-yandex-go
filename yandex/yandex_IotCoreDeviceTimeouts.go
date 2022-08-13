// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type IotCoreDeviceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/iot_core_device#create IotCoreDevice#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/iot_core_device#delete IotCoreDevice#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/iot_core_device#update IotCoreDevice#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

