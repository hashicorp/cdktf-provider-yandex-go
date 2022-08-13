// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type KubernetesClusterMasterZonal struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#subnet_id KubernetesCluster#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#zone KubernetesCluster#zone}.
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

