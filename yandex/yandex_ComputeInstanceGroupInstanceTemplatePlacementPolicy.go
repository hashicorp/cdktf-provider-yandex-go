// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ComputeInstanceGroupInstanceTemplatePlacementPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#placement_group_id ComputeInstanceGroup#placement_group_id}.
	PlacementGroupId *string `field:"required" json:"placementGroupId" yaml:"placementGroupId"`
}

