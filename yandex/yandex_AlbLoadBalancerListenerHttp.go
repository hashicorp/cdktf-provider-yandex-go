// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbLoadBalancerListenerHttp struct {
	// handler block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_load_balancer#handler AlbLoadBalancer#handler}
	Handler *AlbLoadBalancerListenerHttpHandler `field:"optional" json:"handler" yaml:"handler"`
	// redirects block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_load_balancer#redirects AlbLoadBalancer#redirects}
	Redirects *AlbLoadBalancerListenerHttpRedirects `field:"optional" json:"redirects" yaml:"redirects"`
}

