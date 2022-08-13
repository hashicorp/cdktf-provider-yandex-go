// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbSqlserverClusterBackupWindowStart struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_sqlserver_cluster#hours MdbSqlserverCluster#hours}.
	Hours *float64 `field:"optional" json:"hours" yaml:"hours"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_sqlserver_cluster#minutes MdbSqlserverCluster#minutes}.
	Minutes *float64 `field:"optional" json:"minutes" yaml:"minutes"`
}

