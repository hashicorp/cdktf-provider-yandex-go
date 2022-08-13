// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type KubernetesNodeGroupInstanceTemplateSchedulingPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group#preemptible KubernetesNodeGroup#preemptible}.
	Preemptible interface{} `field:"optional" json:"preemptible" yaml:"preemptible"`
}

