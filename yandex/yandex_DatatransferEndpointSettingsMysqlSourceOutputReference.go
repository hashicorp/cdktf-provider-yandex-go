// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatatransferEndpointSettingsMysqlSourceOutputReference interface {
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
	Connection() DatatransferEndpointSettingsMysqlSourceConnectionOutputReference
	ConnectionInput() *DatatransferEndpointSettingsMysqlSourceConnection
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	ExcludeTablesRegex() *[]*string
	SetExcludeTablesRegex(val *[]*string)
	ExcludeTablesRegexInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludeTablesRegex() *[]*string
	SetIncludeTablesRegex(val *[]*string)
	IncludeTablesRegexInput() *[]*string
	InternalValue() *DatatransferEndpointSettingsMysqlSource
	SetInternalValue(val *DatatransferEndpointSettingsMysqlSource)
	ObjectTransferSettings() DatatransferEndpointSettingsMysqlSourceObjectTransferSettingsOutputReference
	ObjectTransferSettingsInput() *DatatransferEndpointSettingsMysqlSourceObjectTransferSettings
	Password() DatatransferEndpointSettingsMysqlSourcePasswordOutputReference
	PasswordInput() *DatatransferEndpointSettingsMysqlSourcePassword
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	SecurityGroupsInput() *[]*string
	ServiceDatabase() *string
	SetServiceDatabase(val *string)
	ServiceDatabaseInput() *string
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
	PutConnection(value *DatatransferEndpointSettingsMysqlSourceConnection)
	PutObjectTransferSettings(value *DatatransferEndpointSettingsMysqlSourceObjectTransferSettings)
	PutPassword(value *DatatransferEndpointSettingsMysqlSourcePassword)
	ResetConnection()
	ResetDatabase()
	ResetExcludeTablesRegex()
	ResetIncludeTablesRegex()
	ResetObjectTransferSettings()
	ResetPassword()
	ResetSecurityGroups()
	ResetServiceDatabase()
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

// The jsii proxy struct for DatatransferEndpointSettingsMysqlSourceOutputReference
type jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) Connection() DatatransferEndpointSettingsMysqlSourceConnectionOutputReference {
	var returns DatatransferEndpointSettingsMysqlSourceConnectionOutputReference
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) ConnectionInput() *DatatransferEndpointSettingsMysqlSourceConnection {
	var returns *DatatransferEndpointSettingsMysqlSourceConnection
	_jsii_.Get(
		j,
		"connectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) ExcludeTablesRegex() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeTablesRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) ExcludeTablesRegexInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeTablesRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) IncludeTablesRegex() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeTablesRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) IncludeTablesRegexInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeTablesRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) InternalValue() *DatatransferEndpointSettingsMysqlSource {
	var returns *DatatransferEndpointSettingsMysqlSource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) ObjectTransferSettings() DatatransferEndpointSettingsMysqlSourceObjectTransferSettingsOutputReference {
	var returns DatatransferEndpointSettingsMysqlSourceObjectTransferSettingsOutputReference
	_jsii_.Get(
		j,
		"objectTransferSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) ObjectTransferSettingsInput() *DatatransferEndpointSettingsMysqlSourceObjectTransferSettings {
	var returns *DatatransferEndpointSettingsMysqlSourceObjectTransferSettings
	_jsii_.Get(
		j,
		"objectTransferSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) Password() DatatransferEndpointSettingsMysqlSourcePasswordOutputReference {
	var returns DatatransferEndpointSettingsMysqlSourcePasswordOutputReference
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) PasswordInput() *DatatransferEndpointSettingsMysqlSourcePassword {
	var returns *DatatransferEndpointSettingsMysqlSourcePassword
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) ServiceDatabase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) ServiceDatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceDatabaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) UserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}


func NewDatatransferEndpointSettingsMysqlSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatatransferEndpointSettingsMysqlSourceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsMysqlSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatatransferEndpointSettingsMysqlSourceOutputReference_Override(d DatatransferEndpointSettingsMysqlSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsMysqlSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) SetDatabase(val *string) {
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) SetExcludeTablesRegex(val *[]*string) {
	_jsii_.Set(
		j,
		"excludeTablesRegex",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) SetIncludeTablesRegex(val *[]*string) {
	_jsii_.Set(
		j,
		"includeTablesRegex",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) SetInternalValue(val *DatatransferEndpointSettingsMysqlSource) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) SetSecurityGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) SetServiceDatabase(val *string) {
	_jsii_.Set(
		j,
		"serviceDatabase",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) SetTimezone(val *string) {
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) SetUser(val *string) {
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) PutConnection(value *DatatransferEndpointSettingsMysqlSourceConnection) {
	_jsii_.InvokeVoid(
		d,
		"putConnection",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) PutObjectTransferSettings(value *DatatransferEndpointSettingsMysqlSourceObjectTransferSettings) {
	_jsii_.InvokeVoid(
		d,
		"putObjectTransferSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) PutPassword(value *DatatransferEndpointSettingsMysqlSourcePassword) {
	_jsii_.InvokeVoid(
		d,
		"putPassword",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) ResetConnection() {
	_jsii_.InvokeVoid(
		d,
		"resetConnection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) ResetDatabase() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) ResetExcludeTablesRegex() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludeTablesRegex",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) ResetIncludeTablesRegex() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeTablesRegex",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) ResetObjectTransferSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetObjectTransferSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) ResetServiceDatabase() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceDatabase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) ResetTimezone() {
	_jsii_.InvokeVoid(
		d,
		"resetTimezone",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) ResetUser() {
	_jsii_.InvokeVoid(
		d,
		"resetUser",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

