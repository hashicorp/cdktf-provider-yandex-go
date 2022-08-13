// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbMysqlClusterUserOutputReference interface {
	cdktf.ComplexObject
	AuthenticationPlugin() *string
	SetAuthenticationPlugin(val *string)
	AuthenticationPluginInput() *string
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
	ConnectionLimits() MdbMysqlClusterUserConnectionLimitsOutputReference
	ConnectionLimitsInput() *MdbMysqlClusterUserConnectionLimits
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GlobalPermissions() *[]*string
	SetGlobalPermissions(val *[]*string)
	GlobalPermissionsInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	Permission() MdbMysqlClusterUserPermissionList
	PermissionInput() interface{}
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
	PutConnectionLimits(value *MdbMysqlClusterUserConnectionLimits)
	PutPermission(value interface{})
	ResetAuthenticationPlugin()
	ResetConnectionLimits()
	ResetGlobalPermissions()
	ResetPermission()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MdbMysqlClusterUserOutputReference
type jsiiProxy_MdbMysqlClusterUserOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MdbMysqlClusterUserOutputReference) AuthenticationPlugin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationPlugin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlClusterUserOutputReference) AuthenticationPluginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationPluginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlClusterUserOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlClusterUserOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlClusterUserOutputReference) ConnectionLimits() MdbMysqlClusterUserConnectionLimitsOutputReference {
	var returns MdbMysqlClusterUserConnectionLimitsOutputReference
	_jsii_.Get(
		j,
		"connectionLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlClusterUserOutputReference) ConnectionLimitsInput() *MdbMysqlClusterUserConnectionLimits {
	var returns *MdbMysqlClusterUserConnectionLimits
	_jsii_.Get(
		j,
		"connectionLimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlClusterUserOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlClusterUserOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlClusterUserOutputReference) GlobalPermissions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"globalPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlClusterUserOutputReference) GlobalPermissionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"globalPermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlClusterUserOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlClusterUserOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlClusterUserOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlClusterUserOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlClusterUserOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlClusterUserOutputReference) Permission() MdbMysqlClusterUserPermissionList {
	var returns MdbMysqlClusterUserPermissionList
	_jsii_.Get(
		j,
		"permission",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlClusterUserOutputReference) PermissionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlClusterUserOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlClusterUserOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMdbMysqlClusterUserOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MdbMysqlClusterUserOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MdbMysqlClusterUserOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbMysqlClusterUserOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMdbMysqlClusterUserOutputReference_Override(m MdbMysqlClusterUserOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbMysqlClusterUserOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MdbMysqlClusterUserOutputReference) SetAuthenticationPlugin(val *string) {
	_jsii_.Set(
		j,
		"authenticationPlugin",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlClusterUserOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlClusterUserOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlClusterUserOutputReference) SetGlobalPermissions(val *[]*string) {
	_jsii_.Set(
		j,
		"globalPermissions",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlClusterUserOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlClusterUserOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlClusterUserOutputReference) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlClusterUserOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlClusterUserOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MdbMysqlClusterUserOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlClusterUserOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlClusterUserOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlClusterUserOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlClusterUserOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlClusterUserOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlClusterUserOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlClusterUserOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlClusterUserOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlClusterUserOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlClusterUserOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlClusterUserOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlClusterUserOutputReference) PutConnectionLimits(value *MdbMysqlClusterUserConnectionLimits) {
	_jsii_.InvokeVoid(
		m,
		"putConnectionLimits",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbMysqlClusterUserOutputReference) PutPermission(value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"putPermission",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbMysqlClusterUserOutputReference) ResetAuthenticationPlugin() {
	_jsii_.InvokeVoid(
		m,
		"resetAuthenticationPlugin",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMysqlClusterUserOutputReference) ResetConnectionLimits() {
	_jsii_.InvokeVoid(
		m,
		"resetConnectionLimits",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMysqlClusterUserOutputReference) ResetGlobalPermissions() {
	_jsii_.InvokeVoid(
		m,
		"resetGlobalPermissions",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMysqlClusterUserOutputReference) ResetPermission() {
	_jsii_.InvokeVoid(
		m,
		"resetPermission",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMysqlClusterUserOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlClusterUserOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

