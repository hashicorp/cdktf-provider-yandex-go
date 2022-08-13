// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference interface {
	cdktf.ComplexObject
	AutoCreateTopicsEnable() interface{}
	SetAutoCreateTopicsEnable(val interface{})
	AutoCreateTopicsEnableInput() interface{}
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
	CompressionType() *string
	SetCompressionType(val *string)
	CompressionTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultReplicationFactor() *string
	SetDefaultReplicationFactor(val *string)
	DefaultReplicationFactorInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataYandexMdbKafkaClusterConfigKafkaKafkaConfig
	SetInternalValue(val *DataYandexMdbKafkaClusterConfigKafkaKafkaConfig)
	LogFlushIntervalMessages() *string
	SetLogFlushIntervalMessages(val *string)
	LogFlushIntervalMessagesInput() *string
	LogFlushIntervalMs() *string
	SetLogFlushIntervalMs(val *string)
	LogFlushIntervalMsInput() *string
	LogFlushSchedulerIntervalMs() *string
	SetLogFlushSchedulerIntervalMs(val *string)
	LogFlushSchedulerIntervalMsInput() *string
	LogPreallocate() interface{}
	SetLogPreallocate(val interface{})
	LogPreallocateInput() interface{}
	LogRetentionBytes() *string
	SetLogRetentionBytes(val *string)
	LogRetentionBytesInput() *string
	LogRetentionHours() *string
	SetLogRetentionHours(val *string)
	LogRetentionHoursInput() *string
	LogRetentionMinutes() *string
	SetLogRetentionMinutes(val *string)
	LogRetentionMinutesInput() *string
	LogRetentionMs() *string
	SetLogRetentionMs(val *string)
	LogRetentionMsInput() *string
	LogSegmentBytes() *string
	SetLogSegmentBytes(val *string)
	LogSegmentBytesInput() *string
	NumPartitions() *string
	SetNumPartitions(val *string)
	NumPartitionsInput() *string
	SocketReceiveBufferBytes() *string
	SetSocketReceiveBufferBytes(val *string)
	SocketReceiveBufferBytesInput() *string
	SocketSendBufferBytes() *string
	SetSocketSendBufferBytes(val *string)
	SocketSendBufferBytesInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetAutoCreateTopicsEnable()
	ResetCompressionType()
	ResetDefaultReplicationFactor()
	ResetLogFlushIntervalMessages()
	ResetLogFlushIntervalMs()
	ResetLogFlushSchedulerIntervalMs()
	ResetLogPreallocate()
	ResetLogRetentionBytes()
	ResetLogRetentionHours()
	ResetLogRetentionMinutes()
	ResetLogRetentionMs()
	ResetLogSegmentBytes()
	ResetNumPartitions()
	ResetSocketReceiveBufferBytes()
	ResetSocketSendBufferBytes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference
type jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) AutoCreateTopicsEnable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoCreateTopicsEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) AutoCreateTopicsEnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoCreateTopicsEnableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) CompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) CompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) DefaultReplicationFactor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultReplicationFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) DefaultReplicationFactorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultReplicationFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) InternalValue() *DataYandexMdbKafkaClusterConfigKafkaKafkaConfig {
	var returns *DataYandexMdbKafkaClusterConfigKafkaKafkaConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogFlushIntervalMessages() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logFlushIntervalMessages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogFlushIntervalMessagesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logFlushIntervalMessagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogFlushIntervalMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logFlushIntervalMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogFlushIntervalMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logFlushIntervalMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogFlushSchedulerIntervalMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logFlushSchedulerIntervalMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogFlushSchedulerIntervalMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logFlushSchedulerIntervalMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogPreallocate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logPreallocate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogPreallocateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logPreallocateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogRetentionBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logRetentionBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogRetentionBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logRetentionBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogRetentionHours() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logRetentionHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogRetentionHoursInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logRetentionHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogRetentionMinutes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logRetentionMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogRetentionMinutesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logRetentionMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogRetentionMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logRetentionMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogRetentionMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logRetentionMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogSegmentBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logSegmentBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogSegmentBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logSegmentBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) NumPartitions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"numPartitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) NumPartitionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"numPartitionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SocketReceiveBufferBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"socketReceiveBufferBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SocketReceiveBufferBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"socketReceiveBufferBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SocketSendBufferBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"socketSendBufferBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SocketSendBufferBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"socketSendBufferBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference_Override(d DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetAutoCreateTopicsEnable(val interface{}) {
	_jsii_.Set(
		j,
		"autoCreateTopicsEnable",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetCompressionType(val *string) {
	_jsii_.Set(
		j,
		"compressionType",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetDefaultReplicationFactor(val *string) {
	_jsii_.Set(
		j,
		"defaultReplicationFactor",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetInternalValue(val *DataYandexMdbKafkaClusterConfigKafkaKafkaConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetLogFlushIntervalMessages(val *string) {
	_jsii_.Set(
		j,
		"logFlushIntervalMessages",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetLogFlushIntervalMs(val *string) {
	_jsii_.Set(
		j,
		"logFlushIntervalMs",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetLogFlushSchedulerIntervalMs(val *string) {
	_jsii_.Set(
		j,
		"logFlushSchedulerIntervalMs",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetLogPreallocate(val interface{}) {
	_jsii_.Set(
		j,
		"logPreallocate",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetLogRetentionBytes(val *string) {
	_jsii_.Set(
		j,
		"logRetentionBytes",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetLogRetentionHours(val *string) {
	_jsii_.Set(
		j,
		"logRetentionHours",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetLogRetentionMinutes(val *string) {
	_jsii_.Set(
		j,
		"logRetentionMinutes",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetLogRetentionMs(val *string) {
	_jsii_.Set(
		j,
		"logRetentionMs",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetLogSegmentBytes(val *string) {
	_jsii_.Set(
		j,
		"logSegmentBytes",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetNumPartitions(val *string) {
	_jsii_.Set(
		j,
		"numPartitions",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetSocketReceiveBufferBytes(val *string) {
	_jsii_.Set(
		j,
		"socketReceiveBufferBytes",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetSocketSendBufferBytes(val *string) {
	_jsii_.Set(
		j,
		"socketSendBufferBytes",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ResetAutoCreateTopicsEnable() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoCreateTopicsEnable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ResetCompressionType() {
	_jsii_.InvokeVoid(
		d,
		"resetCompressionType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ResetDefaultReplicationFactor() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultReplicationFactor",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ResetLogFlushIntervalMessages() {
	_jsii_.InvokeVoid(
		d,
		"resetLogFlushIntervalMessages",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ResetLogFlushIntervalMs() {
	_jsii_.InvokeVoid(
		d,
		"resetLogFlushIntervalMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ResetLogFlushSchedulerIntervalMs() {
	_jsii_.InvokeVoid(
		d,
		"resetLogFlushSchedulerIntervalMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ResetLogPreallocate() {
	_jsii_.InvokeVoid(
		d,
		"resetLogPreallocate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ResetLogRetentionBytes() {
	_jsii_.InvokeVoid(
		d,
		"resetLogRetentionBytes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ResetLogRetentionHours() {
	_jsii_.InvokeVoid(
		d,
		"resetLogRetentionHours",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ResetLogRetentionMinutes() {
	_jsii_.InvokeVoid(
		d,
		"resetLogRetentionMinutes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ResetLogRetentionMs() {
	_jsii_.InvokeVoid(
		d,
		"resetLogRetentionMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ResetLogSegmentBytes() {
	_jsii_.InvokeVoid(
		d,
		"resetLogSegmentBytes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ResetNumPartitions() {
	_jsii_.InvokeVoid(
		d,
		"resetNumPartitions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ResetSocketReceiveBufferBytes() {
	_jsii_.InvokeVoid(
		d,
		"resetSocketReceiveBufferBytes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ResetSocketSendBufferBytes() {
	_jsii_.InvokeVoid(
		d,
		"resetSocketSendBufferBytes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

