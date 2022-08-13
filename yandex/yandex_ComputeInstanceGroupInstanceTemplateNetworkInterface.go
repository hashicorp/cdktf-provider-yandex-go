// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ComputeInstanceGroupInstanceTemplateNetworkInterface struct {
	// dns_record block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#dns_record ComputeInstanceGroup#dns_record}
	DnsRecord interface{} `field:"optional" json:"dnsRecord" yaml:"dnsRecord"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#ip_address ComputeInstanceGroup#ip_address}.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#ipv4 ComputeInstanceGroup#ipv4}.
	Ipv4 interface{} `field:"optional" json:"ipv4" yaml:"ipv4"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#ipv6 ComputeInstanceGroup#ipv6}.
	Ipv6 interface{} `field:"optional" json:"ipv6" yaml:"ipv6"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#ipv6_address ComputeInstanceGroup#ipv6_address}.
	Ipv6Address *string `field:"optional" json:"ipv6Address" yaml:"ipv6Address"`
	// ipv6_dns_record block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#ipv6_dns_record ComputeInstanceGroup#ipv6_dns_record}
	Ipv6DnsRecord interface{} `field:"optional" json:"ipv6DnsRecord" yaml:"ipv6DnsRecord"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#nat ComputeInstanceGroup#nat}.
	Nat interface{} `field:"optional" json:"nat" yaml:"nat"`
	// nat_dns_record block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#nat_dns_record ComputeInstanceGroup#nat_dns_record}
	NatDnsRecord interface{} `field:"optional" json:"natDnsRecord" yaml:"natDnsRecord"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#nat_ip_address ComputeInstanceGroup#nat_ip_address}.
	NatIpAddress *string `field:"optional" json:"natIpAddress" yaml:"natIpAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#network_id ComputeInstanceGroup#network_id}.
	NetworkId *string `field:"optional" json:"networkId" yaml:"networkId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#security_group_ids ComputeInstanceGroup#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#subnet_ids ComputeInstanceGroup#subnet_ids}.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

