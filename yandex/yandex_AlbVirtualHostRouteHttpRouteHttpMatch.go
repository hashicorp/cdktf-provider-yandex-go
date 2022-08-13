// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbVirtualHostRouteHttpRouteHttpMatch struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#http_method AlbVirtualHost#http_method}.
	HttpMethod *[]*string `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// path block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#path AlbVirtualHost#path}
	Path *AlbVirtualHostRouteHttpRouteHttpMatchPath `field:"optional" json:"path" yaml:"path"`
}

