// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbKafkaTopicTopicConfigOutputReference interface {
	cdktf.ComplexObject
	CleanupPolicy() *string
	SetCleanupPolicy(val *string)
	CleanupPolicyInput() *string
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
	DeleteRetentionMs() *string
	SetDeleteRetentionMs(val *string)
	DeleteRetentionMsInput() *string
	FileDeleteDelayMs() *string
	SetFileDeleteDelayMs(val *string)
	FileDeleteDelayMsInput() *string
	FlushMessages() *string
	SetFlushMessages(val *string)
	FlushMessagesInput() *string
	FlushMs() *string
	SetFlushMs(val *string)
	FlushMsInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MdbKafkaTopicTopicConfig
	SetInternalValue(val *MdbKafkaTopicTopicConfig)
	MaxMessageBytes() *string
	SetMaxMessageBytes(val *string)
	MaxMessageBytesInput() *string
	MinCompactionLagMs() *string
	SetMinCompactionLagMs(val *string)
	MinCompactionLagMsInput() *string
	MinInsyncReplicas() *string
	SetMinInsyncReplicas(val *string)
	MinInsyncReplicasInput() *string
	Preallocate() interface{}
	SetPreallocate(val interface{})
	PreallocateInput() interface{}
	RetentionBytes() *string
	SetRetentionBytes(val *string)
	RetentionBytesInput() *string
	RetentionMs() *string
	SetRetentionMs(val *string)
	RetentionMsInput() *string
	SegmentBytes() *string
	SetSegmentBytes(val *string)
	SegmentBytesInput() *string
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
	ResetCleanupPolicy()
	ResetCompressionType()
	ResetDeleteRetentionMs()
	ResetFileDeleteDelayMs()
	ResetFlushMessages()
	ResetFlushMs()
	ResetMaxMessageBytes()
	ResetMinCompactionLagMs()
	ResetMinInsyncReplicas()
	ResetPreallocate()
	ResetRetentionBytes()
	ResetRetentionMs()
	ResetSegmentBytes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MdbKafkaTopicTopicConfigOutputReference
type jsiiProxy_MdbKafkaTopicTopicConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) CleanupPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cleanupPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) CleanupPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cleanupPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) CompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) CompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) DeleteRetentionMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteRetentionMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) DeleteRetentionMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteRetentionMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) FileDeleteDelayMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileDeleteDelayMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) FileDeleteDelayMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileDeleteDelayMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) FlushMessages() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flushMessages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) FlushMessagesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flushMessagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) FlushMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flushMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) FlushMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flushMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) InternalValue() *MdbKafkaTopicTopicConfig {
	var returns *MdbKafkaTopicTopicConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) MaxMessageBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxMessageBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) MaxMessageBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxMessageBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) MinCompactionLagMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCompactionLagMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) MinCompactionLagMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCompactionLagMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) MinInsyncReplicas() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minInsyncReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) MinInsyncReplicasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minInsyncReplicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) Preallocate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preallocate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) PreallocateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preallocateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) RetentionBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) RetentionBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) RetentionMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) RetentionMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) SegmentBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) SegmentBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMdbKafkaTopicTopicConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MdbKafkaTopicTopicConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MdbKafkaTopicTopicConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbKafkaTopicTopicConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMdbKafkaTopicTopicConfigOutputReference_Override(m MdbKafkaTopicTopicConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbKafkaTopicTopicConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) SetCleanupPolicy(val *string) {
	_jsii_.Set(
		j,
		"cleanupPolicy",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) SetCompressionType(val *string) {
	_jsii_.Set(
		j,
		"compressionType",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) SetDeleteRetentionMs(val *string) {
	_jsii_.Set(
		j,
		"deleteRetentionMs",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) SetFileDeleteDelayMs(val *string) {
	_jsii_.Set(
		j,
		"fileDeleteDelayMs",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) SetFlushMessages(val *string) {
	_jsii_.Set(
		j,
		"flushMessages",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) SetFlushMs(val *string) {
	_jsii_.Set(
		j,
		"flushMs",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) SetInternalValue(val *MdbKafkaTopicTopicConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) SetMaxMessageBytes(val *string) {
	_jsii_.Set(
		j,
		"maxMessageBytes",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) SetMinCompactionLagMs(val *string) {
	_jsii_.Set(
		j,
		"minCompactionLagMs",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) SetMinInsyncReplicas(val *string) {
	_jsii_.Set(
		j,
		"minInsyncReplicas",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) SetPreallocate(val interface{}) {
	_jsii_.Set(
		j,
		"preallocate",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) SetRetentionBytes(val *string) {
	_jsii_.Set(
		j,
		"retentionBytes",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) SetRetentionMs(val *string) {
	_jsii_.Set(
		j,
		"retentionMs",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) SetSegmentBytes(val *string) {
	_jsii_.Set(
		j,
		"segmentBytes",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) ResetCleanupPolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetCleanupPolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) ResetCompressionType() {
	_jsii_.InvokeVoid(
		m,
		"resetCompressionType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) ResetDeleteRetentionMs() {
	_jsii_.InvokeVoid(
		m,
		"resetDeleteRetentionMs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) ResetFileDeleteDelayMs() {
	_jsii_.InvokeVoid(
		m,
		"resetFileDeleteDelayMs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) ResetFlushMessages() {
	_jsii_.InvokeVoid(
		m,
		"resetFlushMessages",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) ResetFlushMs() {
	_jsii_.InvokeVoid(
		m,
		"resetFlushMs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) ResetMaxMessageBytes() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxMessageBytes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) ResetMinCompactionLagMs() {
	_jsii_.InvokeVoid(
		m,
		"resetMinCompactionLagMs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) ResetMinInsyncReplicas() {
	_jsii_.InvokeVoid(
		m,
		"resetMinInsyncReplicas",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) ResetPreallocate() {
	_jsii_.InvokeVoid(
		m,
		"resetPreallocate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) ResetRetentionBytes() {
	_jsii_.InvokeVoid(
		m,
		"resetRetentionBytes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) ResetRetentionMs() {
	_jsii_.InvokeVoid(
		m,
		"resetRetentionMs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) ResetSegmentBytes() {
	_jsii_.InvokeVoid(
		m,
		"resetSegmentBytes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaTopicTopicConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

