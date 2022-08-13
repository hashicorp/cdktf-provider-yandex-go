// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotCoreDeviceConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/iot_core_device#name IotCoreDevice#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/iot_core_device#registry_id IotCoreDevice#registry_id}.
	RegistryId *string `field:"required" json:"registryId" yaml:"registryId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/iot_core_device#aliases IotCoreDevice#aliases}.
	Aliases *map[string]*string `field:"optional" json:"aliases" yaml:"aliases"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/iot_core_device#certificates IotCoreDevice#certificates}.
	Certificates *[]*string `field:"optional" json:"certificates" yaml:"certificates"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/iot_core_device#description IotCoreDevice#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/iot_core_device#id IotCoreDevice#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/iot_core_device#passwords IotCoreDevice#passwords}.
	Passwords *[]*string `field:"optional" json:"passwords" yaml:"passwords"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/iot_core_device#timeouts IotCoreDevice#timeouts}
	Timeouts *IotCoreDeviceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

