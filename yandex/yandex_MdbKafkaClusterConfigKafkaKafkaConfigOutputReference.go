// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbKafkaClusterConfigKafkaKafkaConfigOutputReference interface {
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
	InternalValue() *MdbKafkaClusterConfigKafkaKafkaConfig
	SetInternalValue(val *MdbKafkaClusterConfigKafkaKafkaConfig)
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

// The jsii proxy struct for MdbKafkaClusterConfigKafkaKafkaConfigOutputReference
type jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) AutoCreateTopicsEnable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoCreateTopicsEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) AutoCreateTopicsEnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoCreateTopicsEnableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) CompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) CompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) DefaultReplicationFactor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultReplicationFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) DefaultReplicationFactorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultReplicationFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) InternalValue() *MdbKafkaClusterConfigKafkaKafkaConfig {
	var returns *MdbKafkaClusterConfigKafkaKafkaConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogFlushIntervalMessages() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logFlushIntervalMessages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogFlushIntervalMessagesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logFlushIntervalMessagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogFlushIntervalMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logFlushIntervalMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogFlushIntervalMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logFlushIntervalMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogFlushSchedulerIntervalMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logFlushSchedulerIntervalMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogFlushSchedulerIntervalMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logFlushSchedulerIntervalMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogPreallocate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logPreallocate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogPreallocateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logPreallocateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogRetentionBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logRetentionBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogRetentionBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logRetentionBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogRetentionHours() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logRetentionHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogRetentionHoursInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logRetentionHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogRetentionMinutes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logRetentionMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogRetentionMinutesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logRetentionMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogRetentionMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logRetentionMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogRetentionMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logRetentionMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogSegmentBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logSegmentBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) LogSegmentBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logSegmentBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) NumPartitions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"numPartitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) NumPartitionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"numPartitionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SocketReceiveBufferBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"socketReceiveBufferBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SocketReceiveBufferBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"socketReceiveBufferBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SocketSendBufferBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"socketSendBufferBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SocketSendBufferBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"socketSendBufferBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMdbKafkaClusterConfigKafkaKafkaConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MdbKafkaClusterConfigKafkaKafkaConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbKafkaClusterConfigKafkaKafkaConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMdbKafkaClusterConfigKafkaKafkaConfigOutputReference_Override(m MdbKafkaClusterConfigKafkaKafkaConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbKafkaClusterConfigKafkaKafkaConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetAutoCreateTopicsEnable(val interface{}) {
	_jsii_.Set(
		j,
		"autoCreateTopicsEnable",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetCompressionType(val *string) {
	_jsii_.Set(
		j,
		"compressionType",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetDefaultReplicationFactor(val *string) {
	_jsii_.Set(
		j,
		"defaultReplicationFactor",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetInternalValue(val *MdbKafkaClusterConfigKafkaKafkaConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetLogFlushIntervalMessages(val *string) {
	_jsii_.Set(
		j,
		"logFlushIntervalMessages",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetLogFlushIntervalMs(val *string) {
	_jsii_.Set(
		j,
		"logFlushIntervalMs",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetLogFlushSchedulerIntervalMs(val *string) {
	_jsii_.Set(
		j,
		"logFlushSchedulerIntervalMs",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetLogPreallocate(val interface{}) {
	_jsii_.Set(
		j,
		"logPreallocate",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetLogRetentionBytes(val *string) {
	_jsii_.Set(
		j,
		"logRetentionBytes",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetLogRetentionHours(val *string) {
	_jsii_.Set(
		j,
		"logRetentionHours",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetLogRetentionMinutes(val *string) {
	_jsii_.Set(
		j,
		"logRetentionMinutes",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetLogRetentionMs(val *string) {
	_jsii_.Set(
		j,
		"logRetentionMs",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetLogSegmentBytes(val *string) {
	_jsii_.Set(
		j,
		"logSegmentBytes",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetNumPartitions(val *string) {
	_jsii_.Set(
		j,
		"numPartitions",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetSocketReceiveBufferBytes(val *string) {
	_jsii_.Set(
		j,
		"socketReceiveBufferBytes",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetSocketSendBufferBytes(val *string) {
	_jsii_.Set(
		j,
		"socketSendBufferBytes",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ResetAutoCreateTopicsEnable() {
	_jsii_.InvokeVoid(
		m,
		"resetAutoCreateTopicsEnable",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ResetCompressionType() {
	_jsii_.InvokeVoid(
		m,
		"resetCompressionType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ResetDefaultReplicationFactor() {
	_jsii_.InvokeVoid(
		m,
		"resetDefaultReplicationFactor",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ResetLogFlushIntervalMessages() {
	_jsii_.InvokeVoid(
		m,
		"resetLogFlushIntervalMessages",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ResetLogFlushIntervalMs() {
	_jsii_.InvokeVoid(
		m,
		"resetLogFlushIntervalMs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ResetLogFlushSchedulerIntervalMs() {
	_jsii_.InvokeVoid(
		m,
		"resetLogFlushSchedulerIntervalMs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ResetLogPreallocate() {
	_jsii_.InvokeVoid(
		m,
		"resetLogPreallocate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ResetLogRetentionBytes() {
	_jsii_.InvokeVoid(
		m,
		"resetLogRetentionBytes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ResetLogRetentionHours() {
	_jsii_.InvokeVoid(
		m,
		"resetLogRetentionHours",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ResetLogRetentionMinutes() {
	_jsii_.InvokeVoid(
		m,
		"resetLogRetentionMinutes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ResetLogRetentionMs() {
	_jsii_.InvokeVoid(
		m,
		"resetLogRetentionMs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ResetLogSegmentBytes() {
	_jsii_.InvokeVoid(
		m,
		"resetLogSegmentBytes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ResetNumPartitions() {
	_jsii_.InvokeVoid(
		m,
		"resetNumPartitions",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ResetSocketReceiveBufferBytes() {
	_jsii_.InvokeVoid(
		m,
		"resetSocketReceiveBufferBytes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ResetSocketSendBufferBytes() {
	_jsii_.InvokeVoid(
		m,
		"resetSocketSendBufferBytes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaKafkaConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

