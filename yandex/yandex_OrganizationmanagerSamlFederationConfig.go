// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrganizationmanagerSamlFederationConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/organizationmanager_saml_federation#issuer OrganizationmanagerSamlFederation#issuer}.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/organizationmanager_saml_federation#name OrganizationmanagerSamlFederation#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/organizationmanager_saml_federation#organization_id OrganizationmanagerSamlFederation#organization_id}.
	OrganizationId *string `field:"required" json:"organizationId" yaml:"organizationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/organizationmanager_saml_federation#sso_binding OrganizationmanagerSamlFederation#sso_binding}.
	SsoBinding *string `field:"required" json:"ssoBinding" yaml:"ssoBinding"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/organizationmanager_saml_federation#sso_url OrganizationmanagerSamlFederation#sso_url}.
	SsoUrl *string `field:"required" json:"ssoUrl" yaml:"ssoUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/organizationmanager_saml_federation#auto_create_account_on_login OrganizationmanagerSamlFederation#auto_create_account_on_login}.
	AutoCreateAccountOnLogin interface{} `field:"optional" json:"autoCreateAccountOnLogin" yaml:"autoCreateAccountOnLogin"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/organizationmanager_saml_federation#case_insensitive_name_ids OrganizationmanagerSamlFederation#case_insensitive_name_ids}.
	CaseInsensitiveNameIds interface{} `field:"optional" json:"caseInsensitiveNameIds" yaml:"caseInsensitiveNameIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/organizationmanager_saml_federation#cookie_max_age OrganizationmanagerSamlFederation#cookie_max_age}.
	CookieMaxAge *string `field:"optional" json:"cookieMaxAge" yaml:"cookieMaxAge"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/organizationmanager_saml_federation#description OrganizationmanagerSamlFederation#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/organizationmanager_saml_federation#id OrganizationmanagerSamlFederation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/organizationmanager_saml_federation#labels OrganizationmanagerSamlFederation#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// security_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/organizationmanager_saml_federation#security_settings OrganizationmanagerSamlFederation#security_settings}
	SecuritySettings *OrganizationmanagerSamlFederationSecuritySettings `field:"optional" json:"securitySettings" yaml:"securitySettings"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/organizationmanager_saml_federation#timeouts OrganizationmanagerSamlFederation#timeouts}
	Timeouts *OrganizationmanagerSamlFederationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

