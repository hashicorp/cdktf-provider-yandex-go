// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type KubernetesClusterMasterMaintenancePolicyMaintenanceWindow struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#duration KubernetesCluster#duration}.
	Duration *string `field:"required" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#start_time KubernetesCluster#start_time}.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#day KubernetesCluster#day}.
	Day *string `field:"optional" json:"day" yaml:"day"`
}

