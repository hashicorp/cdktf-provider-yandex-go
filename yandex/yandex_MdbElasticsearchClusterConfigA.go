// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbElasticsearchClusterConfigA struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_elasticsearch_cluster#admin_password MdbElasticsearchCluster#admin_password}.
	AdminPassword *string `field:"required" json:"adminPassword" yaml:"adminPassword"`
	// data_node block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_elasticsearch_cluster#data_node MdbElasticsearchCluster#data_node}
	DataNode *MdbElasticsearchClusterConfigDataNode `field:"required" json:"dataNode" yaml:"dataNode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_elasticsearch_cluster#edition MdbElasticsearchCluster#edition}.
	Edition *string `field:"optional" json:"edition" yaml:"edition"`
	// master_node block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_elasticsearch_cluster#master_node MdbElasticsearchCluster#master_node}
	MasterNode *MdbElasticsearchClusterConfigMasterNode `field:"optional" json:"masterNode" yaml:"masterNode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_elasticsearch_cluster#plugins MdbElasticsearchCluster#plugins}.
	Plugins *[]*string `field:"optional" json:"plugins" yaml:"plugins"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_elasticsearch_cluster#version MdbElasticsearchCluster#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

