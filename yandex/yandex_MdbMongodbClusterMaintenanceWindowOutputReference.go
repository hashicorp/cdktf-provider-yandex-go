// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbMongodbClusterMaintenanceWindowOutputReference interface {
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
	Day() *string
	SetDay(val *string)
	DayInput() *string
	// Experimental.
	Fqn() *string
	Hour() *float64
	SetHour(val *float64)
	HourInput() *float64
	InternalValue() *MdbMongodbClusterMaintenanceWindow
	SetInternalValue(val *MdbMongodbClusterMaintenanceWindow)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetDay()
	ResetHour()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MdbMongodbClusterMaintenanceWindowOutputReference
type jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) Day() *string {
	var returns *string
	_jsii_.Get(
		j,
		"day",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) DayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) Hour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) HourInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) InternalValue() *MdbMongodbClusterMaintenanceWindow {
	var returns *MdbMongodbClusterMaintenanceWindow
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewMdbMongodbClusterMaintenanceWindowOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MdbMongodbClusterMaintenanceWindowOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbMongodbClusterMaintenanceWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMdbMongodbClusterMaintenanceWindowOutputReference_Override(m MdbMongodbClusterMaintenanceWindowOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbMongodbClusterMaintenanceWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) SetDay(val *string) {
	_jsii_.Set(
		j,
		"day",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) SetHour(val *float64) {
	_jsii_.Set(
		j,
		"hour",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) SetInternalValue(val *MdbMongodbClusterMaintenanceWindow) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (m *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) ResetDay() {
	_jsii_.InvokeVoid(
		m,
		"resetDay",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) ResetHour() {
	_jsii_.InvokeVoid(
		m,
		"resetHour",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterMaintenanceWindowOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

