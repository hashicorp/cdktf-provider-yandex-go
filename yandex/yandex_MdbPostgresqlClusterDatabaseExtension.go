// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbPostgresqlClusterDatabaseExtension struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#name MdbPostgresqlCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#version MdbPostgresqlCluster#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

