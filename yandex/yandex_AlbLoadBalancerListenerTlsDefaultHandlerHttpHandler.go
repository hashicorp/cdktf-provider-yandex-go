// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbLoadBalancerListenerTlsDefaultHandlerHttpHandler struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_load_balancer#allow_http10 AlbLoadBalancer#allow_http10}.
	AllowHttp10 interface{} `field:"optional" json:"allowHttp10" yaml:"allowHttp10"`
	// http2_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_load_balancer#http2_options AlbLoadBalancer#http2_options}
	Http2Options *AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerHttp2Options `field:"optional" json:"http2Options" yaml:"http2Options"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_load_balancer#http_router_id AlbLoadBalancer#http_router_id}.
	HttpRouterId *string `field:"optional" json:"httpRouterId" yaml:"httpRouterId"`
}

