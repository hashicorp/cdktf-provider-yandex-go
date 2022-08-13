// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IamServiceAccountKeyConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/iam_service_account_key#service_account_id IamServiceAccountKey#service_account_id}.
	ServiceAccountId *string `field:"required" json:"serviceAccountId" yaml:"serviceAccountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/iam_service_account_key#description IamServiceAccountKey#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/iam_service_account_key#format IamServiceAccountKey#format}.
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/iam_service_account_key#id IamServiceAccountKey#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/iam_service_account_key#key_algorithm IamServiceAccountKey#key_algorithm}.
	KeyAlgorithm *string `field:"optional" json:"keyAlgorithm" yaml:"keyAlgorithm"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/iam_service_account_key#pgp_key IamServiceAccountKey#pgp_key}.
	PgpKey *string `field:"optional" json:"pgpKey" yaml:"pgpKey"`
}

