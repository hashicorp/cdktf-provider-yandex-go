// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbVirtualHostConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#http_router_id AlbVirtualHost#http_router_id}.
	HttpRouterId *string `field:"required" json:"httpRouterId" yaml:"httpRouterId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#name AlbVirtualHost#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#authority AlbVirtualHost#authority}.
	Authority *[]*string `field:"optional" json:"authority" yaml:"authority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#id AlbVirtualHost#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// modify_request_headers block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#modify_request_headers AlbVirtualHost#modify_request_headers}
	ModifyRequestHeaders interface{} `field:"optional" json:"modifyRequestHeaders" yaml:"modifyRequestHeaders"`
	// modify_response_headers block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#modify_response_headers AlbVirtualHost#modify_response_headers}
	ModifyResponseHeaders interface{} `field:"optional" json:"modifyResponseHeaders" yaml:"modifyResponseHeaders"`
	// route block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#route AlbVirtualHost#route}
	Route interface{} `field:"optional" json:"route" yaml:"route"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host#timeouts AlbVirtualHost#timeouts}
	Timeouts *AlbVirtualHostTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

