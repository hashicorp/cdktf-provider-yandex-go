// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexMdbKafkaClusterConfigKafkaKafkaConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#auto_create_topics_enable DataYandexMdbKafkaCluster#auto_create_topics_enable}.
	AutoCreateTopicsEnable interface{} `field:"optional" json:"autoCreateTopicsEnable" yaml:"autoCreateTopicsEnable"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#compression_type DataYandexMdbKafkaCluster#compression_type}.
	CompressionType *string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#default_replication_factor DataYandexMdbKafkaCluster#default_replication_factor}.
	DefaultReplicationFactor *string `field:"optional" json:"defaultReplicationFactor" yaml:"defaultReplicationFactor"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#log_flush_interval_messages DataYandexMdbKafkaCluster#log_flush_interval_messages}.
	LogFlushIntervalMessages *string `field:"optional" json:"logFlushIntervalMessages" yaml:"logFlushIntervalMessages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#log_flush_interval_ms DataYandexMdbKafkaCluster#log_flush_interval_ms}.
	LogFlushIntervalMs *string `field:"optional" json:"logFlushIntervalMs" yaml:"logFlushIntervalMs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#log_flush_scheduler_interval_ms DataYandexMdbKafkaCluster#log_flush_scheduler_interval_ms}.
	LogFlushSchedulerIntervalMs *string `field:"optional" json:"logFlushSchedulerIntervalMs" yaml:"logFlushSchedulerIntervalMs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#log_preallocate DataYandexMdbKafkaCluster#log_preallocate}.
	LogPreallocate interface{} `field:"optional" json:"logPreallocate" yaml:"logPreallocate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#log_retention_bytes DataYandexMdbKafkaCluster#log_retention_bytes}.
	LogRetentionBytes *string `field:"optional" json:"logRetentionBytes" yaml:"logRetentionBytes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#log_retention_hours DataYandexMdbKafkaCluster#log_retention_hours}.
	LogRetentionHours *string `field:"optional" json:"logRetentionHours" yaml:"logRetentionHours"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#log_retention_minutes DataYandexMdbKafkaCluster#log_retention_minutes}.
	LogRetentionMinutes *string `field:"optional" json:"logRetentionMinutes" yaml:"logRetentionMinutes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#log_retention_ms DataYandexMdbKafkaCluster#log_retention_ms}.
	LogRetentionMs *string `field:"optional" json:"logRetentionMs" yaml:"logRetentionMs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#log_segment_bytes DataYandexMdbKafkaCluster#log_segment_bytes}.
	LogSegmentBytes *string `field:"optional" json:"logSegmentBytes" yaml:"logSegmentBytes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#num_partitions DataYandexMdbKafkaCluster#num_partitions}.
	NumPartitions *string `field:"optional" json:"numPartitions" yaml:"numPartitions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#socket_receive_buffer_bytes DataYandexMdbKafkaCluster#socket_receive_buffer_bytes}.
	SocketReceiveBufferBytes *string `field:"optional" json:"socketReceiveBufferBytes" yaml:"socketReceiveBufferBytes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#socket_send_buffer_bytes DataYandexMdbKafkaCluster#socket_send_buffer_bytes}.
	SocketSendBufferBytes *string `field:"optional" json:"socketSendBufferBytes" yaml:"socketSendBufferBytes"`
}

