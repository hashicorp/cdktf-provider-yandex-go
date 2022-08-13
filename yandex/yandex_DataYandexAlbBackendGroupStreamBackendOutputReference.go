// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexAlbBackendGroupStreamBackendOutputReference interface {
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
	EnableProxyProtocol() interface{}
	SetEnableProxyProtocol(val interface{})
	EnableProxyProtocolInput() interface{}
	// Experimental.
	Fqn() *string
	Healthcheck() DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference
	HealthcheckInput() *DataYandexAlbBackendGroupStreamBackendHealthcheck
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LoadBalancingConfig() DataYandexAlbBackendGroupStreamBackendLoadBalancingConfigOutputReference
	LoadBalancingConfigInput() *DataYandexAlbBackendGroupStreamBackendLoadBalancingConfig
	Name() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	TargetGroupIds() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Tls() DataYandexAlbBackendGroupStreamBackendTlsOutputReference
	TlsInput() *DataYandexAlbBackendGroupStreamBackendTls
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
	PutHealthcheck(value *DataYandexAlbBackendGroupStreamBackendHealthcheck)
	PutLoadBalancingConfig(value *DataYandexAlbBackendGroupStreamBackendLoadBalancingConfig)
	PutTls(value *DataYandexAlbBackendGroupStreamBackendTls)
	ResetEnableProxyProtocol()
	ResetHealthcheck()
	ResetLoadBalancingConfig()
	ResetPort()
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

// The jsii proxy struct for DataYandexAlbBackendGroupStreamBackendOutputReference
type jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) EnableProxyProtocol() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableProxyProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) EnableProxyProtocolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableProxyProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) Healthcheck() DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference {
	var returns DataYandexAlbBackendGroupStreamBackendHealthcheckOutputReference
	_jsii_.Get(
		j,
		"healthcheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) HealthcheckInput() *DataYandexAlbBackendGroupStreamBackendHealthcheck {
	var returns *DataYandexAlbBackendGroupStreamBackendHealthcheck
	_jsii_.Get(
		j,
		"healthcheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) LoadBalancingConfig() DataYandexAlbBackendGroupStreamBackendLoadBalancingConfigOutputReference {
	var returns DataYandexAlbBackendGroupStreamBackendLoadBalancingConfigOutputReference
	_jsii_.Get(
		j,
		"loadBalancingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) LoadBalancingConfigInput() *DataYandexAlbBackendGroupStreamBackendLoadBalancingConfig {
	var returns *DataYandexAlbBackendGroupStreamBackendLoadBalancingConfig
	_jsii_.Get(
		j,
		"loadBalancingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) TargetGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) Tls() DataYandexAlbBackendGroupStreamBackendTlsOutputReference {
	var returns DataYandexAlbBackendGroupStreamBackendTlsOutputReference
	_jsii_.Get(
		j,
		"tls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) TlsInput() *DataYandexAlbBackendGroupStreamBackendTls {
	var returns *DataYandexAlbBackendGroupStreamBackendTls
	_jsii_.Get(
		j,
		"tlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) Weight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) WeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightInput",
		&returns,
	)
	return returns
}


func NewDataYandexAlbBackendGroupStreamBackendOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataYandexAlbBackendGroupStreamBackendOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexAlbBackendGroupStreamBackendOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataYandexAlbBackendGroupStreamBackendOutputReference_Override(d DataYandexAlbBackendGroupStreamBackendOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexAlbBackendGroupStreamBackendOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) SetEnableProxyProtocol(val interface{}) {
	_jsii_.Set(
		j,
		"enableProxyProtocol",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) SetWeight(val *float64) {
	_jsii_.Set(
		j,
		"weight",
		val,
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) PutHealthcheck(value *DataYandexAlbBackendGroupStreamBackendHealthcheck) {
	_jsii_.InvokeVoid(
		d,
		"putHealthcheck",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) PutLoadBalancingConfig(value *DataYandexAlbBackendGroupStreamBackendLoadBalancingConfig) {
	_jsii_.InvokeVoid(
		d,
		"putLoadBalancingConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) PutTls(value *DataYandexAlbBackendGroupStreamBackendTls) {
	_jsii_.InvokeVoid(
		d,
		"putTls",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) ResetEnableProxyProtocol() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableProxyProtocol",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) ResetHealthcheck() {
	_jsii_.InvokeVoid(
		d,
		"resetHealthcheck",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) ResetLoadBalancingConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetLoadBalancingConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		d,
		"resetPort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) ResetTls() {
	_jsii_.InvokeVoid(
		d,
		"resetTls",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) ResetWeight() {
	_jsii_.InvokeVoid(
		d,
		"resetWeight",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

