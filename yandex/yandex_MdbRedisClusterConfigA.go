// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbRedisClusterConfigA struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_redis_cluster#password MdbRedisCluster#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_redis_cluster#version MdbRedisCluster#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_redis_cluster#client_output_buffer_limit_normal MdbRedisCluster#client_output_buffer_limit_normal}.
	ClientOutputBufferLimitNormal *string `field:"optional" json:"clientOutputBufferLimitNormal" yaml:"clientOutputBufferLimitNormal"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_redis_cluster#client_output_buffer_limit_pubsub MdbRedisCluster#client_output_buffer_limit_pubsub}.
	ClientOutputBufferLimitPubsub *string `field:"optional" json:"clientOutputBufferLimitPubsub" yaml:"clientOutputBufferLimitPubsub"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_redis_cluster#databases MdbRedisCluster#databases}.
	Databases *float64 `field:"optional" json:"databases" yaml:"databases"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_redis_cluster#maxmemory_policy MdbRedisCluster#maxmemory_policy}.
	MaxmemoryPolicy *string `field:"optional" json:"maxmemoryPolicy" yaml:"maxmemoryPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_redis_cluster#notify_keyspace_events MdbRedisCluster#notify_keyspace_events}.
	NotifyKeyspaceEvents *string `field:"optional" json:"notifyKeyspaceEvents" yaml:"notifyKeyspaceEvents"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_redis_cluster#slowlog_log_slower_than MdbRedisCluster#slowlog_log_slower_than}.
	SlowlogLogSlowerThan *float64 `field:"optional" json:"slowlogLogSlowerThan" yaml:"slowlogLogSlowerThan"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_redis_cluster#slowlog_max_len MdbRedisCluster#slowlog_max_len}.
	SlowlogMaxLen *float64 `field:"optional" json:"slowlogMaxLen" yaml:"slowlogMaxLen"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_redis_cluster#timeout MdbRedisCluster#timeout}.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

