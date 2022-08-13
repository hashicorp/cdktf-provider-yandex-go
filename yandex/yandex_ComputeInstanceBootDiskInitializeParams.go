// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ComputeInstanceBootDiskInitializeParams struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#block_size ComputeInstance#block_size}.
	BlockSize *float64 `field:"optional" json:"blockSize" yaml:"blockSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#description ComputeInstance#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#image_id ComputeInstance#image_id}.
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#name ComputeInstance#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#size ComputeInstance#size}.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#snapshot_id ComputeInstance#snapshot_id}.
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#type ComputeInstance#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

