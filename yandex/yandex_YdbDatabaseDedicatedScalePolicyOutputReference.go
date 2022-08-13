// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type YdbDatabaseDedicatedScalePolicyOutputReference interface {
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
	FixedScale() YdbDatabaseDedicatedScalePolicyFixedScaleOutputReference
	FixedScaleInput() *YdbDatabaseDedicatedScalePolicyFixedScale
	// Experimental.
	Fqn() *string
	InternalValue() *YdbDatabaseDedicatedScalePolicy
	SetInternalValue(val *YdbDatabaseDedicatedScalePolicy)
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
	PutFixedScale(value *YdbDatabaseDedicatedScalePolicyFixedScale)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for YdbDatabaseDedicatedScalePolicyOutputReference
type jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference) FixedScale() YdbDatabaseDedicatedScalePolicyFixedScaleOutputReference {
	var returns YdbDatabaseDedicatedScalePolicyFixedScaleOutputReference
	_jsii_.Get(
		j,
		"fixedScale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference) FixedScaleInput() *YdbDatabaseDedicatedScalePolicyFixedScale {
	var returns *YdbDatabaseDedicatedScalePolicyFixedScale
	_jsii_.Get(
		j,
		"fixedScaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference) InternalValue() *YdbDatabaseDedicatedScalePolicy {
	var returns *YdbDatabaseDedicatedScalePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewYdbDatabaseDedicatedScalePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) YdbDatabaseDedicatedScalePolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.YdbDatabaseDedicatedScalePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewYdbDatabaseDedicatedScalePolicyOutputReference_Override(y YdbDatabaseDedicatedScalePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.YdbDatabaseDedicatedScalePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		y,
	)
}

func (j *jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference) SetInternalValue(val *YdbDatabaseDedicatedScalePolicy) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (y *jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		y,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		y,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		y,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		y,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		y,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		y,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		y,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		y,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		y,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		y,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		y,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		y,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference) PutFixedScale(value *YdbDatabaseDedicatedScalePolicyFixedScale) {
	_jsii_.InvokeVoid(
		y,
		"putFixedScale",
		[]interface{}{value},
	)
}

func (y *jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		y,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseDedicatedScalePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		y,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

