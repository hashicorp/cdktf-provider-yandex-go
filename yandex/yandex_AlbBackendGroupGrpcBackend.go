// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbBackendGroupGrpcBackend struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#name AlbBackendGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#target_group_ids AlbBackendGroup#target_group_ids}.
	TargetGroupIds *[]*string `field:"required" json:"targetGroupIds" yaml:"targetGroupIds"`
	// healthcheck block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#healthcheck AlbBackendGroup#healthcheck}
	Healthcheck *AlbBackendGroupGrpcBackendHealthcheck `field:"optional" json:"healthcheck" yaml:"healthcheck"`
	// load_balancing_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#load_balancing_config AlbBackendGroup#load_balancing_config}
	LoadBalancingConfig *AlbBackendGroupGrpcBackendLoadBalancingConfig `field:"optional" json:"loadBalancingConfig" yaml:"loadBalancingConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#port AlbBackendGroup#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#tls AlbBackendGroup#tls}
	Tls *AlbBackendGroupGrpcBackendTls `field:"optional" json:"tls" yaml:"tls"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#weight AlbBackendGroup#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

