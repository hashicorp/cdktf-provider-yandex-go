// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbBackendGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#create AlbBackendGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#delete AlbBackendGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#update AlbBackendGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

