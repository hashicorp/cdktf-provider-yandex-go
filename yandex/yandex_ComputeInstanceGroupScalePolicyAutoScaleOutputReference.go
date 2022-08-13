// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeInstanceGroupScalePolicyAutoScaleOutputReference interface {
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
	CpuUtilizationTarget() *float64
	SetCpuUtilizationTarget(val *float64)
	CpuUtilizationTargetInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomRule() ComputeInstanceGroupScalePolicyAutoScaleCustomRuleList
	CustomRuleInput() interface{}
	// Experimental.
	Fqn() *string
	InitialSize() *float64
	SetInitialSize(val *float64)
	InitialSizeInput() *float64
	InternalValue() *ComputeInstanceGroupScalePolicyAutoScale
	SetInternalValue(val *ComputeInstanceGroupScalePolicyAutoScale)
	MaxSize() *float64
	SetMaxSize(val *float64)
	MaxSizeInput() *float64
	MeasurementDuration() *float64
	SetMeasurementDuration(val *float64)
	MeasurementDurationInput() *float64
	MinZoneSize() *float64
	SetMinZoneSize(val *float64)
	MinZoneSizeInput() *float64
	StabilizationDuration() *float64
	SetStabilizationDuration(val *float64)
	StabilizationDurationInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WarmupDuration() *float64
	SetWarmupDuration(val *float64)
	WarmupDurationInput() *float64
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
	PutCustomRule(value interface{})
	ResetCpuUtilizationTarget()
	ResetCustomRule()
	ResetMaxSize()
	ResetMinZoneSize()
	ResetStabilizationDuration()
	ResetWarmupDuration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeInstanceGroupScalePolicyAutoScaleOutputReference
type jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) CpuUtilizationTarget() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuUtilizationTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) CpuUtilizationTargetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuUtilizationTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) CustomRule() ComputeInstanceGroupScalePolicyAutoScaleCustomRuleList {
	var returns ComputeInstanceGroupScalePolicyAutoScaleCustomRuleList
	_jsii_.Get(
		j,
		"customRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) CustomRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) InitialSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) InitialSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) InternalValue() *ComputeInstanceGroupScalePolicyAutoScale {
	var returns *ComputeInstanceGroupScalePolicyAutoScale
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) MaxSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) MaxSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) MeasurementDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"measurementDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) MeasurementDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"measurementDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) MinZoneSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minZoneSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) MinZoneSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minZoneSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) StabilizationDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stabilizationDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) StabilizationDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stabilizationDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) WarmupDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"warmupDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) WarmupDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"warmupDurationInput",
		&returns,
	)
	return returns
}


func NewComputeInstanceGroupScalePolicyAutoScaleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeInstanceGroupScalePolicyAutoScaleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.ComputeInstanceGroupScalePolicyAutoScaleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeInstanceGroupScalePolicyAutoScaleOutputReference_Override(c ComputeInstanceGroupScalePolicyAutoScaleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.ComputeInstanceGroupScalePolicyAutoScaleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) SetCpuUtilizationTarget(val *float64) {
	_jsii_.Set(
		j,
		"cpuUtilizationTarget",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) SetInitialSize(val *float64) {
	_jsii_.Set(
		j,
		"initialSize",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) SetInternalValue(val *ComputeInstanceGroupScalePolicyAutoScale) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) SetMaxSize(val *float64) {
	_jsii_.Set(
		j,
		"maxSize",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) SetMeasurementDuration(val *float64) {
	_jsii_.Set(
		j,
		"measurementDuration",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) SetMinZoneSize(val *float64) {
	_jsii_.Set(
		j,
		"minZoneSize",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) SetStabilizationDuration(val *float64) {
	_jsii_.Set(
		j,
		"stabilizationDuration",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) SetWarmupDuration(val *float64) {
	_jsii_.Set(
		j,
		"warmupDuration",
		val,
	)
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) PutCustomRule(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putCustomRule",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) ResetCpuUtilizationTarget() {
	_jsii_.InvokeVoid(
		c,
		"resetCpuUtilizationTarget",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) ResetCustomRule() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomRule",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) ResetMaxSize() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) ResetMinZoneSize() {
	_jsii_.InvokeVoid(
		c,
		"resetMinZoneSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) ResetStabilizationDuration() {
	_jsii_.InvokeVoid(
		c,
		"resetStabilizationDuration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) ResetWarmupDuration() {
	_jsii_.InvokeVoid(
		c,
		"resetWarmupDuration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupScalePolicyAutoScaleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

