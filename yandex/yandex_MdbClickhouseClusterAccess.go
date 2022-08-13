// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbClickhouseClusterAccess struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#data_lens MdbClickhouseCluster#data_lens}.
	DataLens interface{} `field:"optional" json:"dataLens" yaml:"dataLens"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#metrika MdbClickhouseCluster#metrika}.
	Metrika interface{} `field:"optional" json:"metrika" yaml:"metrika"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#serverless MdbClickhouseCluster#serverless}.
	Serverless interface{} `field:"optional" json:"serverless" yaml:"serverless"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#web_sql MdbClickhouseCluster#web_sql}.
	WebSql interface{} `field:"optional" json:"webSql" yaml:"webSql"`
}

