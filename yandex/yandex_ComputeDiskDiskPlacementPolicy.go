// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ComputeDiskDiskPlacementPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_disk#disk_placement_group_id ComputeDisk#disk_placement_group_id}.
	DiskPlacementGroupId *string `field:"required" json:"diskPlacementGroupId" yaml:"diskPlacementGroupId"`
}

