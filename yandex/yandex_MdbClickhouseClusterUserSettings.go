// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbClickhouseClusterUserSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#add_http_cors_header MdbClickhouseCluster#add_http_cors_header}.
	AddHttpCorsHeader interface{} `field:"optional" json:"addHttpCorsHeader" yaml:"addHttpCorsHeader"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#allow_ddl MdbClickhouseCluster#allow_ddl}.
	AllowDdl interface{} `field:"optional" json:"allowDdl" yaml:"allowDdl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#compile MdbClickhouseCluster#compile}.
	Compile interface{} `field:"optional" json:"compile" yaml:"compile"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#compile_expressions MdbClickhouseCluster#compile_expressions}.
	CompileExpressions interface{} `field:"optional" json:"compileExpressions" yaml:"compileExpressions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#connect_timeout MdbClickhouseCluster#connect_timeout}.
	ConnectTimeout *float64 `field:"optional" json:"connectTimeout" yaml:"connectTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#count_distinct_implementation MdbClickhouseCluster#count_distinct_implementation}.
	CountDistinctImplementation *string `field:"optional" json:"countDistinctImplementation" yaml:"countDistinctImplementation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#distinct_overflow_mode MdbClickhouseCluster#distinct_overflow_mode}.
	DistinctOverflowMode *string `field:"optional" json:"distinctOverflowMode" yaml:"distinctOverflowMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#distributed_aggregation_memory_efficient MdbClickhouseCluster#distributed_aggregation_memory_efficient}.
	DistributedAggregationMemoryEfficient interface{} `field:"optional" json:"distributedAggregationMemoryEfficient" yaml:"distributedAggregationMemoryEfficient"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#distributed_ddl_task_timeout MdbClickhouseCluster#distributed_ddl_task_timeout}.
	DistributedDdlTaskTimeout *float64 `field:"optional" json:"distributedDdlTaskTimeout" yaml:"distributedDdlTaskTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#distributed_product_mode MdbClickhouseCluster#distributed_product_mode}.
	DistributedProductMode *string `field:"optional" json:"distributedProductMode" yaml:"distributedProductMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#empty_result_for_aggregation_by_empty_set MdbClickhouseCluster#empty_result_for_aggregation_by_empty_set}.
	EmptyResultForAggregationByEmptySet interface{} `field:"optional" json:"emptyResultForAggregationByEmptySet" yaml:"emptyResultForAggregationByEmptySet"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#enable_http_compression MdbClickhouseCluster#enable_http_compression}.
	EnableHttpCompression interface{} `field:"optional" json:"enableHttpCompression" yaml:"enableHttpCompression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#fallback_to_stale_replicas_for_distributed_queries MdbClickhouseCluster#fallback_to_stale_replicas_for_distributed_queries}.
	FallbackToStaleReplicasForDistributedQueries interface{} `field:"optional" json:"fallbackToStaleReplicasForDistributedQueries" yaml:"fallbackToStaleReplicasForDistributedQueries"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#force_index_by_date MdbClickhouseCluster#force_index_by_date}.
	ForceIndexByDate interface{} `field:"optional" json:"forceIndexByDate" yaml:"forceIndexByDate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#force_primary_key MdbClickhouseCluster#force_primary_key}.
	ForcePrimaryKey interface{} `field:"optional" json:"forcePrimaryKey" yaml:"forcePrimaryKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#group_by_overflow_mode MdbClickhouseCluster#group_by_overflow_mode}.
	GroupByOverflowMode *string `field:"optional" json:"groupByOverflowMode" yaml:"groupByOverflowMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#group_by_two_level_threshold MdbClickhouseCluster#group_by_two_level_threshold}.
	GroupByTwoLevelThreshold *float64 `field:"optional" json:"groupByTwoLevelThreshold" yaml:"groupByTwoLevelThreshold"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#group_by_two_level_threshold_bytes MdbClickhouseCluster#group_by_two_level_threshold_bytes}.
	GroupByTwoLevelThresholdBytes *float64 `field:"optional" json:"groupByTwoLevelThresholdBytes" yaml:"groupByTwoLevelThresholdBytes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#http_connection_timeout MdbClickhouseCluster#http_connection_timeout}.
	HttpConnectionTimeout *float64 `field:"optional" json:"httpConnectionTimeout" yaml:"httpConnectionTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#http_headers_progress_interval MdbClickhouseCluster#http_headers_progress_interval}.
	HttpHeadersProgressInterval *float64 `field:"optional" json:"httpHeadersProgressInterval" yaml:"httpHeadersProgressInterval"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#http_receive_timeout MdbClickhouseCluster#http_receive_timeout}.
	HttpReceiveTimeout *float64 `field:"optional" json:"httpReceiveTimeout" yaml:"httpReceiveTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#http_send_timeout MdbClickhouseCluster#http_send_timeout}.
	HttpSendTimeout *float64 `field:"optional" json:"httpSendTimeout" yaml:"httpSendTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#input_format_defaults_for_omitted_fields MdbClickhouseCluster#input_format_defaults_for_omitted_fields}.
	InputFormatDefaultsForOmittedFields interface{} `field:"optional" json:"inputFormatDefaultsForOmittedFields" yaml:"inputFormatDefaultsForOmittedFields"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#input_format_values_interpret_expressions MdbClickhouseCluster#input_format_values_interpret_expressions}.
	InputFormatValuesInterpretExpressions interface{} `field:"optional" json:"inputFormatValuesInterpretExpressions" yaml:"inputFormatValuesInterpretExpressions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#insert_quorum MdbClickhouseCluster#insert_quorum}.
	InsertQuorum *float64 `field:"optional" json:"insertQuorum" yaml:"insertQuorum"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#insert_quorum_timeout MdbClickhouseCluster#insert_quorum_timeout}.
	InsertQuorumTimeout *float64 `field:"optional" json:"insertQuorumTimeout" yaml:"insertQuorumTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#joined_subquery_requires_alias MdbClickhouseCluster#joined_subquery_requires_alias}.
	JoinedSubqueryRequiresAlias interface{} `field:"optional" json:"joinedSubqueryRequiresAlias" yaml:"joinedSubqueryRequiresAlias"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#join_overflow_mode MdbClickhouseCluster#join_overflow_mode}.
	JoinOverflowMode *string `field:"optional" json:"joinOverflowMode" yaml:"joinOverflowMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#join_use_nulls MdbClickhouseCluster#join_use_nulls}.
	JoinUseNulls interface{} `field:"optional" json:"joinUseNulls" yaml:"joinUseNulls"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#low_cardinality_allow_in_native_format MdbClickhouseCluster#low_cardinality_allow_in_native_format}.
	LowCardinalityAllowInNativeFormat interface{} `field:"optional" json:"lowCardinalityAllowInNativeFormat" yaml:"lowCardinalityAllowInNativeFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_ast_depth MdbClickhouseCluster#max_ast_depth}.
	MaxAstDepth *float64 `field:"optional" json:"maxAstDepth" yaml:"maxAstDepth"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_ast_elements MdbClickhouseCluster#max_ast_elements}.
	MaxAstElements *float64 `field:"optional" json:"maxAstElements" yaml:"maxAstElements"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_block_size MdbClickhouseCluster#max_block_size}.
	MaxBlockSize *float64 `field:"optional" json:"maxBlockSize" yaml:"maxBlockSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_bytes_before_external_group_by MdbClickhouseCluster#max_bytes_before_external_group_by}.
	MaxBytesBeforeExternalGroupBy *float64 `field:"optional" json:"maxBytesBeforeExternalGroupBy" yaml:"maxBytesBeforeExternalGroupBy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_bytes_before_external_sort MdbClickhouseCluster#max_bytes_before_external_sort}.
	MaxBytesBeforeExternalSort *float64 `field:"optional" json:"maxBytesBeforeExternalSort" yaml:"maxBytesBeforeExternalSort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_bytes_in_distinct MdbClickhouseCluster#max_bytes_in_distinct}.
	MaxBytesInDistinct *float64 `field:"optional" json:"maxBytesInDistinct" yaml:"maxBytesInDistinct"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_bytes_in_join MdbClickhouseCluster#max_bytes_in_join}.
	MaxBytesInJoin *float64 `field:"optional" json:"maxBytesInJoin" yaml:"maxBytesInJoin"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_bytes_in_set MdbClickhouseCluster#max_bytes_in_set}.
	MaxBytesInSet *float64 `field:"optional" json:"maxBytesInSet" yaml:"maxBytesInSet"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_bytes_to_read MdbClickhouseCluster#max_bytes_to_read}.
	MaxBytesToRead *float64 `field:"optional" json:"maxBytesToRead" yaml:"maxBytesToRead"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_bytes_to_sort MdbClickhouseCluster#max_bytes_to_sort}.
	MaxBytesToSort *float64 `field:"optional" json:"maxBytesToSort" yaml:"maxBytesToSort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_bytes_to_transfer MdbClickhouseCluster#max_bytes_to_transfer}.
	MaxBytesToTransfer *float64 `field:"optional" json:"maxBytesToTransfer" yaml:"maxBytesToTransfer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_columns_to_read MdbClickhouseCluster#max_columns_to_read}.
	MaxColumnsToRead *float64 `field:"optional" json:"maxColumnsToRead" yaml:"maxColumnsToRead"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_execution_time MdbClickhouseCluster#max_execution_time}.
	MaxExecutionTime *float64 `field:"optional" json:"maxExecutionTime" yaml:"maxExecutionTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_expanded_ast_elements MdbClickhouseCluster#max_expanded_ast_elements}.
	MaxExpandedAstElements *float64 `field:"optional" json:"maxExpandedAstElements" yaml:"maxExpandedAstElements"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_insert_block_size MdbClickhouseCluster#max_insert_block_size}.
	MaxInsertBlockSize *float64 `field:"optional" json:"maxInsertBlockSize" yaml:"maxInsertBlockSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_memory_usage MdbClickhouseCluster#max_memory_usage}.
	MaxMemoryUsage *float64 `field:"optional" json:"maxMemoryUsage" yaml:"maxMemoryUsage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_memory_usage_for_user MdbClickhouseCluster#max_memory_usage_for_user}.
	MaxMemoryUsageForUser *float64 `field:"optional" json:"maxMemoryUsageForUser" yaml:"maxMemoryUsageForUser"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_network_bandwidth MdbClickhouseCluster#max_network_bandwidth}.
	MaxNetworkBandwidth *float64 `field:"optional" json:"maxNetworkBandwidth" yaml:"maxNetworkBandwidth"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_network_bandwidth_for_user MdbClickhouseCluster#max_network_bandwidth_for_user}.
	MaxNetworkBandwidthForUser *float64 `field:"optional" json:"maxNetworkBandwidthForUser" yaml:"maxNetworkBandwidthForUser"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_query_size MdbClickhouseCluster#max_query_size}.
	MaxQuerySize *float64 `field:"optional" json:"maxQuerySize" yaml:"maxQuerySize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_replica_delay_for_distributed_queries MdbClickhouseCluster#max_replica_delay_for_distributed_queries}.
	MaxReplicaDelayForDistributedQueries *float64 `field:"optional" json:"maxReplicaDelayForDistributedQueries" yaml:"maxReplicaDelayForDistributedQueries"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_result_bytes MdbClickhouseCluster#max_result_bytes}.
	MaxResultBytes *float64 `field:"optional" json:"maxResultBytes" yaml:"maxResultBytes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_result_rows MdbClickhouseCluster#max_result_rows}.
	MaxResultRows *float64 `field:"optional" json:"maxResultRows" yaml:"maxResultRows"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_rows_in_distinct MdbClickhouseCluster#max_rows_in_distinct}.
	MaxRowsInDistinct *float64 `field:"optional" json:"maxRowsInDistinct" yaml:"maxRowsInDistinct"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_rows_in_join MdbClickhouseCluster#max_rows_in_join}.
	MaxRowsInJoin *float64 `field:"optional" json:"maxRowsInJoin" yaml:"maxRowsInJoin"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_rows_in_set MdbClickhouseCluster#max_rows_in_set}.
	MaxRowsInSet *float64 `field:"optional" json:"maxRowsInSet" yaml:"maxRowsInSet"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_rows_to_group_by MdbClickhouseCluster#max_rows_to_group_by}.
	MaxRowsToGroupBy *float64 `field:"optional" json:"maxRowsToGroupBy" yaml:"maxRowsToGroupBy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_rows_to_read MdbClickhouseCluster#max_rows_to_read}.
	MaxRowsToRead *float64 `field:"optional" json:"maxRowsToRead" yaml:"maxRowsToRead"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_rows_to_sort MdbClickhouseCluster#max_rows_to_sort}.
	MaxRowsToSort *float64 `field:"optional" json:"maxRowsToSort" yaml:"maxRowsToSort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_rows_to_transfer MdbClickhouseCluster#max_rows_to_transfer}.
	MaxRowsToTransfer *float64 `field:"optional" json:"maxRowsToTransfer" yaml:"maxRowsToTransfer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_temporary_columns MdbClickhouseCluster#max_temporary_columns}.
	MaxTemporaryColumns *float64 `field:"optional" json:"maxTemporaryColumns" yaml:"maxTemporaryColumns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_temporary_non_const_columns MdbClickhouseCluster#max_temporary_non_const_columns}.
	MaxTemporaryNonConstColumns *float64 `field:"optional" json:"maxTemporaryNonConstColumns" yaml:"maxTemporaryNonConstColumns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_threads MdbClickhouseCluster#max_threads}.
	MaxThreads *float64 `field:"optional" json:"maxThreads" yaml:"maxThreads"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#merge_tree_max_bytes_to_use_cache MdbClickhouseCluster#merge_tree_max_bytes_to_use_cache}.
	MergeTreeMaxBytesToUseCache *float64 `field:"optional" json:"mergeTreeMaxBytesToUseCache" yaml:"mergeTreeMaxBytesToUseCache"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#merge_tree_max_rows_to_use_cache MdbClickhouseCluster#merge_tree_max_rows_to_use_cache}.
	MergeTreeMaxRowsToUseCache *float64 `field:"optional" json:"mergeTreeMaxRowsToUseCache" yaml:"mergeTreeMaxRowsToUseCache"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#merge_tree_min_bytes_for_concurrent_read MdbClickhouseCluster#merge_tree_min_bytes_for_concurrent_read}.
	MergeTreeMinBytesForConcurrentRead *float64 `field:"optional" json:"mergeTreeMinBytesForConcurrentRead" yaml:"mergeTreeMinBytesForConcurrentRead"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#merge_tree_min_rows_for_concurrent_read MdbClickhouseCluster#merge_tree_min_rows_for_concurrent_read}.
	MergeTreeMinRowsForConcurrentRead *float64 `field:"optional" json:"mergeTreeMinRowsForConcurrentRead" yaml:"mergeTreeMinRowsForConcurrentRead"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#min_bytes_to_use_direct_io MdbClickhouseCluster#min_bytes_to_use_direct_io}.
	MinBytesToUseDirectIo *float64 `field:"optional" json:"minBytesToUseDirectIo" yaml:"minBytesToUseDirectIo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#min_count_to_compile MdbClickhouseCluster#min_count_to_compile}.
	MinCountToCompile *float64 `field:"optional" json:"minCountToCompile" yaml:"minCountToCompile"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#min_count_to_compile_expression MdbClickhouseCluster#min_count_to_compile_expression}.
	MinCountToCompileExpression *float64 `field:"optional" json:"minCountToCompileExpression" yaml:"minCountToCompileExpression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#min_execution_speed MdbClickhouseCluster#min_execution_speed}.
	MinExecutionSpeed *float64 `field:"optional" json:"minExecutionSpeed" yaml:"minExecutionSpeed"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#min_execution_speed_bytes MdbClickhouseCluster#min_execution_speed_bytes}.
	MinExecutionSpeedBytes *float64 `field:"optional" json:"minExecutionSpeedBytes" yaml:"minExecutionSpeedBytes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#min_insert_block_size_bytes MdbClickhouseCluster#min_insert_block_size_bytes}.
	MinInsertBlockSizeBytes *float64 `field:"optional" json:"minInsertBlockSizeBytes" yaml:"minInsertBlockSizeBytes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#min_insert_block_size_rows MdbClickhouseCluster#min_insert_block_size_rows}.
	MinInsertBlockSizeRows *float64 `field:"optional" json:"minInsertBlockSizeRows" yaml:"minInsertBlockSizeRows"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#output_format_json_quote_64bit_integers MdbClickhouseCluster#output_format_json_quote_64bit_integers}.
	OutputFormatJsonQuote64BitIntegers interface{} `field:"optional" json:"outputFormatJsonQuote64BitIntegers" yaml:"outputFormatJsonQuote64BitIntegers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#output_format_json_quote_denormals MdbClickhouseCluster#output_format_json_quote_denormals}.
	OutputFormatJsonQuoteDenormals interface{} `field:"optional" json:"outputFormatJsonQuoteDenormals" yaml:"outputFormatJsonQuoteDenormals"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#priority MdbClickhouseCluster#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#quota_mode MdbClickhouseCluster#quota_mode}.
	QuotaMode *string `field:"optional" json:"quotaMode" yaml:"quotaMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#readonly MdbClickhouseCluster#readonly}.
	Readonly *float64 `field:"optional" json:"readonly" yaml:"readonly"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#read_overflow_mode MdbClickhouseCluster#read_overflow_mode}.
	ReadOverflowMode *string `field:"optional" json:"readOverflowMode" yaml:"readOverflowMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#receive_timeout MdbClickhouseCluster#receive_timeout}.
	ReceiveTimeout *float64 `field:"optional" json:"receiveTimeout" yaml:"receiveTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#replication_alter_partitions_sync MdbClickhouseCluster#replication_alter_partitions_sync}.
	ReplicationAlterPartitionsSync *float64 `field:"optional" json:"replicationAlterPartitionsSync" yaml:"replicationAlterPartitionsSync"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#result_overflow_mode MdbClickhouseCluster#result_overflow_mode}.
	ResultOverflowMode *string `field:"optional" json:"resultOverflowMode" yaml:"resultOverflowMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#select_sequential_consistency MdbClickhouseCluster#select_sequential_consistency}.
	SelectSequentialConsistency interface{} `field:"optional" json:"selectSequentialConsistency" yaml:"selectSequentialConsistency"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#send_progress_in_http_headers MdbClickhouseCluster#send_progress_in_http_headers}.
	SendProgressInHttpHeaders interface{} `field:"optional" json:"sendProgressInHttpHeaders" yaml:"sendProgressInHttpHeaders"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#send_timeout MdbClickhouseCluster#send_timeout}.
	SendTimeout *float64 `field:"optional" json:"sendTimeout" yaml:"sendTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#set_overflow_mode MdbClickhouseCluster#set_overflow_mode}.
	SetOverflowMode *string `field:"optional" json:"setOverflowMode" yaml:"setOverflowMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#skip_unavailable_shards MdbClickhouseCluster#skip_unavailable_shards}.
	SkipUnavailableShards interface{} `field:"optional" json:"skipUnavailableShards" yaml:"skipUnavailableShards"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#sort_overflow_mode MdbClickhouseCluster#sort_overflow_mode}.
	SortOverflowMode *string `field:"optional" json:"sortOverflowMode" yaml:"sortOverflowMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#timeout_overflow_mode MdbClickhouseCluster#timeout_overflow_mode}.
	TimeoutOverflowMode *string `field:"optional" json:"timeoutOverflowMode" yaml:"timeoutOverflowMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#transfer_overflow_mode MdbClickhouseCluster#transfer_overflow_mode}.
	TransferOverflowMode *string `field:"optional" json:"transferOverflowMode" yaml:"transferOverflowMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#transform_null_in MdbClickhouseCluster#transform_null_in}.
	TransformNullIn interface{} `field:"optional" json:"transformNullIn" yaml:"transformNullIn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#use_uncompressed_cache MdbClickhouseCluster#use_uncompressed_cache}.
	UseUncompressedCache interface{} `field:"optional" json:"useUncompressedCache" yaml:"useUncompressedCache"`
}

