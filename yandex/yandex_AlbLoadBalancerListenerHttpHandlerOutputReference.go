// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbLoadBalancerListenerHttpHandlerOutputReference interface {
	cdktf.ComplexObject
	AllowHttp10() interface{}
	SetAllowHttp10(val interface{})
	AllowHttp10Input() interface{}
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
	Http2Options() AlbLoadBalancerListenerHttpHandlerHttp2OptionsOutputReference
	Http2OptionsInput() *AlbLoadBalancerListenerHttpHandlerHttp2Options
	HttpRouterId() *string
	SetHttpRouterId(val *string)
	HttpRouterIdInput() *string
	InternalValue() *AlbLoadBalancerListenerHttpHandler
	SetInternalValue(val *AlbLoadBalancerListenerHttpHandler)
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
	PutHttp2Options(value *AlbLoadBalancerListenerHttpHandlerHttp2Options)
	ResetAllowHttp10()
	ResetHttp2Options()
	ResetHttpRouterId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlbLoadBalancerListenerHttpHandlerOutputReference
type jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) AllowHttp10() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowHttp10",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) AllowHttp10Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowHttp10Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) Http2Options() AlbLoadBalancerListenerHttpHandlerHttp2OptionsOutputReference {
	var returns AlbLoadBalancerListenerHttpHandlerHttp2OptionsOutputReference
	_jsii_.Get(
		j,
		"http2Options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) Http2OptionsInput() *AlbLoadBalancerListenerHttpHandlerHttp2Options {
	var returns *AlbLoadBalancerListenerHttpHandlerHttp2Options
	_jsii_.Get(
		j,
		"http2OptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) HttpRouterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpRouterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) HttpRouterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpRouterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) InternalValue() *AlbLoadBalancerListenerHttpHandler {
	var returns *AlbLoadBalancerListenerHttpHandler
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAlbLoadBalancerListenerHttpHandlerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlbLoadBalancerListenerHttpHandlerOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbLoadBalancerListenerHttpHandlerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlbLoadBalancerListenerHttpHandlerOutputReference_Override(a AlbLoadBalancerListenerHttpHandlerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbLoadBalancerListenerHttpHandlerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) SetAllowHttp10(val interface{}) {
	_jsii_.Set(
		j,
		"allowHttp10",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) SetHttpRouterId(val *string) {
	_jsii_.Set(
		j,
		"httpRouterId",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) SetInternalValue(val *AlbLoadBalancerListenerHttpHandler) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) PutHttp2Options(value *AlbLoadBalancerListenerHttpHandlerHttp2Options) {
	_jsii_.InvokeVoid(
		a,
		"putHttp2Options",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) ResetAllowHttp10() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowHttp10",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) ResetHttp2Options() {
	_jsii_.InvokeVoid(
		a,
		"resetHttp2Options",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) ResetHttpRouterId() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpRouterId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerHttpHandlerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

