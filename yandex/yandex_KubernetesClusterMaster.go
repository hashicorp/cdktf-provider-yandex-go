// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type KubernetesClusterMaster struct {
	// maintenance_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#maintenance_policy KubernetesCluster#maintenance_policy}
	MaintenancePolicy *KubernetesClusterMasterMaintenancePolicy `field:"optional" json:"maintenancePolicy" yaml:"maintenancePolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#public_ip KubernetesCluster#public_ip}.
	PublicIp interface{} `field:"optional" json:"publicIp" yaml:"publicIp"`
	// regional block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#regional KubernetesCluster#regional}
	Regional *KubernetesClusterMasterRegional `field:"optional" json:"regional" yaml:"regional"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#security_group_ids KubernetesCluster#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#version KubernetesCluster#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
	// zonal block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#zonal KubernetesCluster#zonal}
	Zonal *KubernetesClusterMasterZonal `field:"optional" json:"zonal" yaml:"zonal"`
}

