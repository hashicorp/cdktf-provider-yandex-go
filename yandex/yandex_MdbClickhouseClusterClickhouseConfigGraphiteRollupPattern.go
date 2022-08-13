// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbClickhouseClusterClickhouseConfigGraphiteRollupPattern struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#function MdbClickhouseCluster#function}.
	Function *string `field:"required" json:"function" yaml:"function"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#regexp MdbClickhouseCluster#regexp}.
	Regexp *string `field:"optional" json:"regexp" yaml:"regexp"`
	// retention block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#retention MdbClickhouseCluster#retention}
	Retention interface{} `field:"optional" json:"retention" yaml:"retention"`
}

