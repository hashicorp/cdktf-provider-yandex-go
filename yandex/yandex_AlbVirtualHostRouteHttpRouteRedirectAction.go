// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbVirtualHostRouteHttpRouteRedirectAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#remove_query AlbVirtualHost#remove_query}.
	RemoveQuery interface{} `field:"optional" json:"removeQuery" yaml:"removeQuery"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#replace_host AlbVirtualHost#replace_host}.
	ReplaceHost *string `field:"optional" json:"replaceHost" yaml:"replaceHost"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#replace_path AlbVirtualHost#replace_path}.
	ReplacePath *string `field:"optional" json:"replacePath" yaml:"replacePath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#replace_port AlbVirtualHost#replace_port}.
	ReplacePort *float64 `field:"optional" json:"replacePort" yaml:"replacePort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#replace_prefix AlbVirtualHost#replace_prefix}.
	ReplacePrefix *string `field:"optional" json:"replacePrefix" yaml:"replacePrefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#replace_scheme AlbVirtualHost#replace_scheme}.
	ReplaceScheme *string `field:"optional" json:"replaceScheme" yaml:"replaceScheme"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#response_code AlbVirtualHost#response_code}.
	ResponseCode *string `field:"optional" json:"responseCode" yaml:"responseCode"`
}

