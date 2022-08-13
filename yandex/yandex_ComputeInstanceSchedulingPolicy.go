// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ComputeInstanceSchedulingPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#preemptible ComputeInstance#preemptible}.
	Preemptible interface{} `field:"optional" json:"preemptible" yaml:"preemptible"`
}

