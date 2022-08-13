// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbPostgresqlClusterRestore struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#backup_id MdbPostgresqlCluster#backup_id}.
	BackupId *string `field:"required" json:"backupId" yaml:"backupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#time MdbPostgresqlCluster#time}.
	Time *string `field:"optional" json:"time" yaml:"time"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#time_inclusive MdbPostgresqlCluster#time_inclusive}.
	TimeInclusive interface{} `field:"optional" json:"timeInclusive" yaml:"timeInclusive"`
}

