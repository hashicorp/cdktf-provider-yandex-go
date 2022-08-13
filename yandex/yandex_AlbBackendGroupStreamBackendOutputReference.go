// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbBackendGroupStreamBackendOutputReference interface {
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
	Healthcheck() AlbBackendGroupStreamBackendHealthcheckOutputReference
	HealthcheckInput() *AlbBackendGroupStreamBackendHealthcheck
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LoadBalancingConfig() AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference
	LoadBalancingConfigInput() *AlbBackendGroupStreamBackendLoadBalancingConfig
	Name() *string
	SetName(val *string)
	NameInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
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
	Tls() AlbBackendGroupStreamBackendTlsOutputReference
	TlsInput() *AlbBackendGroupStreamBackendTls
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
	PutHealthcheck(value *AlbBackendGroupStreamBackendHealthcheck)
	PutLoadBalancingConfig(value *AlbBackendGroupStreamBackendLoadBalancingConfig)
	PutTls(value *AlbBackendGroupStreamBackendTls)
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

// The jsii proxy struct for AlbBackendGroupStreamBackendOutputReference
type jsiiProxy_AlbBackendGroupStreamBackendOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) EnableProxyProtocol() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableProxyProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) EnableProxyProtocolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableProxyProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) Healthcheck() AlbBackendGroupStreamBackendHealthcheckOutputReference {
	var returns AlbBackendGroupStreamBackendHealthcheckOutputReference
	_jsii_.Get(
		j,
		"healthcheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) HealthcheckInput() *AlbBackendGroupStreamBackendHealthcheck {
	var returns *AlbBackendGroupStreamBackendHealthcheck
	_jsii_.Get(
		j,
		"healthcheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) LoadBalancingConfig() AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference {
	var returns AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference
	_jsii_.Get(
		j,
		"loadBalancingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) LoadBalancingConfigInput() *AlbBackendGroupStreamBackendLoadBalancingConfig {
	var returns *AlbBackendGroupStreamBackendLoadBalancingConfig
	_jsii_.Get(
		j,
		"loadBalancingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) TargetGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) TargetGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) Tls() AlbBackendGroupStreamBackendTlsOutputReference {
	var returns AlbBackendGroupStreamBackendTlsOutputReference
	_jsii_.Get(
		j,
		"tls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) TlsInput() *AlbBackendGroupStreamBackendTls {
	var returns *AlbBackendGroupStreamBackendTls
	_jsii_.Get(
		j,
		"tlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) Weight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) WeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightInput",
		&returns,
	)
	return returns
}


func NewAlbBackendGroupStreamBackendOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AlbBackendGroupStreamBackendOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AlbBackendGroupStreamBackendOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbBackendGroupStreamBackendOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAlbBackendGroupStreamBackendOutputReference_Override(a AlbBackendGroupStreamBackendOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbBackendGroupStreamBackendOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) SetEnableProxyProtocol(val interface{}) {
	_jsii_.Set(
		j,
		"enableProxyProtocol",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) SetTargetGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"targetGroupIds",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) SetWeight(val *float64) {
	_jsii_.Set(
		j,
		"weight",
		val,
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) PutHealthcheck(value *AlbBackendGroupStreamBackendHealthcheck) {
	_jsii_.InvokeVoid(
		a,
		"putHealthcheck",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) PutLoadBalancingConfig(value *AlbBackendGroupStreamBackendLoadBalancingConfig) {
	_jsii_.InvokeVoid(
		a,
		"putLoadBalancingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) PutTls(value *AlbBackendGroupStreamBackendTls) {
	_jsii_.InvokeVoid(
		a,
		"putTls",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) ResetEnableProxyProtocol() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableProxyProtocol",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) ResetHealthcheck() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthcheck",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) ResetLoadBalancingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLoadBalancingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		a,
		"resetPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) ResetTls() {
	_jsii_.InvokeVoid(
		a,
		"resetTls",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) ResetWeight() {
	_jsii_.InvokeVoid(
		a,
		"resetWeight",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

