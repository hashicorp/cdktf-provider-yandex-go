// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbPostgresqlClusterUserPermission struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#database_name MdbPostgresqlCluster#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
}

