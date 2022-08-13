// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type LbNetworkLoadBalancerAttachedTargetGroupHealthcheckTcpOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/lb_network_load_balancer#port LbNetworkLoadBalancer#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

