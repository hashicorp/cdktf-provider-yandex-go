// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbGreenplumClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#create MdbGreenplumCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#delete MdbGreenplumCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster#update MdbGreenplumCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

