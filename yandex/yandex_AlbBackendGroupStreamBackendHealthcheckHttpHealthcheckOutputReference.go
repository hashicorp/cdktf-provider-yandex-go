// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference interface {
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
	Host() *string
	SetHost(val *string)
	HostInput() *string
	Http2() interface{}
	SetHttp2(val interface{})
	Http2Input() interface{}
	InternalValue() *AlbBackendGroupStreamBackendHealthcheckHttpHealthcheck
	SetInternalValue(val *AlbBackendGroupStreamBackendHealthcheckHttpHealthcheck)
	Path() *string
	SetPath(val *string)
	PathInput() *string
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
	ResetHost()
	ResetHttp2()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference
type jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) Http2() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) Http2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) InternalValue() *AlbBackendGroupStreamBackendHealthcheckHttpHealthcheck {
	var returns *AlbBackendGroupStreamBackendHealthcheckHttpHealthcheck
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference_Override(a AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) SetHost(val *string) {
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) SetHttp2(val interface{}) {
	_jsii_.Set(
		j,
		"http2",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) SetInternalValue(val *AlbBackendGroupStreamBackendHealthcheckHttpHealthcheck) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) SetPath(val *string) {
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) ResetHost() {
	_jsii_.InvokeVoid(
		a,
		"resetHost",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) ResetHttp2() {
	_jsii_.InvokeVoid(
		a,
		"resetHttp2",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

