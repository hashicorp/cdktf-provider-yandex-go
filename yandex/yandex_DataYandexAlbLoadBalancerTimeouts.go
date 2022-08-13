// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexAlbLoadBalancerTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_load_balancer#create DataYandexAlbLoadBalancer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_load_balancer#delete DataYandexAlbLoadBalancer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_load_balancer#update DataYandexAlbLoadBalancer#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

