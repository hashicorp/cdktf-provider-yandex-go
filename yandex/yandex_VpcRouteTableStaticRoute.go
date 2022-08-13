// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type VpcRouteTableStaticRoute struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_route_table#destination_prefix VpcRouteTable#destination_prefix}.
	DestinationPrefix *string `field:"optional" json:"destinationPrefix" yaml:"destinationPrefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_route_table#next_hop_address VpcRouteTable#next_hop_address}.
	NextHopAddress *string `field:"optional" json:"nextHopAddress" yaml:"nextHopAddress"`
}

