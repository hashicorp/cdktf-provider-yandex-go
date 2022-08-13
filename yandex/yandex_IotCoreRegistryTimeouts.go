// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type IotCoreRegistryTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/iot_core_registry#create IotCoreRegistry#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/iot_core_registry#delete IotCoreRegistry#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/iot_core_registry#update IotCoreRegistry#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

