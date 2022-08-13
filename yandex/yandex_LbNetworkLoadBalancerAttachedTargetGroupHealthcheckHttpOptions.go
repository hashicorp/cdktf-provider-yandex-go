// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type LbNetworkLoadBalancerAttachedTargetGroupHealthcheckHttpOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/lb_network_load_balancer#port LbNetworkLoadBalancer#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/lb_network_load_balancer#path LbNetworkLoadBalancer#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

