// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbBackendGroupGrpcBackendTls struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#sni AlbBackendGroup#sni}.
	Sni *string `field:"optional" json:"sni" yaml:"sni"`
	// validation_context block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#validation_context AlbBackendGroup#validation_context}
	ValidationContext *AlbBackendGroupGrpcBackendTlsValidationContext `field:"optional" json:"validationContext" yaml:"validationContext"`
}

