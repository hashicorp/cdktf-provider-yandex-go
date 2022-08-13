// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbVirtualHostRouteHttpRouteDirectResponseAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#body AlbVirtualHost#body}.
	Body *string `field:"optional" json:"body" yaml:"body"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#status AlbVirtualHost#status}.
	Status *float64 `field:"optional" json:"status" yaml:"status"`
}

