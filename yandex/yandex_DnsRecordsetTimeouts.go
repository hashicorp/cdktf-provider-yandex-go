// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DnsRecordsetTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dns_recordset#create DnsRecordset#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dns_recordset#delete DnsRecordset#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dns_recordset#update DnsRecordset#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

