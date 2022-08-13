// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbBackendGroupSessionAffinityConnection struct {
	// Use source IP address.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#source_ip AlbBackendGroup#source_ip}
	SourceIp interface{} `field:"optional" json:"sourceIp" yaml:"sourceIp"`
}

