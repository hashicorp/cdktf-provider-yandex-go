// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexAlbBackendGroupStreamBackendLoadBalancingConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#locality_aware_routing_percent DataYandexAlbBackendGroup#locality_aware_routing_percent}.
	LocalityAwareRoutingPercent *float64 `field:"optional" json:"localityAwareRoutingPercent" yaml:"localityAwareRoutingPercent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#mode DataYandexAlbBackendGroup#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#panic_threshold DataYandexAlbBackendGroup#panic_threshold}.
	PanicThreshold *float64 `field:"optional" json:"panicThreshold" yaml:"panicThreshold"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#strict_locality DataYandexAlbBackendGroup#strict_locality}.
	StrictLocality interface{} `field:"optional" json:"strictLocality" yaml:"strictLocality"`
}

