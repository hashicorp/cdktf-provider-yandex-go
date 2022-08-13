// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexMdbKafkaClusterConfigA struct {
	// kafka block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#kafka DataYandexMdbKafkaCluster#kafka}
	Kafka *DataYandexMdbKafkaClusterConfigKafka `field:"required" json:"kafka" yaml:"kafka"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#version DataYandexMdbKafkaCluster#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#zones DataYandexMdbKafkaCluster#zones}.
	Zones *[]*string `field:"required" json:"zones" yaml:"zones"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#assign_public_ip DataYandexMdbKafkaCluster#assign_public_ip}.
	AssignPublicIp interface{} `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#brokers_count DataYandexMdbKafkaCluster#brokers_count}.
	BrokersCount *float64 `field:"optional" json:"brokersCount" yaml:"brokersCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#schema_registry DataYandexMdbKafkaCluster#schema_registry}.
	SchemaRegistry interface{} `field:"optional" json:"schemaRegistry" yaml:"schemaRegistry"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#unmanaged_topics DataYandexMdbKafkaCluster#unmanaged_topics}.
	UnmanagedTopics interface{} `field:"optional" json:"unmanagedTopics" yaml:"unmanagedTopics"`
	// zookeeper block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#zookeeper DataYandexMdbKafkaCluster#zookeeper}
	Zookeeper *DataYandexMdbKafkaClusterConfigZookeeper `field:"optional" json:"zookeeper" yaml:"zookeeper"`
}

