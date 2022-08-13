// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ComputePlacementGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_placement_group#create ComputePlacementGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_placement_group#delete ComputePlacementGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_placement_group#update ComputePlacementGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

