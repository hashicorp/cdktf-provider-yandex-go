// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type VpcRouteTableTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_route_table#create VpcRouteTable#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_route_table#delete VpcRouteTable#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_route_table#update VpcRouteTable#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

