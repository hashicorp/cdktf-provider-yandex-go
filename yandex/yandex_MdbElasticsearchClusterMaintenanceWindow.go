// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbElasticsearchClusterMaintenanceWindow struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_elasticsearch_cluster#type MdbElasticsearchCluster#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_elasticsearch_cluster#day MdbElasticsearchCluster#day}.
	Day *string `field:"optional" json:"day" yaml:"day"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_elasticsearch_cluster#hour MdbElasticsearchCluster#hour}.
	Hour *float64 `field:"optional" json:"hour" yaml:"hour"`
}

