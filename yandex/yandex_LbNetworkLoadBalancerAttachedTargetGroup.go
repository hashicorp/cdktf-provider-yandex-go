// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type LbNetworkLoadBalancerAttachedTargetGroup struct {
	// healthcheck block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/lb_network_load_balancer#healthcheck LbNetworkLoadBalancer#healthcheck}
	Healthcheck interface{} `field:"required" json:"healthcheck" yaml:"healthcheck"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/lb_network_load_balancer#target_group_id LbNetworkLoadBalancer#target_group_id}.
	TargetGroupId *string `field:"required" json:"targetGroupId" yaml:"targetGroupId"`
}

