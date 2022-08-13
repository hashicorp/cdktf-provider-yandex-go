// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FunctionTriggerLoggingOutputReference interface {
	cdktf.ComplexObject
	BatchCutoff() *string
	SetBatchCutoff(val *string)
	BatchCutoffInput() *string
	BatchSize() *string
	SetBatchSize(val *string)
	BatchSizeInput() *string
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
	GroupId() *string
	SetGroupId(val *string)
	GroupIdInput() *string
	InternalValue() *FunctionTriggerLogging
	SetInternalValue(val *FunctionTriggerLogging)
	Levels() *[]*string
	SetLevels(val *[]*string)
	LevelsInput() *[]*string
	ResourceIds() *[]*string
	SetResourceIds(val *[]*string)
	ResourceIdsInput() *[]*string
	ResourceTypes() *[]*string
	SetResourceTypes(val *[]*string)
	ResourceTypesInput() *[]*string
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
	ResetBatchSize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FunctionTriggerLoggingOutputReference
type jsiiProxy_FunctionTriggerLoggingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FunctionTriggerLoggingOutputReference) BatchCutoff() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchCutoff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerLoggingOutputReference) BatchCutoffInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchCutoffInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerLoggingOutputReference) BatchSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerLoggingOutputReference) BatchSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerLoggingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerLoggingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerLoggingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerLoggingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerLoggingOutputReference) GroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerLoggingOutputReference) GroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerLoggingOutputReference) InternalValue() *FunctionTriggerLogging {
	var returns *FunctionTriggerLogging
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerLoggingOutputReference) Levels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"levels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerLoggingOutputReference) LevelsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"levelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerLoggingOutputReference) ResourceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerLoggingOutputReference) ResourceIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerLoggingOutputReference) ResourceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerLoggingOutputReference) ResourceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerLoggingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerLoggingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFunctionTriggerLoggingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FunctionTriggerLoggingOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FunctionTriggerLoggingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.FunctionTriggerLoggingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFunctionTriggerLoggingOutputReference_Override(f FunctionTriggerLoggingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.FunctionTriggerLoggingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FunctionTriggerLoggingOutputReference) SetBatchCutoff(val *string) {
	_jsii_.Set(
		j,
		"batchCutoff",
		val,
	)
}

func (j *jsiiProxy_FunctionTriggerLoggingOutputReference) SetBatchSize(val *string) {
	_jsii_.Set(
		j,
		"batchSize",
		val,
	)
}

func (j *jsiiProxy_FunctionTriggerLoggingOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FunctionTriggerLoggingOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FunctionTriggerLoggingOutputReference) SetGroupId(val *string) {
	_jsii_.Set(
		j,
		"groupId",
		val,
	)
}

func (j *jsiiProxy_FunctionTriggerLoggingOutputReference) SetInternalValue(val *FunctionTriggerLogging) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FunctionTriggerLoggingOutputReference) SetLevels(val *[]*string) {
	_jsii_.Set(
		j,
		"levels",
		val,
	)
}

func (j *jsiiProxy_FunctionTriggerLoggingOutputReference) SetResourceIds(val *[]*string) {
	_jsii_.Set(
		j,
		"resourceIds",
		val,
	)
}

func (j *jsiiProxy_FunctionTriggerLoggingOutputReference) SetResourceTypes(val *[]*string) {
	_jsii_.Set(
		j,
		"resourceTypes",
		val,
	)
}

func (j *jsiiProxy_FunctionTriggerLoggingOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FunctionTriggerLoggingOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FunctionTriggerLoggingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerLoggingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerLoggingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerLoggingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerLoggingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerLoggingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerLoggingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerLoggingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerLoggingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerLoggingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerLoggingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerLoggingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerLoggingOutputReference) ResetBatchSize() {
	_jsii_.InvokeVoid(
		f,
		"resetBatchSize",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionTriggerLoggingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerLoggingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

