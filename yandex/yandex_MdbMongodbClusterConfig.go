// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbMongodbClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#cluster_config MdbMongodbCluster#cluster_config}
	ClusterConfig *MdbMongodbClusterClusterConfig `field:"required" json:"clusterConfig" yaml:"clusterConfig"`
	// database block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#database MdbMongodbCluster#database}
	Database interface{} `field:"required" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#environment MdbMongodbCluster#environment}.
	Environment *string `field:"required" json:"environment" yaml:"environment"`
	// host block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#host MdbMongodbCluster#host}
	Host interface{} `field:"required" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#name MdbMongodbCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#network_id MdbMongodbCluster#network_id}.
	NetworkId *string `field:"required" json:"networkId" yaml:"networkId"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#resources MdbMongodbCluster#resources}
	Resources *MdbMongodbClusterResources `field:"required" json:"resources" yaml:"resources"`
	// user block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#user MdbMongodbCluster#user}
	User interface{} `field:"required" json:"user" yaml:"user"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#cluster_id MdbMongodbCluster#cluster_id}.
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#deletion_protection MdbMongodbCluster#deletion_protection}.
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#description MdbMongodbCluster#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#folder_id MdbMongodbCluster#folder_id}.
	FolderId *string `field:"optional" json:"folderId" yaml:"folderId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#id MdbMongodbCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#labels MdbMongodbCluster#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#maintenance_window MdbMongodbCluster#maintenance_window}
	MaintenanceWindow *MdbMongodbClusterMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#security_group_ids MdbMongodbCluster#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#timeouts MdbMongodbCluster#timeouts}
	Timeouts *MdbMongodbClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

