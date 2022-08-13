// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbKafkaConnectorConnectorConfigMirrormaker struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_connector#replication_factor MdbKafkaConnector#replication_factor}.
	ReplicationFactor *float64 `field:"required" json:"replicationFactor" yaml:"replicationFactor"`
	// source_cluster block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_connector#source_cluster MdbKafkaConnector#source_cluster}
	SourceCluster *MdbKafkaConnectorConnectorConfigMirrormakerSourceCluster `field:"required" json:"sourceCluster" yaml:"sourceCluster"`
	// target_cluster block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_connector#target_cluster MdbKafkaConnector#target_cluster}
	TargetCluster *MdbKafkaConnectorConnectorConfigMirrormakerTargetCluster `field:"required" json:"targetCluster" yaml:"targetCluster"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_connector#topics MdbKafkaConnector#topics}.
	Topics *string `field:"required" json:"topics" yaml:"topics"`
}

