// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbBackendGroupGrpcBackendHealthcheckStreamHealthcheck struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#receive AlbBackendGroup#receive}.
	Receive *string `field:"optional" json:"receive" yaml:"receive"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#send AlbBackendGroup#send}.
	Send *string `field:"optional" json:"send" yaml:"send"`
}

