// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexAlbBackendGroupStreamBackendHealthcheckStreamHealthcheck struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#receive DataYandexAlbBackendGroup#receive}.
	Receive *string `field:"optional" json:"receive" yaml:"receive"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#send DataYandexAlbBackendGroup#send}.
	Send *string `field:"optional" json:"send" yaml:"send"`
}

