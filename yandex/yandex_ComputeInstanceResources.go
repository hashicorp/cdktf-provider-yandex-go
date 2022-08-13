// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ComputeInstanceResources struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#cores ComputeInstance#cores}.
	Cores *float64 `field:"required" json:"cores" yaml:"cores"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#memory ComputeInstance#memory}.
	Memory *float64 `field:"required" json:"memory" yaml:"memory"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#core_fraction ComputeInstance#core_fraction}.
	CoreFraction *float64 `field:"optional" json:"coreFraction" yaml:"coreFraction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#gpus ComputeInstance#gpus}.
	Gpus *float64 `field:"optional" json:"gpus" yaml:"gpus"`
}

