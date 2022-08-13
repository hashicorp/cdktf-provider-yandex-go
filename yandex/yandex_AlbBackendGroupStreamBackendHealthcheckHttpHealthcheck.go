// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbBackendGroupStreamBackendHealthcheckHttpHealthcheck struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#path AlbBackendGroup#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#host AlbBackendGroup#host}.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#http2 AlbBackendGroup#http2}.
	Http2 interface{} `field:"optional" json:"http2" yaml:"http2"`
}

