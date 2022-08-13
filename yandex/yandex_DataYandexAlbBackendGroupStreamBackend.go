// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexAlbBackendGroupStreamBackend struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#enable_proxy_protocol DataYandexAlbBackendGroup#enable_proxy_protocol}.
	EnableProxyProtocol interface{} `field:"optional" json:"enableProxyProtocol" yaml:"enableProxyProtocol"`
	// healthcheck block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#healthcheck DataYandexAlbBackendGroup#healthcheck}
	Healthcheck *DataYandexAlbBackendGroupStreamBackendHealthcheck `field:"optional" json:"healthcheck" yaml:"healthcheck"`
	// load_balancing_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#load_balancing_config DataYandexAlbBackendGroup#load_balancing_config}
	LoadBalancingConfig *DataYandexAlbBackendGroupStreamBackendLoadBalancingConfig `field:"optional" json:"loadBalancingConfig" yaml:"loadBalancingConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#port DataYandexAlbBackendGroup#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#tls DataYandexAlbBackendGroup#tls}
	Tls *DataYandexAlbBackendGroupStreamBackendTls `field:"optional" json:"tls" yaml:"tls"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#weight DataYandexAlbBackendGroup#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

