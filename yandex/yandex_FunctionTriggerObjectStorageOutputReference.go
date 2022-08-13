// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FunctionTriggerObjectStorageOutputReference interface {
	cdktf.ComplexObject
	BucketId() *string
	SetBucketId(val *string)
	BucketIdInput() *string
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
	Create() interface{}
	SetCreate(val interface{})
	CreateInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Delete() interface{}
	SetDelete(val interface{})
	DeleteInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *FunctionTriggerObjectStorage
	SetInternalValue(val *FunctionTriggerObjectStorage)
	Prefix() *string
	SetPrefix(val *string)
	PrefixInput() *string
	Suffix() *string
	SetSuffix(val *string)
	SuffixInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Update() interface{}
	SetUpdate(val interface{})
	UpdateInput() interface{}
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
	ResetCreate()
	ResetDelete()
	ResetPrefix()
	ResetSuffix()
	ResetUpdate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FunctionTriggerObjectStorageOutputReference
type jsiiProxy_FunctionTriggerObjectStorageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FunctionTriggerObjectStorageOutputReference) BucketId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerObjectStorageOutputReference) BucketIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerObjectStorageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerObjectStorageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerObjectStorageOutputReference) Create() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerObjectStorageOutputReference) CreateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerObjectStorageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerObjectStorageOutputReference) Delete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerObjectStorageOutputReference) DeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerObjectStorageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerObjectStorageOutputReference) InternalValue() *FunctionTriggerObjectStorage {
	var returns *FunctionTriggerObjectStorage
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerObjectStorageOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerObjectStorageOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerObjectStorageOutputReference) Suffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerObjectStorageOutputReference) SuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerObjectStorageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerObjectStorageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerObjectStorageOutputReference) Update() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerObjectStorageOutputReference) UpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewFunctionTriggerObjectStorageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FunctionTriggerObjectStorageOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FunctionTriggerObjectStorageOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.FunctionTriggerObjectStorageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFunctionTriggerObjectStorageOutputReference_Override(f FunctionTriggerObjectStorageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.FunctionTriggerObjectStorageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FunctionTriggerObjectStorageOutputReference) SetBucketId(val *string) {
	_jsii_.Set(
		j,
		"bucketId",
		val,
	)
}

func (j *jsiiProxy_FunctionTriggerObjectStorageOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FunctionTriggerObjectStorageOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FunctionTriggerObjectStorageOutputReference) SetCreate(val interface{}) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_FunctionTriggerObjectStorageOutputReference) SetDelete(val interface{}) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_FunctionTriggerObjectStorageOutputReference) SetInternalValue(val *FunctionTriggerObjectStorage) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FunctionTriggerObjectStorageOutputReference) SetPrefix(val *string) {
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_FunctionTriggerObjectStorageOutputReference) SetSuffix(val *string) {
	_jsii_.Set(
		j,
		"suffix",
		val,
	)
}

func (j *jsiiProxy_FunctionTriggerObjectStorageOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FunctionTriggerObjectStorageOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FunctionTriggerObjectStorageOutputReference) SetUpdate(val interface{}) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (f *jsiiProxy_FunctionTriggerObjectStorageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerObjectStorageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerObjectStorageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerObjectStorageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerObjectStorageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerObjectStorageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerObjectStorageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerObjectStorageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerObjectStorageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerObjectStorageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerObjectStorageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerObjectStorageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerObjectStorageOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		f,
		"resetCreate",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionTriggerObjectStorageOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		f,
		"resetDelete",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionTriggerObjectStorageOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		f,
		"resetPrefix",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionTriggerObjectStorageOutputReference) ResetSuffix() {
	_jsii_.InvokeVoid(
		f,
		"resetSuffix",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionTriggerObjectStorageOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		f,
		"resetUpdate",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionTriggerObjectStorageOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerObjectStorageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

