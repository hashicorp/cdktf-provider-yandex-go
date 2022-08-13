// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbClickhouseClusterClickhouse struct {
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#resources MdbClickhouseCluster#resources}
	Resources *MdbClickhouseClusterClickhouseResources `field:"required" json:"resources" yaml:"resources"`
	// config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#config MdbClickhouseCluster#config}
	Config *MdbClickhouseClusterClickhouseConfig `field:"optional" json:"config" yaml:"config"`
}

