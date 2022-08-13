// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexAlbBackendGroupHttpBackend struct {
	// healthcheck block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#healthcheck DataYandexAlbBackendGroup#healthcheck}
	Healthcheck *DataYandexAlbBackendGroupHttpBackendHealthcheck `field:"optional" json:"healthcheck" yaml:"healthcheck"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#http2 DataYandexAlbBackendGroup#http2}.
	Http2 interface{} `field:"optional" json:"http2" yaml:"http2"`
	// load_balancing_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#load_balancing_config DataYandexAlbBackendGroup#load_balancing_config}
	LoadBalancingConfig *DataYandexAlbBackendGroupHttpBackendLoadBalancingConfig `field:"optional" json:"loadBalancingConfig" yaml:"loadBalancingConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#port DataYandexAlbBackendGroup#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#storage_bucket DataYandexAlbBackendGroup#storage_bucket}.
	StorageBucket *string `field:"optional" json:"storageBucket" yaml:"storageBucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#target_group_ids DataYandexAlbBackendGroup#target_group_ids}.
	TargetGroupIds *[]*string `field:"optional" json:"targetGroupIds" yaml:"targetGroupIds"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#tls DataYandexAlbBackendGroup#tls}
	Tls *DataYandexAlbBackendGroupHttpBackendTls `field:"optional" json:"tls" yaml:"tls"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#weight DataYandexAlbBackendGroup#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

