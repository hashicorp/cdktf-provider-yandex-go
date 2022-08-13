// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbClickhouseClusterUser struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#name MdbClickhouseCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#password MdbClickhouseCluster#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// permission block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#permission MdbClickhouseCluster#permission}
	Permission interface{} `field:"optional" json:"permission" yaml:"permission"`
	// quota block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#quota MdbClickhouseCluster#quota}
	Quota interface{} `field:"optional" json:"quota" yaml:"quota"`
	// settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#settings MdbClickhouseCluster#settings}
	Settings *MdbClickhouseClusterUserSettings `field:"optional" json:"settings" yaml:"settings"`
}

