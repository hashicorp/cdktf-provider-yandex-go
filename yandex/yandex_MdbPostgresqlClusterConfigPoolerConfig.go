// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbPostgresqlClusterConfigPoolerConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#pool_discard MdbPostgresqlCluster#pool_discard}.
	PoolDiscard interface{} `field:"optional" json:"poolDiscard" yaml:"poolDiscard"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#pooling_mode MdbPostgresqlCluster#pooling_mode}.
	PoolingMode *string `field:"optional" json:"poolingMode" yaml:"poolingMode"`
}

