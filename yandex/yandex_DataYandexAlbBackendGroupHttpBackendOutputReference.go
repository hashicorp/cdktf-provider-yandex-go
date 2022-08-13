// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexAlbBackendGroupHttpBackendOutputReference interface {
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
	Healthcheck() DataYandexAlbBackendGroupHttpBackendHealthcheckOutputReference
	HealthcheckInput() *DataYandexAlbBackendGroupHttpBackendHealthcheck
	Http2() interface{}
	SetHttp2(val interface{})
	Http2Input() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LoadBalancingConfig() DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference
	LoadBalancingConfigInput() *DataYandexAlbBackendGroupHttpBackendLoadBalancingConfig
	Name() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	StorageBucket() *string
	SetStorageBucket(val *string)
	StorageBucketInput() *string
	TargetGroupIds() *[]*string
	SetTargetGroupIds(val *[]*string)
	TargetGroupIdsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Tls() DataYandexAlbBackendGroupHttpBackendTlsOutputReference
	TlsInput() *DataYandexAlbBackendGroupHttpBackendTls
	Weight() *float64
	SetWeight(val *float64)
	WeightInput() *float64
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
	PutHealthcheck(value *DataYandexAlbBackendGroupHttpBackendHealthcheck)
	PutLoadBalancingConfig(value *DataYandexAlbBackendGroupHttpBackendLoadBalancingConfig)
	PutTls(value *DataYandexAlbBackendGroupHttpBackendTls)
	ResetHealthcheck()
	ResetHttp2()
	ResetLoadBalancingConfig()
	ResetPort()
	ResetStorageBucket()
	ResetTargetGroupIds()
	ResetTls()
	ResetWeight()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataYandexAlbBackendGroupHttpBackendOutputReference
type jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) Healthcheck() DataYandexAlbBackendGroupHttpBackendHealthcheckOutputReference {
	var returns DataYandexAlbBackendGroupHttpBackendHealthcheckOutputReference
	_jsii_.Get(
		j,
		"healthcheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) HealthcheckInput() *DataYandexAlbBackendGroupHttpBackendHealthcheck {
	var returns *DataYandexAlbBackendGroupHttpBackendHealthcheck
	_jsii_.Get(
		j,
		"healthcheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) Http2() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) Http2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) LoadBalancingConfig() DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference {
	var returns DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference
	_jsii_.Get(
		j,
		"loadBalancingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) LoadBalancingConfigInput() *DataYandexAlbBackendGroupHttpBackendLoadBalancingConfig {
	var returns *DataYandexAlbBackendGroupHttpBackendLoadBalancingConfig
	_jsii_.Get(
		j,
		"loadBalancingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) StorageBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) StorageBucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) TargetGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) TargetGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) Tls() DataYandexAlbBackendGroupHttpBackendTlsOutputReference {
	var returns DataYandexAlbBackendGroupHttpBackendTlsOutputReference
	_jsii_.Get(
		j,
		"tls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) TlsInput() *DataYandexAlbBackendGroupHttpBackendTls {
	var returns *DataYandexAlbBackendGroupHttpBackendTls
	_jsii_.Get(
		j,
		"tlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) Weight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) WeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightInput",
		&returns,
	)
	return returns
}


func NewDataYandexAlbBackendGroupHttpBackendOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataYandexAlbBackendGroupHttpBackendOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexAlbBackendGroupHttpBackendOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataYandexAlbBackendGroupHttpBackendOutputReference_Override(d DataYandexAlbBackendGroupHttpBackendOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexAlbBackendGroupHttpBackendOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) SetHttp2(val interface{}) {
	_jsii_.Set(
		j,
		"http2",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) SetStorageBucket(val *string) {
	_jsii_.Set(
		j,
		"storageBucket",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) SetTargetGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"targetGroupIds",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) SetWeight(val *float64) {
	_jsii_.Set(
		j,
		"weight",
		val,
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) PutHealthcheck(value *DataYandexAlbBackendGroupHttpBackendHealthcheck) {
	_jsii_.InvokeVoid(
		d,
		"putHealthcheck",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) PutLoadBalancingConfig(value *DataYandexAlbBackendGroupHttpBackendLoadBalancingConfig) {
	_jsii_.InvokeVoid(
		d,
		"putLoadBalancingConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) PutTls(value *DataYandexAlbBackendGroupHttpBackendTls) {
	_jsii_.InvokeVoid(
		d,
		"putTls",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) ResetHealthcheck() {
	_jsii_.InvokeVoid(
		d,
		"resetHealthcheck",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) ResetHttp2() {
	_jsii_.InvokeVoid(
		d,
		"resetHttp2",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) ResetLoadBalancingConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetLoadBalancingConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		d,
		"resetPort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) ResetStorageBucket() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageBucket",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) ResetTargetGroupIds() {
	_jsii_.InvokeVoid(
		d,
		"resetTargetGroupIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) ResetTls() {
	_jsii_.InvokeVoid(
		d,
		"resetTls",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) ResetWeight() {
	_jsii_.InvokeVoid(
		d,
		"resetWeight",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

