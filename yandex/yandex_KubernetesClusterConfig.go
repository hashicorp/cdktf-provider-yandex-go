// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesClusterConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// master block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#master KubernetesCluster#master}
	Master *KubernetesClusterMaster `field:"required" json:"master" yaml:"master"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#network_id KubernetesCluster#network_id}.
	NetworkId *string `field:"required" json:"networkId" yaml:"networkId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#node_service_account_id KubernetesCluster#node_service_account_id}.
	NodeServiceAccountId *string `field:"required" json:"nodeServiceAccountId" yaml:"nodeServiceAccountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#service_account_id KubernetesCluster#service_account_id}.
	ServiceAccountId *string `field:"required" json:"serviceAccountId" yaml:"serviceAccountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#cluster_ipv4_range KubernetesCluster#cluster_ipv4_range}.
	ClusterIpv4Range *string `field:"optional" json:"clusterIpv4Range" yaml:"clusterIpv4Range"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#cluster_ipv6_range KubernetesCluster#cluster_ipv6_range}.
	ClusterIpv6Range *string `field:"optional" json:"clusterIpv6Range" yaml:"clusterIpv6Range"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#description KubernetesCluster#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#folder_id KubernetesCluster#folder_id}.
	FolderId *string `field:"optional" json:"folderId" yaml:"folderId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#id KubernetesCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// kms_provider block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#kms_provider KubernetesCluster#kms_provider}
	KmsProvider *KubernetesClusterKmsProvider `field:"optional" json:"kmsProvider" yaml:"kmsProvider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#labels KubernetesCluster#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#name KubernetesCluster#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// network_implementation block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#network_implementation KubernetesCluster#network_implementation}
	NetworkImplementation *KubernetesClusterNetworkImplementation `field:"optional" json:"networkImplementation" yaml:"networkImplementation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#network_policy_provider KubernetesCluster#network_policy_provider}.
	NetworkPolicyProvider *string `field:"optional" json:"networkPolicyProvider" yaml:"networkPolicyProvider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#node_ipv4_cidr_mask_size KubernetesCluster#node_ipv4_cidr_mask_size}.
	NodeIpv4CidrMaskSize *float64 `field:"optional" json:"nodeIpv4CidrMaskSize" yaml:"nodeIpv4CidrMaskSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#release_channel KubernetesCluster#release_channel}.
	ReleaseChannel *string `field:"optional" json:"releaseChannel" yaml:"releaseChannel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#service_ipv4_range KubernetesCluster#service_ipv4_range}.
	ServiceIpv4Range *string `field:"optional" json:"serviceIpv4Range" yaml:"serviceIpv4Range"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#service_ipv6_range KubernetesCluster#service_ipv6_range}.
	ServiceIpv6Range *string `field:"optional" json:"serviceIpv6Range" yaml:"serviceIpv6Range"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster#timeouts KubernetesCluster#timeouts}
	Timeouts *KubernetesClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

