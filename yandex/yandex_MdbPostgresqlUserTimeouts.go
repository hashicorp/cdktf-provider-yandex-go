// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbPostgresqlUserTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_user#create MdbPostgresqlUser#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_user#delete MdbPostgresqlUser#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_user#read MdbPostgresqlUser#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_user#update MdbPostgresqlUser#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

