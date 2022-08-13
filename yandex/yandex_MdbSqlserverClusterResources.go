// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbSqlserverClusterResources struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_sqlserver_cluster#disk_size MdbSqlserverCluster#disk_size}.
	DiskSize *float64 `field:"required" json:"diskSize" yaml:"diskSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_sqlserver_cluster#disk_type_id MdbSqlserverCluster#disk_type_id}.
	DiskTypeId *string `field:"required" json:"diskTypeId" yaml:"diskTypeId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_sqlserver_cluster#resource_preset_id MdbSqlserverCluster#resource_preset_id}.
	ResourcePresetId *string `field:"required" json:"resourcePresetId" yaml:"resourcePresetId"`
}

