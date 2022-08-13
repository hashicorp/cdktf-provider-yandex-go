// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbClickhouseClusterClickhouseConfigMergeTree struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_bytes_to_merge_at_min_space_in_pool MdbClickhouseCluster#max_bytes_to_merge_at_min_space_in_pool}.
	MaxBytesToMergeAtMinSpaceInPool *float64 `field:"optional" json:"maxBytesToMergeAtMinSpaceInPool" yaml:"maxBytesToMergeAtMinSpaceInPool"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#max_replicated_merges_in_queue MdbClickhouseCluster#max_replicated_merges_in_queue}.
	MaxReplicatedMergesInQueue *float64 `field:"optional" json:"maxReplicatedMergesInQueue" yaml:"maxReplicatedMergesInQueue"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#number_of_free_entries_in_pool_to_lower_max_size_of_merge MdbClickhouseCluster#number_of_free_entries_in_pool_to_lower_max_size_of_merge}.
	NumberOfFreeEntriesInPoolToLowerMaxSizeOfMerge *float64 `field:"optional" json:"numberOfFreeEntriesInPoolToLowerMaxSizeOfMerge" yaml:"numberOfFreeEntriesInPoolToLowerMaxSizeOfMerge"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#parts_to_delay_insert MdbClickhouseCluster#parts_to_delay_insert}.
	PartsToDelayInsert *float64 `field:"optional" json:"partsToDelayInsert" yaml:"partsToDelayInsert"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#parts_to_throw_insert MdbClickhouseCluster#parts_to_throw_insert}.
	PartsToThrowInsert *float64 `field:"optional" json:"partsToThrowInsert" yaml:"partsToThrowInsert"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#replicated_deduplication_window MdbClickhouseCluster#replicated_deduplication_window}.
	ReplicatedDeduplicationWindow *float64 `field:"optional" json:"replicatedDeduplicationWindow" yaml:"replicatedDeduplicationWindow"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#replicated_deduplication_window_seconds MdbClickhouseCluster#replicated_deduplication_window_seconds}.
	ReplicatedDeduplicationWindowSeconds *float64 `field:"optional" json:"replicatedDeduplicationWindowSeconds" yaml:"replicatedDeduplicationWindowSeconds"`
}

