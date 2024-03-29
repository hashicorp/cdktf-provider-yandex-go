// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DnsZoneTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dns_zone#create DnsZone#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dns_zone#delete DnsZone#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dns_zone#update DnsZone#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

