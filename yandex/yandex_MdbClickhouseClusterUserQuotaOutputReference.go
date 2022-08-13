// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbClickhouseClusterUserQuotaOutputReference interface {
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
	Errors() *float64
	SetErrors(val *float64)
	ErrorsInput() *float64
	ExecutionTime() *float64
	SetExecutionTime(val *float64)
	ExecutionTimeInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IntervalDuration() *float64
	SetIntervalDuration(val *float64)
	IntervalDurationInput() *float64
	Queries() *float64
	SetQueries(val *float64)
	QueriesInput() *float64
	ReadRows() *float64
	SetReadRows(val *float64)
	ReadRowsInput() *float64
	ResultRows() *float64
	SetResultRows(val *float64)
	ResultRowsInput() *float64
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
	ResetErrors()
	ResetExecutionTime()
	ResetQueries()
	ResetReadRows()
	ResetResultRows()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MdbClickhouseClusterUserQuotaOutputReference
type jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) Errors() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"errors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) ErrorsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"errorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) ExecutionTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) ExecutionTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executionTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) IntervalDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) IntervalDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) Queries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) QueriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) ReadRows() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readRows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) ReadRowsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readRowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) ResultRows() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"resultRows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) ResultRowsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"resultRowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMdbClickhouseClusterUserQuotaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MdbClickhouseClusterUserQuotaOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbClickhouseClusterUserQuotaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMdbClickhouseClusterUserQuotaOutputReference_Override(m MdbClickhouseClusterUserQuotaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbClickhouseClusterUserQuotaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) SetErrors(val *float64) {
	_jsii_.Set(
		j,
		"errors",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) SetExecutionTime(val *float64) {
	_jsii_.Set(
		j,
		"executionTime",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) SetIntervalDuration(val *float64) {
	_jsii_.Set(
		j,
		"intervalDuration",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) SetQueries(val *float64) {
	_jsii_.Set(
		j,
		"queries",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) SetReadRows(val *float64) {
	_jsii_.Set(
		j,
		"readRows",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) SetResultRows(val *float64) {
	_jsii_.Set(
		j,
		"resultRows",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) ResetErrors() {
	_jsii_.InvokeVoid(
		m,
		"resetErrors",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) ResetExecutionTime() {
	_jsii_.InvokeVoid(
		m,
		"resetExecutionTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) ResetQueries() {
	_jsii_.InvokeVoid(
		m,
		"resetQueries",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) ResetReadRows() {
	_jsii_.InvokeVoid(
		m,
		"resetReadRows",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) ResetResultRows() {
	_jsii_.InvokeVoid(
		m,
		"resetResultRows",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterUserQuotaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

