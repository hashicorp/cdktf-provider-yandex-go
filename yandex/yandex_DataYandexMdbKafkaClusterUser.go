// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexMdbKafkaClusterUser struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#name DataYandexMdbKafkaCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#password DataYandexMdbKafkaCluster#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// permission block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#permission DataYandexMdbKafkaCluster#permission}
	Permission interface{} `field:"optional" json:"permission" yaml:"permission"`
}

