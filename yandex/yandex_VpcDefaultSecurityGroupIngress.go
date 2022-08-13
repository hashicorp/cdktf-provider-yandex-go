// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type VpcDefaultSecurityGroupIngress struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_default_security_group#protocol VpcDefaultSecurityGroup#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_default_security_group#description VpcDefaultSecurityGroup#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_default_security_group#from_port VpcDefaultSecurityGroup#from_port}.
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_default_security_group#labels VpcDefaultSecurityGroup#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_default_security_group#port VpcDefaultSecurityGroup#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_default_security_group#predefined_target VpcDefaultSecurityGroup#predefined_target}.
	PredefinedTarget *string `field:"optional" json:"predefinedTarget" yaml:"predefinedTarget"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_default_security_group#security_group_id VpcDefaultSecurityGroup#security_group_id}.
	SecurityGroupId *string `field:"optional" json:"securityGroupId" yaml:"securityGroupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_default_security_group#to_port VpcDefaultSecurityGroup#to_port}.
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_default_security_group#v4_cidr_blocks VpcDefaultSecurityGroup#v4_cidr_blocks}.
	V4CidrBlocks *[]*string `field:"optional" json:"v4CidrBlocks" yaml:"v4CidrBlocks"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_default_security_group#v6_cidr_blocks VpcDefaultSecurityGroup#v6_cidr_blocks}.
	V6CidrBlocks *[]*string `field:"optional" json:"v6CidrBlocks" yaml:"v6CidrBlocks"`
}

