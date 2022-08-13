// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpcSecurityGroupRuleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_security_group_rule#direction VpcSecurityGroupRule#direction}.
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_security_group_rule#security_group_binding VpcSecurityGroupRule#security_group_binding}.
	SecurityGroupBinding *string `field:"required" json:"securityGroupBinding" yaml:"securityGroupBinding"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_security_group_rule#description VpcSecurityGroupRule#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_security_group_rule#from_port VpcSecurityGroupRule#from_port}.
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_security_group_rule#id VpcSecurityGroupRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_security_group_rule#labels VpcSecurityGroupRule#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_security_group_rule#port VpcSecurityGroupRule#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_security_group_rule#predefined_target VpcSecurityGroupRule#predefined_target}.
	PredefinedTarget *string `field:"optional" json:"predefinedTarget" yaml:"predefinedTarget"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_security_group_rule#protocol VpcSecurityGroupRule#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_security_group_rule#security_group_id VpcSecurityGroupRule#security_group_id}.
	SecurityGroupId *string `field:"optional" json:"securityGroupId" yaml:"securityGroupId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_security_group_rule#timeouts VpcSecurityGroupRule#timeouts}
	Timeouts *VpcSecurityGroupRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_security_group_rule#to_port VpcSecurityGroupRule#to_port}.
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_security_group_rule#v4_cidr_blocks VpcSecurityGroupRule#v4_cidr_blocks}.
	V4CidrBlocks *[]*string `field:"optional" json:"v4CidrBlocks" yaml:"v4CidrBlocks"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_security_group_rule#v6_cidr_blocks VpcSecurityGroupRule#v6_cidr_blocks}.
	V6CidrBlocks *[]*string `field:"optional" json:"v6CidrBlocks" yaml:"v6CidrBlocks"`
}

