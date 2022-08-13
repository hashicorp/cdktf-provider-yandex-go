// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbBackendGroupHttpBackend struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#name AlbBackendGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// healthcheck block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#healthcheck AlbBackendGroup#healthcheck}
	Healthcheck *AlbBackendGroupHttpBackendHealthcheck `field:"optional" json:"healthcheck" yaml:"healthcheck"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#http2 AlbBackendGroup#http2}.
	Http2 interface{} `field:"optional" json:"http2" yaml:"http2"`
	// load_balancing_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#load_balancing_config AlbBackendGroup#load_balancing_config}
	LoadBalancingConfig *AlbBackendGroupHttpBackendLoadBalancingConfig `field:"optional" json:"loadBalancingConfig" yaml:"loadBalancingConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#port AlbBackendGroup#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#storage_bucket AlbBackendGroup#storage_bucket}.
	StorageBucket *string `field:"optional" json:"storageBucket" yaml:"storageBucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#target_group_ids AlbBackendGroup#target_group_ids}.
	TargetGroupIds *[]*string `field:"optional" json:"targetGroupIds" yaml:"targetGroupIds"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#tls AlbBackendGroup#tls}
	Tls *AlbBackendGroupHttpBackendTls `field:"optional" json:"tls" yaml:"tls"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#weight AlbBackendGroup#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

