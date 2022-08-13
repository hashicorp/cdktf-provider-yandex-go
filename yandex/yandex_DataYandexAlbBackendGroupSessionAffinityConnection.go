// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexAlbBackendGroupSessionAffinityConnection struct {
	// Use source IP address.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#source_ip DataYandexAlbBackendGroup#source_ip}
	SourceIp interface{} `field:"optional" json:"sourceIp" yaml:"sourceIp"`
}

