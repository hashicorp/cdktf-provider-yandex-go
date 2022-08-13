// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbPostgresqlClusterMaintenanceWindow struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#type MdbPostgresqlCluster#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#day MdbPostgresqlCluster#day}.
	Day *string `field:"optional" json:"day" yaml:"day"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#hour MdbPostgresqlCluster#hour}.
	Hour *float64 `field:"optional" json:"hour" yaml:"hour"`
}

