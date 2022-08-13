// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbVirtualHostTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#create AlbVirtualHost#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#delete AlbVirtualHost#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#update AlbVirtualHost#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

