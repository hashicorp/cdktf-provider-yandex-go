// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbElasticsearchClusterConfigDataNode struct {
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_elasticsearch_cluster#resources MdbElasticsearchCluster#resources}
	Resources *MdbElasticsearchClusterConfigDataNodeResources `field:"required" json:"resources" yaml:"resources"`
}

