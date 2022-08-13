// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ComputeInstanceLocalDisk struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#size_bytes ComputeInstance#size_bytes}.
	SizeBytes *float64 `field:"required" json:"sizeBytes" yaml:"sizeBytes"`
}

