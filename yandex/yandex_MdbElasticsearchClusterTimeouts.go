// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbElasticsearchClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_elasticsearch_cluster#create MdbElasticsearchCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_elasticsearch_cluster#delete MdbElasticsearchCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_elasticsearch_cluster#update MdbElasticsearchCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

