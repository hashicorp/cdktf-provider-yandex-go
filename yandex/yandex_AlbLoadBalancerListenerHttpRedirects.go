// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbLoadBalancerListenerHttpRedirects struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_load_balancer#http_to_https AlbLoadBalancer#http_to_https}.
	HttpToHttps interface{} `field:"optional" json:"httpToHttps" yaml:"httpToHttps"`
}

