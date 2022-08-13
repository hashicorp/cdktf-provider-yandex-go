// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbLoadBalancerListenerEndpointAddress struct {
	// external_ipv4_address block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_load_balancer#external_ipv4_address AlbLoadBalancer#external_ipv4_address}
	ExternalIpv4Address *AlbLoadBalancerListenerEndpointAddressExternalIpv4Address `field:"optional" json:"externalIpv4Address" yaml:"externalIpv4Address"`
	// external_ipv6_address block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_load_balancer#external_ipv6_address AlbLoadBalancer#external_ipv6_address}
	ExternalIpv6Address *AlbLoadBalancerListenerEndpointAddressExternalIpv6Address `field:"optional" json:"externalIpv6Address" yaml:"externalIpv6Address"`
	// internal_ipv4_address block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_load_balancer#internal_ipv4_address AlbLoadBalancer#internal_ipv4_address}
	InternalIpv4Address *AlbLoadBalancerListenerEndpointAddressInternalIpv4Address `field:"optional" json:"internalIpv4Address" yaml:"internalIpv4Address"`
}

