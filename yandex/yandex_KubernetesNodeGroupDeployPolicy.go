// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type KubernetesNodeGroupDeployPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#max_expansion KubernetesNodeGroup#max_expansion}.
	MaxExpansion *float64 `field:"required" json:"maxExpansion" yaml:"maxExpansion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#max_unavailable KubernetesNodeGroup#max_unavailable}.
	MaxUnavailable *float64 `field:"required" json:"maxUnavailable" yaml:"maxUnavailable"`
}

