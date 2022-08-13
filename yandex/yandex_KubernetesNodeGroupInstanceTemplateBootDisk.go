// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type KubernetesNodeGroupInstanceTemplateBootDisk struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#size KubernetesNodeGroup#size}.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#type KubernetesNodeGroup#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

