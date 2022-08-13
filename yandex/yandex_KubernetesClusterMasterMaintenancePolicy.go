// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type KubernetesClusterMasterMaintenancePolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#auto_upgrade KubernetesCluster#auto_upgrade}.
	AutoUpgrade interface{} `field:"required" json:"autoUpgrade" yaml:"autoUpgrade"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#maintenance_window KubernetesCluster#maintenance_window}
	MaintenanceWindow interface{} `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
}

