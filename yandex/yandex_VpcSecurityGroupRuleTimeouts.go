// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type VpcSecurityGroupRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_security_group_rule#create VpcSecurityGroupRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_security_group_rule#delete VpcSecurityGroupRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_security_group_rule#read VpcSecurityGroupRule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_security_group_rule#update VpcSecurityGroupRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

