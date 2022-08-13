// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbGreenplumClusterPoolerConfigOutputReference interface {
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
	InternalValue() *MdbGreenplumClusterPoolerConfig
	SetInternalValue(val *MdbGreenplumClusterPoolerConfig)
	PoolClientIdleTimeout() *float64
	SetPoolClientIdleTimeout(val *float64)
	PoolClientIdleTimeoutInput() *float64
	PoolingMode() *string
	SetPoolingMode(val *string)
	PoolingModeInput() *string
	PoolSize() *float64
	SetPoolSize(val *float64)
	PoolSizeInput() *float64
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
	ResetPoolClientIdleTimeout()
	ResetPoolingMode()
	ResetPoolSize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MdbGreenplumClusterPoolerConfigOutputReference
type jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) InternalValue() *MdbGreenplumClusterPoolerConfig {
	var returns *MdbGreenplumClusterPoolerConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) PoolClientIdleTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"poolClientIdleTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) PoolClientIdleTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"poolClientIdleTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) PoolingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) PoolingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) PoolSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"poolSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) PoolSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"poolSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMdbGreenplumClusterPoolerConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MdbGreenplumClusterPoolerConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbGreenplumClusterPoolerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMdbGreenplumClusterPoolerConfigOutputReference_Override(m MdbGreenplumClusterPoolerConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbGreenplumClusterPoolerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) SetInternalValue(val *MdbGreenplumClusterPoolerConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) SetPoolClientIdleTimeout(val *float64) {
	_jsii_.Set(
		j,
		"poolClientIdleTimeout",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) SetPoolingMode(val *string) {
	_jsii_.Set(
		j,
		"poolingMode",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) SetPoolSize(val *float64) {
	_jsii_.Set(
		j,
		"poolSize",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) ResetPoolClientIdleTimeout() {
	_jsii_.InvokeVoid(
		m,
		"resetPoolClientIdleTimeout",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) ResetPoolingMode() {
	_jsii_.InvokeVoid(
		m,
		"resetPoolingMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) ResetPoolSize() {
	_jsii_.InvokeVoid(
		m,
		"resetPoolSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbGreenplumClusterPoolerConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

