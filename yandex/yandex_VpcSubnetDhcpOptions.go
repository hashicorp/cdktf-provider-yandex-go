// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type VpcSubnetDhcpOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_subnet#domain_name VpcSubnet#domain_name}.
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_subnet#domain_name_servers VpcSubnet#domain_name_servers}.
	DomainNameServers *[]*string `field:"optional" json:"domainNameServers" yaml:"domainNameServers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/vpc_subnet#ntp_servers VpcSubnet#ntp_servers}.
	NtpServers *[]*string `field:"optional" json:"ntpServers" yaml:"ntpServers"`
}

