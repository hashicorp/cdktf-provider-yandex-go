// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbPostgresqlClusterConfigAccessOutputReference interface {
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
	InternalValue() *MdbPostgresqlClusterConfigAccess
	SetInternalValue(val *MdbPostgresqlClusterConfigAccess)
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

// The jsii proxy struct for MdbPostgresqlClusterConfigAccessOutputReference
type jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) DataLens() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataLens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) DataLensInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataLensInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) InternalValue() *MdbPostgresqlClusterConfigAccess {
	var returns *MdbPostgresqlClusterConfigAccess
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) Serverless() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverless",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) ServerlessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverlessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) WebSql() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webSql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) WebSqlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webSqlInput",
		&returns,
	)
	return returns
}


func NewMdbPostgresqlClusterConfigAccessOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MdbPostgresqlClusterConfigAccessOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbPostgresqlClusterConfigAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMdbPostgresqlClusterConfigAccessOutputReference_Override(m MdbPostgresqlClusterConfigAccessOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbPostgresqlClusterConfigAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) SetDataLens(val interface{}) {
	_jsii_.Set(
		j,
		"dataLens",
		val,
	)
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) SetInternalValue(val *MdbPostgresqlClusterConfigAccess) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) SetServerless(val interface{}) {
	_jsii_.Set(
		j,
		"serverless",
		val,
	)
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) SetWebSql(val interface{}) {
	_jsii_.Set(
		j,
		"webSql",
		val,
	)
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) ResetDataLens() {
	_jsii_.InvokeVoid(
		m,
		"resetDataLens",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) ResetServerless() {
	_jsii_.InvokeVoid(
		m,
		"resetServerless",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) ResetWebSql() {
	_jsii_.InvokeVoid(
		m,
		"resetWebSql",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAccessOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

