// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexMdbClickhouseClusterClickhouseConfigOutputReference interface {
	cdktf.ComplexObject
	BackgroundPoolSize() *float64
	BackgroundSchedulePoolSize() *float64
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
	Compression() DataYandexMdbClickhouseClusterClickhouseConfigCompressionList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GeobaseUri() *string
	GraphiteRollup() DataYandexMdbClickhouseClusterClickhouseConfigGraphiteRollupList
	InternalValue() *DataYandexMdbClickhouseClusterClickhouseConfig
	SetInternalValue(val *DataYandexMdbClickhouseClusterClickhouseConfig)
	Kafka() DataYandexMdbClickhouseClusterClickhouseConfigKafkaList
	KafkaTopic() DataYandexMdbClickhouseClusterClickhouseConfigKafkaTopicList
	KeepAliveTimeout() *float64
	LogLevel() *string
	MarkCacheSize() *float64
	MaxConcurrentQueries() *float64
	MaxConnections() *float64
	MaxPartitionSizeToDrop() *float64
	MaxTableSizeToDrop() *float64
	MergeTree() DataYandexMdbClickhouseClusterClickhouseConfigMergeTreeList
	MetricLogEnabled() cdktf.IResolvable
	MetricLogRetentionSize() *float64
	MetricLogRetentionTime() *float64
	PartLogRetentionSize() *float64
	PartLogRetentionTime() *float64
	QueryLogRetentionSize() *float64
	QueryLogRetentionTime() *float64
	QueryThreadLogEnabled() cdktf.IResolvable
	QueryThreadLogRetentionSize() *float64
	QueryThreadLogRetentionTime() *float64
	Rabbitmq() DataYandexMdbClickhouseClusterClickhouseConfigRabbitmqList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TextLogEnabled() cdktf.IResolvable
	TextLogLevel() *string
	TextLogRetentionSize() *float64
	TextLogRetentionTime() *float64
	Timezone() *string
	TraceLogEnabled() cdktf.IResolvable
	TraceLogRetentionSize() *float64
	TraceLogRetentionTime() *float64
	UncompressedCacheSize() *float64
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

// The jsii proxy struct for DataYandexMdbClickhouseClusterClickhouseConfigOutputReference
type jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) BackgroundPoolSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backgroundPoolSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) BackgroundSchedulePoolSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backgroundSchedulePoolSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) Compression() DataYandexMdbClickhouseClusterClickhouseConfigCompressionList {
	var returns DataYandexMdbClickhouseClusterClickhouseConfigCompressionList
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) GeobaseUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geobaseUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) GraphiteRollup() DataYandexMdbClickhouseClusterClickhouseConfigGraphiteRollupList {
	var returns DataYandexMdbClickhouseClusterClickhouseConfigGraphiteRollupList
	_jsii_.Get(
		j,
		"graphiteRollup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) InternalValue() *DataYandexMdbClickhouseClusterClickhouseConfig {
	var returns *DataYandexMdbClickhouseClusterClickhouseConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) Kafka() DataYandexMdbClickhouseClusterClickhouseConfigKafkaList {
	var returns DataYandexMdbClickhouseClusterClickhouseConfigKafkaList
	_jsii_.Get(
		j,
		"kafka",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) KafkaTopic() DataYandexMdbClickhouseClusterClickhouseConfigKafkaTopicList {
	var returns DataYandexMdbClickhouseClusterClickhouseConfigKafkaTopicList
	_jsii_.Get(
		j,
		"kafkaTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) KeepAliveTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepAliveTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) MarkCacheSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"markCacheSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) MaxConcurrentQueries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentQueries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) MaxConnections() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) MaxPartitionSizeToDrop() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPartitionSizeToDrop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) MaxTableSizeToDrop() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTableSizeToDrop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) MergeTree() DataYandexMdbClickhouseClusterClickhouseConfigMergeTreeList {
	var returns DataYandexMdbClickhouseClusterClickhouseConfigMergeTreeList
	_jsii_.Get(
		j,
		"mergeTree",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) MetricLogEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"metricLogEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) MetricLogRetentionSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricLogRetentionSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) MetricLogRetentionTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricLogRetentionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) PartLogRetentionSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partLogRetentionSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) PartLogRetentionTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partLogRetentionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) QueryLogRetentionSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryLogRetentionSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) QueryLogRetentionTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryLogRetentionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) QueryThreadLogEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"queryThreadLogEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) QueryThreadLogRetentionSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryThreadLogRetentionSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) QueryThreadLogRetentionTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryThreadLogRetentionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) Rabbitmq() DataYandexMdbClickhouseClusterClickhouseConfigRabbitmqList {
	var returns DataYandexMdbClickhouseClusterClickhouseConfigRabbitmqList
	_jsii_.Get(
		j,
		"rabbitmq",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) TextLogEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"textLogEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) TextLogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textLogLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) TextLogRetentionSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"textLogRetentionSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) TextLogRetentionTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"textLogRetentionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) TraceLogEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"traceLogEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) TraceLogRetentionSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"traceLogRetentionSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) TraceLogRetentionTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"traceLogRetentionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) UncompressedCacheSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uncompressedCacheSize",
		&returns,
	)
	return returns
}


func NewDataYandexMdbClickhouseClusterClickhouseConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataYandexMdbClickhouseClusterClickhouseConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexMdbClickhouseClusterClickhouseConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataYandexMdbClickhouseClusterClickhouseConfigOutputReference_Override(d DataYandexMdbClickhouseClusterClickhouseConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexMdbClickhouseClusterClickhouseConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) SetInternalValue(val *DataYandexMdbClickhouseClusterClickhouseConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseClusterClickhouseConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

