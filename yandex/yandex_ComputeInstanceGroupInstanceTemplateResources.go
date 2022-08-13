// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ComputeInstanceGroupInstanceTemplateResources struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#cores ComputeInstanceGroup#cores}.
	Cores *float64 `field:"required" json:"cores" yaml:"cores"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#memory ComputeInstanceGroup#memory}.
	Memory *float64 `field:"required" json:"memory" yaml:"memory"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#core_fraction ComputeInstanceGroup#core_fraction}.
	CoreFraction *float64 `field:"optional" json:"coreFraction" yaml:"coreFraction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#gpus ComputeInstanceGroup#gpus}.
	Gpus *float64 `field:"optional" json:"gpus" yaml:"gpus"`
}

