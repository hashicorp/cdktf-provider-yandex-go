// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference interface {
	cdktf.ComplexObject
	AuthSource() *string
	SetAuthSource(val *string)
	AuthSourceInput() *string
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
	InternalValue() *DatatransferEndpointSettingsMongoTargetConnectionConnectionOptions
	SetInternalValue(val *DatatransferEndpointSettingsMongoTargetConnectionConnectionOptions)
	MdbClusterId() *string
	SetMdbClusterId(val *string)
	MdbClusterIdInput() *string
	OnPremise() DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference
	OnPremiseInput() *DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremise
	Password() DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsPasswordOutputReference
	PasswordInput() *DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsPassword
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
	PutOnPremise(value *DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremise)
	PutPassword(value *DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsPassword)
	ResetAuthSource()
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

// The jsii proxy struct for DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference
type jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) AuthSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) AuthSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) InternalValue() *DatatransferEndpointSettingsMongoTargetConnectionConnectionOptions {
	var returns *DatatransferEndpointSettingsMongoTargetConnectionConnectionOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) MdbClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mdbClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) MdbClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mdbClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) OnPremise() DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference {
	var returns DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference
	_jsii_.Get(
		j,
		"onPremise",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) OnPremiseInput() *DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremise {
	var returns *DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremise
	_jsii_.Get(
		j,
		"onPremiseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) Password() DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsPasswordOutputReference {
	var returns DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsPasswordOutputReference
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) PasswordInput() *DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsPassword {
	var returns *DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsPassword
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) UserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}


func NewDatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference_Override(d DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) SetAuthSource(val *string) {
	_jsii_.Set(
		j,
		"authSource",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) SetInternalValue(val *DatatransferEndpointSettingsMongoTargetConnectionConnectionOptions) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) SetMdbClusterId(val *string) {
	_jsii_.Set(
		j,
		"mdbClusterId",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) SetUser(val *string) {
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) PutOnPremise(value *DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremise) {
	_jsii_.InvokeVoid(
		d,
		"putOnPremise",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) PutPassword(value *DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsPassword) {
	_jsii_.InvokeVoid(
		d,
		"putPassword",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) ResetAuthSource() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthSource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) ResetMdbClusterId() {
	_jsii_.InvokeVoid(
		d,
		"resetMdbClusterId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) ResetOnPremise() {
	_jsii_.InvokeVoid(
		d,
		"resetOnPremise",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) ResetUser() {
	_jsii_.InvokeVoid(
		d,
		"resetUser",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

