// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type KubernetesClusterMasterRegional struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#region KubernetesCluster#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// location block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#location KubernetesCluster#location}
	Location interface{} `field:"optional" json:"location" yaml:"location"`
}

