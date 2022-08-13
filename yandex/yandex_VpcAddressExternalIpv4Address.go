// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type VpcAddressExternalIpv4Address struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_address#ddos_protection_provider VpcAddress#ddos_protection_provider}.
	DdosProtectionProvider *string `field:"optional" json:"ddosProtectionProvider" yaml:"ddosProtectionProvider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_address#outgoing_smtp_capability VpcAddress#outgoing_smtp_capability}.
	OutgoingSmtpCapability *string `field:"optional" json:"outgoingSmtpCapability" yaml:"outgoingSmtpCapability"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_address#zone_id VpcAddress#zone_id}.
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

