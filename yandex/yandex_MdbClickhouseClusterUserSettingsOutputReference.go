// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbClickhouseClusterUserSettingsOutputReference interface {
	cdktf.ComplexObject
	AddHttpCorsHeader() interface{}
	SetAddHttpCorsHeader(val interface{})
	AddHttpCorsHeaderInput() interface{}
	AllowDdl() interface{}
	SetAllowDdl(val interface{})
	AllowDdlInput() interface{}
	Compile() interface{}
	SetCompile(val interface{})
	CompileExpressions() interface{}
	SetCompileExpressions(val interface{})
	CompileExpressionsInput() interface{}
	CompileInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ConnectTimeout() *float64
	SetConnectTimeout(val *float64)
	ConnectTimeoutInput() *float64
	CountDistinctImplementation() *string
	SetCountDistinctImplementation(val *string)
	CountDistinctImplementationInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DistinctOverflowMode() *string
	SetDistinctOverflowMode(val *string)
	DistinctOverflowModeInput() *string
	DistributedAggregationMemoryEfficient() interface{}
	SetDistributedAggregationMemoryEfficient(val interface{})
	DistributedAggregationMemoryEfficientInput() interface{}
	DistributedDdlTaskTimeout() *float64
	SetDistributedDdlTaskTimeout(val *float64)
	DistributedDdlTaskTimeoutInput() *float64
	DistributedProductMode() *string
	SetDistributedProductMode(val *string)
	DistributedProductModeInput() *string
	EmptyResultForAggregationByEmptySet() interface{}
	SetEmptyResultForAggregationByEmptySet(val interface{})
	EmptyResultForAggregationByEmptySetInput() interface{}
	EnableHttpCompression() interface{}
	SetEnableHttpCompression(val interface{})
	EnableHttpCompressionInput() interface{}
	FallbackToStaleReplicasForDistributedQueries() interface{}
	SetFallbackToStaleReplicasForDistributedQueries(val interface{})
	FallbackToStaleReplicasForDistributedQueriesInput() interface{}
	ForceIndexByDate() interface{}
	SetForceIndexByDate(val interface{})
	ForceIndexByDateInput() interface{}
	ForcePrimaryKey() interface{}
	SetForcePrimaryKey(val interface{})
	ForcePrimaryKeyInput() interface{}
	// Experimental.
	Fqn() *string
	GroupByOverflowMode() *string
	SetGroupByOverflowMode(val *string)
	GroupByOverflowModeInput() *string
	GroupByTwoLevelThreshold() *float64
	SetGroupByTwoLevelThreshold(val *float64)
	GroupByTwoLevelThresholdBytes() *float64
	SetGroupByTwoLevelThresholdBytes(val *float64)
	GroupByTwoLevelThresholdBytesInput() *float64
	GroupByTwoLevelThresholdInput() *float64
	HttpConnectionTimeout() *float64
	SetHttpConnectionTimeout(val *float64)
	HttpConnectionTimeoutInput() *float64
	HttpHeadersProgressInterval() *float64
	SetHttpHeadersProgressInterval(val *float64)
	HttpHeadersProgressIntervalInput() *float64
	HttpReceiveTimeout() *float64
	SetHttpReceiveTimeout(val *float64)
	HttpReceiveTimeoutInput() *float64
	HttpSendTimeout() *float64
	SetHttpSendTimeout(val *float64)
	HttpSendTimeoutInput() *float64
	InputFormatDefaultsForOmittedFields() interface{}
	SetInputFormatDefaultsForOmittedFields(val interface{})
	InputFormatDefaultsForOmittedFieldsInput() interface{}
	InputFormatValuesInterpretExpressions() interface{}
	SetInputFormatValuesInterpretExpressions(val interface{})
	InputFormatValuesInterpretExpressionsInput() interface{}
	InsertQuorum() *float64
	SetInsertQuorum(val *float64)
	InsertQuorumInput() *float64
	InsertQuorumTimeout() *float64
	SetInsertQuorumTimeout(val *float64)
	InsertQuorumTimeoutInput() *float64
	InternalValue() *MdbClickhouseClusterUserSettings
	SetInternalValue(val *MdbClickhouseClusterUserSettings)
	JoinedSubqueryRequiresAlias() interface{}
	SetJoinedSubqueryRequiresAlias(val interface{})
	JoinedSubqueryRequiresAliasInput() interface{}
	JoinOverflowMode() *string
	SetJoinOverflowMode(val *string)
	JoinOverflowModeInput() *string
	JoinUseNulls() interface{}
	SetJoinUseNulls(val interface{})
	JoinUseNullsInput() interface{}
	LowCardinalityAllowInNativeFormat() interface{}
	SetLowCardinalityAllowInNativeFormat(val interface{})
	LowCardinalityAllowInNativeFormatInput() interface{}
	MaxAstDepth() *float64
	SetMaxAstDepth(val *float64)
	MaxAstDepthInput() *float64
	MaxAstElements() *float64
	SetMaxAstElements(val *float64)
	MaxAstElementsInput() *float64
	MaxBlockSize() *float64
	SetMaxBlockSize(val *float64)
	MaxBlockSizeInput() *float64
	MaxBytesBeforeExternalGroupBy() *float64
	SetMaxBytesBeforeExternalGroupBy(val *float64)
	MaxBytesBeforeExternalGroupByInput() *float64
	MaxBytesBeforeExternalSort() *float64
	SetMaxBytesBeforeExternalSort(val *float64)
	MaxBytesBeforeExternalSortInput() *float64
	MaxBytesInDistinct() *float64
	SetMaxBytesInDistinct(val *float64)
	MaxBytesInDistinctInput() *float64
	MaxBytesInJoin() *float64
	SetMaxBytesInJoin(val *float64)
	MaxBytesInJoinInput() *float64
	MaxBytesInSet() *float64
	SetMaxBytesInSet(val *float64)
	MaxBytesInSetInput() *float64
	MaxBytesToRead() *float64
	SetMaxBytesToRead(val *float64)
	MaxBytesToReadInput() *float64
	MaxBytesToSort() *float64
	SetMaxBytesToSort(val *float64)
	MaxBytesToSortInput() *float64
	MaxBytesToTransfer() *float64
	SetMaxBytesToTransfer(val *float64)
	MaxBytesToTransferInput() *float64
	MaxColumnsToRead() *float64
	SetMaxColumnsToRead(val *float64)
	MaxColumnsToReadInput() *float64
	MaxExecutionTime() *float64
	SetMaxExecutionTime(val *float64)
	MaxExecutionTimeInput() *float64
	MaxExpandedAstElements() *float64
	SetMaxExpandedAstElements(val *float64)
	MaxExpandedAstElementsInput() *float64
	MaxInsertBlockSize() *float64
	SetMaxInsertBlockSize(val *float64)
	MaxInsertBlockSizeInput() *float64
	MaxMemoryUsage() *float64
	SetMaxMemoryUsage(val *float64)
	MaxMemoryUsageForUser() *float64
	SetMaxMemoryUsageForUser(val *float64)
	MaxMemoryUsageForUserInput() *float64
	MaxMemoryUsageInput() *float64
	MaxNetworkBandwidth() *float64
	SetMaxNetworkBandwidth(val *float64)
	MaxNetworkBandwidthForUser() *float64
	SetMaxNetworkBandwidthForUser(val *float64)
	MaxNetworkBandwidthForUserInput() *float64
	MaxNetworkBandwidthInput() *float64
	MaxQuerySize() *float64
	SetMaxQuerySize(val *float64)
	MaxQuerySizeInput() *float64
	MaxReplicaDelayForDistributedQueries() *float64
	SetMaxReplicaDelayForDistributedQueries(val *float64)
	MaxReplicaDelayForDistributedQueriesInput() *float64
	MaxResultBytes() *float64
	SetMaxResultBytes(val *float64)
	MaxResultBytesInput() *float64
	MaxResultRows() *float64
	SetMaxResultRows(val *float64)
	MaxResultRowsInput() *float64
	MaxRowsInDistinct() *float64
	SetMaxRowsInDistinct(val *float64)
	MaxRowsInDistinctInput() *float64
	MaxRowsInJoin() *float64
	SetMaxRowsInJoin(val *float64)
	MaxRowsInJoinInput() *float64
	MaxRowsInSet() *float64
	SetMaxRowsInSet(val *float64)
	MaxRowsInSetInput() *float64
	MaxRowsToGroupBy() *float64
	SetMaxRowsToGroupBy(val *float64)
	MaxRowsToGroupByInput() *float64
	MaxRowsToRead() *float64
	SetMaxRowsToRead(val *float64)
	MaxRowsToReadInput() *float64
	MaxRowsToSort() *float64
	SetMaxRowsToSort(val *float64)
	MaxRowsToSortInput() *float64
	MaxRowsToTransfer() *float64
	SetMaxRowsToTransfer(val *float64)
	MaxRowsToTransferInput() *float64
	MaxTemporaryColumns() *float64
	SetMaxTemporaryColumns(val *float64)
	MaxTemporaryColumnsInput() *float64
	MaxTemporaryNonConstColumns() *float64
	SetMaxTemporaryNonConstColumns(val *float64)
	MaxTemporaryNonConstColumnsInput() *float64
	MaxThreads() *float64
	SetMaxThreads(val *float64)
	MaxThreadsInput() *float64
	MergeTreeMaxBytesToUseCache() *float64
	SetMergeTreeMaxBytesToUseCache(val *float64)
	MergeTreeMaxBytesToUseCacheInput() *float64
	MergeTreeMaxRowsToUseCache() *float64
	SetMergeTreeMaxRowsToUseCache(val *float64)
	MergeTreeMaxRowsToUseCacheInput() *float64
	MergeTreeMinBytesForConcurrentRead() *float64
	SetMergeTreeMinBytesForConcurrentRead(val *float64)
	MergeTreeMinBytesForConcurrentReadInput() *float64
	MergeTreeMinRowsForConcurrentRead() *float64
	SetMergeTreeMinRowsForConcurrentRead(val *float64)
	MergeTreeMinRowsForConcurrentReadInput() *float64
	MinBytesToUseDirectIo() *float64
	SetMinBytesToUseDirectIo(val *float64)
	MinBytesToUseDirectIoInput() *float64
	MinCountToCompile() *float64
	SetMinCountToCompile(val *float64)
	MinCountToCompileExpression() *float64
	SetMinCountToCompileExpression(val *float64)
	MinCountToCompileExpressionInput() *float64
	MinCountToCompileInput() *float64
	MinExecutionSpeed() *float64
	SetMinExecutionSpeed(val *float64)
	MinExecutionSpeedBytes() *float64
	SetMinExecutionSpeedBytes(val *float64)
	MinExecutionSpeedBytesInput() *float64
	MinExecutionSpeedInput() *float64
	MinInsertBlockSizeBytes() *float64
	SetMinInsertBlockSizeBytes(val *float64)
	MinInsertBlockSizeBytesInput() *float64
	MinInsertBlockSizeRows() *float64
	SetMinInsertBlockSizeRows(val *float64)
	MinInsertBlockSizeRowsInput() *float64
	OutputFormatJsonQuote64BitIntegers() interface{}
	SetOutputFormatJsonQuote64BitIntegers(val interface{})
	OutputFormatJsonQuote64BitIntegersInput() interface{}
	OutputFormatJsonQuoteDenormals() interface{}
	SetOutputFormatJsonQuoteDenormals(val interface{})
	OutputFormatJsonQuoteDenormalsInput() interface{}
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	QuotaMode() *string
	SetQuotaMode(val *string)
	QuotaModeInput() *string
	Readonly() *float64
	SetReadonly(val *float64)
	ReadonlyInput() *float64
	ReadOverflowMode() *string
	SetReadOverflowMode(val *string)
	ReadOverflowModeInput() *string
	ReceiveTimeout() *float64
	SetReceiveTimeout(val *float64)
	ReceiveTimeoutInput() *float64
	ReplicationAlterPartitionsSync() *float64
	SetReplicationAlterPartitionsSync(val *float64)
	ReplicationAlterPartitionsSyncInput() *float64
	ResultOverflowMode() *string
	SetResultOverflowMode(val *string)
	ResultOverflowModeInput() *string
	SelectSequentialConsistency() interface{}
	SetSelectSequentialConsistency(val interface{})
	SelectSequentialConsistencyInput() interface{}
	SendProgressInHttpHeaders() interface{}
	SetSendProgressInHttpHeaders(val interface{})
	SendProgressInHttpHeadersInput() interface{}
	SendTimeout() *float64
	SetSendTimeout(val *float64)
	SendTimeoutInput() *float64
	SetOverflowMode() *string
	SetSetOverflowMode(val *string)
	SetOverflowModeInput() *string
	SkipUnavailableShards() interface{}
	SetSkipUnavailableShards(val interface{})
	SkipUnavailableShardsInput() interface{}
	SortOverflowMode() *string
	SetSortOverflowMode(val *string)
	SortOverflowModeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeoutOverflowMode() *string
	SetTimeoutOverflowMode(val *string)
	TimeoutOverflowModeInput() *string
	TransferOverflowMode() *string
	SetTransferOverflowMode(val *string)
	TransferOverflowModeInput() *string
	TransformNullIn() interface{}
	SetTransformNullIn(val interface{})
	TransformNullInInput() interface{}
	UseUncompressedCache() interface{}
	SetUseUncompressedCache(val interface{})
	UseUncompressedCacheInput() interface{}
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetAddHttpCorsHeader()
	ResetAllowDdl()
	ResetCompile()
	ResetCompileExpressions()
	ResetConnectTimeout()
	ResetCountDistinctImplementation()
	ResetDistinctOverflowMode()
	ResetDistributedAggregationMemoryEfficient()
	ResetDistributedDdlTaskTimeout()
	ResetDistributedProductMode()
	ResetEmptyResultForAggregationByEmptySet()
	ResetEnableHttpCompression()
	ResetFallbackToStaleReplicasForDistributedQueries()
	ResetForceIndexByDate()
	ResetForcePrimaryKey()
	ResetGroupByOverflowMode()
	ResetGroupByTwoLevelThreshold()
	ResetGroupByTwoLevelThresholdBytes()
	ResetHttpConnectionTimeout()
	ResetHttpHeadersProgressInterval()
	ResetHttpReceiveTimeout()
	ResetHttpSendTimeout()
	ResetInputFormatDefaultsForOmittedFields()
	ResetInputFormatValuesInterpretExpressions()
	ResetInsertQuorum()
	ResetInsertQuorumTimeout()
	ResetJoinedSubqueryRequiresAlias()
	ResetJoinOverflowMode()
	ResetJoinUseNulls()
	ResetLowCardinalityAllowInNativeFormat()
	ResetMaxAstDepth()
	ResetMaxAstElements()
	ResetMaxBlockSize()
	ResetMaxBytesBeforeExternalGroupBy()
	ResetMaxBytesBeforeExternalSort()
	ResetMaxBytesInDistinct()
	ResetMaxBytesInJoin()
	ResetMaxBytesInSet()
	ResetMaxBytesToRead()
	ResetMaxBytesToSort()
	ResetMaxBytesToTransfer()
	ResetMaxColumnsToRead()
	ResetMaxExecutionTime()
	ResetMaxExpandedAstElements()
	ResetMaxInsertBlockSize()
	ResetMaxMemoryUsage()
	ResetMaxMemoryUsageForUser()
	ResetMaxNetworkBandwidth()
	ResetMaxNetworkBandwidthForUser()
	ResetMaxQuerySize()
	ResetMaxReplicaDelayForDistributedQueries()
	ResetMaxResultBytes()
	ResetMaxResultRows()
	ResetMaxRowsInDistinct()
	ResetMaxRowsInJoin()
	ResetMaxRowsInSet()
	ResetMaxRowsToGroupBy()
	ResetMaxRowsToRead()
	ResetMaxRowsToSort()
	ResetMaxRowsToTransfer()
	ResetMaxTemporaryColumns()
	ResetMaxTemporaryNonConstColumns()
	ResetMaxThreads()
	ResetMergeTreeMaxBytesToUseCache()
	ResetMergeTreeMaxRowsToUseCache()
	ResetMergeTreeMinBytesForConcurrentRead()
	ResetMergeTreeMinRowsForConcurrentRead()
	ResetMinBytesToUseDirectIo()
	ResetMinCountToCompile()
	ResetMinCountToCompileExpression()
	ResetMinExecutionSpeed()
	ResetMinExecutionSpeedBytes()
	ResetMinInsertBlockSizeBytes()
	ResetMinInsertBlockSizeRows()
	ResetOutputFormatJsonQuote64BitIntegers()
	ResetOutputFormatJsonQuoteDenormals()
	ResetPriority()
	ResetQuotaMode()
	ResetReadonly()
	ResetReadOverflowMode()
	ResetReceiveTimeout()
	ResetReplicationAlterPartitionsSync()
	ResetResultOverflowMode()
	ResetSelectSequentialConsistency()
	ResetSendProgressInHttpHeaders()
	ResetSendTimeout()
	ResetSetOverflowMode()
	ResetSkipUnavailableShards()
	ResetSortOverflowMode()
	ResetTimeoutOverflowMode()
	ResetTransferOverflowMode()
	ResetTransformNullIn()
	ResetUseUncompressedCache()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MdbClickhouseClusterUserSettingsOutputReference
type jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) AddHttpCorsHeader() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addHttpCorsHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) AddHttpCorsHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addHttpCorsHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) AllowDdl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowDdl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) AllowDdlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowDdlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) Compile() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) CompileExpressions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compileExpressions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) CompileExpressionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compileExpressionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) CompileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ConnectTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ConnectTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) CountDistinctImplementation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countDistinctImplementation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) CountDistinctImplementationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countDistinctImplementationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) DistinctOverflowMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distinctOverflowMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) DistinctOverflowModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distinctOverflowModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) DistributedAggregationMemoryEfficient() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"distributedAggregationMemoryEfficient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) DistributedAggregationMemoryEfficientInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"distributedAggregationMemoryEfficientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) DistributedDdlTaskTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"distributedDdlTaskTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) DistributedDdlTaskTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"distributedDdlTaskTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) DistributedProductMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributedProductMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) DistributedProductModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributedProductModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) EmptyResultForAggregationByEmptySet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emptyResultForAggregationByEmptySet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) EmptyResultForAggregationByEmptySetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emptyResultForAggregationByEmptySetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) EnableHttpCompression() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHttpCompression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) EnableHttpCompressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHttpCompressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) FallbackToStaleReplicasForDistributedQueries() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fallbackToStaleReplicasForDistributedQueries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) FallbackToStaleReplicasForDistributedQueriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fallbackToStaleReplicasForDistributedQueriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ForceIndexByDate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceIndexByDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ForceIndexByDateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceIndexByDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ForcePrimaryKey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forcePrimaryKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ForcePrimaryKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forcePrimaryKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) GroupByOverflowMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupByOverflowMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) GroupByOverflowModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupByOverflowModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) GroupByTwoLevelThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupByTwoLevelThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) GroupByTwoLevelThresholdBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupByTwoLevelThresholdBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) GroupByTwoLevelThresholdBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupByTwoLevelThresholdBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) GroupByTwoLevelThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupByTwoLevelThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) HttpConnectionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpConnectionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) HttpConnectionTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpConnectionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) HttpHeadersProgressInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpHeadersProgressInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) HttpHeadersProgressIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpHeadersProgressIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) HttpReceiveTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpReceiveTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) HttpReceiveTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpReceiveTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) HttpSendTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpSendTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) HttpSendTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpSendTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) InputFormatDefaultsForOmittedFields() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputFormatDefaultsForOmittedFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) InputFormatDefaultsForOmittedFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputFormatDefaultsForOmittedFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) InputFormatValuesInterpretExpressions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputFormatValuesInterpretExpressions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) InputFormatValuesInterpretExpressionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputFormatValuesInterpretExpressionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) InsertQuorum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"insertQuorum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) InsertQuorumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"insertQuorumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) InsertQuorumTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"insertQuorumTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) InsertQuorumTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"insertQuorumTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) InternalValue() *MdbClickhouseClusterUserSettings {
	var returns *MdbClickhouseClusterUserSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) JoinedSubqueryRequiresAlias() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"joinedSubqueryRequiresAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) JoinedSubqueryRequiresAliasInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"joinedSubqueryRequiresAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) JoinOverflowMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"joinOverflowMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) JoinOverflowModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"joinOverflowModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) JoinUseNulls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"joinUseNulls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) JoinUseNullsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"joinUseNullsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) LowCardinalityAllowInNativeFormat() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lowCardinalityAllowInNativeFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) LowCardinalityAllowInNativeFormatInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lowCardinalityAllowInNativeFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxAstDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAstDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxAstDepthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAstDepthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxAstElements() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAstElements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxAstElementsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAstElementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxBlockSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBlockSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxBlockSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBlockSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxBytesBeforeExternalGroupBy() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBytesBeforeExternalGroupBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxBytesBeforeExternalGroupByInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBytesBeforeExternalGroupByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxBytesBeforeExternalSort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBytesBeforeExternalSort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxBytesBeforeExternalSortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBytesBeforeExternalSortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxBytesInDistinct() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBytesInDistinct",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxBytesInDistinctInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBytesInDistinctInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxBytesInJoin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBytesInJoin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxBytesInJoinInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBytesInJoinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxBytesInSet() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBytesInSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxBytesInSetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBytesInSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxBytesToRead() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBytesToRead",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxBytesToReadInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBytesToReadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxBytesToSort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBytesToSort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxBytesToSortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBytesToSortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxBytesToTransfer() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBytesToTransfer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxBytesToTransferInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBytesToTransferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxColumnsToRead() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxColumnsToRead",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxColumnsToReadInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxColumnsToReadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxExecutionTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxExecutionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxExecutionTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxExecutionTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxExpandedAstElements() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxExpandedAstElements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxExpandedAstElementsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxExpandedAstElementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxInsertBlockSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInsertBlockSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxInsertBlockSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInsertBlockSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxMemoryUsage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMemoryUsage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxMemoryUsageForUser() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMemoryUsageForUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxMemoryUsageForUserInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMemoryUsageForUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxMemoryUsageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMemoryUsageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxNetworkBandwidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNetworkBandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxNetworkBandwidthForUser() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNetworkBandwidthForUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxNetworkBandwidthForUserInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNetworkBandwidthForUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxNetworkBandwidthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNetworkBandwidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxQuerySize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxQuerySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxQuerySizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxQuerySizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxReplicaDelayForDistributedQueries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxReplicaDelayForDistributedQueries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxReplicaDelayForDistributedQueriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxReplicaDelayForDistributedQueriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxResultBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxResultBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxResultBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxResultBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxResultRows() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxResultRows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxResultRowsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxResultRowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxRowsInDistinct() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRowsInDistinct",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxRowsInDistinctInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRowsInDistinctInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxRowsInJoin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRowsInJoin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxRowsInJoinInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRowsInJoinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxRowsInSet() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRowsInSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxRowsInSetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRowsInSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxRowsToGroupBy() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRowsToGroupBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxRowsToGroupByInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRowsToGroupByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxRowsToRead() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRowsToRead",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxRowsToReadInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRowsToReadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxRowsToSort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRowsToSort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxRowsToSortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRowsToSortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxRowsToTransfer() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRowsToTransfer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxRowsToTransferInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRowsToTransferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxTemporaryColumns() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTemporaryColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxTemporaryColumnsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTemporaryColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxTemporaryNonConstColumns() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTemporaryNonConstColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxTemporaryNonConstColumnsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTemporaryNonConstColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxThreads() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxThreads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MaxThreadsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxThreadsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MergeTreeMaxBytesToUseCache() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mergeTreeMaxBytesToUseCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MergeTreeMaxBytesToUseCacheInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mergeTreeMaxBytesToUseCacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MergeTreeMaxRowsToUseCache() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mergeTreeMaxRowsToUseCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MergeTreeMaxRowsToUseCacheInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mergeTreeMaxRowsToUseCacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MergeTreeMinBytesForConcurrentRead() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mergeTreeMinBytesForConcurrentRead",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MergeTreeMinBytesForConcurrentReadInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mergeTreeMinBytesForConcurrentReadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MergeTreeMinRowsForConcurrentRead() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mergeTreeMinRowsForConcurrentRead",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MergeTreeMinRowsForConcurrentReadInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mergeTreeMinRowsForConcurrentReadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MinBytesToUseDirectIo() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minBytesToUseDirectIo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MinBytesToUseDirectIoInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minBytesToUseDirectIoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MinCountToCompile() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCountToCompile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MinCountToCompileExpression() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCountToCompileExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MinCountToCompileExpressionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCountToCompileExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MinCountToCompileInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCountToCompileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MinExecutionSpeed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minExecutionSpeed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MinExecutionSpeedBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minExecutionSpeedBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MinExecutionSpeedBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minExecutionSpeedBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MinExecutionSpeedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minExecutionSpeedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MinInsertBlockSizeBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInsertBlockSizeBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MinInsertBlockSizeBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInsertBlockSizeBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MinInsertBlockSizeRows() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInsertBlockSizeRows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) MinInsertBlockSizeRowsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInsertBlockSizeRowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) OutputFormatJsonQuote64BitIntegers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputFormatJsonQuote64BitIntegers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) OutputFormatJsonQuote64BitIntegersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputFormatJsonQuote64BitIntegersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) OutputFormatJsonQuoteDenormals() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputFormatJsonQuoteDenormals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) OutputFormatJsonQuoteDenormalsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputFormatJsonQuoteDenormalsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) QuotaMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quotaMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) QuotaModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quotaModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) Readonly() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ReadonlyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readonlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ReadOverflowMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readOverflowMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ReadOverflowModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readOverflowModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ReceiveTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"receiveTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ReceiveTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"receiveTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ReplicationAlterPartitionsSync() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicationAlterPartitionsSync",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ReplicationAlterPartitionsSyncInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicationAlterPartitionsSyncInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResultOverflowMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultOverflowMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResultOverflowModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultOverflowModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SelectSequentialConsistency() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selectSequentialConsistency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SelectSequentialConsistencyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selectSequentialConsistencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SendProgressInHttpHeaders() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendProgressInHttpHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SendProgressInHttpHeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendProgressInHttpHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SendTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sendTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SendTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sendTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetOverflowMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"setOverflowMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetOverflowModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"setOverflowModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SkipUnavailableShards() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipUnavailableShards",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SkipUnavailableShardsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipUnavailableShardsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SortOverflowMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortOverflowMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SortOverflowModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortOverflowModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) TimeoutOverflowMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutOverflowMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) TimeoutOverflowModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutOverflowModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) TransferOverflowMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transferOverflowMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) TransferOverflowModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transferOverflowModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) TransformNullIn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transformNullIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) TransformNullInInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transformNullInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) UseUncompressedCache() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useUncompressedCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) UseUncompressedCacheInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useUncompressedCacheInput",
		&returns,
	)
	return returns
}


func NewMdbClickhouseClusterUserSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MdbClickhouseClusterUserSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbClickhouseClusterUserSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMdbClickhouseClusterUserSettingsOutputReference_Override(m MdbClickhouseClusterUserSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbClickhouseClusterUserSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetAddHttpCorsHeader(val interface{}) {
	_jsii_.Set(
		j,
		"addHttpCorsHeader",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetAllowDdl(val interface{}) {
	_jsii_.Set(
		j,
		"allowDdl",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetCompile(val interface{}) {
	_jsii_.Set(
		j,
		"compile",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetCompileExpressions(val interface{}) {
	_jsii_.Set(
		j,
		"compileExpressions",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetConnectTimeout(val *float64) {
	_jsii_.Set(
		j,
		"connectTimeout",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetCountDistinctImplementation(val *string) {
	_jsii_.Set(
		j,
		"countDistinctImplementation",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetDistinctOverflowMode(val *string) {
	_jsii_.Set(
		j,
		"distinctOverflowMode",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetDistributedAggregationMemoryEfficient(val interface{}) {
	_jsii_.Set(
		j,
		"distributedAggregationMemoryEfficient",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetDistributedDdlTaskTimeout(val *float64) {
	_jsii_.Set(
		j,
		"distributedDdlTaskTimeout",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetDistributedProductMode(val *string) {
	_jsii_.Set(
		j,
		"distributedProductMode",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetEmptyResultForAggregationByEmptySet(val interface{}) {
	_jsii_.Set(
		j,
		"emptyResultForAggregationByEmptySet",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetEnableHttpCompression(val interface{}) {
	_jsii_.Set(
		j,
		"enableHttpCompression",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetFallbackToStaleReplicasForDistributedQueries(val interface{}) {
	_jsii_.Set(
		j,
		"fallbackToStaleReplicasForDistributedQueries",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetForceIndexByDate(val interface{}) {
	_jsii_.Set(
		j,
		"forceIndexByDate",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetForcePrimaryKey(val interface{}) {
	_jsii_.Set(
		j,
		"forcePrimaryKey",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetGroupByOverflowMode(val *string) {
	_jsii_.Set(
		j,
		"groupByOverflowMode",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetGroupByTwoLevelThreshold(val *float64) {
	_jsii_.Set(
		j,
		"groupByTwoLevelThreshold",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetGroupByTwoLevelThresholdBytes(val *float64) {
	_jsii_.Set(
		j,
		"groupByTwoLevelThresholdBytes",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetHttpConnectionTimeout(val *float64) {
	_jsii_.Set(
		j,
		"httpConnectionTimeout",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetHttpHeadersProgressInterval(val *float64) {
	_jsii_.Set(
		j,
		"httpHeadersProgressInterval",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetHttpReceiveTimeout(val *float64) {
	_jsii_.Set(
		j,
		"httpReceiveTimeout",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetHttpSendTimeout(val *float64) {
	_jsii_.Set(
		j,
		"httpSendTimeout",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetInputFormatDefaultsForOmittedFields(val interface{}) {
	_jsii_.Set(
		j,
		"inputFormatDefaultsForOmittedFields",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetInputFormatValuesInterpretExpressions(val interface{}) {
	_jsii_.Set(
		j,
		"inputFormatValuesInterpretExpressions",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetInsertQuorum(val *float64) {
	_jsii_.Set(
		j,
		"insertQuorum",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetInsertQuorumTimeout(val *float64) {
	_jsii_.Set(
		j,
		"insertQuorumTimeout",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetInternalValue(val *MdbClickhouseClusterUserSettings) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetJoinedSubqueryRequiresAlias(val interface{}) {
	_jsii_.Set(
		j,
		"joinedSubqueryRequiresAlias",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetJoinOverflowMode(val *string) {
	_jsii_.Set(
		j,
		"joinOverflowMode",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetJoinUseNulls(val interface{}) {
	_jsii_.Set(
		j,
		"joinUseNulls",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetLowCardinalityAllowInNativeFormat(val interface{}) {
	_jsii_.Set(
		j,
		"lowCardinalityAllowInNativeFormat",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxAstDepth(val *float64) {
	_jsii_.Set(
		j,
		"maxAstDepth",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxAstElements(val *float64) {
	_jsii_.Set(
		j,
		"maxAstElements",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxBlockSize(val *float64) {
	_jsii_.Set(
		j,
		"maxBlockSize",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxBytesBeforeExternalGroupBy(val *float64) {
	_jsii_.Set(
		j,
		"maxBytesBeforeExternalGroupBy",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxBytesBeforeExternalSort(val *float64) {
	_jsii_.Set(
		j,
		"maxBytesBeforeExternalSort",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxBytesInDistinct(val *float64) {
	_jsii_.Set(
		j,
		"maxBytesInDistinct",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxBytesInJoin(val *float64) {
	_jsii_.Set(
		j,
		"maxBytesInJoin",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxBytesInSet(val *float64) {
	_jsii_.Set(
		j,
		"maxBytesInSet",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxBytesToRead(val *float64) {
	_jsii_.Set(
		j,
		"maxBytesToRead",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxBytesToSort(val *float64) {
	_jsii_.Set(
		j,
		"maxBytesToSort",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxBytesToTransfer(val *float64) {
	_jsii_.Set(
		j,
		"maxBytesToTransfer",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxColumnsToRead(val *float64) {
	_jsii_.Set(
		j,
		"maxColumnsToRead",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxExecutionTime(val *float64) {
	_jsii_.Set(
		j,
		"maxExecutionTime",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxExpandedAstElements(val *float64) {
	_jsii_.Set(
		j,
		"maxExpandedAstElements",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxInsertBlockSize(val *float64) {
	_jsii_.Set(
		j,
		"maxInsertBlockSize",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxMemoryUsage(val *float64) {
	_jsii_.Set(
		j,
		"maxMemoryUsage",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxMemoryUsageForUser(val *float64) {
	_jsii_.Set(
		j,
		"maxMemoryUsageForUser",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxNetworkBandwidth(val *float64) {
	_jsii_.Set(
		j,
		"maxNetworkBandwidth",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxNetworkBandwidthForUser(val *float64) {
	_jsii_.Set(
		j,
		"maxNetworkBandwidthForUser",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxQuerySize(val *float64) {
	_jsii_.Set(
		j,
		"maxQuerySize",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxReplicaDelayForDistributedQueries(val *float64) {
	_jsii_.Set(
		j,
		"maxReplicaDelayForDistributedQueries",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxResultBytes(val *float64) {
	_jsii_.Set(
		j,
		"maxResultBytes",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxResultRows(val *float64) {
	_jsii_.Set(
		j,
		"maxResultRows",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxRowsInDistinct(val *float64) {
	_jsii_.Set(
		j,
		"maxRowsInDistinct",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxRowsInJoin(val *float64) {
	_jsii_.Set(
		j,
		"maxRowsInJoin",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxRowsInSet(val *float64) {
	_jsii_.Set(
		j,
		"maxRowsInSet",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxRowsToGroupBy(val *float64) {
	_jsii_.Set(
		j,
		"maxRowsToGroupBy",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxRowsToRead(val *float64) {
	_jsii_.Set(
		j,
		"maxRowsToRead",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxRowsToSort(val *float64) {
	_jsii_.Set(
		j,
		"maxRowsToSort",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxRowsToTransfer(val *float64) {
	_jsii_.Set(
		j,
		"maxRowsToTransfer",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxTemporaryColumns(val *float64) {
	_jsii_.Set(
		j,
		"maxTemporaryColumns",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxTemporaryNonConstColumns(val *float64) {
	_jsii_.Set(
		j,
		"maxTemporaryNonConstColumns",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMaxThreads(val *float64) {
	_jsii_.Set(
		j,
		"maxThreads",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMergeTreeMaxBytesToUseCache(val *float64) {
	_jsii_.Set(
		j,
		"mergeTreeMaxBytesToUseCache",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMergeTreeMaxRowsToUseCache(val *float64) {
	_jsii_.Set(
		j,
		"mergeTreeMaxRowsToUseCache",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMergeTreeMinBytesForConcurrentRead(val *float64) {
	_jsii_.Set(
		j,
		"mergeTreeMinBytesForConcurrentRead",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMergeTreeMinRowsForConcurrentRead(val *float64) {
	_jsii_.Set(
		j,
		"mergeTreeMinRowsForConcurrentRead",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMinBytesToUseDirectIo(val *float64) {
	_jsii_.Set(
		j,
		"minBytesToUseDirectIo",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMinCountToCompile(val *float64) {
	_jsii_.Set(
		j,
		"minCountToCompile",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMinCountToCompileExpression(val *float64) {
	_jsii_.Set(
		j,
		"minCountToCompileExpression",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMinExecutionSpeed(val *float64) {
	_jsii_.Set(
		j,
		"minExecutionSpeed",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMinExecutionSpeedBytes(val *float64) {
	_jsii_.Set(
		j,
		"minExecutionSpeedBytes",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMinInsertBlockSizeBytes(val *float64) {
	_jsii_.Set(
		j,
		"minInsertBlockSizeBytes",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetMinInsertBlockSizeRows(val *float64) {
	_jsii_.Set(
		j,
		"minInsertBlockSizeRows",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetOutputFormatJsonQuote64BitIntegers(val interface{}) {
	_jsii_.Set(
		j,
		"outputFormatJsonQuote64BitIntegers",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetOutputFormatJsonQuoteDenormals(val interface{}) {
	_jsii_.Set(
		j,
		"outputFormatJsonQuoteDenormals",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetPriority(val *float64) {
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetQuotaMode(val *string) {
	_jsii_.Set(
		j,
		"quotaMode",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetReadonly(val *float64) {
	_jsii_.Set(
		j,
		"readonly",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetReadOverflowMode(val *string) {
	_jsii_.Set(
		j,
		"readOverflowMode",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetReceiveTimeout(val *float64) {
	_jsii_.Set(
		j,
		"receiveTimeout",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetReplicationAlterPartitionsSync(val *float64) {
	_jsii_.Set(
		j,
		"replicationAlterPartitionsSync",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetResultOverflowMode(val *string) {
	_jsii_.Set(
		j,
		"resultOverflowMode",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetSelectSequentialConsistency(val interface{}) {
	_jsii_.Set(
		j,
		"selectSequentialConsistency",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetSendProgressInHttpHeaders(val interface{}) {
	_jsii_.Set(
		j,
		"sendProgressInHttpHeaders",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetSendTimeout(val *float64) {
	_jsii_.Set(
		j,
		"sendTimeout",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetSetOverflowMode(val *string) {
	_jsii_.Set(
		j,
		"setOverflowMode",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetSkipUnavailableShards(val interface{}) {
	_jsii_.Set(
		j,
		"skipUnavailableShards",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetSortOverflowMode(val *string) {
	_jsii_.Set(
		j,
		"sortOverflowMode",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetTimeoutOverflowMode(val *string) {
	_jsii_.Set(
		j,
		"timeoutOverflowMode",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetTransferOverflowMode(val *string) {
	_jsii_.Set(
		j,
		"transferOverflowMode",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetTransformNullIn(val interface{}) {
	_jsii_.Set(
		j,
		"transformNullIn",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) SetUseUncompressedCache(val interface{}) {
	_jsii_.Set(
		j,
		"useUncompressedCache",
		val,
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetAddHttpCorsHeader() {
	_jsii_.InvokeVoid(
		m,
		"resetAddHttpCorsHeader",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetAllowDdl() {
	_jsii_.InvokeVoid(
		m,
		"resetAllowDdl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetCompile() {
	_jsii_.InvokeVoid(
		m,
		"resetCompile",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetCompileExpressions() {
	_jsii_.InvokeVoid(
		m,
		"resetCompileExpressions",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetConnectTimeout() {
	_jsii_.InvokeVoid(
		m,
		"resetConnectTimeout",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetCountDistinctImplementation() {
	_jsii_.InvokeVoid(
		m,
		"resetCountDistinctImplementation",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetDistinctOverflowMode() {
	_jsii_.InvokeVoid(
		m,
		"resetDistinctOverflowMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetDistributedAggregationMemoryEfficient() {
	_jsii_.InvokeVoid(
		m,
		"resetDistributedAggregationMemoryEfficient",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetDistributedDdlTaskTimeout() {
	_jsii_.InvokeVoid(
		m,
		"resetDistributedDdlTaskTimeout",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetDistributedProductMode() {
	_jsii_.InvokeVoid(
		m,
		"resetDistributedProductMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetEmptyResultForAggregationByEmptySet() {
	_jsii_.InvokeVoid(
		m,
		"resetEmptyResultForAggregationByEmptySet",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetEnableHttpCompression() {
	_jsii_.InvokeVoid(
		m,
		"resetEnableHttpCompression",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetFallbackToStaleReplicasForDistributedQueries() {
	_jsii_.InvokeVoid(
		m,
		"resetFallbackToStaleReplicasForDistributedQueries",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetForceIndexByDate() {
	_jsii_.InvokeVoid(
		m,
		"resetForceIndexByDate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetForcePrimaryKey() {
	_jsii_.InvokeVoid(
		m,
		"resetForcePrimaryKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetGroupByOverflowMode() {
	_jsii_.InvokeVoid(
		m,
		"resetGroupByOverflowMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetGroupByTwoLevelThreshold() {
	_jsii_.InvokeVoid(
		m,
		"resetGroupByTwoLevelThreshold",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetGroupByTwoLevelThresholdBytes() {
	_jsii_.InvokeVoid(
		m,
		"resetGroupByTwoLevelThresholdBytes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetHttpConnectionTimeout() {
	_jsii_.InvokeVoid(
		m,
		"resetHttpConnectionTimeout",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetHttpHeadersProgressInterval() {
	_jsii_.InvokeVoid(
		m,
		"resetHttpHeadersProgressInterval",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetHttpReceiveTimeout() {
	_jsii_.InvokeVoid(
		m,
		"resetHttpReceiveTimeout",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetHttpSendTimeout() {
	_jsii_.InvokeVoid(
		m,
		"resetHttpSendTimeout",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetInputFormatDefaultsForOmittedFields() {
	_jsii_.InvokeVoid(
		m,
		"resetInputFormatDefaultsForOmittedFields",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetInputFormatValuesInterpretExpressions() {
	_jsii_.InvokeVoid(
		m,
		"resetInputFormatValuesInterpretExpressions",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetInsertQuorum() {
	_jsii_.InvokeVoid(
		m,
		"resetInsertQuorum",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetInsertQuorumTimeout() {
	_jsii_.InvokeVoid(
		m,
		"resetInsertQuorumTimeout",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetJoinedSubqueryRequiresAlias() {
	_jsii_.InvokeVoid(
		m,
		"resetJoinedSubqueryRequiresAlias",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetJoinOverflowMode() {
	_jsii_.InvokeVoid(
		m,
		"resetJoinOverflowMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetJoinUseNulls() {
	_jsii_.InvokeVoid(
		m,
		"resetJoinUseNulls",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetLowCardinalityAllowInNativeFormat() {
	_jsii_.InvokeVoid(
		m,
		"resetLowCardinalityAllowInNativeFormat",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxAstDepth() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxAstDepth",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxAstElements() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxAstElements",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxBlockSize() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxBlockSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxBytesBeforeExternalGroupBy() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxBytesBeforeExternalGroupBy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxBytesBeforeExternalSort() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxBytesBeforeExternalSort",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxBytesInDistinct() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxBytesInDistinct",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxBytesInJoin() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxBytesInJoin",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxBytesInSet() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxBytesInSet",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxBytesToRead() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxBytesToRead",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxBytesToSort() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxBytesToSort",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxBytesToTransfer() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxBytesToTransfer",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxColumnsToRead() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxColumnsToRead",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxExecutionTime() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxExecutionTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxExpandedAstElements() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxExpandedAstElements",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxInsertBlockSize() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxInsertBlockSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxMemoryUsage() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxMemoryUsage",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxMemoryUsageForUser() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxMemoryUsageForUser",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxNetworkBandwidth() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxNetworkBandwidth",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxNetworkBandwidthForUser() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxNetworkBandwidthForUser",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxQuerySize() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxQuerySize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxReplicaDelayForDistributedQueries() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxReplicaDelayForDistributedQueries",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxResultBytes() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxResultBytes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxResultRows() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxResultRows",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxRowsInDistinct() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxRowsInDistinct",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxRowsInJoin() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxRowsInJoin",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxRowsInSet() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxRowsInSet",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxRowsToGroupBy() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxRowsToGroupBy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxRowsToRead() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxRowsToRead",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxRowsToSort() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxRowsToSort",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxRowsToTransfer() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxRowsToTransfer",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxTemporaryColumns() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxTemporaryColumns",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxTemporaryNonConstColumns() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxTemporaryNonConstColumns",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMaxThreads() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxThreads",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMergeTreeMaxBytesToUseCache() {
	_jsii_.InvokeVoid(
		m,
		"resetMergeTreeMaxBytesToUseCache",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMergeTreeMaxRowsToUseCache() {
	_jsii_.InvokeVoid(
		m,
		"resetMergeTreeMaxRowsToUseCache",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMergeTreeMinBytesForConcurrentRead() {
	_jsii_.InvokeVoid(
		m,
		"resetMergeTreeMinBytesForConcurrentRead",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMergeTreeMinRowsForConcurrentRead() {
	_jsii_.InvokeVoid(
		m,
		"resetMergeTreeMinRowsForConcurrentRead",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMinBytesToUseDirectIo() {
	_jsii_.InvokeVoid(
		m,
		"resetMinBytesToUseDirectIo",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMinCountToCompile() {
	_jsii_.InvokeVoid(
		m,
		"resetMinCountToCompile",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMinCountToCompileExpression() {
	_jsii_.InvokeVoid(
		m,
		"resetMinCountToCompileExpression",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMinExecutionSpeed() {
	_jsii_.InvokeVoid(
		m,
		"resetMinExecutionSpeed",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMinExecutionSpeedBytes() {
	_jsii_.InvokeVoid(
		m,
		"resetMinExecutionSpeedBytes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMinInsertBlockSizeBytes() {
	_jsii_.InvokeVoid(
		m,
		"resetMinInsertBlockSizeBytes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetMinInsertBlockSizeRows() {
	_jsii_.InvokeVoid(
		m,
		"resetMinInsertBlockSizeRows",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetOutputFormatJsonQuote64BitIntegers() {
	_jsii_.InvokeVoid(
		m,
		"resetOutputFormatJsonQuote64BitIntegers",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetOutputFormatJsonQuoteDenormals() {
	_jsii_.InvokeVoid(
		m,
		"resetOutputFormatJsonQuoteDenormals",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		m,
		"resetPriority",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetQuotaMode() {
	_jsii_.InvokeVoid(
		m,
		"resetQuotaMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetReadonly() {
	_jsii_.InvokeVoid(
		m,
		"resetReadonly",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetReadOverflowMode() {
	_jsii_.InvokeVoid(
		m,
		"resetReadOverflowMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetReceiveTimeout() {
	_jsii_.InvokeVoid(
		m,
		"resetReceiveTimeout",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetReplicationAlterPartitionsSync() {
	_jsii_.InvokeVoid(
		m,
		"resetReplicationAlterPartitionsSync",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetResultOverflowMode() {
	_jsii_.InvokeVoid(
		m,
		"resetResultOverflowMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetSelectSequentialConsistency() {
	_jsii_.InvokeVoid(
		m,
		"resetSelectSequentialConsistency",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetSendProgressInHttpHeaders() {
	_jsii_.InvokeVoid(
		m,
		"resetSendProgressInHttpHeaders",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetSendTimeout() {
	_jsii_.InvokeVoid(
		m,
		"resetSendTimeout",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetSetOverflowMode() {
	_jsii_.InvokeVoid(
		m,
		"resetSetOverflowMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetSkipUnavailableShards() {
	_jsii_.InvokeVoid(
		m,
		"resetSkipUnavailableShards",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetSortOverflowMode() {
	_jsii_.InvokeVoid(
		m,
		"resetSortOverflowMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetTimeoutOverflowMode() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeoutOverflowMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetTransferOverflowMode() {
	_jsii_.InvokeVoid(
		m,
		"resetTransferOverflowMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetTransformNullIn() {
	_jsii_.InvokeVoid(
		m,
		"resetTransformNullIn",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ResetUseUncompressedCache() {
	_jsii_.InvokeVoid(
		m,
		"resetUseUncompressedCache",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterUserSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

