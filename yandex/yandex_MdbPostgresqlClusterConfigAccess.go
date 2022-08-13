// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbPostgresqlClusterConfigAccess struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#data_lens MdbPostgresqlCluster#data_lens}.
	DataLens interface{} `field:"optional" json:"dataLens" yaml:"dataLens"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#serverless MdbPostgresqlCluster#serverless}.
	Serverless interface{} `field:"optional" json:"serverless" yaml:"serverless"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#web_sql MdbPostgresqlCluster#web_sql}.
	WebSql interface{} `field:"optional" json:"webSql" yaml:"webSql"`
}

