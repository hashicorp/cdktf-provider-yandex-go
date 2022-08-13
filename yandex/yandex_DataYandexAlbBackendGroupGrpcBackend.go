// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexAlbBackendGroupGrpcBackend struct {
	// healthcheck block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#healthcheck DataYandexAlbBackendGroup#healthcheck}
	Healthcheck *DataYandexAlbBackendGroupGrpcBackendHealthcheck `field:"optional" json:"healthcheck" yaml:"healthcheck"`
	// load_balancing_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#load_balancing_config DataYandexAlbBackendGroup#load_balancing_config}
	LoadBalancingConfig *DataYandexAlbBackendGroupGrpcBackendLoadBalancingConfig `field:"optional" json:"loadBalancingConfig" yaml:"loadBalancingConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#port DataYandexAlbBackendGroup#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#tls DataYandexAlbBackendGroup#tls}
	Tls *DataYandexAlbBackendGroupGrpcBackendTls `field:"optional" json:"tls" yaml:"tls"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#weight DataYandexAlbBackendGroup#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

