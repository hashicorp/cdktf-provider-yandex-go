// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbClickhouseClusterClickhouseConfigKafka struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#sasl_mechanism MdbClickhouseCluster#sasl_mechanism}.
	SaslMechanism *string `field:"optional" json:"saslMechanism" yaml:"saslMechanism"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#sasl_password MdbClickhouseCluster#sasl_password}.
	SaslPassword *string `field:"optional" json:"saslPassword" yaml:"saslPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#sasl_username MdbClickhouseCluster#sasl_username}.
	SaslUsername *string `field:"optional" json:"saslUsername" yaml:"saslUsername"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#security_protocol MdbClickhouseCluster#security_protocol}.
	SecurityProtocol *string `field:"optional" json:"securityProtocol" yaml:"securityProtocol"`
}

