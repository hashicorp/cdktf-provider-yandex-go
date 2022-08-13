// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexCdnResourceSslCertificate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/cdn_resource#type DataYandexCdnResource#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/cdn_resource#certificate_manager_id DataYandexCdnResource#certificate_manager_id}.
	CertificateManagerId *string `field:"optional" json:"certificateManagerId" yaml:"certificateManagerId"`
}

