// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbPostgresqlClusterDatabase struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#name MdbPostgresqlCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#owner MdbPostgresqlCluster#owner}.
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// extension block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#extension MdbPostgresqlCluster#extension}
	Extension interface{} `field:"optional" json:"extension" yaml:"extension"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#lc_collate MdbPostgresqlCluster#lc_collate}.
	LcCollate *string `field:"optional" json:"lcCollate" yaml:"lcCollate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#lc_type MdbPostgresqlCluster#lc_type}.
	LcType *string `field:"optional" json:"lcType" yaml:"lcType"`
}

