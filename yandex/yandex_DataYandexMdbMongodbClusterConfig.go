// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexMdbMongodbClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#cluster_config DataYandexMdbMongodbCluster#cluster_config}
	ClusterConfig *DataYandexMdbMongodbClusterClusterConfig `field:"optional" json:"clusterConfig" yaml:"clusterConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#cluster_id DataYandexMdbMongodbCluster#cluster_id}.
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#created_at DataYandexMdbMongodbCluster#created_at}.
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// database block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#database DataYandexMdbMongodbCluster#database}
	Database interface{} `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#deletion_protection DataYandexMdbMongodbCluster#deletion_protection}.
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#description DataYandexMdbMongodbCluster#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#environment DataYandexMdbMongodbCluster#environment}.
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#folder_id DataYandexMdbMongodbCluster#folder_id}.
	FolderId *string `field:"optional" json:"folderId" yaml:"folderId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#health DataYandexMdbMongodbCluster#health}.
	Health *string `field:"optional" json:"health" yaml:"health"`
	// host block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#host DataYandexMdbMongodbCluster#host}
	Host interface{} `field:"optional" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#id DataYandexMdbMongodbCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#labels DataYandexMdbMongodbCluster#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#maintenance_window DataYandexMdbMongodbCluster#maintenance_window}
	MaintenanceWindow *DataYandexMdbMongodbClusterMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#name DataYandexMdbMongodbCluster#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#network_id DataYandexMdbMongodbCluster#network_id}.
	NetworkId *string `field:"optional" json:"networkId" yaml:"networkId"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#resources DataYandexMdbMongodbCluster#resources}
	Resources *DataYandexMdbMongodbClusterResources `field:"optional" json:"resources" yaml:"resources"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#security_group_ids DataYandexMdbMongodbCluster#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#sharded DataYandexMdbMongodbCluster#sharded}.
	Sharded interface{} `field:"optional" json:"sharded" yaml:"sharded"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#status DataYandexMdbMongodbCluster#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// user block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#user DataYandexMdbMongodbCluster#user}
	User interface{} `field:"optional" json:"user" yaml:"user"`
}

