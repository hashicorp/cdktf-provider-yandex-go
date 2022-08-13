// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbVirtualHostRouteGrpcRoute struct {
	// grpc_match block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#grpc_match AlbVirtualHost#grpc_match}
	GrpcMatch interface{} `field:"optional" json:"grpcMatch" yaml:"grpcMatch"`
	// grpc_route_action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#grpc_route_action AlbVirtualHost#grpc_route_action}
	GrpcRouteAction *AlbVirtualHostRouteGrpcRouteGrpcRouteAction `field:"optional" json:"grpcRouteAction" yaml:"grpcRouteAction"`
	// grpc_status_response_action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#grpc_status_response_action AlbVirtualHost#grpc_status_response_action}
	GrpcStatusResponseAction *AlbVirtualHostRouteGrpcRouteGrpcStatusResponseAction `field:"optional" json:"grpcStatusResponseAction" yaml:"grpcStatusResponseAction"`
}

