// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexMdbKafkaClusterConfigKafka struct {
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#resources DataYandexMdbKafkaCluster#resources}
	Resources *DataYandexMdbKafkaClusterConfigKafkaResources `field:"required" json:"resources" yaml:"resources"`
	// kafka_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#kafka_config DataYandexMdbKafkaCluster#kafka_config}
	KafkaConfig *DataYandexMdbKafkaClusterConfigKafkaKafkaConfig `field:"optional" json:"kafkaConfig" yaml:"kafkaConfig"`
}

