// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatatransferEndpointSettingsPostgresSourceOutputReference interface {
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
	Connection() DatatransferEndpointSettingsPostgresSourceConnectionOutputReference
	ConnectionInput() *DatatransferEndpointSettingsPostgresSourceConnection
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	ExcludeTables() *[]*string
	SetExcludeTables(val *[]*string)
	ExcludeTablesInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludeTables() *[]*string
	SetIncludeTables(val *[]*string)
	IncludeTablesInput() *[]*string
	InternalValue() *DatatransferEndpointSettingsPostgresSource
	SetInternalValue(val *DatatransferEndpointSettingsPostgresSource)
	ObjectTransferSettings() DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference
	ObjectTransferSettingsInput() *DatatransferEndpointSettingsPostgresSourceObjectTransferSettings
	Password() DatatransferEndpointSettingsPostgresSourcePasswordOutputReference
	PasswordInput() *DatatransferEndpointSettingsPostgresSourcePassword
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	SecurityGroupsInput() *[]*string
	ServiceSchema() *string
	SetServiceSchema(val *string)
	ServiceSchemaInput() *string
	SlotGigabyteLagLimit() *float64
	SetSlotGigabyteLagLimit(val *float64)
	SlotGigabyteLagLimitInput() *float64
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
	PutConnection(value *DatatransferEndpointSettingsPostgresSourceConnection)
	PutObjectTransferSettings(value *DatatransferEndpointSettingsPostgresSourceObjectTransferSettings)
	PutPassword(value *DatatransferEndpointSettingsPostgresSourcePassword)
	ResetConnection()
	ResetDatabase()
	ResetExcludeTables()
	ResetIncludeTables()
	ResetObjectTransferSettings()
	ResetPassword()
	ResetSecurityGroups()
	ResetServiceSchema()
	ResetSlotGigabyteLagLimit()
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

// The jsii proxy struct for DatatransferEndpointSettingsPostgresSourceOutputReference
type jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) Connection() DatatransferEndpointSettingsPostgresSourceConnectionOutputReference {
	var returns DatatransferEndpointSettingsPostgresSourceConnectionOutputReference
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) ConnectionInput() *DatatransferEndpointSettingsPostgresSourceConnection {
	var returns *DatatransferEndpointSettingsPostgresSourceConnection
	_jsii_.Get(
		j,
		"connectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) ExcludeTables() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeTables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) ExcludeTablesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeTablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) IncludeTables() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeTables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) IncludeTablesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeTablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) InternalValue() *DatatransferEndpointSettingsPostgresSource {
	var returns *DatatransferEndpointSettingsPostgresSource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) ObjectTransferSettings() DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference {
	var returns DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference
	_jsii_.Get(
		j,
		"objectTransferSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) ObjectTransferSettingsInput() *DatatransferEndpointSettingsPostgresSourceObjectTransferSettings {
	var returns *DatatransferEndpointSettingsPostgresSourceObjectTransferSettings
	_jsii_.Get(
		j,
		"objectTransferSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) Password() DatatransferEndpointSettingsPostgresSourcePasswordOutputReference {
	var returns DatatransferEndpointSettingsPostgresSourcePasswordOutputReference
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) PasswordInput() *DatatransferEndpointSettingsPostgresSourcePassword {
	var returns *DatatransferEndpointSettingsPostgresSourcePassword
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) ServiceSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) ServiceSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) SlotGigabyteLagLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"slotGigabyteLagLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) SlotGigabyteLagLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"slotGigabyteLagLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) UserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}


func NewDatatransferEndpointSettingsPostgresSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatatransferEndpointSettingsPostgresSourceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsPostgresSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatatransferEndpointSettingsPostgresSourceOutputReference_Override(d DatatransferEndpointSettingsPostgresSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsPostgresSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) SetDatabase(val *string) {
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) SetExcludeTables(val *[]*string) {
	_jsii_.Set(
		j,
		"excludeTables",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) SetIncludeTables(val *[]*string) {
	_jsii_.Set(
		j,
		"includeTables",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) SetInternalValue(val *DatatransferEndpointSettingsPostgresSource) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) SetSecurityGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) SetServiceSchema(val *string) {
	_jsii_.Set(
		j,
		"serviceSchema",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) SetSlotGigabyteLagLimit(val *float64) {
	_jsii_.Set(
		j,
		"slotGigabyteLagLimit",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) SetUser(val *string) {
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) PutConnection(value *DatatransferEndpointSettingsPostgresSourceConnection) {
	_jsii_.InvokeVoid(
		d,
		"putConnection",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) PutObjectTransferSettings(value *DatatransferEndpointSettingsPostgresSourceObjectTransferSettings) {
	_jsii_.InvokeVoid(
		d,
		"putObjectTransferSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) PutPassword(value *DatatransferEndpointSettingsPostgresSourcePassword) {
	_jsii_.InvokeVoid(
		d,
		"putPassword",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) ResetConnection() {
	_jsii_.InvokeVoid(
		d,
		"resetConnection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) ResetDatabase() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) ResetExcludeTables() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludeTables",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) ResetIncludeTables() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeTables",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) ResetObjectTransferSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetObjectTransferSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) ResetServiceSchema() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceSchema",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) ResetSlotGigabyteLagLimit() {
	_jsii_.InvokeVoid(
		d,
		"resetSlotGigabyteLagLimit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) ResetUser() {
	_jsii_.InvokeVoid(
		d,
		"resetUser",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

