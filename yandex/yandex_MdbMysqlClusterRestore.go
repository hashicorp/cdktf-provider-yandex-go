// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbMysqlClusterRestore struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#backup_id MdbMysqlCluster#backup_id}.
	BackupId *string `field:"required" json:"backupId" yaml:"backupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#time MdbMysqlCluster#time}.
	Time *string `field:"optional" json:"time" yaml:"time"`
}

