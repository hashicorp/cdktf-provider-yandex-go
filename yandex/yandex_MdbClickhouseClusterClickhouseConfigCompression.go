// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbClickhouseClusterClickhouseConfigCompression struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#method MdbClickhouseCluster#method}.
	Method *string `field:"required" json:"method" yaml:"method"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#min_part_size MdbClickhouseCluster#min_part_size}.
	MinPartSize *float64 `field:"required" json:"minPartSize" yaml:"minPartSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#min_part_size_ratio MdbClickhouseCluster#min_part_size_ratio}.
	MinPartSizeRatio *float64 `field:"required" json:"minPartSizeRatio" yaml:"minPartSizeRatio"`
}

