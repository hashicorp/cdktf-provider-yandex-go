// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type VpcNetworkTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_network#create VpcNetwork#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_network#delete VpcNetwork#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_network#update VpcNetwork#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

