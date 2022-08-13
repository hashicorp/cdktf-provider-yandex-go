// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbClickhouseClusterAccessOutputReference interface {
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
	DataLens() interface{}
	SetDataLens(val interface{})
	DataLensInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *MdbClickhouseClusterAccess
	SetInternalValue(val *MdbClickhouseClusterAccess)
	Metrika() interface{}
	SetMetrika(val interface{})
	MetrikaInput() interface{}
	Serverless() interface{}
	SetServerless(val interface{})
	ServerlessInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WebSql() interface{}
	SetWebSql(val interface{})
	WebSqlInput() interface{}
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
	ResetDataLens()
	ResetMetrika()
	ResetServerless()
	ResetWebSql()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MdbClickhouseClusterAccessOutputReference
type jsiiProxy_MdbClickhouseClusterAccessOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MdbClickhouseClusterAccessOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterAccessOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterAccessOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterAccessOutputReference) DataLens() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataLens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterAccessOutputReference) DataLensInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataLensInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterAccessOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterAccessOutputReference) InternalValue() *MdbClickhouseClusterAccess {
	var returns *MdbClickhouseClusterAccess
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterAccessOutputReference) Metrika() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metrika",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterAccessOutputReference) MetrikaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metrikaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterAccessOutputReference) Serverless() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverless",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterAccessOutputReference) ServerlessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverlessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterAccessOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterAccessOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterAccessOutputReference) WebSql() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webSql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterAccessOutputReference) WebSqlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webSqlInput",
		&returns,
	)
	return returns
}


func NewMdbClickhouseClusterAccessOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MdbClickhouseClusterAccessOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MdbClickhouseClusterAccessOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbClickhouseClusterAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMdbClickhouseClusterAccessOutputReference_Override(m MdbClickhouseClusterAccessOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbClickhouseClusterAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterAccessOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterAccessOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterAccessOutputReference) SetDataLens(val interface{}) {
	_jsii_.Set(
		j,
		"dataLens",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterAccessOutputReference) SetInternalValue(val *MdbClickhouseClusterAccess) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterAccessOutputReference) SetMetrika(val interface{}) {
	_jsii_.Set(
		j,
		"metrika",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterAccessOutputReference) SetServerless(val interface{}) {
	_jsii_.Set(
		j,
		"serverless",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterAccessOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterAccessOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterAccessOutputReference) SetWebSql(val interface{}) {
	_jsii_.Set(
		j,
		"webSql",
		val,
	)
}

func (m *jsiiProxy_MdbClickhouseClusterAccessOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterAccessOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterAccessOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterAccessOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterAccessOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterAccessOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterAccessOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterAccessOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterAccessOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterAccessOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterAccessOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterAccessOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterAccessOutputReference) ResetDataLens() {
	_jsii_.InvokeVoid(
		m,
		"resetDataLens",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterAccessOutputReference) ResetMetrika() {
	_jsii_.InvokeVoid(
		m,
		"resetMetrika",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterAccessOutputReference) ResetServerless() {
	_jsii_.InvokeVoid(
		m,
		"resetServerless",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterAccessOutputReference) ResetWebSql() {
	_jsii_.InvokeVoid(
		m,
		"resetWebSql",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterAccessOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterAccessOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

