// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbBackendGroupSessionAffinityOutputReference interface {
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
	Connection() AlbBackendGroupSessionAffinityConnectionOutputReference
	ConnectionInput() *AlbBackendGroupSessionAffinityConnection
	Cookie() AlbBackendGroupSessionAffinityCookieOutputReference
	CookieInput() *AlbBackendGroupSessionAffinityCookie
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Header() AlbBackendGroupSessionAffinityHeaderOutputReference
	HeaderInput() *AlbBackendGroupSessionAffinityHeader
	InternalValue() *AlbBackendGroupSessionAffinity
	SetInternalValue(val *AlbBackendGroupSessionAffinity)
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
	PutConnection(value *AlbBackendGroupSessionAffinityConnection)
	PutCookie(value *AlbBackendGroupSessionAffinityCookie)
	PutHeader(value *AlbBackendGroupSessionAffinityHeader)
	ResetConnection()
	ResetCookie()
	ResetHeader()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlbBackendGroupSessionAffinityOutputReference
type jsiiProxy_AlbBackendGroupSessionAffinityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) Connection() AlbBackendGroupSessionAffinityConnectionOutputReference {
	var returns AlbBackendGroupSessionAffinityConnectionOutputReference
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) ConnectionInput() *AlbBackendGroupSessionAffinityConnection {
	var returns *AlbBackendGroupSessionAffinityConnection
	_jsii_.Get(
		j,
		"connectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) Cookie() AlbBackendGroupSessionAffinityCookieOutputReference {
	var returns AlbBackendGroupSessionAffinityCookieOutputReference
	_jsii_.Get(
		j,
		"cookie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) CookieInput() *AlbBackendGroupSessionAffinityCookie {
	var returns *AlbBackendGroupSessionAffinityCookie
	_jsii_.Get(
		j,
		"cookieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) Header() AlbBackendGroupSessionAffinityHeaderOutputReference {
	var returns AlbBackendGroupSessionAffinityHeaderOutputReference
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) HeaderInput() *AlbBackendGroupSessionAffinityHeader {
	var returns *AlbBackendGroupSessionAffinityHeader
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) InternalValue() *AlbBackendGroupSessionAffinity {
	var returns *AlbBackendGroupSessionAffinity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAlbBackendGroupSessionAffinityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlbBackendGroupSessionAffinityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AlbBackendGroupSessionAffinityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbBackendGroupSessionAffinityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlbBackendGroupSessionAffinityOutputReference_Override(a AlbBackendGroupSessionAffinityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbBackendGroupSessionAffinityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) SetInternalValue(val *AlbBackendGroupSessionAffinity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) PutConnection(value *AlbBackendGroupSessionAffinityConnection) {
	_jsii_.InvokeVoid(
		a,
		"putConnection",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) PutCookie(value *AlbBackendGroupSessionAffinityCookie) {
	_jsii_.InvokeVoid(
		a,
		"putCookie",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) PutHeader(value *AlbBackendGroupSessionAffinityHeader) {
	_jsii_.InvokeVoid(
		a,
		"putHeader",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) ResetConnection() {
	_jsii_.InvokeVoid(
		a,
		"resetConnection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) ResetCookie() {
	_jsii_.InvokeVoid(
		a,
		"resetCookie",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) ResetHeader() {
	_jsii_.InvokeVoid(
		a,
		"resetHeader",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupSessionAffinityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

