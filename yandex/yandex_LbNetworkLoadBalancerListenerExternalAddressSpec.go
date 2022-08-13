// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type LbNetworkLoadBalancerListenerExternalAddressSpec struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/lb_network_load_balancer#address LbNetworkLoadBalancer#address}.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/lb_network_load_balancer#ip_version LbNetworkLoadBalancer#ip_version}.
	IpVersion *string `field:"optional" json:"ipVersion" yaml:"ipVersion"`
}

