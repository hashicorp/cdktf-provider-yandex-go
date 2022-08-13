// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocClusterConfig struct {
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
	// cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#cluster_config DataprocCluster#cluster_config}
	ClusterConfig *DataprocClusterClusterConfig `field:"required" json:"clusterConfig" yaml:"clusterConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#name DataprocCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#service_account_id DataprocCluster#service_account_id}.
	ServiceAccountId *string `field:"required" json:"serviceAccountId" yaml:"serviceAccountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#bucket DataprocCluster#bucket}.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#deletion_protection DataprocCluster#deletion_protection}.
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#description DataprocCluster#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#folder_id DataprocCluster#folder_id}.
	FolderId *string `field:"optional" json:"folderId" yaml:"folderId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#host_group_ids DataprocCluster#host_group_ids}.
	HostGroupIds *[]*string `field:"optional" json:"hostGroupIds" yaml:"hostGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#id DataprocCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#labels DataprocCluster#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#security_group_ids DataprocCluster#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#timeouts DataprocCluster#timeouts}
	Timeouts *DataprocClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#ui_proxy DataprocCluster#ui_proxy}.
	UiProxy interface{} `field:"optional" json:"uiProxy" yaml:"uiProxy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#zone_id DataprocCluster#zone_id}.
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

