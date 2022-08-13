// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataprocClusterClusterConfigSubclusterSpecResources struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#disk_size DataprocCluster#disk_size}.
	DiskSize *float64 `field:"required" json:"diskSize" yaml:"diskSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#resource_preset_id DataprocCluster#resource_preset_id}.
	ResourcePresetId *string `field:"required" json:"resourcePresetId" yaml:"resourcePresetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#disk_type_id DataprocCluster#disk_type_id}.
	DiskTypeId *string `field:"optional" json:"diskTypeId" yaml:"diskTypeId"`
}

