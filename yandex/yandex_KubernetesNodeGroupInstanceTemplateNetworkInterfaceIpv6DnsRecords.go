// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv6DnsRecords struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#fqdn KubernetesNodeGroup#fqdn}.
	Fqdn *string `field:"required" json:"fqdn" yaml:"fqdn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#dns_zone_id KubernetesNodeGroup#dns_zone_id}.
	DnsZoneId *string `field:"optional" json:"dnsZoneId" yaml:"dnsZoneId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#ptr KubernetesNodeGroup#ptr}.
	Ptr interface{} `field:"optional" json:"ptr" yaml:"ptr"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#ttl KubernetesNodeGroup#ttl}.
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

