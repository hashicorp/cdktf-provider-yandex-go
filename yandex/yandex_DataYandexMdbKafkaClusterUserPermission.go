// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexMdbKafkaClusterUserPermission struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#role DataYandexMdbKafkaCluster#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#topic_name DataYandexMdbKafkaCluster#topic_name}.
	TopicName *string `field:"required" json:"topicName" yaml:"topicName"`
}

