// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbVirtualHostRouteHttpRoute struct {
	// direct_response_action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#direct_response_action AlbVirtualHost#direct_response_action}
	DirectResponseAction *AlbVirtualHostRouteHttpRouteDirectResponseAction `field:"optional" json:"directResponseAction" yaml:"directResponseAction"`
	// http_match block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#http_match AlbVirtualHost#http_match}
	HttpMatch interface{} `field:"optional" json:"httpMatch" yaml:"httpMatch"`
	// http_route_action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#http_route_action AlbVirtualHost#http_route_action}
	HttpRouteAction *AlbVirtualHostRouteHttpRouteHttpRouteAction `field:"optional" json:"httpRouteAction" yaml:"httpRouteAction"`
	// redirect_action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#redirect_action AlbVirtualHost#redirect_action}
	RedirectAction *AlbVirtualHostRouteHttpRouteRedirectAction `field:"optional" json:"redirectAction" yaml:"redirectAction"`
}

