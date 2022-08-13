// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type VpcSecurityGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_security_group#create VpcSecurityGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_security_group#delete VpcSecurityGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_security_group#update VpcSecurityGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

