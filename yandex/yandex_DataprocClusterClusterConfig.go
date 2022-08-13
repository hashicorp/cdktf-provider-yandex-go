// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataprocClusterClusterConfig struct {
	// subcluster_spec block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#subcluster_spec DataprocCluster#subcluster_spec}
	SubclusterSpec interface{} `field:"required" json:"subclusterSpec" yaml:"subclusterSpec"`
	// hadoop block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#hadoop DataprocCluster#hadoop}
	Hadoop *DataprocClusterClusterConfigHadoop `field:"optional" json:"hadoop" yaml:"hadoop"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#version_id DataprocCluster#version_id}.
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

