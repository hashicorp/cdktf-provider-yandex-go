// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbMysqlClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#environment MdbMysqlCluster#environment}.
	Environment *string `field:"required" json:"environment" yaml:"environment"`
	// host block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#host MdbMysqlCluster#host}
	Host interface{} `field:"required" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#name MdbMysqlCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#network_id MdbMysqlCluster#network_id}.
	NetworkId *string `field:"required" json:"networkId" yaml:"networkId"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#resources MdbMysqlCluster#resources}
	Resources *MdbMysqlClusterResources `field:"required" json:"resources" yaml:"resources"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#version MdbMysqlCluster#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// access block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#access MdbMysqlCluster#access}
	Access *MdbMysqlClusterAccess `field:"optional" json:"access" yaml:"access"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#allow_regeneration_host MdbMysqlCluster#allow_regeneration_host}.
	AllowRegenerationHost interface{} `field:"optional" json:"allowRegenerationHost" yaml:"allowRegenerationHost"`
	// backup_window_start block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#backup_window_start MdbMysqlCluster#backup_window_start}
	BackupWindowStart *MdbMysqlClusterBackupWindowStart `field:"optional" json:"backupWindowStart" yaml:"backupWindowStart"`
	// database block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#database MdbMysqlCluster#database}
	Database interface{} `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#deletion_protection MdbMysqlCluster#deletion_protection}.
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#description MdbMysqlCluster#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#folder_id MdbMysqlCluster#folder_id}.
	FolderId *string `field:"optional" json:"folderId" yaml:"folderId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#host_group_ids MdbMysqlCluster#host_group_ids}.
	HostGroupIds *[]*string `field:"optional" json:"hostGroupIds" yaml:"hostGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#id MdbMysqlCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#labels MdbMysqlCluster#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#maintenance_window MdbMysqlCluster#maintenance_window}
	MaintenanceWindow *MdbMysqlClusterMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#mysql_config MdbMysqlCluster#mysql_config}.
	MysqlConfig *map[string]*string `field:"optional" json:"mysqlConfig" yaml:"mysqlConfig"`
	// performance_diagnostics block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#performance_diagnostics MdbMysqlCluster#performance_diagnostics}
	PerformanceDiagnostics *MdbMysqlClusterPerformanceDiagnostics `field:"optional" json:"performanceDiagnostics" yaml:"performanceDiagnostics"`
	// restore block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#restore MdbMysqlCluster#restore}
	Restore *MdbMysqlClusterRestore `field:"optional" json:"restore" yaml:"restore"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#security_group_ids MdbMysqlCluster#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#timeouts MdbMysqlCluster#timeouts}
	Timeouts *MdbMysqlClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// user block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#user MdbMysqlCluster#user}
	User interface{} `field:"optional" json:"user" yaml:"user"`
}

