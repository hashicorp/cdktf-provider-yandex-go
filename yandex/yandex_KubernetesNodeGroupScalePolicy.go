// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type KubernetesNodeGroupScalePolicy struct {
	// auto_scale block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#auto_scale KubernetesNodeGroup#auto_scale}
	AutoScale *KubernetesNodeGroupScalePolicyAutoScale `field:"optional" json:"autoScale" yaml:"autoScale"`
	// fixed_scale block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#fixed_scale KubernetesNodeGroup#fixed_scale}
	FixedScale *KubernetesNodeGroupScalePolicyFixedScale `field:"optional" json:"fixedScale" yaml:"fixedScale"`
}

