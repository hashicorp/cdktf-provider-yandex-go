// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexMdbMongodbClusterMaintenanceWindow struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#day DataYandexMdbMongodbCluster#day}.
	Day *string `field:"optional" json:"day" yaml:"day"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#hour DataYandexMdbMongodbCluster#hour}.
	Hour *float64 `field:"optional" json:"hour" yaml:"hour"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#type DataYandexMdbMongodbCluster#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

