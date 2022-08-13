// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrganizationmanagerOrganizationIamBindingConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/organizationmanager_organization_iam_binding#members OrganizationmanagerOrganizationIamBinding#members}.
	Members *[]*string `field:"required" json:"members" yaml:"members"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/organizationmanager_organization_iam_binding#organization_id OrganizationmanagerOrganizationIamBinding#organization_id}.
	OrganizationId *string `field:"required" json:"organizationId" yaml:"organizationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/organizationmanager_organization_iam_binding#role OrganizationmanagerOrganizationIamBinding#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/organizationmanager_organization_iam_binding#id OrganizationmanagerOrganizationIamBinding#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/organizationmanager_organization_iam_binding#sleep_after OrganizationmanagerOrganizationIamBinding#sleep_after}.
	SleepAfter *float64 `field:"optional" json:"sleepAfter" yaml:"sleepAfter"`
}

