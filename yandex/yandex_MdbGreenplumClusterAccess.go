// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbGreenplumClusterAccess struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#data_lens MdbGreenplumCluster#data_lens}.
	DataLens interface{} `field:"optional" json:"dataLens" yaml:"dataLens"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#web_sql MdbGreenplumCluster#web_sql}.
	WebSql interface{} `field:"optional" json:"webSql" yaml:"webSql"`
}

