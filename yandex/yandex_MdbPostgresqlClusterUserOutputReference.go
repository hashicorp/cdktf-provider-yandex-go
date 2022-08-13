// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbPostgresqlClusterUserOutputReference interface {
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
	ConnLimit() *float64
	SetConnLimit(val *float64)
	ConnLimitInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Grants() *[]*string
	SetGrants(val *[]*string)
	GrantsInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Login() interface{}
	SetLogin(val interface{})
	LoginInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	Permission() MdbPostgresqlClusterUserPermissionList
	PermissionInput() interface{}
	Settings() *map[string]*string
	SetSettings(val *map[string]*string)
	SettingsInput() *map[string]*string
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
	PutPermission(value interface{})
	ResetConnLimit()
	ResetGrants()
	ResetLogin()
	ResetPermission()
	ResetSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MdbPostgresqlClusterUserOutputReference
type jsiiProxy_MdbPostgresqlClusterUserOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) ConnLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) ConnLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) Grants() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"grants",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) GrantsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"grantsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) Login() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"login",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) LoginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) Permission() MdbPostgresqlClusterUserPermissionList {
	var returns MdbPostgresqlClusterUserPermissionList
	_jsii_.Get(
		j,
		"permission",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) PermissionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) Settings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) SettingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"settingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMdbPostgresqlClusterUserOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MdbPostgresqlClusterUserOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MdbPostgresqlClusterUserOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbPostgresqlClusterUserOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMdbPostgresqlClusterUserOutputReference_Override(m MdbPostgresqlClusterUserOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbPostgresqlClusterUserOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) SetConnLimit(val *float64) {
	_jsii_.Set(
		j,
		"connLimit",
		val,
	)
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) SetGrants(val *[]*string) {
	_jsii_.Set(
		j,
		"grants",
		val,
	)
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) SetLogin(val interface{}) {
	_jsii_.Set(
		j,
		"login",
		val,
	)
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) SetSettings(val *map[string]*string) {
	_jsii_.Set(
		j,
		"settings",
		val,
	)
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbPostgresqlClusterUserOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MdbPostgresqlClusterUserOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterUserOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterUserOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterUserOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterUserOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterUserOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterUserOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterUserOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterUserOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterUserOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterUserOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterUserOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterUserOutputReference) PutPermission(value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"putPermission",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbPostgresqlClusterUserOutputReference) ResetConnLimit() {
	_jsii_.InvokeVoid(
		m,
		"resetConnLimit",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbPostgresqlClusterUserOutputReference) ResetGrants() {
	_jsii_.InvokeVoid(
		m,
		"resetGrants",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbPostgresqlClusterUserOutputReference) ResetLogin() {
	_jsii_.InvokeVoid(
		m,
		"resetLogin",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbPostgresqlClusterUserOutputReference) ResetPermission() {
	_jsii_.InvokeVoid(
		m,
		"resetPermission",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbPostgresqlClusterUserOutputReference) ResetSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbPostgresqlClusterUserOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterUserOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

