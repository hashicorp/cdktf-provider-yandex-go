// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbPostgresqlClusterConfigBackupWindowStart struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#hours MdbPostgresqlCluster#hours}.
	Hours *float64 `field:"optional" json:"hours" yaml:"hours"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#minutes MdbPostgresqlCluster#minutes}.
	Minutes *float64 `field:"optional" json:"minutes" yaml:"minutes"`
}

