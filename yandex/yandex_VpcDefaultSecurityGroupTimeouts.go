// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type VpcDefaultSecurityGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_default_security_group#create VpcDefaultSecurityGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_default_security_group#delete VpcDefaultSecurityGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_default_security_group#read VpcDefaultSecurityGroup#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_default_security_group#update VpcDefaultSecurityGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

