// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type CdnResourceSslCertificate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#type CdnResource#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#certificate_manager_id CdnResource#certificate_manager_id}.
	CertificateManagerId *string `field:"optional" json:"certificateManagerId" yaml:"certificateManagerId"`
}

