// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type KubernetesNodeGroupInstanceTemplateNetworkInterface struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#subnet_ids KubernetesNodeGroup#subnet_ids}.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#ipv4 KubernetesNodeGroup#ipv4}.
	Ipv4 interface{} `field:"optional" json:"ipv4" yaml:"ipv4"`
	// ipv4_dns_records block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#ipv4_dns_records KubernetesNodeGroup#ipv4_dns_records}
	Ipv4DnsRecords interface{} `field:"optional" json:"ipv4DnsRecords" yaml:"ipv4DnsRecords"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#ipv6 KubernetesNodeGroup#ipv6}.
	Ipv6 interface{} `field:"optional" json:"ipv6" yaml:"ipv6"`
	// ipv6_dns_records block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#ipv6_dns_records KubernetesNodeGroup#ipv6_dns_records}
	Ipv6DnsRecords interface{} `field:"optional" json:"ipv6DnsRecords" yaml:"ipv6DnsRecords"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#nat KubernetesNodeGroup#nat}.
	Nat interface{} `field:"optional" json:"nat" yaml:"nat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#security_group_ids KubernetesNodeGroup#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
}

