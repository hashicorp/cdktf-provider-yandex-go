// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexAlbBackendGroupSessionAffinityCookie struct {
	// TTL for the cookie (if not set, session cookie will be used).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#ttl DataYandexAlbBackendGroup#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
}

