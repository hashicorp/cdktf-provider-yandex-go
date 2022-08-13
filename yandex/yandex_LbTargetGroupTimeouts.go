// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type LbTargetGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/lb_target_group#create LbTargetGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/lb_target_group#delete LbTargetGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/lb_target_group#update LbTargetGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

