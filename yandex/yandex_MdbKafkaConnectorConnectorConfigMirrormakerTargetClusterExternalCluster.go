// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalCluster struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_connector#bootstrap_servers MdbKafkaConnector#bootstrap_servers}.
	BootstrapServers *string `field:"required" json:"bootstrapServers" yaml:"bootstrapServers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_connector#sasl_mechanism MdbKafkaConnector#sasl_mechanism}.
	SaslMechanism *string `field:"optional" json:"saslMechanism" yaml:"saslMechanism"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_connector#sasl_password MdbKafkaConnector#sasl_password}.
	SaslPassword *string `field:"optional" json:"saslPassword" yaml:"saslPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_connector#sasl_username MdbKafkaConnector#sasl_username}.
	SaslUsername *string `field:"optional" json:"saslUsername" yaml:"saslUsername"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_connector#security_protocol MdbKafkaConnector#security_protocol}.
	SecurityProtocol *string `field:"optional" json:"securityProtocol" yaml:"securityProtocol"`
}

