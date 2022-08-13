// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeInstanceConfig struct {
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
	// boot_disk block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#boot_disk ComputeInstance#boot_disk}
	BootDisk *ComputeInstanceBootDisk `field:"required" json:"bootDisk" yaml:"bootDisk"`
	// network_interface block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#network_interface ComputeInstance#network_interface}
	NetworkInterface interface{} `field:"required" json:"networkInterface" yaml:"networkInterface"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#resources ComputeInstance#resources}
	Resources *ComputeInstanceResources `field:"required" json:"resources" yaml:"resources"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#allow_recreate ComputeInstance#allow_recreate}.
	AllowRecreate interface{} `field:"optional" json:"allowRecreate" yaml:"allowRecreate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#allow_stopping_for_update ComputeInstance#allow_stopping_for_update}.
	AllowStoppingForUpdate interface{} `field:"optional" json:"allowStoppingForUpdate" yaml:"allowStoppingForUpdate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#description ComputeInstance#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#folder_id ComputeInstance#folder_id}.
	FolderId *string `field:"optional" json:"folderId" yaml:"folderId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#hostname ComputeInstance#hostname}.
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#id ComputeInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#labels ComputeInstance#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// local_disk block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#local_disk ComputeInstance#local_disk}
	LocalDisk interface{} `field:"optional" json:"localDisk" yaml:"localDisk"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#metadata ComputeInstance#metadata}.
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#name ComputeInstance#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#network_acceleration_type ComputeInstance#network_acceleration_type}.
	NetworkAccelerationType *string `field:"optional" json:"networkAccelerationType" yaml:"networkAccelerationType"`
	// placement_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#placement_policy ComputeInstance#placement_policy}
	PlacementPolicy *ComputeInstancePlacementPolicy `field:"optional" json:"placementPolicy" yaml:"placementPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#platform_id ComputeInstance#platform_id}.
	PlatformId *string `field:"optional" json:"platformId" yaml:"platformId"`
	// scheduling_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#scheduling_policy ComputeInstance#scheduling_policy}
	SchedulingPolicy *ComputeInstanceSchedulingPolicy `field:"optional" json:"schedulingPolicy" yaml:"schedulingPolicy"`
	// secondary_disk block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#secondary_disk ComputeInstance#secondary_disk}
	SecondaryDisk interface{} `field:"optional" json:"secondaryDisk" yaml:"secondaryDisk"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#service_account_id ComputeInstance#service_account_id}.
	ServiceAccountId *string `field:"optional" json:"serviceAccountId" yaml:"serviceAccountId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#timeouts ComputeInstance#timeouts}
	Timeouts *ComputeInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance#zone ComputeInstance#zone}.
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

