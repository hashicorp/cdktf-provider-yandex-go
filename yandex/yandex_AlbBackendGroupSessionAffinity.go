// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbBackendGroupSessionAffinity struct {
	// connection block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#connection AlbBackendGroup#connection}
	Connection *AlbBackendGroupSessionAffinityConnection `field:"optional" json:"connection" yaml:"connection"`
	// cookie block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#cookie AlbBackendGroup#cookie}
	Cookie *AlbBackendGroupSessionAffinityCookie `field:"optional" json:"cookie" yaml:"cookie"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#header AlbBackendGroup#header}
	Header *AlbBackendGroupSessionAffinityHeader `field:"optional" json:"header" yaml:"header"`
}

