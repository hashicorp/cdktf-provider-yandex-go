// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexMdbMongodbClusterResources struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#disk_size DataYandexMdbMongodbCluster#disk_size}.
	DiskSize *float64 `field:"optional" json:"diskSize" yaml:"diskSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#disk_type_id DataYandexMdbMongodbCluster#disk_type_id}.
	DiskTypeId *string `field:"optional" json:"diskTypeId" yaml:"diskTypeId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#resource_preset_id DataYandexMdbMongodbCluster#resource_preset_id}.
	ResourcePresetId *string `field:"optional" json:"resourcePresetId" yaml:"resourcePresetId"`
}

