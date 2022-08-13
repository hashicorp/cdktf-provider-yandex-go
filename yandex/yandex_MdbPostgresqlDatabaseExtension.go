// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbPostgresqlDatabaseExtension struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_database#name MdbPostgresqlDatabase#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_database#version MdbPostgresqlDatabase#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

