// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference interface {
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
	Function() *string
	SetFunction(val *string)
	FunctionInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Regexp() *string
	SetRegexp(val *string)
	RegexpInput() *string
	Retention() MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternRetentionList
	RetentionInput() interface{}
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
	PutRetention(value interface{})
	ResetRegexp()
	ResetRetention()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference
type jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) Function() *string {
	var returns *string
	_jsii_.Get(
		j,
		"function",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) FunctionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) Regexp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regexp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) RegexpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regexpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) Retention() MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternRetentionList {
	var returns MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternRetentionList
	_jsii_.Get(
		j,
		"retention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) RetentionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference_Override(m MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) SetFunction(val *string) {
	_jsii_.Set(
		j,
		"function",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) SetRegexp(val *string) {
	_jsii_.Set(
		j,
		"regexp",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) PutRetention(value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"putRetention",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) ResetRegexp() {
	_jsii_.InvokeVoid(
		m,
		"resetRegexp",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) ResetRetention() {
	_jsii_.InvokeVoid(
		m,
		"resetRetention",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

