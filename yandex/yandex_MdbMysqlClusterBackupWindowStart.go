// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbMysqlClusterBackupWindowStart struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#hours MdbMysqlCluster#hours}.
	Hours *float64 `field:"optional" json:"hours" yaml:"hours"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#minutes MdbMysqlCluster#minutes}.
	Minutes *float64 `field:"optional" json:"minutes" yaml:"minutes"`
}

