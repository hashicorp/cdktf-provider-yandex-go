// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbPostgresqlClusterConfigResources struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#disk_size MdbPostgresqlCluster#disk_size}.
	DiskSize *float64 `field:"required" json:"diskSize" yaml:"diskSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#resource_preset_id MdbPostgresqlCluster#resource_preset_id}.
	ResourcePresetId *string `field:"required" json:"resourcePresetId" yaml:"resourcePresetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#disk_type_id MdbPostgresqlCluster#disk_type_id}.
	DiskTypeId *string `field:"optional" json:"diskTypeId" yaml:"diskTypeId"`
}

