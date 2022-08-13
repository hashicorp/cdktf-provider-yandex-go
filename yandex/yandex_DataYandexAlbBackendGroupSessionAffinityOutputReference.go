// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexAlbBackendGroupSessionAffinityOutputReference interface {
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
	Connection() DataYandexAlbBackendGroupSessionAffinityConnectionOutputReference
	ConnectionInput() *DataYandexAlbBackendGroupSessionAffinityConnection
	Cookie() DataYandexAlbBackendGroupSessionAffinityCookieOutputReference
	CookieInput() *DataYandexAlbBackendGroupSessionAffinityCookie
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Header() DataYandexAlbBackendGroupSessionAffinityHeaderOutputReference
	HeaderInput() *DataYandexAlbBackendGroupSessionAffinityHeader
	InternalValue() *DataYandexAlbBackendGroupSessionAffinity
	SetInternalValue(val *DataYandexAlbBackendGroupSessionAffinity)
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
	PutConnection(value *DataYandexAlbBackendGroupSessionAffinityConnection)
	PutCookie(value *DataYandexAlbBackendGroupSessionAffinityCookie)
	PutHeader(value *DataYandexAlbBackendGroupSessionAffinityHeader)
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

// The jsii proxy struct for DataYandexAlbBackendGroupSessionAffinityOutputReference
type jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) Connection() DataYandexAlbBackendGroupSessionAffinityConnectionOutputReference {
	var returns DataYandexAlbBackendGroupSessionAffinityConnectionOutputReference
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) ConnectionInput() *DataYandexAlbBackendGroupSessionAffinityConnection {
	var returns *DataYandexAlbBackendGroupSessionAffinityConnection
	_jsii_.Get(
		j,
		"connectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) Cookie() DataYandexAlbBackendGroupSessionAffinityCookieOutputReference {
	var returns DataYandexAlbBackendGroupSessionAffinityCookieOutputReference
	_jsii_.Get(
		j,
		"cookie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) CookieInput() *DataYandexAlbBackendGroupSessionAffinityCookie {
	var returns *DataYandexAlbBackendGroupSessionAffinityCookie
	_jsii_.Get(
		j,
		"cookieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) Header() DataYandexAlbBackendGroupSessionAffinityHeaderOutputReference {
	var returns DataYandexAlbBackendGroupSessionAffinityHeaderOutputReference
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) HeaderInput() *DataYandexAlbBackendGroupSessionAffinityHeader {
	var returns *DataYandexAlbBackendGroupSessionAffinityHeader
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) InternalValue() *DataYandexAlbBackendGroupSessionAffinity {
	var returns *DataYandexAlbBackendGroupSessionAffinity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataYandexAlbBackendGroupSessionAffinityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataYandexAlbBackendGroupSessionAffinityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexAlbBackendGroupSessionAffinityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataYandexAlbBackendGroupSessionAffinityOutputReference_Override(d DataYandexAlbBackendGroupSessionAffinityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexAlbBackendGroupSessionAffinityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) SetInternalValue(val *DataYandexAlbBackendGroupSessionAffinity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) PutConnection(value *DataYandexAlbBackendGroupSessionAffinityConnection) {
	_jsii_.InvokeVoid(
		d,
		"putConnection",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) PutCookie(value *DataYandexAlbBackendGroupSessionAffinityCookie) {
	_jsii_.InvokeVoid(
		d,
		"putCookie",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) PutHeader(value *DataYandexAlbBackendGroupSessionAffinityHeader) {
	_jsii_.InvokeVoid(
		d,
		"putHeader",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) ResetConnection() {
	_jsii_.InvokeVoid(
		d,
		"resetConnection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) ResetCookie() {
	_jsii_.InvokeVoid(
		d,
		"resetCookie",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) ResetHeader() {
	_jsii_.InvokeVoid(
		d,
		"resetHeader",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupSessionAffinityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

