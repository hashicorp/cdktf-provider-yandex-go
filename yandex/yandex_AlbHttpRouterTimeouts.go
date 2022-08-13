// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbHttpRouterTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_http_router#create AlbHttpRouter#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_http_router#delete AlbHttpRouter#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_http_router#update AlbHttpRouter#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

