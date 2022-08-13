// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbBackendGroupSessionAffinityCookie struct {
	// Name of the HTTP cookie.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#name AlbBackendGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// TTL for the cookie (if not set, session cookie will be used).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#ttl AlbBackendGroup#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
}

