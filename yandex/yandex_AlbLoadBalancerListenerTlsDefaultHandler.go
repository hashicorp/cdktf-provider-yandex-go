// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbLoadBalancerListenerTlsDefaultHandler struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_load_balancer#certificate_ids AlbLoadBalancer#certificate_ids}.
	CertificateIds *[]*string `field:"required" json:"certificateIds" yaml:"certificateIds"`
	// http_handler block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_load_balancer#http_handler AlbLoadBalancer#http_handler}
	HttpHandler *AlbLoadBalancerListenerTlsDefaultHandlerHttpHandler `field:"optional" json:"httpHandler" yaml:"httpHandler"`
	// stream_handler block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_load_balancer#stream_handler AlbLoadBalancer#stream_handler}
	StreamHandler *AlbLoadBalancerListenerTlsDefaultHandlerStreamHandler `field:"optional" json:"streamHandler" yaml:"streamHandler"`
}

