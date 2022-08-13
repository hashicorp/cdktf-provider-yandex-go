// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbLoadBalancerListener struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_load_balancer#name AlbLoadBalancer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// endpoint block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_load_balancer#endpoint AlbLoadBalancer#endpoint}
	Endpoint interface{} `field:"optional" json:"endpoint" yaml:"endpoint"`
	// http block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_load_balancer#http AlbLoadBalancer#http}
	Http *AlbLoadBalancerListenerHttp `field:"optional" json:"http" yaml:"http"`
	// stream block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_load_balancer#stream AlbLoadBalancer#stream}
	Stream *AlbLoadBalancerListenerStream `field:"optional" json:"stream" yaml:"stream"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_load_balancer#tls AlbLoadBalancer#tls}
	Tls *AlbLoadBalancerListenerTls `field:"optional" json:"tls" yaml:"tls"`
}

