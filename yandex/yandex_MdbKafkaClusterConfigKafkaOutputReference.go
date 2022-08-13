// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbKafkaClusterConfigKafkaOutputReference interface {
	cdktf.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *MdbKafkaClusterConfigKafka
	SetInternalValue(val *MdbKafkaClusterConfigKafka)
	KafkaConfig() MdbKafkaClusterConfigKafkaKafkaConfigOutputReference
	KafkaConfigInput() *MdbKafkaClusterConfigKafkaKafkaConfig
	Resources() MdbKafkaClusterConfigKafkaResourcesOutputReference
	ResourcesInput() *MdbKafkaClusterConfigKafkaResources
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
	PutKafkaConfig(value *MdbKafkaClusterConfigKafkaKafkaConfig)
	PutResources(value *MdbKafkaClusterConfigKafkaResources)
	ResetKafkaConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MdbKafkaClusterConfigKafkaOutputReference
type jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) InternalValue() *MdbKafkaClusterConfigKafka {
	var returns *MdbKafkaClusterConfigKafka
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) KafkaConfig() MdbKafkaClusterConfigKafkaKafkaConfigOutputReference {
	var returns MdbKafkaClusterConfigKafkaKafkaConfigOutputReference
	_jsii_.Get(
		j,
		"kafkaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) KafkaConfigInput() *MdbKafkaClusterConfigKafkaKafkaConfig {
	var returns *MdbKafkaClusterConfigKafkaKafkaConfig
	_jsii_.Get(
		j,
		"kafkaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) Resources() MdbKafkaClusterConfigKafkaResourcesOutputReference {
	var returns MdbKafkaClusterConfigKafkaResourcesOutputReference
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) ResourcesInput() *MdbKafkaClusterConfigKafkaResources {
	var returns *MdbKafkaClusterConfigKafkaResources
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMdbKafkaClusterConfigKafkaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MdbKafkaClusterConfigKafkaOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbKafkaClusterConfigKafkaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMdbKafkaClusterConfigKafkaOutputReference_Override(m MdbKafkaClusterConfigKafkaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbKafkaClusterConfigKafkaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) SetInternalValue(val *MdbKafkaClusterConfigKafka) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) PutKafkaConfig(value *MdbKafkaClusterConfigKafkaKafkaConfig) {
	_jsii_.InvokeVoid(
		m,
		"putKafkaConfig",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) PutResources(value *MdbKafkaClusterConfigKafkaResources) {
	_jsii_.InvokeVoid(
		m,
		"putResources",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) ResetKafkaConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetKafkaConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

