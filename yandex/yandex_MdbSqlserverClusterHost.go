// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbSqlserverClusterHost struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_sqlserver_cluster#zone MdbSqlserverCluster#zone}.
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_sqlserver_cluster#assign_public_ip MdbSqlserverCluster#assign_public_ip}.
	AssignPublicIp interface{} `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_sqlserver_cluster#subnet_id MdbSqlserverCluster#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

