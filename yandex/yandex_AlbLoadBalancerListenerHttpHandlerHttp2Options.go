// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbLoadBalancerListenerHttpHandlerHttp2Options struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_load_balancer#max_concurrent_streams AlbLoadBalancer#max_concurrent_streams}.
	MaxConcurrentStreams *float64 `field:"optional" json:"maxConcurrentStreams" yaml:"maxConcurrentStreams"`
}

