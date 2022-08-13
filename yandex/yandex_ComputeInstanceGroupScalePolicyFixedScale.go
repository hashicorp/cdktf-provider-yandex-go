// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ComputeInstanceGroupScalePolicyFixedScale struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#size ComputeInstanceGroup#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
}

