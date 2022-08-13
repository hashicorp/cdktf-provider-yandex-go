// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbLoadBalancerListenerTlsDefaultHandlerOutputReference interface {
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
	HttpHandler() AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference
	HttpHandlerInput() *AlbLoadBalancerListenerTlsDefaultHandlerHttpHandler
	InternalValue() *AlbLoadBalancerListenerTlsDefaultHandler
	SetInternalValue(val *AlbLoadBalancerListenerTlsDefaultHandler)
	StreamHandler() AlbLoadBalancerListenerTlsDefaultHandlerStreamHandlerOutputReference
	StreamHandlerInput() *AlbLoadBalancerListenerTlsDefaultHandlerStreamHandler
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
	PutHttpHandler(value *AlbLoadBalancerListenerTlsDefaultHandlerHttpHandler)
	PutStreamHandler(value *AlbLoadBalancerListenerTlsDefaultHandlerStreamHandler)
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

// The jsii proxy struct for AlbLoadBalancerListenerTlsDefaultHandlerOutputReference
type jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) CertificateIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"certificateIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) CertificateIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"certificateIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) HttpHandler() AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference {
	var returns AlbLoadBalancerListenerTlsDefaultHandlerHttpHandlerOutputReference
	_jsii_.Get(
		j,
		"httpHandler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) HttpHandlerInput() *AlbLoadBalancerListenerTlsDefaultHandlerHttpHandler {
	var returns *AlbLoadBalancerListenerTlsDefaultHandlerHttpHandler
	_jsii_.Get(
		j,
		"httpHandlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) InternalValue() *AlbLoadBalancerListenerTlsDefaultHandler {
	var returns *AlbLoadBalancerListenerTlsDefaultHandler
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) StreamHandler() AlbLoadBalancerListenerTlsDefaultHandlerStreamHandlerOutputReference {
	var returns AlbLoadBalancerListenerTlsDefaultHandlerStreamHandlerOutputReference
	_jsii_.Get(
		j,
		"streamHandler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) StreamHandlerInput() *AlbLoadBalancerListenerTlsDefaultHandlerStreamHandler {
	var returns *AlbLoadBalancerListenerTlsDefaultHandlerStreamHandler
	_jsii_.Get(
		j,
		"streamHandlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAlbLoadBalancerListenerTlsDefaultHandlerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlbLoadBalancerListenerTlsDefaultHandlerOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbLoadBalancerListenerTlsDefaultHandlerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlbLoadBalancerListenerTlsDefaultHandlerOutputReference_Override(a AlbLoadBalancerListenerTlsDefaultHandlerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbLoadBalancerListenerTlsDefaultHandlerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) SetCertificateIds(val *[]*string) {
	_jsii_.Set(
		j,
		"certificateIds",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) SetInternalValue(val *AlbLoadBalancerListenerTlsDefaultHandler) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) PutHttpHandler(value *AlbLoadBalancerListenerTlsDefaultHandlerHttpHandler) {
	_jsii_.InvokeVoid(
		a,
		"putHttpHandler",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) PutStreamHandler(value *AlbLoadBalancerListenerTlsDefaultHandlerStreamHandler) {
	_jsii_.InvokeVoid(
		a,
		"putStreamHandler",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) ResetHttpHandler() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpHandler",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) ResetStreamHandler() {
	_jsii_.InvokeVoid(
		a,
		"resetStreamHandler",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerTlsDefaultHandlerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

