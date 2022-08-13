// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type CdnResourceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#create CdnResource#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#delete CdnResource#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#update CdnResource#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

