// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ComputeInstanceGroupInstanceTemplateNetworkInterfaceNatDnsRecord struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#fqdn ComputeInstanceGroup#fqdn}.
	Fqdn *string `field:"required" json:"fqdn" yaml:"fqdn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#dns_zone_id ComputeInstanceGroup#dns_zone_id}.
	DnsZoneId *string `field:"optional" json:"dnsZoneId" yaml:"dnsZoneId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#ptr ComputeInstanceGroup#ptr}.
	Ptr interface{} `field:"optional" json:"ptr" yaml:"ptr"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#ttl ComputeInstanceGroup#ttl}.
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

