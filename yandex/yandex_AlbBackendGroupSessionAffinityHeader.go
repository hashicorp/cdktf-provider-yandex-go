// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbBackendGroupSessionAffinityHeader struct {
	// The name of the request header that will be used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#header_name AlbBackendGroup#header_name}
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
}

