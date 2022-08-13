// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexAlbBackendGroupGrpcBackendTlsValidationContext struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#trusted_ca_bytes DataYandexAlbBackendGroup#trusted_ca_bytes}.
	TrustedCaBytes *string `field:"optional" json:"trustedCaBytes" yaml:"trustedCaBytes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#trusted_ca_id DataYandexAlbBackendGroup#trusted_ca_id}.
	TrustedCaId *string `field:"optional" json:"trustedCaId" yaml:"trustedCaId"`
}

