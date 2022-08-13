// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexComputeInstancePlacementPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/compute_instance#host_affinity_rules DataYandexComputeInstance#host_affinity_rules}.
	HostAffinityRules interface{} `field:"optional" json:"hostAffinityRules" yaml:"hostAffinityRules"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/compute_instance#placement_group_id DataYandexComputeInstance#placement_group_id}.
	PlacementGroupId *string `field:"optional" json:"placementGroupId" yaml:"placementGroupId"`
}

