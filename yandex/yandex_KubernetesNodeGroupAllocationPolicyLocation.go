// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type KubernetesNodeGroupAllocationPolicyLocation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#subnet_id KubernetesNodeGroup#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#zone KubernetesNodeGroup#zone}.
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

