// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexMdbPostgresqlDatabaseExtension struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_postgresql_database#name DataYandexMdbPostgresqlDatabase#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_postgresql_database#version DataYandexMdbPostgresqlDatabase#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

