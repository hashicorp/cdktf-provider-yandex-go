// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbVirtualHostRoute struct {
	// grpc_route block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#grpc_route AlbVirtualHost#grpc_route}
	GrpcRoute *AlbVirtualHostRouteGrpcRoute `field:"optional" json:"grpcRoute" yaml:"grpcRoute"`
	// http_route block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#http_route AlbVirtualHost#http_route}
	HttpRoute *AlbVirtualHostRouteHttpRoute `field:"optional" json:"httpRoute" yaml:"httpRoute"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#name AlbVirtualHost#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

