// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbMysqlUserConnectionLimitsOutputReference interface {
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
	InternalValue() *MdbMysqlUserConnectionLimits
	SetInternalValue(val *MdbMysqlUserConnectionLimits)
	MaxConnectionsPerHour() *float64
	SetMaxConnectionsPerHour(val *float64)
	MaxConnectionsPerHourInput() *float64
	MaxQuestionsPerHour() *float64
	SetMaxQuestionsPerHour(val *float64)
	MaxQuestionsPerHourInput() *float64
	MaxUpdatesPerHour() *float64
	SetMaxUpdatesPerHour(val *float64)
	MaxUpdatesPerHourInput() *float64
	MaxUserConnections() *float64
	SetMaxUserConnections(val *float64)
	MaxUserConnectionsInput() *float64
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
	ResetMaxConnectionsPerHour()
	ResetMaxQuestionsPerHour()
	ResetMaxUpdatesPerHour()
	ResetMaxUserConnections()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MdbMysqlUserConnectionLimitsOutputReference
type jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) InternalValue() *MdbMysqlUserConnectionLimits {
	var returns *MdbMysqlUserConnectionLimits
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) MaxConnectionsPerHour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionsPerHour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) MaxConnectionsPerHourInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionsPerHourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) MaxQuestionsPerHour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxQuestionsPerHour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) MaxQuestionsPerHourInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxQuestionsPerHourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) MaxUpdatesPerHour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUpdatesPerHour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) MaxUpdatesPerHourInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUpdatesPerHourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) MaxUserConnections() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUserConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) MaxUserConnectionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUserConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMdbMysqlUserConnectionLimitsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MdbMysqlUserConnectionLimitsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbMysqlUserConnectionLimitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMdbMysqlUserConnectionLimitsOutputReference_Override(m MdbMysqlUserConnectionLimitsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbMysqlUserConnectionLimitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) SetInternalValue(val *MdbMysqlUserConnectionLimits) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) SetMaxConnectionsPerHour(val *float64) {
	_jsii_.Set(
		j,
		"maxConnectionsPerHour",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) SetMaxQuestionsPerHour(val *float64) {
	_jsii_.Set(
		j,
		"maxQuestionsPerHour",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) SetMaxUpdatesPerHour(val *float64) {
	_jsii_.Set(
		j,
		"maxUpdatesPerHour",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) SetMaxUserConnections(val *float64) {
	_jsii_.Set(
		j,
		"maxUserConnections",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) ResetMaxConnectionsPerHour() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxConnectionsPerHour",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) ResetMaxQuestionsPerHour() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxQuestionsPerHour",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) ResetMaxUpdatesPerHour() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxUpdatesPerHour",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) ResetMaxUserConnections() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxUserConnections",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlUserConnectionLimitsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

