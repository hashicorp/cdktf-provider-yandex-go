// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexOrganizationmanagerSamlFederationUserAccountConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/organizationmanager_saml_federation_user_account#federation_id DataYandexOrganizationmanagerSamlFederationUserAccount#federation_id}.
	FederationId *string `field:"required" json:"federationId" yaml:"federationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/organizationmanager_saml_federation_user_account#name_id DataYandexOrganizationmanagerSamlFederationUserAccount#name_id}.
	NameId *string `field:"required" json:"nameId" yaml:"nameId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/organizationmanager_saml_federation_user_account#id DataYandexOrganizationmanagerSamlFederationUserAccount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

