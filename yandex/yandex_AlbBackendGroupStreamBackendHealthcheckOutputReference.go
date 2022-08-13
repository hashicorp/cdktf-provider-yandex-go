// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbBackendGroupStreamBackendHealthcheckOutputReference interface {
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
	GrpcHealthcheck() AlbBackendGroupStreamBackendHealthcheckGrpcHealthcheckOutputReference
	GrpcHealthcheckInput() *AlbBackendGroupStreamBackendHealthcheckGrpcHealthcheck
	HealthcheckPort() *float64
	SetHealthcheckPort(val *float64)
	HealthcheckPortInput() *float64
	HealthyThreshold() *float64
	SetHealthyThreshold(val *float64)
	HealthyThresholdInput() *float64
	HttpHealthcheck() AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference
	HttpHealthcheckInput() *AlbBackendGroupStreamBackendHealthcheckHttpHealthcheck
	InternalValue() *AlbBackendGroupStreamBackendHealthcheck
	SetInternalValue(val *AlbBackendGroupStreamBackendHealthcheck)
	Interval() *string
	SetInterval(val *string)
	IntervalInput() *string
	IntervalJitterPercent() *float64
	SetIntervalJitterPercent(val *float64)
	IntervalJitterPercentInput() *float64
	StreamHealthcheck() AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference
	StreamHealthcheckInput() *AlbBackendGroupStreamBackendHealthcheckStreamHealthcheck
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
	PutGrpcHealthcheck(value *AlbBackendGroupStreamBackendHealthcheckGrpcHealthcheck)
	PutHttpHealthcheck(value *AlbBackendGroupStreamBackendHealthcheckHttpHealthcheck)
	PutStreamHealthcheck(value *AlbBackendGroupStreamBackendHealthcheckStreamHealthcheck)
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

// The jsii proxy struct for AlbBackendGroupStreamBackendHealthcheckOutputReference
type jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) GrpcHealthcheck() AlbBackendGroupStreamBackendHealthcheckGrpcHealthcheckOutputReference {
	var returns AlbBackendGroupStreamBackendHealthcheckGrpcHealthcheckOutputReference
	_jsii_.Get(
		j,
		"grpcHealthcheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) GrpcHealthcheckInput() *AlbBackendGroupStreamBackendHealthcheckGrpcHealthcheck {
	var returns *AlbBackendGroupStreamBackendHealthcheckGrpcHealthcheck
	_jsii_.Get(
		j,
		"grpcHealthcheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) HealthcheckPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthcheckPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) HealthcheckPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthcheckPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) HealthyThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthyThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) HealthyThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthyThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) HttpHealthcheck() AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference {
	var returns AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference
	_jsii_.Get(
		j,
		"httpHealthcheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) HttpHealthcheckInput() *AlbBackendGroupStreamBackendHealthcheckHttpHealthcheck {
	var returns *AlbBackendGroupStreamBackendHealthcheckHttpHealthcheck
	_jsii_.Get(
		j,
		"httpHealthcheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) InternalValue() *AlbBackendGroupStreamBackendHealthcheck {
	var returns *AlbBackendGroupStreamBackendHealthcheck
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) Interval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) IntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) IntervalJitterPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalJitterPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) IntervalJitterPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalJitterPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) StreamHealthcheck() AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference {
	var returns AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference
	_jsii_.Get(
		j,
		"streamHealthcheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) StreamHealthcheckInput() *AlbBackendGroupStreamBackendHealthcheckStreamHealthcheck {
	var returns *AlbBackendGroupStreamBackendHealthcheckStreamHealthcheck
	_jsii_.Get(
		j,
		"streamHealthcheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) TimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) UnhealthyThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) UnhealthyThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyThresholdInput",
		&returns,
	)
	return returns
}


func NewAlbBackendGroupStreamBackendHealthcheckOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlbBackendGroupStreamBackendHealthcheckOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbBackendGroupStreamBackendHealthcheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlbBackendGroupStreamBackendHealthcheckOutputReference_Override(a AlbBackendGroupStreamBackendHealthcheckOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbBackendGroupStreamBackendHealthcheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) SetHealthcheckPort(val *float64) {
	_jsii_.Set(
		j,
		"healthcheckPort",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) SetHealthyThreshold(val *float64) {
	_jsii_.Set(
		j,
		"healthyThreshold",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) SetInternalValue(val *AlbBackendGroupStreamBackendHealthcheck) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) SetInterval(val *string) {
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) SetIntervalJitterPercent(val *float64) {
	_jsii_.Set(
		j,
		"intervalJitterPercent",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) SetTimeout(val *string) {
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) SetUnhealthyThreshold(val *float64) {
	_jsii_.Set(
		j,
		"unhealthyThreshold",
		val,
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) PutGrpcHealthcheck(value *AlbBackendGroupStreamBackendHealthcheckGrpcHealthcheck) {
	_jsii_.InvokeVoid(
		a,
		"putGrpcHealthcheck",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) PutHttpHealthcheck(value *AlbBackendGroupStreamBackendHealthcheckHttpHealthcheck) {
	_jsii_.InvokeVoid(
		a,
		"putHttpHealthcheck",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) PutStreamHealthcheck(value *AlbBackendGroupStreamBackendHealthcheckStreamHealthcheck) {
	_jsii_.InvokeVoid(
		a,
		"putStreamHealthcheck",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) ResetGrpcHealthcheck() {
	_jsii_.InvokeVoid(
		a,
		"resetGrpcHealthcheck",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) ResetHealthcheckPort() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthcheckPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) ResetHealthyThreshold() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthyThreshold",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) ResetHttpHealthcheck() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpHealthcheck",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) ResetIntervalJitterPercent() {
	_jsii_.InvokeVoid(
		a,
		"resetIntervalJitterPercent",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) ResetStreamHealthcheck() {
	_jsii_.InvokeVoid(
		a,
		"resetStreamHealthcheck",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) ResetUnhealthyThreshold() {
	_jsii_.InvokeVoid(
		a,
		"resetUnhealthyThreshold",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendHealthcheckOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

