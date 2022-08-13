// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbGreenplumClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#assign_public_ip MdbGreenplumCluster#assign_public_ip}.
	AssignPublicIp interface{} `field:"required" json:"assignPublicIp" yaml:"assignPublicIp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#environment MdbGreenplumCluster#environment}.
	Environment *string `field:"required" json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#master_host_count MdbGreenplumCluster#master_host_count}.
	MasterHostCount *float64 `field:"required" json:"masterHostCount" yaml:"masterHostCount"`
	// master_subcluster block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#master_subcluster MdbGreenplumCluster#master_subcluster}
	MasterSubcluster *MdbGreenplumClusterMasterSubcluster `field:"required" json:"masterSubcluster" yaml:"masterSubcluster"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#name MdbGreenplumCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#network_id MdbGreenplumCluster#network_id}.
	NetworkId *string `field:"required" json:"networkId" yaml:"networkId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#segment_host_count MdbGreenplumCluster#segment_host_count}.
	SegmentHostCount *float64 `field:"required" json:"segmentHostCount" yaml:"segmentHostCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#segment_in_host MdbGreenplumCluster#segment_in_host}.
	SegmentInHost *float64 `field:"required" json:"segmentInHost" yaml:"segmentInHost"`
	// segment_subcluster block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#segment_subcluster MdbGreenplumCluster#segment_subcluster}
	SegmentSubcluster *MdbGreenplumClusterSegmentSubcluster `field:"required" json:"segmentSubcluster" yaml:"segmentSubcluster"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#subnet_id MdbGreenplumCluster#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#user_name MdbGreenplumCluster#user_name}.
	UserName *string `field:"required" json:"userName" yaml:"userName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#user_password MdbGreenplumCluster#user_password}.
	UserPassword *string `field:"required" json:"userPassword" yaml:"userPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#version MdbGreenplumCluster#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#zone MdbGreenplumCluster#zone}.
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// access block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#access MdbGreenplumCluster#access}
	Access *MdbGreenplumClusterAccess `field:"optional" json:"access" yaml:"access"`
	// backup_window_start block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#backup_window_start MdbGreenplumCluster#backup_window_start}
	BackupWindowStart *MdbGreenplumClusterBackupWindowStart `field:"optional" json:"backupWindowStart" yaml:"backupWindowStart"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#deletion_protection MdbGreenplumCluster#deletion_protection}.
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#description MdbGreenplumCluster#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#folder_id MdbGreenplumCluster#folder_id}.
	FolderId *string `field:"optional" json:"folderId" yaml:"folderId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#greenplum_config MdbGreenplumCluster#greenplum_config}.
	GreenplumConfig *map[string]*string `field:"optional" json:"greenplumConfig" yaml:"greenplumConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#id MdbGreenplumCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#labels MdbGreenplumCluster#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// pooler_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#pooler_config MdbGreenplumCluster#pooler_config}
	PoolerConfig *MdbGreenplumClusterPoolerConfig `field:"optional" json:"poolerConfig" yaml:"poolerConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#security_group_ids MdbGreenplumCluster#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#timeouts MdbGreenplumCluster#timeouts}
	Timeouts *MdbGreenplumClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

