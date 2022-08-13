// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbRedisClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_redis_cluster#create MdbRedisCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_redis_cluster#delete MdbRedisCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_redis_cluster#update MdbRedisCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

