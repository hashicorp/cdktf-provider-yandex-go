// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexComputeInstancePlacementPolicyHostAffinityRules struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/compute_instance#key DataYandexComputeInstance#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/compute_instance#op DataYandexComputeInstance#op}.
	Op *string `field:"optional" json:"op" yaml:"op"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/compute_instance#values DataYandexComputeInstance#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

