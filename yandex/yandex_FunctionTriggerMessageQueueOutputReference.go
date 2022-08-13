// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FunctionTriggerMessageQueueOutputReference interface {
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
	InternalValue() *FunctionTriggerMessageQueue
	SetInternalValue(val *FunctionTriggerMessageQueue)
	QueueId() *string
	SetQueueId(val *string)
	QueueIdInput() *string
	ServiceAccountId() *string
	SetServiceAccountId(val *string)
	ServiceAccountIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VisibilityTimeout() *string
	SetVisibilityTimeout(val *string)
	VisibilityTimeoutInput() *string
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
	ResetVisibilityTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FunctionTriggerMessageQueueOutputReference
type jsiiProxy_FunctionTriggerMessageQueueOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FunctionTriggerMessageQueueOutputReference) BatchCutoff() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchCutoff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerMessageQueueOutputReference) BatchCutoffInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchCutoffInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerMessageQueueOutputReference) BatchSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerMessageQueueOutputReference) BatchSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerMessageQueueOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerMessageQueueOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerMessageQueueOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerMessageQueueOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerMessageQueueOutputReference) InternalValue() *FunctionTriggerMessageQueue {
	var returns *FunctionTriggerMessageQueue
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerMessageQueueOutputReference) QueueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerMessageQueueOutputReference) QueueIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerMessageQueueOutputReference) ServiceAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerMessageQueueOutputReference) ServiceAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerMessageQueueOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerMessageQueueOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerMessageQueueOutputReference) VisibilityTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibilityTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTriggerMessageQueueOutputReference) VisibilityTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibilityTimeoutInput",
		&returns,
	)
	return returns
}


func NewFunctionTriggerMessageQueueOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FunctionTriggerMessageQueueOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FunctionTriggerMessageQueueOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.FunctionTriggerMessageQueueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFunctionTriggerMessageQueueOutputReference_Override(f FunctionTriggerMessageQueueOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.FunctionTriggerMessageQueueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FunctionTriggerMessageQueueOutputReference) SetBatchCutoff(val *string) {
	_jsii_.Set(
		j,
		"batchCutoff",
		val,
	)
}

func (j *jsiiProxy_FunctionTriggerMessageQueueOutputReference) SetBatchSize(val *string) {
	_jsii_.Set(
		j,
		"batchSize",
		val,
	)
}

func (j *jsiiProxy_FunctionTriggerMessageQueueOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FunctionTriggerMessageQueueOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FunctionTriggerMessageQueueOutputReference) SetInternalValue(val *FunctionTriggerMessageQueue) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FunctionTriggerMessageQueueOutputReference) SetQueueId(val *string) {
	_jsii_.Set(
		j,
		"queueId",
		val,
	)
}

func (j *jsiiProxy_FunctionTriggerMessageQueueOutputReference) SetServiceAccountId(val *string) {
	_jsii_.Set(
		j,
		"serviceAccountId",
		val,
	)
}

func (j *jsiiProxy_FunctionTriggerMessageQueueOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FunctionTriggerMessageQueueOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FunctionTriggerMessageQueueOutputReference) SetVisibilityTimeout(val *string) {
	_jsii_.Set(
		j,
		"visibilityTimeout",
		val,
	)
}

func (f *jsiiProxy_FunctionTriggerMessageQueueOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerMessageQueueOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerMessageQueueOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerMessageQueueOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerMessageQueueOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerMessageQueueOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerMessageQueueOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerMessageQueueOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerMessageQueueOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerMessageQueueOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerMessageQueueOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerMessageQueueOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerMessageQueueOutputReference) ResetBatchSize() {
	_jsii_.InvokeVoid(
		f,
		"resetBatchSize",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionTriggerMessageQueueOutputReference) ResetVisibilityTimeout() {
	_jsii_.InvokeVoid(
		f,
		"resetVisibilityTimeout",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionTriggerMessageQueueOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTriggerMessageQueueOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

