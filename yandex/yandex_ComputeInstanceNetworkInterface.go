// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ComputeInstanceNetworkInterface struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#subnet_id ComputeInstance#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// dns_record block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#dns_record ComputeInstance#dns_record}
	DnsRecord interface{} `field:"optional" json:"dnsRecord" yaml:"dnsRecord"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#ip_address ComputeInstance#ip_address}.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#ipv4 ComputeInstance#ipv4}.
	Ipv4 interface{} `field:"optional" json:"ipv4" yaml:"ipv4"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#ipv6 ComputeInstance#ipv6}.
	Ipv6 interface{} `field:"optional" json:"ipv6" yaml:"ipv6"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#ipv6_address ComputeInstance#ipv6_address}.
	Ipv6Address *string `field:"optional" json:"ipv6Address" yaml:"ipv6Address"`
	// ipv6_dns_record block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#ipv6_dns_record ComputeInstance#ipv6_dns_record}
	Ipv6DnsRecord interface{} `field:"optional" json:"ipv6DnsRecord" yaml:"ipv6DnsRecord"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#nat ComputeInstance#nat}.
	Nat interface{} `field:"optional" json:"nat" yaml:"nat"`
	// nat_dns_record block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#nat_dns_record ComputeInstance#nat_dns_record}
	NatDnsRecord interface{} `field:"optional" json:"natDnsRecord" yaml:"natDnsRecord"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#nat_ip_address ComputeInstance#nat_ip_address}.
	NatIpAddress *string `field:"optional" json:"natIpAddress" yaml:"natIpAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#security_group_ids ComputeInstance#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
}

