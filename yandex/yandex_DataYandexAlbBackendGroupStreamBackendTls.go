// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexAlbBackendGroupStreamBackendTls struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#sni DataYandexAlbBackendGroup#sni}.
	Sni *string `field:"optional" json:"sni" yaml:"sni"`
	// validation_context block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#validation_context DataYandexAlbBackendGroup#validation_context}
	ValidationContext *DataYandexAlbBackendGroupStreamBackendTlsValidationContext `field:"optional" json:"validationContext" yaml:"validationContext"`
}

