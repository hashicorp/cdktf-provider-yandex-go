// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbKafkaClusterTopic struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#name MdbKafkaCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#partitions MdbKafkaCluster#partitions}.
	Partitions *float64 `field:"required" json:"partitions" yaml:"partitions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#replication_factor MdbKafkaCluster#replication_factor}.
	ReplicationFactor *float64 `field:"required" json:"replicationFactor" yaml:"replicationFactor"`
	// topic_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#topic_config MdbKafkaCluster#topic_config}
	TopicConfig *MdbKafkaClusterTopicTopicConfig `field:"optional" json:"topicConfig" yaml:"topicConfig"`
}

