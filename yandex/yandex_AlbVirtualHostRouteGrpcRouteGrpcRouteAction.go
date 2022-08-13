// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbVirtualHostRouteGrpcRouteGrpcRouteAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#backend_group_id AlbVirtualHost#backend_group_id}.
	BackendGroupId *string `field:"required" json:"backendGroupId" yaml:"backendGroupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#auto_host_rewrite AlbVirtualHost#auto_host_rewrite}.
	AutoHostRewrite interface{} `field:"optional" json:"autoHostRewrite" yaml:"autoHostRewrite"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#host_rewrite AlbVirtualHost#host_rewrite}.
	HostRewrite *string `field:"optional" json:"hostRewrite" yaml:"hostRewrite"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#idle_timeout AlbVirtualHost#idle_timeout}.
	IdleTimeout *string `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#max_timeout AlbVirtualHost#max_timeout}.
	MaxTimeout *string `field:"optional" json:"maxTimeout" yaml:"maxTimeout"`
}

