// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexMdbClickhouseClusterUserSettingsOutputReference interface {
	cdktf.ComplexObject
	AddHttpCorsHeader() cdktf.IResolvable
	AllowDdl() cdktf.IResolvable
	Compile() cdktf.IResolvable
	CompileExpressions() cdktf.IResolvable
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
	CountDistinctImplementation() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DistinctOverflowMode() *string
	DistributedAggregationMemoryEfficient() cdktf.IResolvable
	DistributedDdlTaskTimeout() *float64
	DistributedProductMode() *string
	EmptyResultForAggregationByEmptySet() cdktf.IResolvable
	EnableHttpCompression() cdktf.IResolvable
	FallbackToStaleReplicasForDistributedQueries() cdktf.IResolvable
	ForceIndexByDate() cdktf.IResolvable
	ForcePrimaryKey() cdktf.IResolvable
	// Experimental.
	Fqn() *string
	GroupByOverflowMode() *string
	GroupByTwoLevelThreshold() *float64
	GroupByTwoLevelThresholdBytes() *float64
	HttpConnectionTimeout() *float64
	HttpHeadersProgressInterval() *float64
	HttpReceiveTimeout() *float64
	HttpSendTimeout() *float64
	InputFormatDefaultsForOmittedFields() cdktf.IResolvable
	InputFormatValuesInterpretExpressions() cdktf.IResolvable
	InsertQuorum() *float64
	InsertQuorumTimeout() *float64
	InternalValue() *DataYandexMdbClickhouseClusterUserSettings
	SetInternalValue(val *DataYandexMdbClickhouseClusterUserSettings)
	JoinedSubqueryRequiresAlias() cdktf.IResolvable
	JoinOverflowMode() *string
	JoinUseNulls() cdktf.IResolvable
	LowCardinalityAllowInNativeFormat() cdktf.IResolvable
	MaxAstDepth() *float64
	MaxAstElements() *float64
	MaxBlockSize() *float64
	MaxBytesBeforeExternalGroupBy() *float64
	MaxBytesBeforeExternalSort() *float64
	MaxBytesInDistinct() *float64
	MaxBytesInJoin() *float64
	MaxBytesInSet() *float64
	MaxBytesToRead() *float64
	MaxBytesToSort() *float64
	MaxBytesToTransfer() *float64
	MaxColumnsToRead() *float64
	MaxExecutionTime() *float64
	MaxExpandedAstElements() *float64
	MaxInsertBlockSize() *float64
	MaxMemoryUsage() *float64
	MaxMemoryUsageForUser() *float64
	MaxNetworkBandwidth() *float64
	MaxNetworkBandwidthForUser() *float64
	MaxQuerySize() *float64
	MaxReplicaDelayForDistributedQueries() *float64
	MaxResultBytes() *float64
	MaxResultRows() *float64
	MaxRowsInDistinct() *float64
	MaxRowsInJoin() *float64
	MaxRowsInSet() *float64
	MaxRowsToGroupBy() *float64
	MaxRowsToRead() *float64
	MaxRowsToSort() *float64
	MaxRowsToTransfer() *float64
	MaxTemporaryColumns() *float64
	MaxTemporaryNonConstColumns() *float64
	MaxThreads() *float64
	MergeTreeMaxBytesToUseCache() *float64
	MergeTreeMaxRowsToUseCache() *float64
	MergeTreeMinBytesForConcurrentRead() *float64
	MergeTreeMinRowsForConcurrentRead() *float64
	MinBytesToUseDirectIo() *float64
	MinCountToCompile() *float64
	MinCountToCompileExpression() *float64
	MinExecutionSpeed() *float64
	MinExecutionSpeedBytes() *float64
	MinInsertBlockSizeBytes() *float64
	MinInsertBlockSizeRows() *float64
	OutputFormatJsonQuote64BitIntegers() cdktf.IResolvable
	OutputFormatJsonQuoteDenormals() cdktf.IResolvable
	Priority() *float64
	QuotaMode() *string
	Readonly() *float64
	ReadOverflowMode() *string
	ReceiveTimeout() *float64
	ReplicationAlterPartitionsSync() *float64
	ResultOverflowMode() *string
	SelectSequentialConsistency() cdktf.IResolvable
	SendProgressInHttpHeaders() cdktf.IResolvable
	SendTimeout() *float64
	SetOverflowMode() *string
	SkipUnavailableShards() cdktf.IResolvable
	SortOverflowMode() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeoutOverflowMode() *string
	TransferOverflowMode() *string
	TransformNullIn() cdktf.IResolvable
	UseUncompressedCache() cdktf.IResolvable
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataYandexMdbClickhouseClusterUserSettingsOutputReference
type jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) AddHttpCorsHeader() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"addHttpCorsHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) AllowDdl() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"allowDdl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) Compile() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"compile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) CompileExpressions() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"compileExpressions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) ConnectTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) CountDistinctImplementation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countDistinctImplementation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) DistinctOverflowMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distinctOverflowMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) DistributedAggregationMemoryEfficient() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"distributedAggregationMemoryEfficient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) DistributedDdlTaskTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"distributedDdlTaskTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) DistributedProductMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributedProductMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) EmptyResultForAggregationByEmptySet() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"emptyResultForAggregationByEmptySet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) EnableHttpCompression() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableHttpCompression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) FallbackToStaleReplicasForDistributedQueries() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"fallbackToStaleReplicasForDistributedQueries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) ForceIndexByDate() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"forceIndexByDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) ForcePrimaryKey() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"forcePrimaryKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) GroupByOverflowMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupByOverflowMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) GroupByTwoLevelThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupByTwoLevelThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) GroupByTwoLevelThresholdBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupByTwoLevelThresholdBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) HttpConnectionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpConnectionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) HttpHeadersProgressInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpHeadersProgressInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) HttpReceiveTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpReceiveTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) HttpSendTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpSendTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) InputFormatDefaultsForOmittedFields() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"inputFormatDefaultsForOmittedFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) InputFormatValuesInterpretExpressions() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"inputFormatValuesInterpretExpressions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) InsertQuorum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"insertQuorum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) InsertQuorumTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"insertQuorumTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) InternalValue() *DataYandexMdbClickhouseClusterUserSettings {
	var returns *DataYandexMdbClickhouseClusterUserSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) JoinedSubqueryRequiresAlias() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"joinedSubqueryRequiresAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) JoinOverflowMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"joinOverflowMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) JoinUseNulls() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"joinUseNulls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) LowCardinalityAllowInNativeFormat() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"lowCardinalityAllowInNativeFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxAstDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAstDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxAstElements() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAstElements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxBlockSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBlockSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxBytesBeforeExternalGroupBy() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBytesBeforeExternalGroupBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxBytesBeforeExternalSort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBytesBeforeExternalSort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxBytesInDistinct() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBytesInDistinct",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxBytesInJoin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBytesInJoin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxBytesInSet() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBytesInSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxBytesToRead() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBytesToRead",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxBytesToSort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBytesToSort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxBytesToTransfer() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBytesToTransfer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxColumnsToRead() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxColumnsToRead",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxExecutionTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxExecutionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxExpandedAstElements() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxExpandedAstElements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxInsertBlockSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInsertBlockSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxMemoryUsage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMemoryUsage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxMemoryUsageForUser() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMemoryUsageForUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxNetworkBandwidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNetworkBandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxNetworkBandwidthForUser() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNetworkBandwidthForUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxQuerySize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxQuerySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxReplicaDelayForDistributedQueries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxReplicaDelayForDistributedQueries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxResultBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxResultBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxResultRows() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxResultRows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxRowsInDistinct() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRowsInDistinct",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxRowsInJoin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRowsInJoin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxRowsInSet() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRowsInSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxRowsToGroupBy() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRowsToGroupBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxRowsToRead() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRowsToRead",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxRowsToSort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRowsToSort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxRowsToTransfer() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRowsToTransfer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxTemporaryColumns() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTemporaryColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxTemporaryNonConstColumns() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTemporaryNonConstColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MaxThreads() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxThreads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MergeTreeMaxBytesToUseCache() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mergeTreeMaxBytesToUseCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MergeTreeMaxRowsToUseCache() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mergeTreeMaxRowsToUseCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MergeTreeMinBytesForConcurrentRead() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mergeTreeMinBytesForConcurrentRead",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MergeTreeMinRowsForConcurrentRead() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mergeTreeMinRowsForConcurrentRead",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MinBytesToUseDirectIo() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minBytesToUseDirectIo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MinCountToCompile() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCountToCompile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MinCountToCompileExpression() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCountToCompileExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MinExecutionSpeed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minExecutionSpeed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MinExecutionSpeedBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minExecutionSpeedBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MinInsertBlockSizeBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInsertBlockSizeBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) MinInsertBlockSizeRows() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInsertBlockSizeRows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) OutputFormatJsonQuote64BitIntegers() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"outputFormatJsonQuote64BitIntegers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) OutputFormatJsonQuoteDenormals() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"outputFormatJsonQuoteDenormals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) QuotaMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quotaMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) Readonly() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) ReadOverflowMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readOverflowMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) ReceiveTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"receiveTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) ReplicationAlterPartitionsSync() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicationAlterPartitionsSync",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) ResultOverflowMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultOverflowMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) SelectSequentialConsistency() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"selectSequentialConsistency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) SendProgressInHttpHeaders() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"sendProgressInHttpHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) SendTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sendTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) SetOverflowMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"setOverflowMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) SkipUnavailableShards() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"skipUnavailableShards",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) SortOverflowMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortOverflowMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) TimeoutOverflowMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutOverflowMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) TransferOverflowMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transferOverflowMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) TransformNullIn() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"transformNullIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) UseUncompressedCache() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"useUncompressedCache",
		&returns,
	)
	return returns
}


func NewDataYandexMdbClickhouseClusterUserSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataYandexMdbClickhouseClusterUserSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexMdbClickhouseClusterUserSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataYandexMdbClickhouseClusterUserSettingsOutputReference_Override(d DataYandexMdbClickhouseClusterUserSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexMdbClickhouseClusterUserSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) SetInternalValue(val *DataYandexMdbClickhouseClusterUserSettings) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseClusterUserSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

