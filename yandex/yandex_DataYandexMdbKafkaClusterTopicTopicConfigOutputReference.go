// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexMdbKafkaClusterTopicTopicConfigOutputReference interface {
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
	InternalValue() *DataYandexMdbKafkaClusterTopicTopicConfig
	SetInternalValue(val *DataYandexMdbKafkaClusterTopicTopicConfig)
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

// The jsii proxy struct for DataYandexMdbKafkaClusterTopicTopicConfigOutputReference
type jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) CleanupPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cleanupPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) CleanupPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cleanupPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) CompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) CompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) DeleteRetentionMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteRetentionMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) DeleteRetentionMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteRetentionMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) FileDeleteDelayMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileDeleteDelayMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) FileDeleteDelayMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileDeleteDelayMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) FlushMessages() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flushMessages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) FlushMessagesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flushMessagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) FlushMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flushMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) FlushMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flushMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) InternalValue() *DataYandexMdbKafkaClusterTopicTopicConfig {
	var returns *DataYandexMdbKafkaClusterTopicTopicConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) MaxMessageBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxMessageBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) MaxMessageBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxMessageBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) MinCompactionLagMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCompactionLagMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) MinCompactionLagMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCompactionLagMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) MinInsyncReplicas() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minInsyncReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) MinInsyncReplicasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minInsyncReplicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) Preallocate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preallocate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) PreallocateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preallocateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) RetentionBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) RetentionBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) RetentionMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) RetentionMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) SegmentBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) SegmentBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataYandexMdbKafkaClusterTopicTopicConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataYandexMdbKafkaClusterTopicTopicConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexMdbKafkaClusterTopicTopicConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataYandexMdbKafkaClusterTopicTopicConfigOutputReference_Override(d DataYandexMdbKafkaClusterTopicTopicConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexMdbKafkaClusterTopicTopicConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) SetCleanupPolicy(val *string) {
	_jsii_.Set(
		j,
		"cleanupPolicy",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) SetCompressionType(val *string) {
	_jsii_.Set(
		j,
		"compressionType",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) SetDeleteRetentionMs(val *string) {
	_jsii_.Set(
		j,
		"deleteRetentionMs",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) SetFileDeleteDelayMs(val *string) {
	_jsii_.Set(
		j,
		"fileDeleteDelayMs",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) SetFlushMessages(val *string) {
	_jsii_.Set(
		j,
		"flushMessages",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) SetFlushMs(val *string) {
	_jsii_.Set(
		j,
		"flushMs",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) SetInternalValue(val *DataYandexMdbKafkaClusterTopicTopicConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) SetMaxMessageBytes(val *string) {
	_jsii_.Set(
		j,
		"maxMessageBytes",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) SetMinCompactionLagMs(val *string) {
	_jsii_.Set(
		j,
		"minCompactionLagMs",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) SetMinInsyncReplicas(val *string) {
	_jsii_.Set(
		j,
		"minInsyncReplicas",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) SetPreallocate(val interface{}) {
	_jsii_.Set(
		j,
		"preallocate",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) SetRetentionBytes(val *string) {
	_jsii_.Set(
		j,
		"retentionBytes",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) SetRetentionMs(val *string) {
	_jsii_.Set(
		j,
		"retentionMs",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) SetSegmentBytes(val *string) {
	_jsii_.Set(
		j,
		"segmentBytes",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) ResetCleanupPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetCleanupPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) ResetCompressionType() {
	_jsii_.InvokeVoid(
		d,
		"resetCompressionType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) ResetDeleteRetentionMs() {
	_jsii_.InvokeVoid(
		d,
		"resetDeleteRetentionMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) ResetFileDeleteDelayMs() {
	_jsii_.InvokeVoid(
		d,
		"resetFileDeleteDelayMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) ResetFlushMessages() {
	_jsii_.InvokeVoid(
		d,
		"resetFlushMessages",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) ResetFlushMs() {
	_jsii_.InvokeVoid(
		d,
		"resetFlushMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) ResetMaxMessageBytes() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxMessageBytes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) ResetMinCompactionLagMs() {
	_jsii_.InvokeVoid(
		d,
		"resetMinCompactionLagMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) ResetMinInsyncReplicas() {
	_jsii_.InvokeVoid(
		d,
		"resetMinInsyncReplicas",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) ResetPreallocate() {
	_jsii_.InvokeVoid(
		d,
		"resetPreallocate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) ResetRetentionBytes() {
	_jsii_.InvokeVoid(
		d,
		"resetRetentionBytes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) ResetRetentionMs() {
	_jsii_.InvokeVoid(
		d,
		"resetRetentionMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) ResetSegmentBytes() {
	_jsii_.InvokeVoid(
		d,
		"resetSegmentBytes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterTopicTopicConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

