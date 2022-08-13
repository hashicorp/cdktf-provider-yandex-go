// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexMdbKafkaClusterTopicTopicConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#cleanup_policy DataYandexMdbKafkaCluster#cleanup_policy}.
	CleanupPolicy *string `field:"optional" json:"cleanupPolicy" yaml:"cleanupPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#compression_type DataYandexMdbKafkaCluster#compression_type}.
	CompressionType *string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#delete_retention_ms DataYandexMdbKafkaCluster#delete_retention_ms}.
	DeleteRetentionMs *string `field:"optional" json:"deleteRetentionMs" yaml:"deleteRetentionMs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#file_delete_delay_ms DataYandexMdbKafkaCluster#file_delete_delay_ms}.
	FileDeleteDelayMs *string `field:"optional" json:"fileDeleteDelayMs" yaml:"fileDeleteDelayMs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#flush_messages DataYandexMdbKafkaCluster#flush_messages}.
	FlushMessages *string `field:"optional" json:"flushMessages" yaml:"flushMessages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#flush_ms DataYandexMdbKafkaCluster#flush_ms}.
	FlushMs *string `field:"optional" json:"flushMs" yaml:"flushMs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#max_message_bytes DataYandexMdbKafkaCluster#max_message_bytes}.
	MaxMessageBytes *string `field:"optional" json:"maxMessageBytes" yaml:"maxMessageBytes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#min_compaction_lag_ms DataYandexMdbKafkaCluster#min_compaction_lag_ms}.
	MinCompactionLagMs *string `field:"optional" json:"minCompactionLagMs" yaml:"minCompactionLagMs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#min_insync_replicas DataYandexMdbKafkaCluster#min_insync_replicas}.
	MinInsyncReplicas *string `field:"optional" json:"minInsyncReplicas" yaml:"minInsyncReplicas"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#preallocate DataYandexMdbKafkaCluster#preallocate}.
	Preallocate interface{} `field:"optional" json:"preallocate" yaml:"preallocate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#retention_bytes DataYandexMdbKafkaCluster#retention_bytes}.
	RetentionBytes *string `field:"optional" json:"retentionBytes" yaml:"retentionBytes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#retention_ms DataYandexMdbKafkaCluster#retention_ms}.
	RetentionMs *string `field:"optional" json:"retentionMs" yaml:"retentionMs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster#segment_bytes DataYandexMdbKafkaCluster#segment_bytes}.
	SegmentBytes *string `field:"optional" json:"segmentBytes" yaml:"segmentBytes"`
}

