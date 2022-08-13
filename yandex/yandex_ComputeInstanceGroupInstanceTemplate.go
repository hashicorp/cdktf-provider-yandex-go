// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ComputeInstanceGroupInstanceTemplate struct {
	// boot_disk block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#boot_disk ComputeInstanceGroup#boot_disk}
	BootDisk *ComputeInstanceGroupInstanceTemplateBootDisk `field:"required" json:"bootDisk" yaml:"bootDisk"`
	// network_interface block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#network_interface ComputeInstanceGroup#network_interface}
	NetworkInterface interface{} `field:"required" json:"networkInterface" yaml:"networkInterface"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#resources ComputeInstanceGroup#resources}
	Resources *ComputeInstanceGroupInstanceTemplateResources `field:"required" json:"resources" yaml:"resources"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#description ComputeInstanceGroup#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#hostname ComputeInstanceGroup#hostname}.
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#labels ComputeInstanceGroup#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#metadata ComputeInstanceGroup#metadata}.
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#name ComputeInstanceGroup#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// network_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#network_settings ComputeInstanceGroup#network_settings}
	NetworkSettings interface{} `field:"optional" json:"networkSettings" yaml:"networkSettings"`
	// placement_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#placement_policy ComputeInstanceGroup#placement_policy}
	PlacementPolicy *ComputeInstanceGroupInstanceTemplatePlacementPolicy `field:"optional" json:"placementPolicy" yaml:"placementPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#platform_id ComputeInstanceGroup#platform_id}.
	PlatformId *string `field:"optional" json:"platformId" yaml:"platformId"`
	// scheduling_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#scheduling_policy ComputeInstanceGroup#scheduling_policy}
	SchedulingPolicy *ComputeInstanceGroupInstanceTemplateSchedulingPolicy `field:"optional" json:"schedulingPolicy" yaml:"schedulingPolicy"`
	// secondary_disk block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#secondary_disk ComputeInstanceGroup#secondary_disk}
	SecondaryDisk interface{} `field:"optional" json:"secondaryDisk" yaml:"secondaryDisk"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#service_account_id ComputeInstanceGroup#service_account_id}.
	ServiceAccountId *string `field:"optional" json:"serviceAccountId" yaml:"serviceAccountId"`
}

