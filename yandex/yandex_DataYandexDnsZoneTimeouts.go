// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexDnsZoneTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/dns_zone#create DataYandexDnsZone#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/dns_zone#delete DataYandexDnsZone#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/dns_zone#update DataYandexDnsZone#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
