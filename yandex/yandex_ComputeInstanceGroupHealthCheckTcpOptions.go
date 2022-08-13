// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ComputeInstanceGroupHealthCheckTcpOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#port ComputeInstanceGroup#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

