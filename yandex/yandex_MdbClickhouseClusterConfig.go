// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbClickhouseClusterConfig struct {
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
	// clickhouse block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#clickhouse MdbClickhouseCluster#clickhouse}
	Clickhouse *MdbClickhouseClusterClickhouse `field:"required" json:"clickhouse" yaml:"clickhouse"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#environment MdbClickhouseCluster#environment}.
	Environment *string `field:"required" json:"environment" yaml:"environment"`
	// host block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#host MdbClickhouseCluster#host}
	Host interface{} `field:"required" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#name MdbClickhouseCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#network_id MdbClickhouseCluster#network_id}.
	NetworkId *string `field:"required" json:"networkId" yaml:"networkId"`
	// access block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#access MdbClickhouseCluster#access}
	Access *MdbClickhouseClusterAccess `field:"optional" json:"access" yaml:"access"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#admin_password MdbClickhouseCluster#admin_password}.
	AdminPassword *string `field:"optional" json:"adminPassword" yaml:"adminPassword"`
	// backup_window_start block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#backup_window_start MdbClickhouseCluster#backup_window_start}
	BackupWindowStart *MdbClickhouseClusterBackupWindowStart `field:"optional" json:"backupWindowStart" yaml:"backupWindowStart"`
	// cloud_storage block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#cloud_storage MdbClickhouseCluster#cloud_storage}
	CloudStorage *MdbClickhouseClusterCloudStorage `field:"optional" json:"cloudStorage" yaml:"cloudStorage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#copy_schema_on_new_hosts MdbClickhouseCluster#copy_schema_on_new_hosts}.
	CopySchemaOnNewHosts interface{} `field:"optional" json:"copySchemaOnNewHosts" yaml:"copySchemaOnNewHosts"`
	// database block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#database MdbClickhouseCluster#database}
	Database interface{} `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#deletion_protection MdbClickhouseCluster#deletion_protection}.
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#description MdbClickhouseCluster#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#embedded_keeper MdbClickhouseCluster#embedded_keeper}.
	EmbeddedKeeper interface{} `field:"optional" json:"embeddedKeeper" yaml:"embeddedKeeper"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#folder_id MdbClickhouseCluster#folder_id}.
	FolderId *string `field:"optional" json:"folderId" yaml:"folderId"`
	// format_schema block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#format_schema MdbClickhouseCluster#format_schema}
	FormatSchema interface{} `field:"optional" json:"formatSchema" yaml:"formatSchema"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#id MdbClickhouseCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#labels MdbClickhouseCluster#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#maintenance_window MdbClickhouseCluster#maintenance_window}
	MaintenanceWindow *MdbClickhouseClusterMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// ml_model block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#ml_model MdbClickhouseCluster#ml_model}
	MlModel interface{} `field:"optional" json:"mlModel" yaml:"mlModel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#security_group_ids MdbClickhouseCluster#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#service_account_id MdbClickhouseCluster#service_account_id}.
	ServiceAccountId *string `field:"optional" json:"serviceAccountId" yaml:"serviceAccountId"`
	// shard_group block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#shard_group MdbClickhouseCluster#shard_group}
	ShardGroup interface{} `field:"optional" json:"shardGroup" yaml:"shardGroup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#sql_database_management MdbClickhouseCluster#sql_database_management}.
	SqlDatabaseManagement interface{} `field:"optional" json:"sqlDatabaseManagement" yaml:"sqlDatabaseManagement"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#sql_user_management MdbClickhouseCluster#sql_user_management}.
	SqlUserManagement interface{} `field:"optional" json:"sqlUserManagement" yaml:"sqlUserManagement"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#timeouts MdbClickhouseCluster#timeouts}
	Timeouts *MdbClickhouseClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// user block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#user MdbClickhouseCluster#user}
	User interface{} `field:"optional" json:"user" yaml:"user"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#version MdbClickhouseCluster#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
	// zookeeper block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#zookeeper MdbClickhouseCluster#zookeeper}
	Zookeeper *MdbClickhouseClusterZookeeper `field:"optional" json:"zookeeper" yaml:"zookeeper"`
}

