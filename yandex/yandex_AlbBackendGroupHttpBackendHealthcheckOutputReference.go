// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbBackendGroupHttpBackendHealthcheckOutputReference interface {
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
	GrpcHealthcheck() AlbBackendGroupHttpBackendHealthcheckGrpcHealthcheckOutputReference
	GrpcHealthcheckInput() *AlbBackendGroupHttpBackendHealthcheckGrpcHealthcheck
	HealthcheckPort() *float64
	SetHealthcheckPort(val *float64)
	HealthcheckPortInput() *float64
	HealthyThreshold() *float64
	SetHealthyThreshold(val *float64)
	HealthyThresholdInput() *float64
	HttpHealthcheck() AlbBackendGroupHttpBackendHealthcheckHttpHealthcheckOutputReference
	HttpHealthcheckInput() *AlbBackendGroupHttpBackendHealthcheckHttpHealthcheck
	InternalValue() *AlbBackendGroupHttpBackendHealthcheck
	SetInternalValue(val *AlbBackendGroupHttpBackendHealthcheck)
	Interval() *string
	SetInterval(val *string)
	IntervalInput() *string
	IntervalJitterPercent() *float64
	SetIntervalJitterPercent(val *float64)
	IntervalJitterPercentInput() *float64
	StreamHealthcheck() AlbBackendGroupHttpBackendHealthcheckStreamHealthcheckOutputReference
	StreamHealthcheckInput() *AlbBackendGroupHttpBackendHealthcheckStreamHealthcheck
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() *string
	SetTimeout(val *string)
	TimeoutInput() *string
	UnhealthyThreshold() *float64
	SetUnhealthyThreshold(val *float64)
	UnhealthyThresholdInput() *float64
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
	PutGrpcHealthcheck(value *AlbBackendGroupHttpBackendHealthcheckGrpcHealthcheck)
	PutHttpHealthcheck(value *AlbBackendGroupHttpBackendHealthcheckHttpHealthcheck)
	PutStreamHealthcheck(value *AlbBackendGroupHttpBackendHealthcheckStreamHealthcheck)
	ResetGrpcHealthcheck()
	ResetHealthcheckPort()
	ResetHealthyThreshold()
	ResetHttpHealthcheck()
	ResetIntervalJitterPercent()
	ResetStreamHealthcheck()
	ResetUnhealthyThreshold()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlbBackendGroupHttpBackendHealthcheckOutputReference
type jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) GrpcHealthcheck() AlbBackendGroupHttpBackendHealthcheckGrpcHealthcheckOutputReference {
	var returns AlbBackendGroupHttpBackendHealthcheckGrpcHealthcheckOutputReference
	_jsii_.Get(
		j,
		"grpcHealthcheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) GrpcHealthcheckInput() *AlbBackendGroupHttpBackendHealthcheckGrpcHealthcheck {
	var returns *AlbBackendGroupHttpBackendHealthcheckGrpcHealthcheck
	_jsii_.Get(
		j,
		"grpcHealthcheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) HealthcheckPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthcheckPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) HealthcheckPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthcheckPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) HealthyThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthyThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) HealthyThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthyThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) HttpHealthcheck() AlbBackendGroupHttpBackendHealthcheckHttpHealthcheckOutputReference {
	var returns AlbBackendGroupHttpBackendHealthcheckHttpHealthcheckOutputReference
	_jsii_.Get(
		j,
		"httpHealthcheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) HttpHealthcheckInput() *AlbBackendGroupHttpBackendHealthcheckHttpHealthcheck {
	var returns *AlbBackendGroupHttpBackendHealthcheckHttpHealthcheck
	_jsii_.Get(
		j,
		"httpHealthcheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) InternalValue() *AlbBackendGroupHttpBackendHealthcheck {
	var returns *AlbBackendGroupHttpBackendHealthcheck
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) Interval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) IntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) IntervalJitterPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalJitterPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) IntervalJitterPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalJitterPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) StreamHealthcheck() AlbBackendGroupHttpBackendHealthcheckStreamHealthcheckOutputReference {
	var returns AlbBackendGroupHttpBackendHealthcheckStreamHealthcheckOutputReference
	_jsii_.Get(
		j,
		"streamHealthcheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) StreamHealthcheckInput() *AlbBackendGroupHttpBackendHealthcheckStreamHealthcheck {
	var returns *AlbBackendGroupHttpBackendHealthcheckStreamHealthcheck
	_jsii_.Get(
		j,
		"streamHealthcheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) TimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) UnhealthyThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) UnhealthyThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyThresholdInput",
		&returns,
	)
	return returns
}


func NewAlbBackendGroupHttpBackendHealthcheckOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlbBackendGroupHttpBackendHealthcheckOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbBackendGroupHttpBackendHealthcheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlbBackendGroupHttpBackendHealthcheckOutputReference_Override(a AlbBackendGroupHttpBackendHealthcheckOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbBackendGroupHttpBackendHealthcheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) SetHealthcheckPort(val *float64) {
	_jsii_.Set(
		j,
		"healthcheckPort",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) SetHealthyThreshold(val *float64) {
	_jsii_.Set(
		j,
		"healthyThreshold",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) SetInternalValue(val *AlbBackendGroupHttpBackendHealthcheck) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) SetInterval(val *string) {
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) SetIntervalJitterPercent(val *float64) {
	_jsii_.Set(
		j,
		"intervalJitterPercent",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) SetTimeout(val *string) {
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) SetUnhealthyThreshold(val *float64) {
	_jsii_.Set(
		j,
		"unhealthyThreshold",
		val,
	)
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) PutGrpcHealthcheck(value *AlbBackendGroupHttpBackendHealthcheckGrpcHealthcheck) {
	_jsii_.InvokeVoid(
		a,
		"putGrpcHealthcheck",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) PutHttpHealthcheck(value *AlbBackendGroupHttpBackendHealthcheckHttpHealthcheck) {
	_jsii_.InvokeVoid(
		a,
		"putHttpHealthcheck",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) PutStreamHealthcheck(value *AlbBackendGroupHttpBackendHealthcheckStreamHealthcheck) {
	_jsii_.InvokeVoid(
		a,
		"putStreamHealthcheck",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) ResetGrpcHealthcheck() {
	_jsii_.InvokeVoid(
		a,
		"resetGrpcHealthcheck",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) ResetHealthcheckPort() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthcheckPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) ResetHealthyThreshold() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthyThreshold",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) ResetHttpHealthcheck() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpHealthcheck",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) ResetIntervalJitterPercent() {
	_jsii_.InvokeVoid(
		a,
		"resetIntervalJitterPercent",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) ResetStreamHealthcheck() {
	_jsii_.InvokeVoid(
		a,
		"resetStreamHealthcheck",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) ResetUnhealthyThreshold() {
	_jsii_.InvokeVoid(
		a,
		"resetUnhealthyThreshold",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendHealthcheckOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

