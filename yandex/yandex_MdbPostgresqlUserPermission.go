// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbPostgresqlUserPermission struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_user#database_name MdbPostgresqlUser#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
}

