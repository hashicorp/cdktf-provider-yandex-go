// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbKafkaTopicTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_topic#create MdbKafkaTopic#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_topic#delete MdbKafkaTopic#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_topic#read MdbKafkaTopic#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_topic#update MdbKafkaTopic#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

