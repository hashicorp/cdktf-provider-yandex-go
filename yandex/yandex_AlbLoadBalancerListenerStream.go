// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbLoadBalancerListenerStream struct {
	// handler block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_load_balancer#handler AlbLoadBalancer#handler}
	Handler *AlbLoadBalancerListenerStreamHandler `field:"optional" json:"handler" yaml:"handler"`
}

