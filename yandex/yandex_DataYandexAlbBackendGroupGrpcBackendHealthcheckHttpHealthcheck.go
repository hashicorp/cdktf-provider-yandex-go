// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexAlbBackendGroupGrpcBackendHealthcheckHttpHealthcheck struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#host DataYandexAlbBackendGroup#host}.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#http2 DataYandexAlbBackendGroup#http2}.
	Http2 interface{} `field:"optional" json:"http2" yaml:"http2"`
}

