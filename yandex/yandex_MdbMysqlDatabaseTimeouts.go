// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbMysqlDatabaseTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_database#create MdbMysqlDatabase#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_database#delete MdbMysqlDatabase#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_database#read MdbMysqlDatabase#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_database#update MdbMysqlDatabase#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

