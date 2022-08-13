// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ComputeInstanceGroupAllocationPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#zones ComputeInstanceGroup#zones}.
	Zones *[]*string `field:"required" json:"zones" yaml:"zones"`
}

