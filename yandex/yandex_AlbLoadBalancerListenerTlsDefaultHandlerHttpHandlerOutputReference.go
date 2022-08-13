// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference interface {
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
	Http2Options() AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerHttp2OptionsOutputReference
	Http2OptionsInput() *AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerHttp2Options
	HttpRouterId() *string
	SetHttpRouterId(val *string)
	HttpRouterIdInput() *string
	InternalValue() *AlbLoadBalancerListenerTlsDefaultHandlerHttpHandler
	SetInternalValue(val *AlbLoadBalancerListenerTlsDefaultHandlerHttpHandler)
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
	PutHttp2Options(value *AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerHttp2Options)
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

// The jsii proxy struct for AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference
type jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) AllowHttp10() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowHttp10",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) AllowHttp10Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowHttp10Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) Http2Options() AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerHttp2OptionsOutputReference {
	var returns AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerHttp2OptionsOutputReference
	_jsii_.Get(
		j,
		"http2Options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) Http2OptionsInput() *AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerHttp2Options {
	var returns *AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerHttp2Options
	_jsii_.Get(
		j,
		"http2OptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) HttpRouterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpRouterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) HttpRouterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpRouterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) InternalValue() *AlbLoadBalancerListenerTlsDefaultHandlerHttpHandler {
	var returns *AlbLoadBalancerListenerTlsDefaultHandlerHttpHandler
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference_Override(a AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) SetAllowHttp10(val interface{}) {
	_jsii_.Set(
		j,
		"allowHttp10",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) SetHttpRouterId(val *string) {
	_jsii_.Set(
		j,
		"httpRouterId",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) SetInternalValue(val *AlbLoadBalancerListenerTlsDefaultHandlerHttpHandler) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) PutHttp2Options(value *AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerHttp2Options) {
	_jsii_.InvokeVoid(
		a,
		"putHttp2Options",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) ResetAllowHttp10() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowHttp10",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) ResetHttp2Options() {
	_jsii_.InvokeVoid(
		a,
		"resetHttp2Options",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) ResetHttpRouterId() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpRouterId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

