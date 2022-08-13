// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbRedisClusterMaintenanceWindow struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_redis_cluster#type MdbRedisCluster#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_redis_cluster#day MdbRedisCluster#day}.
	Day *string `field:"optional" json:"day" yaml:"day"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_redis_cluster#hour MdbRedisCluster#hour}.
	Hour *float64 `field:"optional" json:"hour" yaml:"hour"`
}

