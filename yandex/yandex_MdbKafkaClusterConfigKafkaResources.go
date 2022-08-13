// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbKafkaClusterConfigKafkaResources struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#disk_size MdbKafkaCluster#disk_size}.
	DiskSize *float64 `field:"required" json:"diskSize" yaml:"diskSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#disk_type_id MdbKafkaCluster#disk_type_id}.
	DiskTypeId *string `field:"required" json:"diskTypeId" yaml:"diskTypeId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#resource_preset_id MdbKafkaCluster#resource_preset_id}.
	ResourcePresetId *string `field:"required" json:"resourcePresetId" yaml:"resourcePresetId"`
}

