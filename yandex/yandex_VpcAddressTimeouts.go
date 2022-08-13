// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type VpcAddressTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_address#create VpcAddress#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_address#delete VpcAddress#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_address#update VpcAddress#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

