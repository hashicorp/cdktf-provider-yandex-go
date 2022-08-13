// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference interface {
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
	InternalValue() *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptions
	SetInternalValue(val *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptions)
	MdbClusterId() *string
	SetMdbClusterId(val *string)
	MdbClusterIdInput() *string
	OnPremise() DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference
	OnPremiseInput() *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremise
	Password() DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsPasswordOutputReference
	PasswordInput() *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsPassword
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
	PutOnPremise(value *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremise)
	PutPassword(value *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsPassword)
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

// The jsii proxy struct for DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference
type jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) AuthSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) AuthSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) InternalValue() *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptions {
	var returns *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) MdbClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mdbClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) MdbClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mdbClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) OnPremise() DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference {
	var returns DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference
	_jsii_.Get(
		j,
		"onPremise",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) OnPremiseInput() *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremise {
	var returns *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremise
	_jsii_.Get(
		j,
		"onPremiseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) Password() DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsPasswordOutputReference {
	var returns DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsPasswordOutputReference
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) PasswordInput() *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsPassword {
	var returns *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsPassword
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) UserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}


func NewDatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference_Override(d DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) SetAuthSource(val *string) {
	_jsii_.Set(
		j,
		"authSource",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) SetInternalValue(val *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptions) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) SetMdbClusterId(val *string) {
	_jsii_.Set(
		j,
		"mdbClusterId",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) SetUser(val *string) {
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) PutOnPremise(value *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremise) {
	_jsii_.InvokeVoid(
		d,
		"putOnPremise",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) PutPassword(value *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsPassword) {
	_jsii_.InvokeVoid(
		d,
		"putPassword",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) ResetAuthSource() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthSource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) ResetMdbClusterId() {
	_jsii_.InvokeVoid(
		d,
		"resetMdbClusterId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) ResetOnPremise() {
	_jsii_.InvokeVoid(
		d,
		"resetOnPremise",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) ResetUser() {
	_jsii_.InvokeVoid(
		d,
		"resetUser",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

