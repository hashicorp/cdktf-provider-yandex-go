// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbClickhouseClusterClickhouseConfigGraphiteRollup struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#name MdbClickhouseCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// pattern block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#pattern MdbClickhouseCluster#pattern}
	Pattern interface{} `field:"optional" json:"pattern" yaml:"pattern"`
}

