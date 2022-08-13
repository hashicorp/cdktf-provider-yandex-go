// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbLoadBalancerListenerTlsSniHandler struct {
	// handler block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_load_balancer#handler AlbLoadBalancer#handler}
	Handler *AlbLoadBalancerListenerTlsSniHandlerHandler `field:"required" json:"handler" yaml:"handler"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_load_balancer#name AlbLoadBalancer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_load_balancer#server_names AlbLoadBalancer#server_names}.
	ServerNames *[]*string `field:"required" json:"serverNames" yaml:"serverNames"`
}

