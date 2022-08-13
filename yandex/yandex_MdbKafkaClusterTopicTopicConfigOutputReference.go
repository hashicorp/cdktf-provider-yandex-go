// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbKafkaClusterTopicTopicConfigOutputReference interface {
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
	InternalValue() *MdbKafkaClusterTopicTopicConfig
	SetInternalValue(val *MdbKafkaClusterTopicTopicConfig)
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

// The jsii proxy struct for MdbKafkaClusterTopicTopicConfigOutputReference
type jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) CleanupPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cleanupPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) CleanupPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cleanupPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) CompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) CompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) DeleteRetentionMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteRetentionMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) DeleteRetentionMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteRetentionMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) FileDeleteDelayMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileDeleteDelayMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) FileDeleteDelayMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileDeleteDelayMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) FlushMessages() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flushMessages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) FlushMessagesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flushMessagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) FlushMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flushMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) FlushMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flushMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) InternalValue() *MdbKafkaClusterTopicTopicConfig {
	var returns *MdbKafkaClusterTopicTopicConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) MaxMessageBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxMessageBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) MaxMessageBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxMessageBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) MinCompactionLagMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCompactionLagMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) MinCompactionLagMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCompactionLagMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) MinInsyncReplicas() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minInsyncReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) MinInsyncReplicasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minInsyncReplicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) Preallocate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preallocate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) PreallocateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preallocateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) RetentionBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) RetentionBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) RetentionMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) RetentionMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) SegmentBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) SegmentBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMdbKafkaClusterTopicTopicConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MdbKafkaClusterTopicTopicConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbKafkaClusterTopicTopicConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMdbKafkaClusterTopicTopicConfigOutputReference_Override(m MdbKafkaClusterTopicTopicConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbKafkaClusterTopicTopicConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) SetCleanupPolicy(val *string) {
	_jsii_.Set(
		j,
		"cleanupPolicy",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) SetCompressionType(val *string) {
	_jsii_.Set(
		j,
		"compressionType",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) SetDeleteRetentionMs(val *string) {
	_jsii_.Set(
		j,
		"deleteRetentionMs",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) SetFileDeleteDelayMs(val *string) {
	_jsii_.Set(
		j,
		"fileDeleteDelayMs",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) SetFlushMessages(val *string) {
	_jsii_.Set(
		j,
		"flushMessages",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) SetFlushMs(val *string) {
	_jsii_.Set(
		j,
		"flushMs",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) SetInternalValue(val *MdbKafkaClusterTopicTopicConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) SetMaxMessageBytes(val *string) {
	_jsii_.Set(
		j,
		"maxMessageBytes",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) SetMinCompactionLagMs(val *string) {
	_jsii_.Set(
		j,
		"minCompactionLagMs",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) SetMinInsyncReplicas(val *string) {
	_jsii_.Set(
		j,
		"minInsyncReplicas",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) SetPreallocate(val interface{}) {
	_jsii_.Set(
		j,
		"preallocate",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) SetRetentionBytes(val *string) {
	_jsii_.Set(
		j,
		"retentionBytes",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) SetRetentionMs(val *string) {
	_jsii_.Set(
		j,
		"retentionMs",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) SetSegmentBytes(val *string) {
	_jsii_.Set(
		j,
		"segmentBytes",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) ResetCleanupPolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetCleanupPolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) ResetCompressionType() {
	_jsii_.InvokeVoid(
		m,
		"resetCompressionType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) ResetDeleteRetentionMs() {
	_jsii_.InvokeVoid(
		m,
		"resetDeleteRetentionMs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) ResetFileDeleteDelayMs() {
	_jsii_.InvokeVoid(
		m,
		"resetFileDeleteDelayMs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) ResetFlushMessages() {
	_jsii_.InvokeVoid(
		m,
		"resetFlushMessages",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) ResetFlushMs() {
	_jsii_.InvokeVoid(
		m,
		"resetFlushMs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) ResetMaxMessageBytes() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxMessageBytes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) ResetMinCompactionLagMs() {
	_jsii_.InvokeVoid(
		m,
		"resetMinCompactionLagMs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) ResetMinInsyncReplicas() {
	_jsii_.InvokeVoid(
		m,
		"resetMinInsyncReplicas",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) ResetPreallocate() {
	_jsii_.InvokeVoid(
		m,
		"resetPreallocate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) ResetRetentionBytes() {
	_jsii_.InvokeVoid(
		m,
		"resetRetentionBytes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) ResetRetentionMs() {
	_jsii_.InvokeVoid(
		m,
		"resetRetentionMs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) ResetSegmentBytes() {
	_jsii_.InvokeVoid(
		m,
		"resetSegmentBytes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterTopicTopicConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

