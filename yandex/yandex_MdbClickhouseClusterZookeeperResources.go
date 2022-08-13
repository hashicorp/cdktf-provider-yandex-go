// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbClickhouseClusterZookeeperResources struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#disk_size MdbClickhouseCluster#disk_size}.
	DiskSize *float64 `field:"optional" json:"diskSize" yaml:"diskSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#disk_type_id MdbClickhouseCluster#disk_type_id}.
	DiskTypeId *string `field:"optional" json:"diskTypeId" yaml:"diskTypeId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#resource_preset_id MdbClickhouseCluster#resource_preset_id}.
	ResourcePresetId *string `field:"optional" json:"resourcePresetId" yaml:"resourcePresetId"`
}

