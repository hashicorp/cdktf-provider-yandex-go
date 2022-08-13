// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexMdbGreenplumClusterPoolerConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_greenplum_cluster#pool_client_idle_timeout DataYandexMdbGreenplumCluster#pool_client_idle_timeout}.
	PoolClientIdleTimeout *float64 `field:"optional" json:"poolClientIdleTimeout" yaml:"poolClientIdleTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_greenplum_cluster#pooling_mode DataYandexMdbGreenplumCluster#pooling_mode}.
	PoolingMode *string `field:"optional" json:"poolingMode" yaml:"poolingMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_greenplum_cluster#pool_size DataYandexMdbGreenplumCluster#pool_size}.
	PoolSize *float64 `field:"optional" json:"poolSize" yaml:"poolSize"`
}

