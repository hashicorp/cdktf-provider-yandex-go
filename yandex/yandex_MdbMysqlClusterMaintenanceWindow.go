// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbMysqlClusterMaintenanceWindow struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#type MdbMysqlCluster#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#day MdbMysqlCluster#day}.
	Day *string `field:"optional" json:"day" yaml:"day"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#hour MdbMysqlCluster#hour}.
	Hour *float64 `field:"optional" json:"hour" yaml:"hour"`
}

