// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbVirtualHostModifyRequestHeadersOutputReference interface {
	cdktf.ComplexObject
	Append() *string
	SetAppend(val *string)
	AppendInput() *string
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Remove() interface{}
	SetRemove(val interface{})
	RemoveInput() interface{}
	Replace() *string
	SetReplace(val *string)
	ReplaceInput() *string
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
	ResetAppend()
	ResetRemove()
	ResetReplace()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlbVirtualHostModifyRequestHeadersOutputReference
type jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) Append() *string {
	var returns *string
	_jsii_.Get(
		j,
		"append",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) AppendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) Remove() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) RemoveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) Replace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) ReplaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAlbVirtualHostModifyRequestHeadersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AlbVirtualHostModifyRequestHeadersOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbVirtualHostModifyRequestHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAlbVirtualHostModifyRequestHeadersOutputReference_Override(a AlbVirtualHostModifyRequestHeadersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbVirtualHostModifyRequestHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) SetAppend(val *string) {
	_jsii_.Set(
		j,
		"append",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) SetRemove(val interface{}) {
	_jsii_.Set(
		j,
		"remove",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) SetReplace(val *string) {
	_jsii_.Set(
		j,
		"replace",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) ResetAppend() {
	_jsii_.InvokeVoid(
		a,
		"resetAppend",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) ResetRemove() {
	_jsii_.InvokeVoid(
		a,
		"resetRemove",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) ResetReplace() {
	_jsii_.InvokeVoid(
		a,
		"resetReplace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostModifyRequestHeadersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

