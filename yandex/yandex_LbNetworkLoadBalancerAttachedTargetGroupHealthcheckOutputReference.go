// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference interface {
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
	HealthyThreshold() *float64
	SetHealthyThreshold(val *float64)
	HealthyThresholdInput() *float64
	HttpOptions() LbNetworkLoadBalancerAttachedTargetGroupHealthcheckHttpOptionsOutputReference
	HttpOptionsInput() *LbNetworkLoadBalancerAttachedTargetGroupHealthcheckHttpOptions
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Interval() *float64
	SetInterval(val *float64)
	IntervalInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	TcpOptions() LbNetworkLoadBalancerAttachedTargetGroupHealthcheckTcpOptionsOutputReference
	TcpOptionsInput() *LbNetworkLoadBalancerAttachedTargetGroupHealthcheckTcpOptions
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() *float64
	SetTimeout(val *float64)
	TimeoutInput() *float64
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
	PutHttpOptions(value *LbNetworkLoadBalancerAttachedTargetGroupHealthcheckHttpOptions)
	PutTcpOptions(value *LbNetworkLoadBalancerAttachedTargetGroupHealthcheckTcpOptions)
	ResetHealthyThreshold()
	ResetHttpOptions()
	ResetInterval()
	ResetTcpOptions()
	ResetTimeout()
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

// The jsii proxy struct for LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference
type jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) HealthyThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthyThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) HealthyThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthyThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) HttpOptions() LbNetworkLoadBalancerAttachedTargetGroupHealthcheckHttpOptionsOutputReference {
	var returns LbNetworkLoadBalancerAttachedTargetGroupHealthcheckHttpOptionsOutputReference
	_jsii_.Get(
		j,
		"httpOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) HttpOptionsInput() *LbNetworkLoadBalancerAttachedTargetGroupHealthcheckHttpOptions {
	var returns *LbNetworkLoadBalancerAttachedTargetGroupHealthcheckHttpOptions
	_jsii_.Get(
		j,
		"httpOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) Interval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) IntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) TcpOptions() LbNetworkLoadBalancerAttachedTargetGroupHealthcheckTcpOptionsOutputReference {
	var returns LbNetworkLoadBalancerAttachedTargetGroupHealthcheckTcpOptionsOutputReference
	_jsii_.Get(
		j,
		"tcpOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) TcpOptionsInput() *LbNetworkLoadBalancerAttachedTargetGroupHealthcheckTcpOptions {
	var returns *LbNetworkLoadBalancerAttachedTargetGroupHealthcheckTcpOptions
	_jsii_.Get(
		j,
		"tcpOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) UnhealthyThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) UnhealthyThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyThresholdInput",
		&returns,
	)
	return returns
}


func NewLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference_Override(l LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) SetHealthyThreshold(val *float64) {
	_jsii_.Set(
		j,
		"healthyThreshold",
		val,
	)
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) SetInterval(val *float64) {
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) SetTimeout(val *float64) {
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) SetUnhealthyThreshold(val *float64) {
	_jsii_.Set(
		j,
		"unhealthyThreshold",
		val,
	)
}

func (l *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) PutHttpOptions(value *LbNetworkLoadBalancerAttachedTargetGroupHealthcheckHttpOptions) {
	_jsii_.InvokeVoid(
		l,
		"putHttpOptions",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) PutTcpOptions(value *LbNetworkLoadBalancerAttachedTargetGroupHealthcheckTcpOptions) {
	_jsii_.InvokeVoid(
		l,
		"putTcpOptions",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) ResetHealthyThreshold() {
	_jsii_.InvokeVoid(
		l,
		"resetHealthyThreshold",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) ResetHttpOptions() {
	_jsii_.InvokeVoid(
		l,
		"resetHttpOptions",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) ResetInterval() {
	_jsii_.InvokeVoid(
		l,
		"resetInterval",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) ResetTcpOptions() {
	_jsii_.InvokeVoid(
		l,
		"resetTcpOptions",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeout",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) ResetUnhealthyThreshold() {
	_jsii_.InvokeVoid(
		l,
		"resetUnhealthyThreshold",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

