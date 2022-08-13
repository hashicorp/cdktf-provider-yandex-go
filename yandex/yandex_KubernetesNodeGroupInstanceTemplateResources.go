// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type KubernetesNodeGroupInstanceTemplateResources struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#core_fraction KubernetesNodeGroup#core_fraction}.
	CoreFraction *float64 `field:"optional" json:"coreFraction" yaml:"coreFraction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#cores KubernetesNodeGroup#cores}.
	Cores *float64 `field:"optional" json:"cores" yaml:"cores"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#gpus KubernetesNodeGroup#gpus}.
	Gpus *float64 `field:"optional" json:"gpus" yaml:"gpus"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#memory KubernetesNodeGroup#memory}.
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
}

