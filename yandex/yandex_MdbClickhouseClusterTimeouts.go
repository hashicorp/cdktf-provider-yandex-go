// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbClickhouseClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#create MdbClickhouseCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#delete MdbClickhouseCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#update MdbClickhouseCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

