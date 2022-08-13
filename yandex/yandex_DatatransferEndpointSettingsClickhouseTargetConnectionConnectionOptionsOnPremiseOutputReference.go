// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference interface {
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
	HttpPort() *float64
	SetHttpPort(val *float64)
	HttpPortInput() *float64
	InternalValue() *DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremise
	SetInternalValue(val *DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremise)
	NativePort() *float64
	SetNativePort(val *float64)
	NativePortInput() *float64
	Shards() DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseShardsList
	ShardsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TlsMode() DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseTlsModeOutputReference
	TlsModeInput() *DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseTlsMode
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
	PutShards(value interface{})
	PutTlsMode(value *DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseTlsMode)
	ResetHttpPort()
	ResetNativePort()
	ResetShards()
	ResetTlsMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference
type jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) HttpPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) HttpPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) InternalValue() *DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremise {
	var returns *DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremise
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) NativePort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nativePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) NativePortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nativePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) Shards() DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseShardsList {
	var returns DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseShardsList
	_jsii_.Get(
		j,
		"shards",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) ShardsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shardsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) TlsMode() DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseTlsModeOutputReference {
	var returns DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseTlsModeOutputReference
	_jsii_.Get(
		j,
		"tlsMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) TlsModeInput() *DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseTlsMode {
	var returns *DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseTlsMode
	_jsii_.Get(
		j,
		"tlsModeInput",
		&returns,
	)
	return returns
}


func NewDatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference_Override(d DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) SetHttpPort(val *float64) {
	_jsii_.Set(
		j,
		"httpPort",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) SetInternalValue(val *DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremise) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) SetNativePort(val *float64) {
	_jsii_.Set(
		j,
		"nativePort",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) PutShards(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putShards",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) PutTlsMode(value *DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseTlsMode) {
	_jsii_.InvokeVoid(
		d,
		"putTlsMode",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) ResetHttpPort() {
	_jsii_.InvokeVoid(
		d,
		"resetHttpPort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) ResetNativePort() {
	_jsii_.InvokeVoid(
		d,
		"resetNativePort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) ResetShards() {
	_jsii_.InvokeVoid(
		d,
		"resetShards",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) ResetTlsMode() {
	_jsii_.InvokeVoid(
		d,
		"resetTlsMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

