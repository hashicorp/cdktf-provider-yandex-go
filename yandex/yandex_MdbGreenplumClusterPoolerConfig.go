// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbGreenplumClusterPoolerConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#pool_client_idle_timeout MdbGreenplumCluster#pool_client_idle_timeout}.
	PoolClientIdleTimeout *float64 `field:"optional" json:"poolClientIdleTimeout" yaml:"poolClientIdleTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#pooling_mode MdbGreenplumCluster#pooling_mode}.
	PoolingMode *string `field:"optional" json:"poolingMode" yaml:"poolingMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#pool_size MdbGreenplumCluster#pool_size}.
	PoolSize *float64 `field:"optional" json:"poolSize" yaml:"poolSize"`
}

