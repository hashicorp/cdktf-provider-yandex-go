// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type VpcSubnetTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_subnet#create VpcSubnet#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_subnet#delete VpcSubnet#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_subnet#update VpcSubnet#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

