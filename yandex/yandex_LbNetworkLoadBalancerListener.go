// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type LbNetworkLoadBalancerListener struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/lb_network_load_balancer#name LbNetworkLoadBalancer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/lb_network_load_balancer#port LbNetworkLoadBalancer#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// external_address_spec block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/lb_network_load_balancer#external_address_spec LbNetworkLoadBalancer#external_address_spec}
	ExternalAddressSpec *LbNetworkLoadBalancerListenerExternalAddressSpec `field:"optional" json:"externalAddressSpec" yaml:"externalAddressSpec"`
	// internal_address_spec block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/lb_network_load_balancer#internal_address_spec LbNetworkLoadBalancer#internal_address_spec}
	InternalAddressSpec *LbNetworkLoadBalancerListenerInternalAddressSpec `field:"optional" json:"internalAddressSpec" yaml:"internalAddressSpec"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/lb_network_load_balancer#protocol LbNetworkLoadBalancer#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/lb_network_load_balancer#target_port LbNetworkLoadBalancer#target_port}.
	TargetPort *float64 `field:"optional" json:"targetPort" yaml:"targetPort"`
}

