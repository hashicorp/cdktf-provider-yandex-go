// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference interface {
	cdktf.ComplexObject
	Address() *string
	SetAddress(val *string)
	AddressInput() *string
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
	InternalValue() *LbNetworkLoadBalancerListenerInternalAddressSpec
	SetInternalValue(val *LbNetworkLoadBalancerListenerInternalAddressSpec)
	IpVersion() *string
	SetIpVersion(val *string)
	IpVersionInput() *string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
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
	ResetAddress()
	ResetIpVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference
type jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) InternalValue() *LbNetworkLoadBalancerListenerInternalAddressSpec {
	var returns *LbNetworkLoadBalancerListenerInternalAddressSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) IpVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) IpVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLbNetworkLoadBalancerListenerInternalAddressSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLbNetworkLoadBalancerListenerInternalAddressSpecOutputReference_Override(l LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) SetAddress(val *string) {
	_jsii_.Set(
		j,
		"address",
		val,
	)
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) SetInternalValue(val *LbNetworkLoadBalancerListenerInternalAddressSpec) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) SetIpVersion(val *string) {
	_jsii_.Set(
		j,
		"ipVersion",
		val,
	)
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) ResetAddress() {
	_jsii_.InvokeVoid(
		l,
		"resetAddress",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) ResetIpVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetIpVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNetworkLoadBalancerListenerInternalAddressSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

