// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type KubernetesNodeGroupInstanceTemplatePlacementPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#placement_group_id KubernetesNodeGroup#placement_group_id}.
	PlacementGroupId *string `field:"required" json:"placementGroupId" yaml:"placementGroupId"`
}

