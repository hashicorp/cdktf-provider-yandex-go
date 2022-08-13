// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ComputeInstanceGroupInstanceTemplateSecondaryDisk struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#device_name ComputeInstanceGroup#device_name}.
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#disk_id ComputeInstanceGroup#disk_id}.
	DiskId *string `field:"optional" json:"diskId" yaml:"diskId"`
	// initialize_params block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#initialize_params ComputeInstanceGroup#initialize_params}
	InitializeParams *ComputeInstanceGroupInstanceTemplateSecondaryDiskInitializeParams `field:"optional" json:"initializeParams" yaml:"initializeParams"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#mode ComputeInstanceGroup#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

