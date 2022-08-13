// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ComputeInstanceNetworkInterfaceNatDnsRecord struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#fqdn ComputeInstance#fqdn}.
	Fqdn *string `field:"required" json:"fqdn" yaml:"fqdn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#dns_zone_id ComputeInstance#dns_zone_id}.
	DnsZoneId *string `field:"optional" json:"dnsZoneId" yaml:"dnsZoneId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#ptr ComputeInstance#ptr}.
	Ptr interface{} `field:"optional" json:"ptr" yaml:"ptr"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#ttl ComputeInstance#ttl}.
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

