// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexComputeDiskDiskPlacementPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/compute_disk#disk_placement_group_id DataYandexComputeDisk#disk_placement_group_id}.
	DiskPlacementGroupId *string `field:"required" json:"diskPlacementGroupId" yaml:"diskPlacementGroupId"`
}

