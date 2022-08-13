// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexAlbBackendGroupGrpcBackendOutputReference interface {
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
	Healthcheck() DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference
	HealthcheckInput() *DataYandexAlbBackendGroupGrpcBackendHealthcheck
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LoadBalancingConfig() DataYandexAlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference
	LoadBalancingConfigInput() *DataYandexAlbBackendGroupGrpcBackendLoadBalancingConfig
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
	Tls() DataYandexAlbBackendGroupGrpcBackendTlsOutputReference
	TlsInput() *DataYandexAlbBackendGroupGrpcBackendTls
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
	PutHealthcheck(value *DataYandexAlbBackendGroupGrpcBackendHealthcheck)
	PutLoadBalancingConfig(value *DataYandexAlbBackendGroupGrpcBackendLoadBalancingConfig)
	PutTls(value *DataYandexAlbBackendGroupGrpcBackendTls)
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

// The jsii proxy struct for DataYandexAlbBackendGroupGrpcBackendOutputReference
type jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) Healthcheck() DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference {
	var returns DataYandexAlbBackendGroupGrpcBackendHealthcheckOutputReference
	_jsii_.Get(
		j,
		"healthcheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) HealthcheckInput() *DataYandexAlbBackendGroupGrpcBackendHealthcheck {
	var returns *DataYandexAlbBackendGroupGrpcBackendHealthcheck
	_jsii_.Get(
		j,
		"healthcheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) LoadBalancingConfig() DataYandexAlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference {
	var returns DataYandexAlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference
	_jsii_.Get(
		j,
		"loadBalancingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) LoadBalancingConfigInput() *DataYandexAlbBackendGroupGrpcBackendLoadBalancingConfig {
	var returns *DataYandexAlbBackendGroupGrpcBackendLoadBalancingConfig
	_jsii_.Get(
		j,
		"loadBalancingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) TargetGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) Tls() DataYandexAlbBackendGroupGrpcBackendTlsOutputReference {
	var returns DataYandexAlbBackendGroupGrpcBackendTlsOutputReference
	_jsii_.Get(
		j,
		"tls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) TlsInput() *DataYandexAlbBackendGroupGrpcBackendTls {
	var returns *DataYandexAlbBackendGroupGrpcBackendTls
	_jsii_.Get(
		j,
		"tlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) Weight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) WeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightInput",
		&returns,
	)
	return returns
}


func NewDataYandexAlbBackendGroupGrpcBackendOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataYandexAlbBackendGroupGrpcBackendOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexAlbBackendGroupGrpcBackendOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataYandexAlbBackendGroupGrpcBackendOutputReference_Override(d DataYandexAlbBackendGroupGrpcBackendOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexAlbBackendGroupGrpcBackendOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) SetWeight(val *float64) {
	_jsii_.Set(
		j,
		"weight",
		val,
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) PutHealthcheck(value *DataYandexAlbBackendGroupGrpcBackendHealthcheck) {
	_jsii_.InvokeVoid(
		d,
		"putHealthcheck",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) PutLoadBalancingConfig(value *DataYandexAlbBackendGroupGrpcBackendLoadBalancingConfig) {
	_jsii_.InvokeVoid(
		d,
		"putLoadBalancingConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) PutTls(value *DataYandexAlbBackendGroupGrpcBackendTls) {
	_jsii_.InvokeVoid(
		d,
		"putTls",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) ResetHealthcheck() {
	_jsii_.InvokeVoid(
		d,
		"resetHealthcheck",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) ResetLoadBalancingConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetLoadBalancingConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		d,
		"resetPort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) ResetTls() {
	_jsii_.InvokeVoid(
		d,
		"resetTls",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) ResetWeight() {
	_jsii_.InvokeVoid(
		d,
		"resetWeight",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupGrpcBackendOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

