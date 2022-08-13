// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeInstanceGroupScalePolicyOutputReference interface {
	cdktf.ComplexObject
	AutoScale() ComputeInstanceGroupScalePolicyAutoScaleOutputReference
	AutoScaleInput() *ComputeInstanceGroupScalePolicyAutoScale
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
	FixedScale() ComputeInstanceGroupScalePolicyFixedScaleOutputReference
	FixedScaleInput() *ComputeInstanceGroupScalePolicyFixedScale
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeInstanceGroupScalePolicy
	SetInternalValue(val *ComputeInstanceGroupScalePolicy)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TestAutoScale() ComputeInstanceGroupScalePolicyTestAutoScaleOutputReference
	TestAutoScaleInput() *ComputeInstanceGroupScalePolicyTestAutoScale
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
	PutAutoScale(value *ComputeInstanceGroupScalePolicyAutoScale)
	PutFixedScale(value *ComputeInstanceGroupScalePolicyFixedScale)
	PutTestAutoScale(value *ComputeInstanceGroupScalePolicyTestAutoScale)
	ResetAutoScale()
	ResetFixedScale()
	ResetTestAutoScale()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeInstanceGroupScalePolicyOutputReference
type jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) AutoScale() ComputeInstanceGroupScalePolicyAutoScaleOutputReference {
	var returns ComputeInstanceGroupScalePolicyAutoScaleOutputReference
	_jsii_.Get(
		j,
		"autoScale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) AutoScaleInput() *ComputeInstanceGroupScalePolicyAutoScale {
	var returns *ComputeInstanceGroupScalePolicyAutoScale
	_jsii_.Get(
		j,
		"autoScaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) FixedScale() ComputeInstanceGroupScalePolicyFixedScaleOutputReference {
	var returns ComputeInstanceGroupScalePolicyFixedScaleOutputReference
	_jsii_.Get(
		j,
		"fixedScale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) FixedScaleInput() *ComputeInstanceGroupScalePolicyFixedScale {
	var returns *ComputeInstanceGroupScalePolicyFixedScale
	_jsii_.Get(
		j,
		"fixedScaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) InternalValue() *ComputeInstanceGroupScalePolicy {
	var returns *ComputeInstanceGroupScalePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) TestAutoScale() ComputeInstanceGroupScalePolicyTestAutoScaleOutputReference {
	var returns ComputeInstanceGroupScalePolicyTestAutoScaleOutputReference
	_jsii_.Get(
		j,
		"testAutoScale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) TestAutoScaleInput() *ComputeInstanceGroupScalePolicyTestAutoScale {
	var returns *ComputeInstanceGroupScalePolicyTestAutoScale
	_jsii_.Get(
		j,
		"testAutoScaleInput",
		&returns,
	)
	return returns
}


func NewComputeInstanceGroupScalePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeInstanceGroupScalePolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.ComputeInstanceGroupScalePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeInstanceGroupScalePolicyOutputReference_Override(c ComputeInstanceGroupScalePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.ComputeInstanceGroupScalePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) SetInternalValue(val *ComputeInstanceGroupScalePolicy) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) PutAutoScale(value *ComputeInstanceGroupScalePolicyAutoScale) {
	_jsii_.InvokeVoid(
		c,
		"putAutoScale",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) PutFixedScale(value *ComputeInstanceGroupScalePolicyFixedScale) {
	_jsii_.InvokeVoid(
		c,
		"putFixedScale",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) PutTestAutoScale(value *ComputeInstanceGroupScalePolicyTestAutoScale) {
	_jsii_.InvokeVoid(
		c,
		"putTestAutoScale",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) ResetAutoScale() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoScale",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) ResetFixedScale() {
	_jsii_.InvokeVoid(
		c,
		"resetFixedScale",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) ResetTestAutoScale() {
	_jsii_.InvokeVoid(
		c,
		"resetTestAutoScale",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

