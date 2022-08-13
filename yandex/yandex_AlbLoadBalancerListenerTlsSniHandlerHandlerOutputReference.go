// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference interface {
	cdktf.ComplexObject
	CertificateIds() *[]*string
	SetCertificateIds(val *[]*string)
	CertificateIdsInput() *[]*string
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
	HttpHandler() AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference
	HttpHandlerInput() *AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandler
	InternalValue() *AlbLoadBalancerListenerTlsSniHandlerHandler
	SetInternalValue(val *AlbLoadBalancerListenerTlsSniHandlerHandler)
	StreamHandler() AlbLoadBalancerListenerTlsSniHandlerHandlerStreamHandlerOutputReference
	StreamHandlerInput() *AlbLoadBalancerListenerTlsSniHandlerHandlerStreamHandler
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
	PutHttpHandler(value *AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandler)
	PutStreamHandler(value *AlbLoadBalancerListenerTlsSniHandlerHandlerStreamHandler)
	ResetHttpHandler()
	ResetStreamHandler()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference
type jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) CertificateIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"certificateIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) CertificateIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"certificateIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) HttpHandler() AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference {
	var returns AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference
	_jsii_.Get(
		j,
		"httpHandler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) HttpHandlerInput() *AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandler {
	var returns *AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandler
	_jsii_.Get(
		j,
		"httpHandlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) InternalValue() *AlbLoadBalancerListenerTlsSniHandlerHandler {
	var returns *AlbLoadBalancerListenerTlsSniHandlerHandler
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) StreamHandler() AlbLoadBalancerListenerTlsSniHandlerHandlerStreamHandlerOutputReference {
	var returns AlbLoadBalancerListenerTlsSniHandlerHandlerStreamHandlerOutputReference
	_jsii_.Get(
		j,
		"streamHandler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) StreamHandlerInput() *AlbLoadBalancerListenerTlsSniHandlerHandlerStreamHandler {
	var returns *AlbLoadBalancerListenerTlsSniHandlerHandlerStreamHandler
	_jsii_.Get(
		j,
		"streamHandlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference_Override(a AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) SetCertificateIds(val *[]*string) {
	_jsii_.Set(
		j,
		"certificateIds",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) SetInternalValue(val *AlbLoadBalancerListenerTlsSniHandlerHandler) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) PutHttpHandler(value *AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandler) {
	_jsii_.InvokeVoid(
		a,
		"putHttpHandler",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) PutStreamHandler(value *AlbLoadBalancerListenerTlsSniHandlerHandlerStreamHandler) {
	_jsii_.InvokeVoid(
		a,
		"putStreamHandler",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) ResetHttpHandler() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpHandler",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) ResetStreamHandler() {
	_jsii_.InvokeVoid(
		a,
		"resetStreamHandler",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsSniHandlerHandlerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

