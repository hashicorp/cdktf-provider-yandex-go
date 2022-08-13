// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type CdnOriginGroupOrigin struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_origin_group#source CdnOriginGroup#source}.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_origin_group#backup CdnOriginGroup#backup}.
	Backup interface{} `field:"optional" json:"backup" yaml:"backup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_origin_group#enabled CdnOriginGroup#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

