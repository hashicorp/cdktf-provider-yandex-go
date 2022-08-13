// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ServerlessContainerTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/serverless_container#create ServerlessContainer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/serverless_container#delete ServerlessContainer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/serverless_container#update ServerlessContainer#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

