// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbLoadBalancerListenerTls struct {
	// default_handler block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_load_balancer#default_handler AlbLoadBalancer#default_handler}
	DefaultHandler *AlbLoadBalancerListenerTlsDefaultHandler `field:"required" json:"defaultHandler" yaml:"defaultHandler"`
	// sni_handler block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_load_balancer#sni_handler AlbLoadBalancer#sni_handler}
	SniHandler interface{} `field:"optional" json:"sniHandler" yaml:"sniHandler"`
}

