// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbKafkaClusterUser struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#name MdbKafkaCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#password MdbKafkaCluster#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// permission block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#permission MdbKafkaCluster#permission}
	Permission interface{} `field:"optional" json:"permission" yaml:"permission"`
}

