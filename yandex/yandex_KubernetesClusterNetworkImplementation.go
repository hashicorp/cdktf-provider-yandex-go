// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type KubernetesClusterNetworkImplementation struct {
	// cilium block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#cilium KubernetesCluster#cilium}
	Cilium *KubernetesClusterNetworkImplementationCilium `field:"optional" json:"cilium" yaml:"cilium"`
}

