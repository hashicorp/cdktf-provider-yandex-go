// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbKafkaTopicTopicConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_topic#cleanup_policy MdbKafkaTopic#cleanup_policy}.
	CleanupPolicy *string `field:"optional" json:"cleanupPolicy" yaml:"cleanupPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_topic#compression_type MdbKafkaTopic#compression_type}.
	CompressionType *string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_topic#delete_retention_ms MdbKafkaTopic#delete_retention_ms}.
	DeleteRetentionMs *string `field:"optional" json:"deleteRetentionMs" yaml:"deleteRetentionMs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_topic#file_delete_delay_ms MdbKafkaTopic#file_delete_delay_ms}.
	FileDeleteDelayMs *string `field:"optional" json:"fileDeleteDelayMs" yaml:"fileDeleteDelayMs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_topic#flush_messages MdbKafkaTopic#flush_messages}.
	FlushMessages *string `field:"optional" json:"flushMessages" yaml:"flushMessages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_topic#flush_ms MdbKafkaTopic#flush_ms}.
	FlushMs *string `field:"optional" json:"flushMs" yaml:"flushMs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_topic#max_message_bytes MdbKafkaTopic#max_message_bytes}.
	MaxMessageBytes *string `field:"optional" json:"maxMessageBytes" yaml:"maxMessageBytes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_topic#min_compaction_lag_ms MdbKafkaTopic#min_compaction_lag_ms}.
	MinCompactionLagMs *string `field:"optional" json:"minCompactionLagMs" yaml:"minCompactionLagMs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_topic#min_insync_replicas MdbKafkaTopic#min_insync_replicas}.
	MinInsyncReplicas *string `field:"optional" json:"minInsyncReplicas" yaml:"minInsyncReplicas"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_topic#preallocate MdbKafkaTopic#preallocate}.
	Preallocate interface{} `field:"optional" json:"preallocate" yaml:"preallocate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_topic#retention_bytes MdbKafkaTopic#retention_bytes}.
	RetentionBytes *string `field:"optional" json:"retentionBytes" yaml:"retentionBytes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_topic#retention_ms MdbKafkaTopic#retention_ms}.
	RetentionMs *string `field:"optional" json:"retentionMs" yaml:"retentionMs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_kafka_topic#segment_bytes MdbKafkaTopic#segment_bytes}.
	SegmentBytes *string `field:"optional" json:"segmentBytes" yaml:"segmentBytes"`
}

