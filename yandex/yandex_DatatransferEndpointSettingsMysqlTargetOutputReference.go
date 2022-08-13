// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatatransferEndpointSettingsMysqlTargetOutputReference interface {
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
	Connection() DatatransferEndpointSettingsMysqlTargetConnectionOutputReference
	ConnectionInput() *DatatransferEndpointSettingsMysqlTargetConnection
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
	InternalValue() *DatatransferEndpointSettingsMysqlTarget
	SetInternalValue(val *DatatransferEndpointSettingsMysqlTarget)
	Password() DatatransferEndpointSettingsMysqlTargetPasswordOutputReference
	PasswordInput() *DatatransferEndpointSettingsMysqlTargetPassword
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	SecurityGroupsInput() *[]*string
	SkipConstraintChecks() interface{}
	SetSkipConstraintChecks(val interface{})
	SkipConstraintChecksInput() interface{}
	SqlMode() *string
	SetSqlMode(val *string)
	SqlModeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timezone() *string
	SetTimezone(val *string)
	TimezoneInput() *string
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
	PutConnection(value *DatatransferEndpointSettingsMysqlTargetConnection)
	PutPassword(value *DatatransferEndpointSettingsMysqlTargetPassword)
	ResetConnection()
	ResetDatabase()
	ResetPassword()
	ResetSecurityGroups()
	ResetSkipConstraintChecks()
	ResetSqlMode()
	ResetTimezone()
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

// The jsii proxy struct for DatatransferEndpointSettingsMysqlTargetOutputReference
type jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) Connection() DatatransferEndpointSettingsMysqlTargetConnectionOutputReference {
	var returns DatatransferEndpointSettingsMysqlTargetConnectionOutputReference
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) ConnectionInput() *DatatransferEndpointSettingsMysqlTargetConnection {
	var returns *DatatransferEndpointSettingsMysqlTargetConnection
	_jsii_.Get(
		j,
		"connectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) InternalValue() *DatatransferEndpointSettingsMysqlTarget {
	var returns *DatatransferEndpointSettingsMysqlTarget
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) Password() DatatransferEndpointSettingsMysqlTargetPasswordOutputReference {
	var returns DatatransferEndpointSettingsMysqlTargetPasswordOutputReference
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) PasswordInput() *DatatransferEndpointSettingsMysqlTargetPassword {
	var returns *DatatransferEndpointSettingsMysqlTargetPassword
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) SkipConstraintChecks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipConstraintChecks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) SkipConstraintChecksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipConstraintChecksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) SqlMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) SqlModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) UserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}


func NewDatatransferEndpointSettingsMysqlTargetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatatransferEndpointSettingsMysqlTargetOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsMysqlTargetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatatransferEndpointSettingsMysqlTargetOutputReference_Override(d DatatransferEndpointSettingsMysqlTargetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsMysqlTargetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) SetDatabase(val *string) {
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) SetInternalValue(val *DatatransferEndpointSettingsMysqlTarget) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) SetSecurityGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) SetSkipConstraintChecks(val interface{}) {
	_jsii_.Set(
		j,
		"skipConstraintChecks",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) SetSqlMode(val *string) {
	_jsii_.Set(
		j,
		"sqlMode",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) SetTimezone(val *string) {
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) SetUser(val *string) {
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) PutConnection(value *DatatransferEndpointSettingsMysqlTargetConnection) {
	_jsii_.InvokeVoid(
		d,
		"putConnection",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) PutPassword(value *DatatransferEndpointSettingsMysqlTargetPassword) {
	_jsii_.InvokeVoid(
		d,
		"putPassword",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) ResetConnection() {
	_jsii_.InvokeVoid(
		d,
		"resetConnection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) ResetDatabase() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) ResetSkipConstraintChecks() {
	_jsii_.InvokeVoid(
		d,
		"resetSkipConstraintChecks",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) ResetSqlMode() {
	_jsii_.InvokeVoid(
		d,
		"resetSqlMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) ResetTimezone() {
	_jsii_.InvokeVoid(
		d,
		"resetTimezone",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) ResetUser() {
	_jsii_.InvokeVoid(
		d,
		"resetUser",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

