// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference interface {
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
	GrpcHealthcheck() DataYandexAlbBackendGroupGrpcBackendHealthcheckGrpcHealthcheckOutputReference
	GrpcHealthcheckInput() *DataYandexAlbBackendGroupGrpcBackendHealthcheckGrpcHealthcheck
	HealthcheckPort() *float64
	SetHealthcheckPort(val *float64)
	HealthcheckPortInput() *float64
	HealthyThreshold() *float64
	SetHealthyThreshold(val *float64)
	HealthyThresholdInput() *float64
	HttpHealthcheck() DataYandexAlbBackendGroupGrpcBackendHealthcheckHttpHealthcheckOutputReference
	HttpHealthcheckInput() *DataYandexAlbBackendGroupGrpcBackendHealthcheckHttpHealthcheck
	InternalValue() *DataYandexAlbBackendGroupGrpcBackendHealthcheck
	SetInternalValue(val *DataYandexAlbBackendGroupGrpcBackendHealthcheck)
	Interval() *string
	IntervalJitterPercent() *float64
	SetIntervalJitterPercent(val *float64)
	IntervalJitterPercentInput() *float64
	StreamHealthcheck() DataYandexAlbBackendGroupGrpcBackendHealthcheckStreamHealthcheckOutputReference
	StreamHealthcheckInput() *DataYandexAlbBackendGroupGrpcBackendHealthcheckStreamHealthcheck
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
	PutGrpcHealthcheck(value *DataYandexAlbBackendGroupGrpcBackendHealthcheckGrpcHealthcheck)
	PutHttpHealthcheck(value *DataYandexAlbBackendGroupGrpcBackendHealthcheckHttpHealthcheck)
	PutStreamHealthcheck(value *DataYandexAlbBackendGroupGrpcBackendHealthcheckStreamHealthcheck)
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

// The jsii proxy struct for DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference
type jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) GrpcHealthcheck() DataYandexAlbBackendGroupGrpcBackendHealthcheckGrpcHealthcheckOutputReference {
	var returns DataYandexAlbBackendGroupGrpcBackendHealthcheckGrpcHealthcheckOutputReference
	_jsii_.Get(
		j,
		"grpcHealthcheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) GrpcHealthcheckInput() *DataYandexAlbBackendGroupGrpcBackendHealthcheckGrpcHealthcheck {
	var returns *DataYandexAlbBackendGroupGrpcBackendHealthcheckGrpcHealthcheck
	_jsii_.Get(
		j,
		"grpcHealthcheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) HealthcheckPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthcheckPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) HealthcheckPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthcheckPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) HealthyThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthyThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) HealthyThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthyThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) HttpHealthcheck() DataYandexAlbBackendGroupGrpcBackendHealthcheckHttpHealthcheckOutputReference {
	var returns DataYandexAlbBackendGroupGrpcBackendHealthcheckHttpHealthcheckOutputReference
	_jsii_.Get(
		j,
		"httpHealthcheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) HttpHealthcheckInput() *DataYandexAlbBackendGroupGrpcBackendHealthcheckHttpHealthcheck {
	var returns *DataYandexAlbBackendGroupGrpcBackendHealthcheckHttpHealthcheck
	_jsii_.Get(
		j,
		"httpHealthcheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) InternalValue() *DataYandexAlbBackendGroupGrpcBackendHealthcheck {
	var returns *DataYandexAlbBackendGroupGrpcBackendHealthcheck
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) Interval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) IntervalJitterPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalJitterPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) IntervalJitterPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalJitterPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) StreamHealthcheck() DataYandexAlbBackendGroupGrpcBackendHealthcheckStreamHealthcheckOutputReference {
	var returns DataYandexAlbBackendGroupGrpcBackendHealthcheckStreamHealthcheckOutputReference
	_jsii_.Get(
		j,
		"streamHealthcheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) StreamHealthcheckInput() *DataYandexAlbBackendGroupGrpcBackendHealthcheckStreamHealthcheck {
	var returns *DataYandexAlbBackendGroupGrpcBackendHealthcheckStreamHealthcheck
	_jsii_.Get(
		j,
		"streamHealthcheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) UnhealthyThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) UnhealthyThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyThresholdInput",
		&returns,
	)
	return returns
}


func NewDataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference_Override(d DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) SetHealthcheckPort(val *float64) {
	_jsii_.Set(
		j,
		"healthcheckPort",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) SetHealthyThreshold(val *float64) {
	_jsii_.Set(
		j,
		"healthyThreshold",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) SetInternalValue(val *DataYandexAlbBackendGroupGrpcBackendHealthcheck) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) SetIntervalJitterPercent(val *float64) {
	_jsii_.Set(
		j,
		"intervalJitterPercent",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) SetUnhealthyThreshold(val *float64) {
	_jsii_.Set(
		j,
		"unhealthyThreshold",
		val,
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) PutGrpcHealthcheck(value *DataYandexAlbBackendGroupGrpcBackendHealthcheckGrpcHealthcheck) {
	_jsii_.InvokeVoid(
		d,
		"putGrpcHealthcheck",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) PutHttpHealthcheck(value *DataYandexAlbBackendGroupGrpcBackendHealthcheckHttpHealthcheck) {
	_jsii_.InvokeVoid(
		d,
		"putHttpHealthcheck",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) PutStreamHealthcheck(value *DataYandexAlbBackendGroupGrpcBackendHealthcheckStreamHealthcheck) {
	_jsii_.InvokeVoid(
		d,
		"putStreamHealthcheck",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) ResetGrpcHealthcheck() {
	_jsii_.InvokeVoid(
		d,
		"resetGrpcHealthcheck",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) ResetHealthcheckPort() {
	_jsii_.InvokeVoid(
		d,
		"resetHealthcheckPort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) ResetHealthyThreshold() {
	_jsii_.InvokeVoid(
		d,
		"resetHealthyThreshold",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) ResetHttpHealthcheck() {
	_jsii_.InvokeVoid(
		d,
		"resetHttpHealthcheck",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) ResetIntervalJitterPercent() {
	_jsii_.InvokeVoid(
		d,
		"resetIntervalJitterPercent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) ResetStreamHealthcheck() {
	_jsii_.InvokeVoid(
		d,
		"resetStreamHealthcheck",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) ResetUnhealthyThreshold() {
	_jsii_.InvokeVoid(
		d,
		"resetUnhealthyThreshold",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

