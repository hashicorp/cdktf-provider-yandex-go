// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbMysqlClusterAccess struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#data_lens MdbMysqlCluster#data_lens}.
	DataLens interface{} `field:"optional" json:"dataLens" yaml:"dataLens"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#web_sql MdbMysqlCluster#web_sql}.
	WebSql interface{} `field:"optional" json:"webSql" yaml:"webSql"`
}

