// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeInstanceGroupLoadBalancerOutputReference interface {
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
	InternalValue() *ComputeInstanceGroupLoadBalancer
	SetInternalValue(val *ComputeInstanceGroupLoadBalancer)
	MaxOpeningTrafficDuration() *float64
	SetMaxOpeningTrafficDuration(val *float64)
	MaxOpeningTrafficDurationInput() *float64
	StatusMessage() *string
	TargetGroupDescription() *string
	SetTargetGroupDescription(val *string)
	TargetGroupDescriptionInput() *string
	TargetGroupId() *string
	TargetGroupLabels() *map[string]*string
	SetTargetGroupLabels(val *map[string]*string)
	TargetGroupLabelsInput() *map[string]*string
	TargetGroupName() *string
	SetTargetGroupName(val *string)
	TargetGroupNameInput() *string
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
	ResetMaxOpeningTrafficDuration()
	ResetTargetGroupDescription()
	ResetTargetGroupLabels()
	ResetTargetGroupName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeInstanceGroupLoadBalancerOutputReference
type jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) InternalValue() *ComputeInstanceGroupLoadBalancer {
	var returns *ComputeInstanceGroupLoadBalancer
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) MaxOpeningTrafficDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxOpeningTrafficDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) MaxOpeningTrafficDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxOpeningTrafficDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) StatusMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) TargetGroupDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) TargetGroupDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) TargetGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) TargetGroupLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"targetGroupLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) TargetGroupLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"targetGroupLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) TargetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) TargetGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeInstanceGroupLoadBalancerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeInstanceGroupLoadBalancerOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.ComputeInstanceGroupLoadBalancerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeInstanceGroupLoadBalancerOutputReference_Override(c ComputeInstanceGroupLoadBalancerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.ComputeInstanceGroupLoadBalancerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) SetInternalValue(val *ComputeInstanceGroupLoadBalancer) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) SetMaxOpeningTrafficDuration(val *float64) {
	_jsii_.Set(
		j,
		"maxOpeningTrafficDuration",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) SetTargetGroupDescription(val *string) {
	_jsii_.Set(
		j,
		"targetGroupDescription",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) SetTargetGroupLabels(val *map[string]*string) {
	_jsii_.Set(
		j,
		"targetGroupLabels",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) SetTargetGroupName(val *string) {
	_jsii_.Set(
		j,
		"targetGroupName",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) ResetMaxOpeningTrafficDuration() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxOpeningTrafficDuration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) ResetTargetGroupDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetGroupDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) ResetTargetGroupLabels() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetGroupLabels",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) ResetTargetGroupName() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetGroupName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupLoadBalancerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

