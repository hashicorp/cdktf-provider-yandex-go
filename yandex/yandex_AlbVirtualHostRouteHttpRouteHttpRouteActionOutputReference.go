// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference interface {
	cdktf.ComplexObject
	AutoHostRewrite() interface{}
	SetAutoHostRewrite(val interface{})
	AutoHostRewriteInput() interface{}
	BackendGroupId() *string
	SetBackendGroupId(val *string)
	BackendGroupIdInput() *string
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
	HostRewrite() *string
	SetHostRewrite(val *string)
	HostRewriteInput() *string
	IdleTimeout() *string
	SetIdleTimeout(val *string)
	IdleTimeoutInput() *string
	InternalValue() *AlbVirtualHostRouteHttpRouteHttpRouteAction
	SetInternalValue(val *AlbVirtualHostRouteHttpRouteHttpRouteAction)
	PrefixRewrite() *string
	SetPrefixRewrite(val *string)
	PrefixRewriteInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() *string
	SetTimeout(val *string)
	TimeoutInput() *string
	UpgradeTypes() *[]*string
	SetUpgradeTypes(val *[]*string)
	UpgradeTypesInput() *[]*string
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
	ResetAutoHostRewrite()
	ResetHostRewrite()
	ResetIdleTimeout()
	ResetPrefixRewrite()
	ResetTimeout()
	ResetUpgradeTypes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference
type jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) AutoHostRewrite() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHostRewrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) AutoHostRewriteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHostRewriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) BackendGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) BackendGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) HostRewrite() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostRewrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) HostRewriteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostRewriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) IdleTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idleTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) IdleTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idleTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) InternalValue() *AlbVirtualHostRouteHttpRouteHttpRouteAction {
	var returns *AlbVirtualHostRouteHttpRouteHttpRouteAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) PrefixRewrite() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixRewrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) PrefixRewriteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixRewriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) TimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) UpgradeTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"upgradeTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) UpgradeTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"upgradeTypesInput",
		&returns,
	)
	return returns
}


func NewAlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference_Override(a AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) SetAutoHostRewrite(val interface{}) {
	_jsii_.Set(
		j,
		"autoHostRewrite",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) SetBackendGroupId(val *string) {
	_jsii_.Set(
		j,
		"backendGroupId",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) SetHostRewrite(val *string) {
	_jsii_.Set(
		j,
		"hostRewrite",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) SetIdleTimeout(val *string) {
	_jsii_.Set(
		j,
		"idleTimeout",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) SetInternalValue(val *AlbVirtualHostRouteHttpRouteHttpRouteAction) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) SetPrefixRewrite(val *string) {
	_jsii_.Set(
		j,
		"prefixRewrite",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) SetTimeout(val *string) {
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) SetUpgradeTypes(val *[]*string) {
	_jsii_.Set(
		j,
		"upgradeTypes",
		val,
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) ResetAutoHostRewrite() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoHostRewrite",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) ResetHostRewrite() {
	_jsii_.InvokeVoid(
		a,
		"resetHostRewrite",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) ResetIdleTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetIdleTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) ResetPrefixRewrite() {
	_jsii_.InvokeVoid(
		a,
		"resetPrefixRewrite",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) ResetUpgradeTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetUpgradeTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

