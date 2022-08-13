// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexIamPolicyBinding struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/iam_policy#members DataYandexIamPolicy#members}.
	Members *[]*string `field:"required" json:"members" yaml:"members"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/iam_policy#role DataYandexIamPolicy#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
}

