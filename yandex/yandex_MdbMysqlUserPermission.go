// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbMysqlUserPermission struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_user#database_name MdbMysqlUser#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_user#roles MdbMysqlUser#roles}.
	Roles *[]*string `field:"optional" json:"roles" yaml:"roles"`
}

