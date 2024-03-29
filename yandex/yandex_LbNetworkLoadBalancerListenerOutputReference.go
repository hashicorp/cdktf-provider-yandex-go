// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbNetworkLoadBalancerListenerOutputReference interface {
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
	ExternalAddressSpec() LbNetworkLoadBalancerListenerExternalAddressSpecOutputReference
	ExternalAddressSpecInput() *LbNetworkLoadBalancerListenerExternalAddressSpec
	// Experimental.
	Fqn() *string
	InternalAddressSpec() LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference
	InternalAddressSpecInput() *LbNetworkLoadBalancerListenerInternalAddressSpec
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	TargetPort() *float64
	SetTargetPort(val *float64)
	TargetPortInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutExternalAddressSpec(value *LbNetworkLoadBalancerListenerExternalAddressSpec)
	PutInternalAddressSpec(value *LbNetworkLoadBalancerListenerInternalAddressSpec)
	ResetExternalAddressSpec()
	ResetInternalAddressSpec()
	ResetProtocol()
	ResetTargetPort()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LbNetworkLoadBalancerListenerOutputReference
type jsiiProxy_LbNetworkLoadBalancerListenerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) ExternalAddressSpec() LbNetworkLoadBalancerListenerExternalAddressSpecOutputReference {
	var returns LbNetworkLoadBalancerListenerExternalAddressSpecOutputReference
	_jsii_.Get(
		j,
		"externalAddressSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) ExternalAddressSpecInput() *LbNetworkLoadBalancerListenerExternalAddressSpec {
	var returns *LbNetworkLoadBalancerListenerExternalAddressSpec
	_jsii_.Get(
		j,
		"externalAddressSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) InternalAddressSpec() LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference {
	var returns LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference
	_jsii_.Get(
		j,
		"internalAddressSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) InternalAddressSpecInput() *LbNetworkLoadBalancerListenerInternalAddressSpec {
	var returns *LbNetworkLoadBalancerListenerInternalAddressSpec
	_jsii_.Get(
		j,
		"internalAddressSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) TargetPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) TargetPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLbNetworkLoadBalancerListenerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LbNetworkLoadBalancerListenerOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LbNetworkLoadBalancerListenerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.LbNetworkLoadBalancerListenerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLbNetworkLoadBalancerListenerOutputReference_Override(l LbNetworkLoadBalancerListenerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.LbNetworkLoadBalancerListenerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) SetProtocol(val *string) {
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) SetTargetPort(val *float64) {
	_jsii_.Set(
		j,
		"targetPort",
		val,
	)
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) PutExternalAddressSpec(value *LbNetworkLoadBalancerListenerExternalAddressSpec) {
	_jsii_.InvokeVoid(
		l,
		"putExternalAddressSpec",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) PutInternalAddressSpec(value *LbNetworkLoadBalancerListenerInternalAddressSpec) {
	_jsii_.InvokeVoid(
		l,
		"putInternalAddressSpec",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) ResetExternalAddressSpec() {
	_jsii_.InvokeVoid(
		l,
		"resetExternalAddressSpec",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) ResetInternalAddressSpec() {
	_jsii_.InvokeVoid(
		l,
		"resetInternalAddressSpec",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) ResetProtocol() {
	_jsii_.InvokeVoid(
		l,
		"resetProtocol",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) ResetTargetPort() {
	_jsii_.InvokeVoid(
		l,
		"resetTargetPort",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

