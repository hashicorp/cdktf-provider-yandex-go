// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbKafkaClusterConfigA struct {
	// kafka block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#kafka MdbKafkaCluster#kafka}
	Kafka *MdbKafkaClusterConfigKafka `field:"required" json:"kafka" yaml:"kafka"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#version MdbKafkaCluster#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#zones MdbKafkaCluster#zones}.
	Zones *[]*string `field:"required" json:"zones" yaml:"zones"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#assign_public_ip MdbKafkaCluster#assign_public_ip}.
	AssignPublicIp interface{} `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#brokers_count MdbKafkaCluster#brokers_count}.
	BrokersCount *float64 `field:"optional" json:"brokersCount" yaml:"brokersCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#schema_registry MdbKafkaCluster#schema_registry}.
	SchemaRegistry interface{} `field:"optional" json:"schemaRegistry" yaml:"schemaRegistry"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#unmanaged_topics MdbKafkaCluster#unmanaged_topics}.
	UnmanagedTopics interface{} `field:"optional" json:"unmanagedTopics" yaml:"unmanagedTopics"`
	// zookeeper block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#zookeeper MdbKafkaCluster#zookeeper}
	Zookeeper *MdbKafkaClusterConfigZookeeper `field:"optional" json:"zookeeper" yaml:"zookeeper"`
}

