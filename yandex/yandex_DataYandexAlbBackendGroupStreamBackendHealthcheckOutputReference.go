// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference interface {
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
	GrpcHealthcheck() DataYandexAlbBackendGroupStreamBackendHealthcheckGrpcHealthcheckOutputReference
	GrpcHealthcheckInput() *DataYandexAlbBackendGroupStreamBackendHealthcheckGrpcHealthcheck
	HealthcheckPort() *float64
	SetHealthcheckPort(val *float64)
	HealthcheckPortInput() *float64
	HealthyThreshold() *float64
	SetHealthyThreshold(val *float64)
	HealthyThresholdInput() *float64
	HttpHealthcheck() DataYandexAlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference
	HttpHealthcheckInput() *DataYandexAlbBackendGroupStreamBackendHealthcheckHttpHealthcheck
	InternalValue() *DataYandexAlbBackendGroupStreamBackendHealthcheck
	SetInternalValue(val *DataYandexAlbBackendGroupStreamBackendHealthcheck)
	Interval() *string
	IntervalJitterPercent() *float64
	SetIntervalJitterPercent(val *float64)
	IntervalJitterPercentInput() *float64
	StreamHealthcheck() DataYandexAlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference
	StreamHealthcheckInput() *DataYandexAlbBackendGroupStreamBackendHealthcheckStreamHealthcheck
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() *string
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
	PutGrpcHealthcheck(value *DataYandexAlbBackendGroupStreamBackendHealthcheckGrpcHealthcheck)
	PutHttpHealthcheck(value *DataYandexAlbBackendGroupStreamBackendHealthcheckHttpHealthcheck)
	PutStreamHealthcheck(value *DataYandexAlbBackendGroupStreamBackendHealthcheckStreamHealthcheck)
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

// The jsii proxy struct for DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference
type jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) GrpcHealthcheck() DataYandexAlbBackendGroupStreamBackendHealthcheckGrpcHealthcheckOutputReference {
	var returns DataYandexAlbBackendGroupStreamBackendHealthcheckGrpcHealthcheckOutputReference
	_jsii_.Get(
		j,
		"grpcHealthcheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) GrpcHealthcheckInput() *DataYandexAlbBackendGroupStreamBackendHealthcheckGrpcHealthcheck {
	var returns *DataYandexAlbBackendGroupStreamBackendHealthcheckGrpcHealthcheck
	_jsii_.Get(
		j,
		"grpcHealthcheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) HealthcheckPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthcheckPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) HealthcheckPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthcheckPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) HealthyThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthyThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) HealthyThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthyThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) HttpHealthcheck() DataYandexAlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference {
	var returns DataYandexAlbBackendGroupStreamBackendHealthcheckHttpHealthcheckOutputReference
	_jsii_.Get(
		j,
		"httpHealthcheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) HttpHealthcheckInput() *DataYandexAlbBackendGroupStreamBackendHealthcheckHttpHealthcheck {
	var returns *DataYandexAlbBackendGroupStreamBackendHealthcheckHttpHealthcheck
	_jsii_.Get(
		j,
		"httpHealthcheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) InternalValue() *DataYandexAlbBackendGroupStreamBackendHealthcheck {
	var returns *DataYandexAlbBackendGroupStreamBackendHealthcheck
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) Interval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) IntervalJitterPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalJitterPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) IntervalJitterPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalJitterPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) StreamHealthcheck() DataYandexAlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference {
	var returns DataYandexAlbBackendGroupStreamBackendHealthcheckStreamHealthcheckOutputReference
	_jsii_.Get(
		j,
		"streamHealthcheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) StreamHealthcheckInput() *DataYandexAlbBackendGroupStreamBackendHealthcheckStreamHealthcheck {
	var returns *DataYandexAlbBackendGroupStreamBackendHealthcheckStreamHealthcheck
	_jsii_.Get(
		j,
		"streamHealthcheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) UnhealthyThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) UnhealthyThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyThresholdInput",
		&returns,
	)
	return returns
}


func NewDataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference_Override(d DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) SetHealthcheckPort(val *float64) {
	_jsii_.Set(
		j,
		"healthcheckPort",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) SetHealthyThreshold(val *float64) {
	_jsii_.Set(
		j,
		"healthyThreshold",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) SetInternalValue(val *DataYandexAlbBackendGroupStreamBackendHealthcheck) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) SetIntervalJitterPercent(val *float64) {
	_jsii_.Set(
		j,
		"intervalJitterPercent",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) SetUnhealthyThreshold(val *float64) {
	_jsii_.Set(
		j,
		"unhealthyThreshold",
		val,
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) PutGrpcHealthcheck(value *DataYandexAlbBackendGroupStreamBackendHealthcheckGrpcHealthcheck) {
	_jsii_.InvokeVoid(
		d,
		"putGrpcHealthcheck",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) PutHttpHealthcheck(value *DataYandexAlbBackendGroupStreamBackendHealthcheckHttpHealthcheck) {
	_jsii_.InvokeVoid(
		d,
		"putHttpHealthcheck",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) PutStreamHealthcheck(value *DataYandexAlbBackendGroupStreamBackendHealthcheckStreamHealthcheck) {
	_jsii_.InvokeVoid(
		d,
		"putStreamHealthcheck",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) ResetGrpcHealthcheck() {
	_jsii_.InvokeVoid(
		d,
		"resetGrpcHealthcheck",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) ResetHealthcheckPort() {
	_jsii_.InvokeVoid(
		d,
		"resetHealthcheckPort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) ResetHealthyThreshold() {
	_jsii_.InvokeVoid(
		d,
		"resetHealthyThreshold",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) ResetHttpHealthcheck() {
	_jsii_.InvokeVoid(
		d,
		"resetHttpHealthcheck",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) ResetIntervalJitterPercent() {
	_jsii_.InvokeVoid(
		d,
		"resetIntervalJitterPercent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) ResetStreamHealthcheck() {
	_jsii_.InvokeVoid(
		d,
		"resetStreamHealthcheck",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) ResetUnhealthyThreshold() {
	_jsii_.InvokeVoid(
		d,
		"resetUnhealthyThreshold",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

