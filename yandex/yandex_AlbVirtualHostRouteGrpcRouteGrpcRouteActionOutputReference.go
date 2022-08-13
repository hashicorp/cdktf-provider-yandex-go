// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference interface {
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
	InternalValue() *AlbVirtualHostRouteGrpcRouteGrpcRouteAction
	SetInternalValue(val *AlbVirtualHostRouteGrpcRouteGrpcRouteAction)
	MaxTimeout() *string
	SetMaxTimeout(val *string)
	MaxTimeoutInput() *string
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
	ResetAutoHostRewrite()
	ResetHostRewrite()
	ResetIdleTimeout()
	ResetMaxTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference
type jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) AutoHostRewrite() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHostRewrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) AutoHostRewriteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHostRewriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) BackendGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) BackendGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) HostRewrite() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostRewrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) HostRewriteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostRewriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) IdleTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idleTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) IdleTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idleTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) InternalValue() *AlbVirtualHostRouteGrpcRouteGrpcRouteAction {
	var returns *AlbVirtualHostRouteGrpcRouteGrpcRouteAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) MaxTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) MaxTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference_Override(a AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) SetAutoHostRewrite(val interface{}) {
	_jsii_.Set(
		j,
		"autoHostRewrite",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) SetBackendGroupId(val *string) {
	_jsii_.Set(
		j,
		"backendGroupId",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) SetHostRewrite(val *string) {
	_jsii_.Set(
		j,
		"hostRewrite",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) SetIdleTimeout(val *string) {
	_jsii_.Set(
		j,
		"idleTimeout",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) SetInternalValue(val *AlbVirtualHostRouteGrpcRouteGrpcRouteAction) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) SetMaxTimeout(val *string) {
	_jsii_.Set(
		j,
		"maxTimeout",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) ResetAutoHostRewrite() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoHostRewrite",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) ResetHostRewrite() {
	_jsii_.InvokeVoid(
		a,
		"resetHostRewrite",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) ResetIdleTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetIdleTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) ResetMaxTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

