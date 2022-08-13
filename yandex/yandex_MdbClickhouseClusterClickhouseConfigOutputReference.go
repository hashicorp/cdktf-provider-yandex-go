// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbClickhouseClusterClickhouseConfigOutputReference interface {
	cdktf.ComplexObject
	BackgroundPoolSize() *float64
	SetBackgroundPoolSize(val *float64)
	BackgroundPoolSizeInput() *float64
	BackgroundSchedulePoolSize() *float64
	SetBackgroundSchedulePoolSize(val *float64)
	BackgroundSchedulePoolSizeInput() *float64
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
	Compression() MdbClickhouseClusterClickhouseConfigCompressionList
	CompressionInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GeobaseUri() *string
	SetGeobaseUri(val *string)
	GeobaseUriInput() *string
	GraphiteRollup() MdbClickhouseClusterClickhouseConfigGraphiteRollupList
	GraphiteRollupInput() interface{}
	InternalValue() *MdbClickhouseClusterClickhouseConfig
	SetInternalValue(val *MdbClickhouseClusterClickhouseConfig)
	Kafka() MdbClickhouseClusterClickhouseConfigKafkaOutputReference
	KafkaInput() *MdbClickhouseClusterClickhouseConfigKafka
	KafkaTopic() MdbClickhouseClusterClickhouseConfigKafkaTopicList
	KafkaTopicInput() interface{}
	KeepAliveTimeout() *float64
	SetKeepAliveTimeout(val *float64)
	KeepAliveTimeoutInput() *float64
	LogLevel() *string
	SetLogLevel(val *string)
	LogLevelInput() *string
	MarkCacheSize() *float64
	SetMarkCacheSize(val *float64)
	MarkCacheSizeInput() *float64
	MaxConcurrentQueries() *float64
	SetMaxConcurrentQueries(val *float64)
	MaxConcurrentQueriesInput() *float64
	MaxConnections() *float64
	SetMaxConnections(val *float64)
	MaxConnectionsInput() *float64
	MaxPartitionSizeToDrop() *float64
	SetMaxPartitionSizeToDrop(val *float64)
	MaxPartitionSizeToDropInput() *float64
	MaxTableSizeToDrop() *float64
	SetMaxTableSizeToDrop(val *float64)
	MaxTableSizeToDropInput() *float64
	MergeTree() MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference
	MergeTreeInput() *MdbClickhouseClusterClickhouseConfigMergeTree
	MetricLogEnabled() interface{}
	SetMetricLogEnabled(val interface{})
	MetricLogEnabledInput() interface{}
	MetricLogRetentionSize() *float64
	SetMetricLogRetentionSize(val *float64)
	MetricLogRetentionSizeInput() *float64
	MetricLogRetentionTime() *float64
	SetMetricLogRetentionTime(val *float64)
	MetricLogRetentionTimeInput() *float64
	PartLogRetentionSize() *float64
	SetPartLogRetentionSize(val *float64)
	PartLogRetentionSizeInput() *float64
	PartLogRetentionTime() *float64
	SetPartLogRetentionTime(val *float64)
	PartLogRetentionTimeInput() *float64
	QueryLogRetentionSize() *float64
	SetQueryLogRetentionSize(val *float64)
	QueryLogRetentionSizeInput() *float64
	QueryLogRetentionTime() *float64
	SetQueryLogRetentionTime(val *float64)
	QueryLogRetentionTimeInput() *float64
	QueryThreadLogEnabled() interface{}
	SetQueryThreadLogEnabled(val interface{})
	QueryThreadLogEnabledInput() interface{}
	QueryThreadLogRetentionSize() *float64
	SetQueryThreadLogRetentionSize(val *float64)
	QueryThreadLogRetentionSizeInput() *float64
	QueryThreadLogRetentionTime() *float64
	SetQueryThreadLogRetentionTime(val *float64)
	QueryThreadLogRetentionTimeInput() *float64
	Rabbitmq() MdbClickhouseClusterClickhouseConfigRabbitmqOutputReference
	RabbitmqInput() *MdbClickhouseClusterClickhouseConfigRabbitmq
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TextLogEnabled() interface{}
	SetTextLogEnabled(val interface{})
	TextLogEnabledInput() interface{}
	TextLogLevel() *string
	SetTextLogLevel(val *string)
	TextLogLevelInput() *string
	TextLogRetentionSize() *float64
	SetTextLogRetentionSize(val *float64)
	TextLogRetentionSizeInput() *float64
	TextLogRetentionTime() *float64
	SetTextLogRetentionTime(val *float64)
	TextLogRetentionTimeInput() *float64
	Timezone() *string
	SetTimezone(val *string)
	TimezoneInput() *string
	TraceLogEnabled() interface{}
	SetTraceLogEnabled(val interface{})
	TraceLogEnabledInput() interface{}
	TraceLogRetentionSize() *float64
	SetTraceLogRetentionSize(val *float64)
	TraceLogRetentionSizeInput() *float64
	TraceLogRetentionTime() *float64
	SetTraceLogRetentionTime(val *float64)
	TraceLogRetentionTimeInput() *float64
	UncompressedCacheSize() *float64
	SetUncompressedCacheSize(val *float64)
	UncompressedCacheSizeInput() *float64
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
	PutCompression(value interface{})
	PutGraphiteRollup(value interface{})
	PutKafka(value *MdbClickhouseClusterClickhouseConfigKafka)
	PutKafkaTopic(value interface{})
	PutMergeTree(value *MdbClickhouseClusterClickhouseConfigMergeTree)
	PutRabbitmq(value *MdbClickhouseClusterClickhouseConfigRabbitmq)
	ResetBackgroundPoolSize()
	ResetBackgroundSchedulePoolSize()
	ResetCompression()
	ResetGeobaseUri()
	ResetGraphiteRollup()
	ResetKafka()
	ResetKafkaTopic()
	ResetKeepAliveTimeout()
	ResetLogLevel()
	ResetMarkCacheSize()
	ResetMaxConcurrentQueries()
	ResetMaxConnections()
	ResetMaxPartitionSizeToDrop()
	ResetMaxTableSizeToDrop()
	ResetMergeTree()
	ResetMetricLogEnabled()
	ResetMetricLogRetentionSize()
	ResetMetricLogRetentionTime()
	ResetPartLogRetentionSize()
	ResetPartLogRetentionTime()
	ResetQueryLogRetentionSize()
	ResetQueryLogRetentionTime()
	ResetQueryThreadLogEnabled()
	ResetQueryThreadLogRetentionSize()
	ResetQueryThreadLogRetentionTime()
	ResetRabbitmq()
	ResetTextLogEnabled()
	ResetTextLogLevel()
	ResetTextLogRetentionSize()
	ResetTextLogRetentionTime()
	ResetTimezone()
	ResetTraceLogEnabled()
	ResetTraceLogRetentionSize()
	ResetTraceLogRetentionTime()
	ResetUncompressedCacheSize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MdbClickhouseClusterClickhouseConfigOutputReference
type jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) BackgroundPoolSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backgroundPoolSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) BackgroundPoolSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backgroundPoolSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) BackgroundSchedulePoolSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backgroundSchedulePoolSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) BackgroundSchedulePoolSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backgroundSchedulePoolSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) Compression() MdbClickhouseClusterClickhouseConfigCompressionList {
	var returns MdbClickhouseClusterClickhouseConfigCompressionList
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) CompressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) GeobaseUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geobaseUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) GeobaseUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geobaseUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) GraphiteRollup() MdbClickhouseClusterClickhouseConfigGraphiteRollupList {
	var returns MdbClickhouseClusterClickhouseConfigGraphiteRollupList
	_jsii_.Get(
		j,
		"graphiteRollup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) GraphiteRollupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"graphiteRollupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) InternalValue() *MdbClickhouseClusterClickhouseConfig {
	var returns *MdbClickhouseClusterClickhouseConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) Kafka() MdbClickhouseClusterClickhouseConfigKafkaOutputReference {
	var returns MdbClickhouseClusterClickhouseConfigKafkaOutputReference
	_jsii_.Get(
		j,
		"kafka",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) KafkaInput() *MdbClickhouseClusterClickhouseConfigKafka {
	var returns *MdbClickhouseClusterClickhouseConfigKafka
	_jsii_.Get(
		j,
		"kafkaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) KafkaTopic() MdbClickhouseClusterClickhouseConfigKafkaTopicList {
	var returns MdbClickhouseClusterClickhouseConfigKafkaTopicList
	_jsii_.Get(
		j,
		"kafkaTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) KafkaTopicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kafkaTopicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) KeepAliveTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepAliveTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) KeepAliveTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepAliveTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) MarkCacheSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"markCacheSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) MarkCacheSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"markCacheSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) MaxConcurrentQueries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentQueries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) MaxConcurrentQueriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentQueriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) MaxConnections() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) MaxConnectionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) MaxPartitionSizeToDrop() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPartitionSizeToDrop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) MaxPartitionSizeToDropInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPartitionSizeToDropInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) MaxTableSizeToDrop() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTableSizeToDrop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) MaxTableSizeToDropInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTableSizeToDropInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) MergeTree() MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference {
	var returns MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference
	_jsii_.Get(
		j,
		"mergeTree",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) MergeTreeInput() *MdbClickhouseClusterClickhouseConfigMergeTree {
	var returns *MdbClickhouseClusterClickhouseConfigMergeTree
	_jsii_.Get(
		j,
		"mergeTreeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) MetricLogEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricLogEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) MetricLogEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricLogEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) MetricLogRetentionSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricLogRetentionSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) MetricLogRetentionSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricLogRetentionSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) MetricLogRetentionTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricLogRetentionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) MetricLogRetentionTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricLogRetentionTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) PartLogRetentionSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partLogRetentionSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) PartLogRetentionSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partLogRetentionSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) PartLogRetentionTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partLogRetentionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) PartLogRetentionTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partLogRetentionTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) QueryLogRetentionSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryLogRetentionSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) QueryLogRetentionSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryLogRetentionSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) QueryLogRetentionTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryLogRetentionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) QueryLogRetentionTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryLogRetentionTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) QueryThreadLogEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryThreadLogEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) QueryThreadLogEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryThreadLogEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) QueryThreadLogRetentionSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryThreadLogRetentionSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) QueryThreadLogRetentionSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryThreadLogRetentionSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) QueryThreadLogRetentionTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryThreadLogRetentionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) QueryThreadLogRetentionTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryThreadLogRetentionTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) Rabbitmq() MdbClickhouseClusterClickhouseConfigRabbitmqOutputReference {
	var returns MdbClickhouseClusterClickhouseConfigRabbitmqOutputReference
	_jsii_.Get(
		j,
		"rabbitmq",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) RabbitmqInput() *MdbClickhouseClusterClickhouseConfigRabbitmq {
	var returns *MdbClickhouseClusterClickhouseConfigRabbitmq
	_jsii_.Get(
		j,
		"rabbitmqInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) TextLogEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"textLogEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) TextLogEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"textLogEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) TextLogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textLogLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) TextLogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textLogLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) TextLogRetentionSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"textLogRetentionSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) TextLogRetentionSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"textLogRetentionSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) TextLogRetentionTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"textLogRetentionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) TextLogRetentionTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"textLogRetentionTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) TraceLogEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"traceLogEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) TraceLogEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"traceLogEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) TraceLogRetentionSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"traceLogRetentionSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) TraceLogRetentionSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"traceLogRetentionSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) TraceLogRetentionTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"traceLogRetentionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) TraceLogRetentionTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"traceLogRetentionTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) UncompressedCacheSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uncompressedCacheSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) UncompressedCacheSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uncompressedCacheSizeInput",
		&returns,
	)
	return returns
}


func NewMdbClickhouseClusterClickhouseConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MdbClickhouseClusterClickhouseConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbClickhouseClusterClickhouseConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMdbClickhouseClusterClickhouseConfigOutputReference_Override(m MdbClickhouseClusterClickhouseConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbClickhouseClusterClickhouseConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetBackgroundPoolSize(val *float64) {
	_jsii_.Set(
		j,
		"backgroundPoolSize",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetBackgroundSchedulePoolSize(val *float64) {
	_jsii_.Set(
		j,
		"backgroundSchedulePoolSize",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetGeobaseUri(val *string) {
	_jsii_.Set(
		j,
		"geobaseUri",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetInternalValue(val *MdbClickhouseClusterClickhouseConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetKeepAliveTimeout(val *float64) {
	_jsii_.Set(
		j,
		"keepAliveTimeout",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetLogLevel(val *string) {
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetMarkCacheSize(val *float64) {
	_jsii_.Set(
		j,
		"markCacheSize",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetMaxConcurrentQueries(val *float64) {
	_jsii_.Set(
		j,
		"maxConcurrentQueries",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetMaxConnections(val *float64) {
	_jsii_.Set(
		j,
		"maxConnections",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetMaxPartitionSizeToDrop(val *float64) {
	_jsii_.Set(
		j,
		"maxPartitionSizeToDrop",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetMaxTableSizeToDrop(val *float64) {
	_jsii_.Set(
		j,
		"maxTableSizeToDrop",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetMetricLogEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"metricLogEnabled",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetMetricLogRetentionSize(val *float64) {
	_jsii_.Set(
		j,
		"metricLogRetentionSize",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetMetricLogRetentionTime(val *float64) {
	_jsii_.Set(
		j,
		"metricLogRetentionTime",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetPartLogRetentionSize(val *float64) {
	_jsii_.Set(
		j,
		"partLogRetentionSize",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetPartLogRetentionTime(val *float64) {
	_jsii_.Set(
		j,
		"partLogRetentionTime",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetQueryLogRetentionSize(val *float64) {
	_jsii_.Set(
		j,
		"queryLogRetentionSize",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetQueryLogRetentionTime(val *float64) {
	_jsii_.Set(
		j,
		"queryLogRetentionTime",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetQueryThreadLogEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"queryThreadLogEnabled",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetQueryThreadLogRetentionSize(val *float64) {
	_jsii_.Set(
		j,
		"queryThreadLogRetentionSize",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetQueryThreadLogRetentionTime(val *float64) {
	_jsii_.Set(
		j,
		"queryThreadLogRetentionTime",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetTextLogEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"textLogEnabled",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetTextLogLevel(val *string) {
	_jsii_.Set(
		j,
		"textLogLevel",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetTextLogRetentionSize(val *float64) {
	_jsii_.Set(
		j,
		"textLogRetentionSize",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetTextLogRetentionTime(val *float64) {
	_jsii_.Set(
		j,
		"textLogRetentionTime",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetTimezone(val *string) {
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetTraceLogEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"traceLogEnabled",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetTraceLogRetentionSize(val *float64) {
	_jsii_.Set(
		j,
		"traceLogRetentionSize",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetTraceLogRetentionTime(val *float64) {
	_jsii_.Set(
		j,
		"traceLogRetentionTime",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) SetUncompressedCacheSize(val *float64) {
	_jsii_.Set(
		j,
		"uncompressedCacheSize",
		val,
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) PutCompression(value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"putCompression",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) PutGraphiteRollup(value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"putGraphiteRollup",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) PutKafka(value *MdbClickhouseClusterClickhouseConfigKafka) {
	_jsii_.InvokeVoid(
		m,
		"putKafka",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) PutKafkaTopic(value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"putKafkaTopic",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) PutMergeTree(value *MdbClickhouseClusterClickhouseConfigMergeTree) {
	_jsii_.InvokeVoid(
		m,
		"putMergeTree",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) PutRabbitmq(value *MdbClickhouseClusterClickhouseConfigRabbitmq) {
	_jsii_.InvokeVoid(
		m,
		"putRabbitmq",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetBackgroundPoolSize() {
	_jsii_.InvokeVoid(
		m,
		"resetBackgroundPoolSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetBackgroundSchedulePoolSize() {
	_jsii_.InvokeVoid(
		m,
		"resetBackgroundSchedulePoolSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetCompression() {
	_jsii_.InvokeVoid(
		m,
		"resetCompression",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetGeobaseUri() {
	_jsii_.InvokeVoid(
		m,
		"resetGeobaseUri",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetGraphiteRollup() {
	_jsii_.InvokeVoid(
		m,
		"resetGraphiteRollup",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetKafka() {
	_jsii_.InvokeVoid(
		m,
		"resetKafka",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetKafkaTopic() {
	_jsii_.InvokeVoid(
		m,
		"resetKafkaTopic",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetKeepAliveTimeout() {
	_jsii_.InvokeVoid(
		m,
		"resetKeepAliveTimeout",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetLogLevel() {
	_jsii_.InvokeVoid(
		m,
		"resetLogLevel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetMarkCacheSize() {
	_jsii_.InvokeVoid(
		m,
		"resetMarkCacheSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetMaxConcurrentQueries() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxConcurrentQueries",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetMaxConnections() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxConnections",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetMaxPartitionSizeToDrop() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxPartitionSizeToDrop",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetMaxTableSizeToDrop() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxTableSizeToDrop",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetMergeTree() {
	_jsii_.InvokeVoid(
		m,
		"resetMergeTree",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetMetricLogEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetMetricLogEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetMetricLogRetentionSize() {
	_jsii_.InvokeVoid(
		m,
		"resetMetricLogRetentionSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetMetricLogRetentionTime() {
	_jsii_.InvokeVoid(
		m,
		"resetMetricLogRetentionTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetPartLogRetentionSize() {
	_jsii_.InvokeVoid(
		m,
		"resetPartLogRetentionSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetPartLogRetentionTime() {
	_jsii_.InvokeVoid(
		m,
		"resetPartLogRetentionTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetQueryLogRetentionSize() {
	_jsii_.InvokeVoid(
		m,
		"resetQueryLogRetentionSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetQueryLogRetentionTime() {
	_jsii_.InvokeVoid(
		m,
		"resetQueryLogRetentionTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetQueryThreadLogEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetQueryThreadLogEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetQueryThreadLogRetentionSize() {
	_jsii_.InvokeVoid(
		m,
		"resetQueryThreadLogRetentionSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetQueryThreadLogRetentionTime() {
	_jsii_.InvokeVoid(
		m,
		"resetQueryThreadLogRetentionTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetRabbitmq() {
	_jsii_.InvokeVoid(
		m,
		"resetRabbitmq",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetTextLogEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetTextLogEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetTextLogLevel() {
	_jsii_.InvokeVoid(
		m,
		"resetTextLogLevel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetTextLogRetentionSize() {
	_jsii_.InvokeVoid(
		m,
		"resetTextLogRetentionSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetTextLogRetentionTime() {
	_jsii_.InvokeVoid(
		m,
		"resetTextLogRetentionTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetTimezone() {
	_jsii_.InvokeVoid(
		m,
		"resetTimezone",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetTraceLogEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetTraceLogEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetTraceLogRetentionSize() {
	_jsii_.InvokeVoid(
		m,
		"resetTraceLogRetentionSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetTraceLogRetentionTime() {
	_jsii_.InvokeVoid(
		m,
		"resetTraceLogRetentionTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ResetUncompressedCacheSize() {
	_jsii_.InvokeVoid(
		m,
		"resetUncompressedCacheSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

