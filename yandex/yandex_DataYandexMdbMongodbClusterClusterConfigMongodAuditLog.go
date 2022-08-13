// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexMdbMongodbClusterClusterConfigMongodAuditLog struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#filter DataYandexMdbMongodbCluster#filter}.
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#runtime_configuration DataYandexMdbMongodbCluster#runtime_configuration}.
	RuntimeConfiguration interface{} `field:"optional" json:"runtimeConfiguration" yaml:"runtimeConfiguration"`
}

