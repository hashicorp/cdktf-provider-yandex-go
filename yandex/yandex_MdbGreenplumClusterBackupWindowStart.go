// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbGreenplumClusterBackupWindowStart struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#hours MdbGreenplumCluster#hours}.
	Hours *float64 `field:"optional" json:"hours" yaml:"hours"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#minutes MdbGreenplumCluster#minutes}.
	Minutes *float64 `field:"optional" json:"minutes" yaml:"minutes"`
}

