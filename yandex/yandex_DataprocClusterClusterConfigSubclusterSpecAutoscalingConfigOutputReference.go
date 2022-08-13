// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference interface {
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
	DecommissionTimeout() *float64
	SetDecommissionTimeout(val *float64)
	DecommissionTimeoutInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *DataprocClusterClusterConfigSubclusterSpecAutoscalingConfig
	SetInternalValue(val *DataprocClusterClusterConfigSubclusterSpecAutoscalingConfig)
	MaxHostsCount() *float64
	SetMaxHostsCount(val *float64)
	MaxHostsCountInput() *float64
	MeasurementDuration() *float64
	SetMeasurementDuration(val *float64)
	MeasurementDurationInput() *float64
	Preemptible() interface{}
	SetPreemptible(val interface{})
	PreemptibleInput() interface{}
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
	ResetCpuUtilizationTarget()
	ResetDecommissionTimeout()
	ResetMeasurementDuration()
	ResetPreemptible()
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

// The jsii proxy struct for DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference
type jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) CpuUtilizationTarget() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuUtilizationTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) CpuUtilizationTargetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuUtilizationTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) DecommissionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"decommissionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) DecommissionTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"decommissionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) InternalValue() *DataprocClusterClusterConfigSubclusterSpecAutoscalingConfig {
	var returns *DataprocClusterClusterConfigSubclusterSpecAutoscalingConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) MaxHostsCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxHostsCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) MaxHostsCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxHostsCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) MeasurementDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"measurementDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) MeasurementDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"measurementDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) Preemptible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preemptible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) PreemptibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preemptibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) StabilizationDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stabilizationDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) StabilizationDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stabilizationDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) WarmupDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"warmupDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) WarmupDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"warmupDurationInput",
		&returns,
	)
	return returns
}


func NewDataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference_Override(d DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) SetCpuUtilizationTarget(val *float64) {
	_jsii_.Set(
		j,
		"cpuUtilizationTarget",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) SetDecommissionTimeout(val *float64) {
	_jsii_.Set(
		j,
		"decommissionTimeout",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) SetInternalValue(val *DataprocClusterClusterConfigSubclusterSpecAutoscalingConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) SetMaxHostsCount(val *float64) {
	_jsii_.Set(
		j,
		"maxHostsCount",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) SetMeasurementDuration(val *float64) {
	_jsii_.Set(
		j,
		"measurementDuration",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) SetPreemptible(val interface{}) {
	_jsii_.Set(
		j,
		"preemptible",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) SetStabilizationDuration(val *float64) {
	_jsii_.Set(
		j,
		"stabilizationDuration",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) SetWarmupDuration(val *float64) {
	_jsii_.Set(
		j,
		"warmupDuration",
		val,
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) ResetCpuUtilizationTarget() {
	_jsii_.InvokeVoid(
		d,
		"resetCpuUtilizationTarget",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) ResetDecommissionTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetDecommissionTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) ResetMeasurementDuration() {
	_jsii_.InvokeVoid(
		d,
		"resetMeasurementDuration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) ResetPreemptible() {
	_jsii_.InvokeVoid(
		d,
		"resetPreemptible",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) ResetStabilizationDuration() {
	_jsii_.InvokeVoid(
		d,
		"resetStabilizationDuration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) ResetWarmupDuration() {
	_jsii_.InvokeVoid(
		d,
		"resetWarmupDuration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

