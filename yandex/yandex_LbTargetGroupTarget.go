// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type LbTargetGroupTarget struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/lb_target_group#address LbTargetGroup#address}.
	Address *string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/lb_target_group#subnet_id LbTargetGroup#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}

