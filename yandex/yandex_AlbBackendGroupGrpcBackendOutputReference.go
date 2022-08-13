// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbBackendGroupGrpcBackendOutputReference interface {
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
	Healthcheck() AlbBackendGroupGrpcBackendHealthcheckOutputReference
	HealthcheckInput() *AlbBackendGroupGrpcBackendHealthcheck
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LoadBalancingConfig() AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference
	LoadBalancingConfigInput() *AlbBackendGroupGrpcBackendLoadBalancingConfig
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
	Tls() AlbBackendGroupGrpcBackendTlsOutputReference
	TlsInput() *AlbBackendGroupGrpcBackendTls
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
	PutHealthcheck(value *AlbBackendGroupGrpcBackendHealthcheck)
	PutLoadBalancingConfig(value *AlbBackendGroupGrpcBackendLoadBalancingConfig)
	PutTls(value *AlbBackendGroupGrpcBackendTls)
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

// The jsii proxy struct for AlbBackendGroupGrpcBackendOutputReference
type jsiiProxy_AlbBackendGroupGrpcBackendOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) Healthcheck() AlbBackendGroupGrpcBackendHealthcheckOutputReference {
	var returns AlbBackendGroupGrpcBackendHealthcheckOutputReference
	_jsii_.Get(
		j,
		"healthcheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) HealthcheckInput() *AlbBackendGroupGrpcBackendHealthcheck {
	var returns *AlbBackendGroupGrpcBackendHealthcheck
	_jsii_.Get(
		j,
		"healthcheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) LoadBalancingConfig() AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference {
	var returns AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference
	_jsii_.Get(
		j,
		"loadBalancingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) LoadBalancingConfigInput() *AlbBackendGroupGrpcBackendLoadBalancingConfig {
	var returns *AlbBackendGroupGrpcBackendLoadBalancingConfig
	_jsii_.Get(
		j,
		"loadBalancingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) TargetGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) TargetGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) Tls() AlbBackendGroupGrpcBackendTlsOutputReference {
	var returns AlbBackendGroupGrpcBackendTlsOutputReference
	_jsii_.Get(
		j,
		"tls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) TlsInput() *AlbBackendGroupGrpcBackendTls {
	var returns *AlbBackendGroupGrpcBackendTls
	_jsii_.Get(
		j,
		"tlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) Weight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) WeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightInput",
		&returns,
	)
	return returns
}


func NewAlbBackendGroupGrpcBackendOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AlbBackendGroupGrpcBackendOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AlbBackendGroupGrpcBackendOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbBackendGroupGrpcBackendOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAlbBackendGroupGrpcBackendOutputReference_Override(a AlbBackendGroupGrpcBackendOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbBackendGroupGrpcBackendOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) SetTargetGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"targetGroupIds",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) SetWeight(val *float64) {
	_jsii_.Set(
		j,
		"weight",
		val,
	)
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) PutHealthcheck(value *AlbBackendGroupGrpcBackendHealthcheck) {
	_jsii_.InvokeVoid(
		a,
		"putHealthcheck",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) PutLoadBalancingConfig(value *AlbBackendGroupGrpcBackendLoadBalancingConfig) {
	_jsii_.InvokeVoid(
		a,
		"putLoadBalancingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) PutTls(value *AlbBackendGroupGrpcBackendTls) {
	_jsii_.InvokeVoid(
		a,
		"putTls",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) ResetHealthcheck() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthcheck",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) ResetLoadBalancingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLoadBalancingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		a,
		"resetPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) ResetTls() {
	_jsii_.InvokeVoid(
		a,
		"resetTls",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) ResetWeight() {
	_jsii_.InvokeVoid(
		a,
		"resetWeight",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

