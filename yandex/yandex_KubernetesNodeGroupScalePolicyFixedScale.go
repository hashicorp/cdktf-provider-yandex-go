// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type KubernetesNodeGroupScalePolicyFixedScale struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#size KubernetesNodeGroup#size}.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

