// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbSqlserverClusterUserPermission struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_sqlserver_cluster#database_name MdbSqlserverCluster#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_sqlserver_cluster#roles MdbSqlserverCluster#roles}.
	Roles *[]*string `field:"optional" json:"roles" yaml:"roles"`
}

