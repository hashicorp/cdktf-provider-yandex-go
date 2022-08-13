// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbClickhouseClusterClickhouseConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#background_pool_size MdbClickhouseCluster#background_pool_size}.
	BackgroundPoolSize *float64 `field:"optional" json:"backgroundPoolSize" yaml:"backgroundPoolSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#background_schedule_pool_size MdbClickhouseCluster#background_schedule_pool_size}.
	BackgroundSchedulePoolSize *float64 `field:"optional" json:"backgroundSchedulePoolSize" yaml:"backgroundSchedulePoolSize"`
	// compression block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#compression MdbClickhouseCluster#compression}
	Compression interface{} `field:"optional" json:"compression" yaml:"compression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#geobase_uri MdbClickhouseCluster#geobase_uri}.
	GeobaseUri *string `field:"optional" json:"geobaseUri" yaml:"geobaseUri"`
	// graphite_rollup block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#graphite_rollup MdbClickhouseCluster#graphite_rollup}
	GraphiteRollup interface{} `field:"optional" json:"graphiteRollup" yaml:"graphiteRollup"`
	// kafka block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#kafka MdbClickhouseCluster#kafka}
	Kafka *MdbClickhouseClusterClickhouseConfigKafka `field:"optional" json:"kafka" yaml:"kafka"`
	// kafka_topic block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#kafka_topic MdbClickhouseCluster#kafka_topic}
	KafkaTopic interface{} `field:"optional" json:"kafkaTopic" yaml:"kafkaTopic"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#keep_alive_timeout MdbClickhouseCluster#keep_alive_timeout}.
	KeepAliveTimeout *float64 `field:"optional" json:"keepAliveTimeout" yaml:"keepAliveTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#log_level MdbClickhouseCluster#log_level}.
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#mark_cache_size MdbClickhouseCluster#mark_cache_size}.
	MarkCacheSize *float64 `field:"optional" json:"markCacheSize" yaml:"markCacheSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_concurrent_queries MdbClickhouseCluster#max_concurrent_queries}.
	MaxConcurrentQueries *float64 `field:"optional" json:"maxConcurrentQueries" yaml:"maxConcurrentQueries"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_connections MdbClickhouseCluster#max_connections}.
	MaxConnections *float64 `field:"optional" json:"maxConnections" yaml:"maxConnections"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_partition_size_to_drop MdbClickhouseCluster#max_partition_size_to_drop}.
	MaxPartitionSizeToDrop *float64 `field:"optional" json:"maxPartitionSizeToDrop" yaml:"maxPartitionSizeToDrop"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_table_size_to_drop MdbClickhouseCluster#max_table_size_to_drop}.
	MaxTableSizeToDrop *float64 `field:"optional" json:"maxTableSizeToDrop" yaml:"maxTableSizeToDrop"`
	// merge_tree block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#merge_tree MdbClickhouseCluster#merge_tree}
	MergeTree *MdbClickhouseClusterClickhouseConfigMergeTree `field:"optional" json:"mergeTree" yaml:"mergeTree"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#metric_log_enabled MdbClickhouseCluster#metric_log_enabled}.
	MetricLogEnabled interface{} `field:"optional" json:"metricLogEnabled" yaml:"metricLogEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#metric_log_retention_size MdbClickhouseCluster#metric_log_retention_size}.
	MetricLogRetentionSize *float64 `field:"optional" json:"metricLogRetentionSize" yaml:"metricLogRetentionSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#metric_log_retention_time MdbClickhouseCluster#metric_log_retention_time}.
	MetricLogRetentionTime *float64 `field:"optional" json:"metricLogRetentionTime" yaml:"metricLogRetentionTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#part_log_retention_size MdbClickhouseCluster#part_log_retention_size}.
	PartLogRetentionSize *float64 `field:"optional" json:"partLogRetentionSize" yaml:"partLogRetentionSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#part_log_retention_time MdbClickhouseCluster#part_log_retention_time}.
	PartLogRetentionTime *float64 `field:"optional" json:"partLogRetentionTime" yaml:"partLogRetentionTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#query_log_retention_size MdbClickhouseCluster#query_log_retention_size}.
	QueryLogRetentionSize *float64 `field:"optional" json:"queryLogRetentionSize" yaml:"queryLogRetentionSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#query_log_retention_time MdbClickhouseCluster#query_log_retention_time}.
	QueryLogRetentionTime *float64 `field:"optional" json:"queryLogRetentionTime" yaml:"queryLogRetentionTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#query_thread_log_enabled MdbClickhouseCluster#query_thread_log_enabled}.
	QueryThreadLogEnabled interface{} `field:"optional" json:"queryThreadLogEnabled" yaml:"queryThreadLogEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#query_thread_log_retention_size MdbClickhouseCluster#query_thread_log_retention_size}.
	QueryThreadLogRetentionSize *float64 `field:"optional" json:"queryThreadLogRetentionSize" yaml:"queryThreadLogRetentionSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#query_thread_log_retention_time MdbClickhouseCluster#query_thread_log_retention_time}.
	QueryThreadLogRetentionTime *float64 `field:"optional" json:"queryThreadLogRetentionTime" yaml:"queryThreadLogRetentionTime"`
	// rabbitmq block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#rabbitmq MdbClickhouseCluster#rabbitmq}
	Rabbitmq *MdbClickhouseClusterClickhouseConfigRabbitmq `field:"optional" json:"rabbitmq" yaml:"rabbitmq"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#text_log_enabled MdbClickhouseCluster#text_log_enabled}.
	TextLogEnabled interface{} `field:"optional" json:"textLogEnabled" yaml:"textLogEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#text_log_level MdbClickhouseCluster#text_log_level}.
	TextLogLevel *string `field:"optional" json:"textLogLevel" yaml:"textLogLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#text_log_retention_size MdbClickhouseCluster#text_log_retention_size}.
	TextLogRetentionSize *float64 `field:"optional" json:"textLogRetentionSize" yaml:"textLogRetentionSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#text_log_retention_time MdbClickhouseCluster#text_log_retention_time}.
	TextLogRetentionTime *float64 `field:"optional" json:"textLogRetentionTime" yaml:"textLogRetentionTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#timezone MdbClickhouseCluster#timezone}.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#trace_log_enabled MdbClickhouseCluster#trace_log_enabled}.
	TraceLogEnabled interface{} `field:"optional" json:"traceLogEnabled" yaml:"traceLogEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#trace_log_retention_size MdbClickhouseCluster#trace_log_retention_size}.
	TraceLogRetentionSize *float64 `field:"optional" json:"traceLogRetentionSize" yaml:"traceLogRetentionSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#trace_log_retention_time MdbClickhouseCluster#trace_log_retention_time}.
	TraceLogRetentionTime *float64 `field:"optional" json:"traceLogRetentionTime" yaml:"traceLogRetentionTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#uncompressed_cache_size MdbClickhouseCluster#uncompressed_cache_size}.
	UncompressedCacheSize *float64 `field:"optional" json:"uncompressedCacheSize" yaml:"uncompressedCacheSize"`
}

