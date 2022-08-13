// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2Options
	SetInternalValue(val *AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2Options)
	MaxConcurrentStreams() *float64
	SetMaxConcurrentStreams(val *float64)
	MaxConcurrentStreamsInput() *float64
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
	ResetMaxConcurrentStreams()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference
type jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference) InternalValue() *AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2Options {
	var returns *AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2Options
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference) MaxConcurrentStreams() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentStreams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference) MaxConcurrentStreamsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentStreamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference_Override(a AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference) SetInternalValue(val *AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2Options) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference) SetMaxConcurrentStreams(val *float64) {
	_jsii_.Set(
		j,
		"maxConcurrentStreams",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference) ResetMaxConcurrentStreams() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxConcurrentStreams",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

