// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbKafkaClusterConfigKafka struct {
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#resources MdbKafkaCluster#resources}
	Resources *MdbKafkaClusterConfigKafkaResources `field:"required" json:"resources" yaml:"resources"`
	// kafka_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#kafka_config MdbKafkaCluster#kafka_config}
	KafkaConfig *MdbKafkaClusterConfigKafkaKafkaConfig `field:"optional" json:"kafkaConfig" yaml:"kafkaConfig"`
}

