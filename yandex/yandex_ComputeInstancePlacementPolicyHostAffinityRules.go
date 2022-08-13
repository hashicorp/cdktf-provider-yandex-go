// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ComputeInstancePlacementPolicyHostAffinityRules struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#key ComputeInstance#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#op ComputeInstance#op}.
	Op *string `field:"optional" json:"op" yaml:"op"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#values ComputeInstance#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

