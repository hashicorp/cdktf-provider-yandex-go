// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type CdnOriginGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_origin_group#create CdnOriginGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_origin_group#delete CdnOriginGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_origin_group#update CdnOriginGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

