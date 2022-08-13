// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type StorageBucketHttps struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#certificate_id StorageBucket#certificate_id}.
	CertificateId *string `field:"required" json:"certificateId" yaml:"certificateId"`
}

