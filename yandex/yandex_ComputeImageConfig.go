// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeImageConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_image#description ComputeImage#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_image#family ComputeImage#family}.
	Family *string `field:"optional" json:"family" yaml:"family"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_image#folder_id ComputeImage#folder_id}.
	FolderId *string `field:"optional" json:"folderId" yaml:"folderId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_image#id ComputeImage#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_image#labels ComputeImage#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_image#min_disk_size ComputeImage#min_disk_size}.
	MinDiskSize *float64 `field:"optional" json:"minDiskSize" yaml:"minDiskSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_image#name ComputeImage#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_image#os_type ComputeImage#os_type}.
	OsType *string `field:"optional" json:"osType" yaml:"osType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_image#pooled ComputeImage#pooled}.
	Pooled interface{} `field:"optional" json:"pooled" yaml:"pooled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_image#product_ids ComputeImage#product_ids}.
	ProductIds *[]*string `field:"optional" json:"productIds" yaml:"productIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_image#source_disk ComputeImage#source_disk}.
	SourceDisk *string `field:"optional" json:"sourceDisk" yaml:"sourceDisk"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_image#source_family ComputeImage#source_family}.
	SourceFamily *string `field:"optional" json:"sourceFamily" yaml:"sourceFamily"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_image#source_image ComputeImage#source_image}.
	SourceImage *string `field:"optional" json:"sourceImage" yaml:"sourceImage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_image#source_snapshot ComputeImage#source_snapshot}.
	SourceSnapshot *string `field:"optional" json:"sourceSnapshot" yaml:"sourceSnapshot"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_image#source_url ComputeImage#source_url}.
	SourceUrl *string `field:"optional" json:"sourceUrl" yaml:"sourceUrl"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_image#timeouts ComputeImage#timeouts}
	Timeouts *ComputeImageTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

