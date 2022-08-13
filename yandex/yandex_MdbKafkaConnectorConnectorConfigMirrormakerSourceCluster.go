// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbKafkaConnectorConnectorConfigMirrormakerSourceCluster struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_connector#alias MdbKafkaConnector#alias}.
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// external_cluster block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_connector#external_cluster MdbKafkaConnector#external_cluster}
	ExternalCluster interface{} `field:"optional" json:"externalCluster" yaml:"externalCluster"`
	// this_cluster block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_connector#this_cluster MdbKafkaConnector#this_cluster}
	ThisCluster interface{} `field:"optional" json:"thisCluster" yaml:"thisCluster"`
}

