// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbKafkaClusterTopicTopicConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#cleanup_policy MdbKafkaCluster#cleanup_policy}.
	CleanupPolicy *string `field:"optional" json:"cleanupPolicy" yaml:"cleanupPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#compression_type MdbKafkaCluster#compression_type}.
	CompressionType *string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#delete_retention_ms MdbKafkaCluster#delete_retention_ms}.
	DeleteRetentionMs *string `field:"optional" json:"deleteRetentionMs" yaml:"deleteRetentionMs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#file_delete_delay_ms MdbKafkaCluster#file_delete_delay_ms}.
	FileDeleteDelayMs *string `field:"optional" json:"fileDeleteDelayMs" yaml:"fileDeleteDelayMs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#flush_messages MdbKafkaCluster#flush_messages}.
	FlushMessages *string `field:"optional" json:"flushMessages" yaml:"flushMessages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#flush_ms MdbKafkaCluster#flush_ms}.
	FlushMs *string `field:"optional" json:"flushMs" yaml:"flushMs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#max_message_bytes MdbKafkaCluster#max_message_bytes}.
	MaxMessageBytes *string `field:"optional" json:"maxMessageBytes" yaml:"maxMessageBytes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#min_compaction_lag_ms MdbKafkaCluster#min_compaction_lag_ms}.
	MinCompactionLagMs *string `field:"optional" json:"minCompactionLagMs" yaml:"minCompactionLagMs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#min_insync_replicas MdbKafkaCluster#min_insync_replicas}.
	MinInsyncReplicas *string `field:"optional" json:"minInsyncReplicas" yaml:"minInsyncReplicas"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#preallocate MdbKafkaCluster#preallocate}.
	Preallocate interface{} `field:"optional" json:"preallocate" yaml:"preallocate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#retention_bytes MdbKafkaCluster#retention_bytes}.
	RetentionBytes *string `field:"optional" json:"retentionBytes" yaml:"retentionBytes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#retention_ms MdbKafkaCluster#retention_ms}.
	RetentionMs *string `field:"optional" json:"retentionMs" yaml:"retentionMs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_cluster#segment_bytes MdbKafkaCluster#segment_bytes}.
	SegmentBytes *string `field:"optional" json:"segmentBytes" yaml:"segmentBytes"`
}

