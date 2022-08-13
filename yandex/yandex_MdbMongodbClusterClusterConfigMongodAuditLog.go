// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbMongodbClusterClusterConfigMongodAuditLog struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#filter MdbMongodbCluster#filter}.
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#runtime_configuration MdbMongodbCluster#runtime_configuration}.
	RuntimeConfiguration interface{} `field:"optional" json:"runtimeConfiguration" yaml:"runtimeConfiguration"`
}

