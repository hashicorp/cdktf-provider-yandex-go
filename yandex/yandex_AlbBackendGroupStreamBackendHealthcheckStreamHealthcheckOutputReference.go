// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference interface {
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
	InternalValue() *AlbBackendGroupStreamBackendHealthcheckStreamHealthcheck
	SetInternalValue(val *AlbBackendGroupStreamBackendHealthcheckStreamHealthcheck)
	Receive() *string
	SetReceive(val *string)
	ReceiveInput() *string
	Send() *string
	SetSend(val *string)
	SendInput() *string
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
	ResetReceive()
	ResetSend()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference
type jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) InternalValue() *AlbBackendGroupStreamBackendHealthcheckStreamHealthcheck {
	var returns *AlbBackendGroupStreamBackendHealthcheckStreamHealthcheck
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) Receive() *string {
	var returns *string
	_jsii_.Get(
		j,
		"receive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) ReceiveInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"receiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) Send() *string {
	var returns *string
	_jsii_.Get(
		j,
		"send",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) SendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference_Override(a AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) SetInternalValue(val *AlbBackendGroupStreamBackendHealthcheckStreamHealthcheck) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) SetReceive(val *string) {
	_jsii_.Set(
		j,
		"receive",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) SetSend(val *string) {
	_jsii_.Set(
		j,
		"send",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) ResetReceive() {
	_jsii_.InvokeVoid(
		a,
		"resetReceive",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) ResetSend() {
	_jsii_.InvokeVoid(
		a,
		"resetSend",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

