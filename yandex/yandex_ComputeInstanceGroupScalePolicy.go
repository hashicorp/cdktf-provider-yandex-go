// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ComputeInstanceGroupScalePolicy struct {
	// auto_scale block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#auto_scale ComputeInstanceGroup#auto_scale}
	AutoScale *ComputeInstanceGroupScalePolicyAutoScale `field:"optional" json:"autoScale" yaml:"autoScale"`
	// fixed_scale block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#fixed_scale ComputeInstanceGroup#fixed_scale}
	FixedScale *ComputeInstanceGroupScalePolicyFixedScale `field:"optional" json:"fixedScale" yaml:"fixedScale"`
	// test_auto_scale block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#test_auto_scale ComputeInstanceGroup#test_auto_scale}
	TestAutoScale *ComputeInstanceGroupScalePolicyTestAutoScale `field:"optional" json:"testAutoScale" yaml:"testAutoScale"`
}

