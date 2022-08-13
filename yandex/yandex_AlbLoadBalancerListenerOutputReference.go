// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbLoadBalancerListenerOutputReference interface {
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
	Endpoint() AlbLoadBalancerListenerEndpointList
	EndpointInput() interface{}
	// Experimental.
	Fqn() *string
	Http() AlbLoadBalancerListenerHttpOutputReference
	HttpInput() *AlbLoadBalancerListenerHttp
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Stream() AlbLoadBalancerListenerStreamOutputReference
	StreamInput() *AlbLoadBalancerListenerStream
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Tls() AlbLoadBalancerListenerTlsOutputReference
	TlsInput() *AlbLoadBalancerListenerTls
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
	PutEndpoint(value interface{})
	PutHttp(value *AlbLoadBalancerListenerHttp)
	PutStream(value *AlbLoadBalancerListenerStream)
	PutTls(value *AlbLoadBalancerListenerTls)
	ResetEndpoint()
	ResetHttp()
	ResetStream()
	ResetTls()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlbLoadBalancerListenerOutputReference
type jsiiProxy_AlbLoadBalancerListenerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlbLoadBalancerListenerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerOutputReference) Endpoint() AlbLoadBalancerListenerEndpointList {
	var returns AlbLoadBalancerListenerEndpointList
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerOutputReference) EndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerOutputReference) Http() AlbLoadBalancerListenerHttpOutputReference {
	var returns AlbLoadBalancerListenerHttpOutputReference
	_jsii_.Get(
		j,
		"http",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerOutputReference) HttpInput() *AlbLoadBalancerListenerHttp {
	var returns *AlbLoadBalancerListenerHttp
	_jsii_.Get(
		j,
		"httpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerOutputReference) Stream() AlbLoadBalancerListenerStreamOutputReference {
	var returns AlbLoadBalancerListenerStreamOutputReference
	_jsii_.Get(
		j,
		"stream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerOutputReference) StreamInput() *AlbLoadBalancerListenerStream {
	var returns *AlbLoadBalancerListenerStream
	_jsii_.Get(
		j,
		"streamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerOutputReference) Tls() AlbLoadBalancerListenerTlsOutputReference {
	var returns AlbLoadBalancerListenerTlsOutputReference
	_jsii_.Get(
		j,
		"tls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerOutputReference) TlsInput() *AlbLoadBalancerListenerTls {
	var returns *AlbLoadBalancerListenerTls
	_jsii_.Get(
		j,
		"tlsInput",
		&returns,
	)
	return returns
}


func NewAlbLoadBalancerListenerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AlbLoadBalancerListenerOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AlbLoadBalancerListenerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbLoadBalancerListenerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAlbLoadBalancerListenerOutputReference_Override(a AlbLoadBalancerListenerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbLoadBalancerListenerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerOutputReference) PutEndpoint(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putEndpoint",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerOutputReference) PutHttp(value *AlbLoadBalancerListenerHttp) {
	_jsii_.InvokeVoid(
		a,
		"putHttp",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerOutputReference) PutStream(value *AlbLoadBalancerListenerStream) {
	_jsii_.InvokeVoid(
		a,
		"putStream",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerOutputReference) PutTls(value *AlbLoadBalancerListenerTls) {
	_jsii_.InvokeVoid(
		a,
		"putTls",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerOutputReference) ResetEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerOutputReference) ResetHttp() {
	_jsii_.InvokeVoid(
		a,
		"resetHttp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerOutputReference) ResetStream() {
	_jsii_.InvokeVoid(
		a,
		"resetStream",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerOutputReference) ResetTls() {
	_jsii_.InvokeVoid(
		a,
		"resetTls",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

