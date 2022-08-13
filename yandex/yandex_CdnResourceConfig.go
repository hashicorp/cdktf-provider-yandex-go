// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CdnResourceConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#active CdnResource#active}.
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#cname CdnResource#cname}.
	Cname *string `field:"optional" json:"cname" yaml:"cname"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#folder_id CdnResource#folder_id}.
	FolderId *string `field:"optional" json:"folderId" yaml:"folderId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#id CdnResource#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#options CdnResource#options}
	Options *CdnResourceOptions `field:"optional" json:"options" yaml:"options"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#origin_group_id CdnResource#origin_group_id}.
	OriginGroupId *float64 `field:"optional" json:"originGroupId" yaml:"originGroupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#origin_group_name CdnResource#origin_group_name}.
	OriginGroupName *string `field:"optional" json:"originGroupName" yaml:"originGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#origin_protocol CdnResource#origin_protocol}.
	OriginProtocol *string `field:"optional" json:"originProtocol" yaml:"originProtocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#secondary_hostnames CdnResource#secondary_hostnames}.
	SecondaryHostnames *[]*string `field:"optional" json:"secondaryHostnames" yaml:"secondaryHostnames"`
	// ssl_certificate block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#ssl_certificate CdnResource#ssl_certificate}
	SslCertificate *CdnResourceSslCertificate `field:"optional" json:"sslCertificate" yaml:"sslCertificate"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#timeouts CdnResource#timeouts}
	Timeouts *CdnResourceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#updated_at CdnResource#updated_at}.
	UpdatedAt *string `field:"optional" json:"updatedAt" yaml:"updatedAt"`
}

