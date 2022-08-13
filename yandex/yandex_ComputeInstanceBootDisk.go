// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ComputeInstanceBootDisk struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#auto_delete ComputeInstance#auto_delete}.
	AutoDelete interface{} `field:"optional" json:"autoDelete" yaml:"autoDelete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#device_name ComputeInstance#device_name}.
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#disk_id ComputeInstance#disk_id}.
	DiskId *string `field:"optional" json:"diskId" yaml:"diskId"`
	// initialize_params block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#initialize_params ComputeInstance#initialize_params}
	InitializeParams *ComputeInstanceBootDiskInitializeParams `field:"optional" json:"initializeParams" yaml:"initializeParams"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#mode ComputeInstance#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

