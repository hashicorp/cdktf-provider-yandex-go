// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type KubernetesNodeGroupInstanceTemplate struct {
	// boot_disk block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#boot_disk KubernetesNodeGroup#boot_disk}
	BootDisk *KubernetesNodeGroupInstanceTemplateBootDisk `field:"optional" json:"bootDisk" yaml:"bootDisk"`
	// container_runtime block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#container_runtime KubernetesNodeGroup#container_runtime}
	ContainerRuntime *KubernetesNodeGroupInstanceTemplateContainerRuntime `field:"optional" json:"containerRuntime" yaml:"containerRuntime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#labels KubernetesNodeGroup#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#metadata KubernetesNodeGroup#metadata}.
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#name KubernetesNodeGroup#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#nat KubernetesNodeGroup#nat}.
	Nat interface{} `field:"optional" json:"nat" yaml:"nat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#network_acceleration_type KubernetesNodeGroup#network_acceleration_type}.
	NetworkAccelerationType *string `field:"optional" json:"networkAccelerationType" yaml:"networkAccelerationType"`
	// network_interface block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#network_interface KubernetesNodeGroup#network_interface}
	NetworkInterface interface{} `field:"optional" json:"networkInterface" yaml:"networkInterface"`
	// placement_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#placement_policy KubernetesNodeGroup#placement_policy}
	PlacementPolicy *KubernetesNodeGroupInstanceTemplatePlacementPolicy `field:"optional" json:"placementPolicy" yaml:"placementPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#platform_id KubernetesNodeGroup#platform_id}.
	PlatformId *string `field:"optional" json:"platformId" yaml:"platformId"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#resources KubernetesNodeGroup#resources}
	Resources *KubernetesNodeGroupInstanceTemplateResources `field:"optional" json:"resources" yaml:"resources"`
	// scheduling_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#scheduling_policy KubernetesNodeGroup#scheduling_policy}
	SchedulingPolicy *KubernetesNodeGroupInstanceTemplateSchedulingPolicy `field:"optional" json:"schedulingPolicy" yaml:"schedulingPolicy"`
}

