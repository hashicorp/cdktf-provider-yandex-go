// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbRedisClusterHost struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_redis_cluster#zone MdbRedisCluster#zone}.
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_redis_cluster#assign_public_ip MdbRedisCluster#assign_public_ip}.
	AssignPublicIp interface{} `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_redis_cluster#replica_priority MdbRedisCluster#replica_priority}.
	ReplicaPriority *float64 `field:"optional" json:"replicaPriority" yaml:"replicaPriority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_redis_cluster#shard_name MdbRedisCluster#shard_name}.
	ShardName *string `field:"optional" json:"shardName" yaml:"shardName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_redis_cluster#subnet_id MdbRedisCluster#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

