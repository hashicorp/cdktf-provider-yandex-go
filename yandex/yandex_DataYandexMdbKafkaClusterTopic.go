// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexMdbKafkaClusterTopic struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#cluster_id DataYandexMdbKafkaCluster#cluster_id}.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#name DataYandexMdbKafkaCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#partitions DataYandexMdbKafkaCluster#partitions}.
	Partitions *float64 `field:"required" json:"partitions" yaml:"partitions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#replication_factor DataYandexMdbKafkaCluster#replication_factor}.
	ReplicationFactor *float64 `field:"required" json:"replicationFactor" yaml:"replicationFactor"`
	// topic_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#topic_config DataYandexMdbKafkaCluster#topic_config}
	TopicConfig *DataYandexMdbKafkaClusterTopicTopicConfig `field:"optional" json:"topicConfig" yaml:"topicConfig"`
}

