// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type IamServiceAccountTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/iam_service_account#create IamServiceAccount#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/iam_service_account#delete IamServiceAccount#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/iam_service_account#update IamServiceAccount#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

