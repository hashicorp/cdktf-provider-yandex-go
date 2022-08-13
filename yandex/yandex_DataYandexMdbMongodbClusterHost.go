// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexMdbMongodbClusterHost struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#assign_public_ip DataYandexMdbMongodbCluster#assign_public_ip}.
	AssignPublicIp interface{} `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#health DataYandexMdbMongodbCluster#health}.
	Health *string `field:"optional" json:"health" yaml:"health"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#name DataYandexMdbMongodbCluster#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#role DataYandexMdbMongodbCluster#role}.
	Role *string `field:"optional" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#shard_name DataYandexMdbMongodbCluster#shard_name}.
	ShardName *string `field:"optional" json:"shardName" yaml:"shardName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#subnet_id DataYandexMdbMongodbCluster#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#type DataYandexMdbMongodbCluster#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#zone_id DataYandexMdbMongodbCluster#zone_id}.
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

