// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type KubernetesClusterKmsProvider struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#key_id KubernetesCluster#key_id}.
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
}

