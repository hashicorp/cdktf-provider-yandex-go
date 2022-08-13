// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbVirtualHostModifyResponseHeaders struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#name AlbVirtualHost#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#append AlbVirtualHost#append}.
	Append *string `field:"optional" json:"append" yaml:"append"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#remove AlbVirtualHost#remove}.
	Remove interface{} `field:"optional" json:"remove" yaml:"remove"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#replace AlbVirtualHost#replace}.
	Replace *string `field:"optional" json:"replace" yaml:"replace"`
}

