// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type LbNetworkLoadBalancerTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/lb_network_load_balancer#create LbNetworkLoadBalancer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/lb_network_load_balancer#delete LbNetworkLoadBalancer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/lb_network_load_balancer#update LbNetworkLoadBalancer#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

