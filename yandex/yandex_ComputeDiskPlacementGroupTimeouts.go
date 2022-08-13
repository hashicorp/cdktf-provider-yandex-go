// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ComputeDiskPlacementGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_disk_placement_group#create ComputeDiskPlacementGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_disk_placement_group#delete ComputeDiskPlacementGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_disk_placement_group#update ComputeDiskPlacementGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

