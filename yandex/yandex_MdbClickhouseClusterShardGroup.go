// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbClickhouseClusterShardGroup struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#name MdbClickhouseCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#shard_names MdbClickhouseCluster#shard_names}.
	ShardNames *[]*string `field:"required" json:"shardNames" yaml:"shardNames"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#description MdbClickhouseCluster#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

