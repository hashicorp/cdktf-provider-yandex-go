// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type LbNetworkLoadBalancerAttachedTargetGroupHealthcheck struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/lb_network_load_balancer#name LbNetworkLoadBalancer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/lb_network_load_balancer#healthy_threshold LbNetworkLoadBalancer#healthy_threshold}.
	HealthyThreshold *float64 `field:"optional" json:"healthyThreshold" yaml:"healthyThreshold"`
	// http_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/lb_network_load_balancer#http_options LbNetworkLoadBalancer#http_options}
	HttpOptions *LbNetworkLoadBalancerAttachedTargetGroupHealthcheckHttpOptions `field:"optional" json:"httpOptions" yaml:"httpOptions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/lb_network_load_balancer#interval LbNetworkLoadBalancer#interval}.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// tcp_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/lb_network_load_balancer#tcp_options LbNetworkLoadBalancer#tcp_options}
	TcpOptions *LbNetworkLoadBalancerAttachedTargetGroupHealthcheckTcpOptions `field:"optional" json:"tcpOptions" yaml:"tcpOptions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/lb_network_load_balancer#timeout LbNetworkLoadBalancer#timeout}.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/lb_network_load_balancer#unhealthy_threshold LbNetworkLoadBalancer#unhealthy_threshold}.
	UnhealthyThreshold *float64 `field:"optional" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

