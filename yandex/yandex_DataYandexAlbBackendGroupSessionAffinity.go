// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexAlbBackendGroupSessionAffinity struct {
	// connection block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#connection DataYandexAlbBackendGroup#connection}
	Connection *DataYandexAlbBackendGroupSessionAffinityConnection `field:"optional" json:"connection" yaml:"connection"`
	// cookie block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#cookie DataYandexAlbBackendGroup#cookie}
	Cookie *DataYandexAlbBackendGroupSessionAffinityCookie `field:"optional" json:"cookie" yaml:"cookie"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#header DataYandexAlbBackendGroup#header}
	Header *DataYandexAlbBackendGroupSessionAffinityHeader `field:"optional" json:"header" yaml:"header"`
}

