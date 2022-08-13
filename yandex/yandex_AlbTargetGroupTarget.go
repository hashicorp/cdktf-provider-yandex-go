// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbTargetGroupTarget struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_target_group#ip_address AlbTargetGroup#ip_address}.
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_target_group#private_ipv4_address AlbTargetGroup#private_ipv4_address}.
	PrivateIpv4Address interface{} `field:"optional" json:"privateIpv4Address" yaml:"privateIpv4Address"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_target_group#subnet_id AlbTargetGroup#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

