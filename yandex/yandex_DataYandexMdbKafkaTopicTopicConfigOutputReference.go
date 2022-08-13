// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexMdbKafkaTopicTopicConfigOutputReference interface {
	cdktf.ComplexObject
	CleanupPolicy() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeleteRetentionMs() *string
	FileDeleteDelayMs() *string
	FlushMessages() *string
	FlushMs() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataYandexMdbKafkaTopicTopicConfig
	SetInternalValue(val *DataYandexMdbKafkaTopicTopicConfig)
	MaxMessageBytes() *string
	MinCompactionLagMs() *string
	MinInsyncReplicas() *string
	Preallocate() cdktf.IResolvable
	RetentionBytes() *string
	RetentionMs() *string
	SegmentBytes() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataYandexMdbKafkaTopicTopicConfigOutputReference
type jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) CleanupPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cleanupPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) CompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) DeleteRetentionMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteRetentionMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) FileDeleteDelayMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileDeleteDelayMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) FlushMessages() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flushMessages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) FlushMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flushMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) InternalValue() *DataYandexMdbKafkaTopicTopicConfig {
	var returns *DataYandexMdbKafkaTopicTopicConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) MaxMessageBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxMessageBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) MinCompactionLagMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCompactionLagMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) MinInsyncReplicas() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minInsyncReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) Preallocate() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"preallocate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) RetentionBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) RetentionMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) SegmentBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataYandexMdbKafkaTopicTopicConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataYandexMdbKafkaTopicTopicConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexMdbKafkaTopicTopicConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataYandexMdbKafkaTopicTopicConfigOutputReference_Override(d DataYandexMdbKafkaTopicTopicConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexMdbKafkaTopicTopicConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) SetInternalValue(val *DataYandexMdbKafkaTopicTopicConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaTopicTopicConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

