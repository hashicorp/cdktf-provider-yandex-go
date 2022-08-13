// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type KubernetesNodeGroupMaintenancePolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#auto_repair KubernetesNodeGroup#auto_repair}.
	AutoRepair interface{} `field:"required" json:"autoRepair" yaml:"autoRepair"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#auto_upgrade KubernetesNodeGroup#auto_upgrade}.
	AutoUpgrade interface{} `field:"required" json:"autoUpgrade" yaml:"autoUpgrade"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#maintenance_window KubernetesNodeGroup#maintenance_window}
	MaintenanceWindow interface{} `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
}

