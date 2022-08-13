// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesNodeGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#cluster_id KubernetesNodeGroup#cluster_id}.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// instance_template block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#instance_template KubernetesNodeGroup#instance_template}
	InstanceTemplate *KubernetesNodeGroupInstanceTemplate `field:"required" json:"instanceTemplate" yaml:"instanceTemplate"`
	// scale_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#scale_policy KubernetesNodeGroup#scale_policy}
	ScalePolicy *KubernetesNodeGroupScalePolicy `field:"required" json:"scalePolicy" yaml:"scalePolicy"`
	// allocation_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#allocation_policy KubernetesNodeGroup#allocation_policy}
	AllocationPolicy *KubernetesNodeGroupAllocationPolicy `field:"optional" json:"allocationPolicy" yaml:"allocationPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#allowed_unsafe_sysctls KubernetesNodeGroup#allowed_unsafe_sysctls}.
	AllowedUnsafeSysctls *[]*string `field:"optional" json:"allowedUnsafeSysctls" yaml:"allowedUnsafeSysctls"`
	// deploy_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#deploy_policy KubernetesNodeGroup#deploy_policy}
	DeployPolicy *KubernetesNodeGroupDeployPolicy `field:"optional" json:"deployPolicy" yaml:"deployPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#description KubernetesNodeGroup#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#id KubernetesNodeGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#labels KubernetesNodeGroup#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// maintenance_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#maintenance_policy KubernetesNodeGroup#maintenance_policy}
	MaintenancePolicy *KubernetesNodeGroupMaintenancePolicy `field:"optional" json:"maintenancePolicy" yaml:"maintenancePolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#name KubernetesNodeGroup#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#node_labels KubernetesNodeGroup#node_labels}.
	NodeLabels *map[string]*string `field:"optional" json:"nodeLabels" yaml:"nodeLabels"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#node_taints KubernetesNodeGroup#node_taints}.
	NodeTaints *[]*string `field:"optional" json:"nodeTaints" yaml:"nodeTaints"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#timeouts KubernetesNodeGroup#timeouts}
	Timeouts *KubernetesNodeGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#version KubernetesNodeGroup#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

