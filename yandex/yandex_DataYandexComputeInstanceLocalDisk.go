// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexComputeInstanceLocalDisk struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/compute_instance#size_bytes DataYandexComputeInstance#size_bytes}.
	SizeBytes *float64 `field:"required" json:"sizeBytes" yaml:"sizeBytes"`
}

