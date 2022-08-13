// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbBackendGroupHttpBackendOutputReference interface {
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
	Healthcheck() AlbBackendGroupHttpBackendHealthcheckOutputReference
	HealthcheckInput() *AlbBackendGroupHttpBackendHealthcheck
	Http2() interface{}
	SetHttp2(val interface{})
	Http2Input() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LoadBalancingConfig() AlbBackendGroupHttpBackendLoadBalancingConfigOutputReference
	LoadBalancingConfigInput() *AlbBackendGroupHttpBackendLoadBalancingConfig
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	Tls() AlbBackendGroupHttpBackendTlsOutputReference
	TlsInput() *AlbBackendGroupHttpBackendTls
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
	PutHealthcheck(value *AlbBackendGroupHttpBackendHealthcheck)
	PutLoadBalancingConfig(value *AlbBackendGroupHttpBackendLoadBalancingConfig)
	PutTls(value *AlbBackendGroupHttpBackendTls)
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

// The jsii proxy struct for AlbBackendGroupHttpBackendOutputReference
type jsiiProxy_AlbBackendGroupHttpBackendOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) Healthcheck() AlbBackendGroupHttpBackendHealthcheckOutputReference {
	var returns AlbBackendGroupHttpBackendHealthcheckOutputReference
	_jsii_.Get(
		j,
		"healthcheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) HealthcheckInput() *AlbBackendGroupHttpBackendHealthcheck {
	var returns *AlbBackendGroupHttpBackendHealthcheck
	_jsii_.Get(
		j,
		"healthcheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) Http2() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) Http2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) LoadBalancingConfig() AlbBackendGroupHttpBackendLoadBalancingConfigOutputReference {
	var returns AlbBackendGroupHttpBackendLoadBalancingConfigOutputReference
	_jsii_.Get(
		j,
		"loadBalancingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) LoadBalancingConfigInput() *AlbBackendGroupHttpBackendLoadBalancingConfig {
	var returns *AlbBackendGroupHttpBackendLoadBalancingConfig
	_jsii_.Get(
		j,
		"loadBalancingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) StorageBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) StorageBucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) TargetGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) TargetGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) Tls() AlbBackendGroupHttpBackendTlsOutputReference {
	var returns AlbBackendGroupHttpBackendTlsOutputReference
	_jsii_.Get(
		j,
		"tls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) TlsInput() *AlbBackendGroupHttpBackendTls {
	var returns *AlbBackendGroupHttpBackendTls
	_jsii_.Get(
		j,
		"tlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) Weight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) WeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightInput",
		&returns,
	)
	return returns
}


func NewAlbBackendGroupHttpBackendOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AlbBackendGroupHttpBackendOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AlbBackendGroupHttpBackendOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbBackendGroupHttpBackendOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAlbBackendGroupHttpBackendOutputReference_Override(a AlbBackendGroupHttpBackendOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbBackendGroupHttpBackendOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) SetHttp2(val interface{}) {
	_jsii_.Set(
		j,
		"http2",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) SetStorageBucket(val *string) {
	_jsii_.Set(
		j,
		"storageBucket",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) SetTargetGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"targetGroupIds",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) SetWeight(val *float64) {
	_jsii_.Set(
		j,
		"weight",
		val,
	)
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) PutHealthcheck(value *AlbBackendGroupHttpBackendHealthcheck) {
	_jsii_.InvokeVoid(
		a,
		"putHealthcheck",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) PutLoadBalancingConfig(value *AlbBackendGroupHttpBackendLoadBalancingConfig) {
	_jsii_.InvokeVoid(
		a,
		"putLoadBalancingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) PutTls(value *AlbBackendGroupHttpBackendTls) {
	_jsii_.InvokeVoid(
		a,
		"putTls",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) ResetHealthcheck() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthcheck",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) ResetHttp2() {
	_jsii_.InvokeVoid(
		a,
		"resetHttp2",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) ResetLoadBalancingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLoadBalancingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		a,
		"resetPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) ResetStorageBucket() {
	_jsii_.InvokeVoid(
		a,
		"resetStorageBucket",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) ResetTargetGroupIds() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetGroupIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) ResetTls() {
	_jsii_.InvokeVoid(
		a,
		"resetTls",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) ResetWeight() {
	_jsii_.InvokeVoid(
		a,
		"resetWeight",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

