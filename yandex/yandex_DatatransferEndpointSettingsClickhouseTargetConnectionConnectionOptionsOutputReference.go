// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference interface {
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
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptions
	SetInternalValue(val *DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptions)
	MdbClusterId() *string
	SetMdbClusterId(val *string)
	MdbClusterIdInput() *string
	OnPremise() DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference
	OnPremiseInput() *DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremise
	Password() DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsPasswordOutputReference
	PasswordInput() *DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsPassword
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	User() *string
	SetUser(val *string)
	UserInput() *string
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
	PutOnPremise(value *DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremise)
	PutPassword(value *DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsPassword)
	ResetDatabase()
	ResetMdbClusterId()
	ResetOnPremise()
	ResetPassword()
	ResetUser()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference
type jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) InternalValue() *DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptions {
	var returns *DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) MdbClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mdbClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) MdbClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mdbClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) OnPremise() DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference {
	var returns DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference
	_jsii_.Get(
		j,
		"onPremise",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) OnPremiseInput() *DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremise {
	var returns *DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremise
	_jsii_.Get(
		j,
		"onPremiseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) Password() DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsPasswordOutputReference {
	var returns DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsPasswordOutputReference
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) PasswordInput() *DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsPassword {
	var returns *DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsPassword
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) UserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}


func NewDatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference_Override(d DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) SetDatabase(val *string) {
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) SetInternalValue(val *DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptions) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) SetMdbClusterId(val *string) {
	_jsii_.Set(
		j,
		"mdbClusterId",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) SetUser(val *string) {
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) PutOnPremise(value *DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremise) {
	_jsii_.InvokeVoid(
		d,
		"putOnPremise",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) PutPassword(value *DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsPassword) {
	_jsii_.InvokeVoid(
		d,
		"putPassword",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) ResetDatabase() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) ResetMdbClusterId() {
	_jsii_.InvokeVoid(
		d,
		"resetMdbClusterId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) ResetOnPremise() {
	_jsii_.InvokeVoid(
		d,
		"resetOnPremise",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) ResetUser() {
	_jsii_.InvokeVoid(
		d,
		"resetUser",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

