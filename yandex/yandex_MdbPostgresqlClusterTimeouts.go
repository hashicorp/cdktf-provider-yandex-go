// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbPostgresqlClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#create MdbPostgresqlCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#delete MdbPostgresqlCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#update MdbPostgresqlCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

