// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbPostgresqlDatabaseTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_database#create MdbPostgresqlDatabase#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_database#delete MdbPostgresqlDatabase#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_database#read MdbPostgresqlDatabase#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_database#update MdbPostgresqlDatabase#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

