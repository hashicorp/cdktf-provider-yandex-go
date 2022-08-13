// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbSqlserverClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_sqlserver_cluster#create MdbSqlserverCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_sqlserver_cluster#delete MdbSqlserverCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_sqlserver_cluster#update MdbSqlserverCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

