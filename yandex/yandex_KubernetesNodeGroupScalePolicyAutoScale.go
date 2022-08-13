// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type KubernetesNodeGroupScalePolicyAutoScale struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#initial KubernetesNodeGroup#initial}.
	Initial *float64 `field:"required" json:"initial" yaml:"initial"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#max KubernetesNodeGroup#max}.
	Max *float64 `field:"required" json:"max" yaml:"max"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#min KubernetesNodeGroup#min}.
	Min *float64 `field:"required" json:"min" yaml:"min"`
}

