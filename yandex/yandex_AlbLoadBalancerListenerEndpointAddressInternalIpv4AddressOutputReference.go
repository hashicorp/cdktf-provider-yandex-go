// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference interface {
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
	InternalValue() *AlbLoadBalancerListenerEndpointAddressInternalIpv4Address
	SetInternalValue(val *AlbLoadBalancerListenerEndpointAddressInternalIpv4Address)
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
	ResetSubnetId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference
type jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) InternalValue() *AlbLoadBalancerListenerEndpointAddressInternalIpv4Address {
	var returns *AlbLoadBalancerListenerEndpointAddressInternalIpv4Address
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference_Override(a AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) SetAddress(val *string) {
	_jsii_.Set(
		j,
		"address",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) SetInternalValue(val *AlbLoadBalancerListenerEndpointAddressInternalIpv4Address) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) ResetAddress() {
	_jsii_.InvokeVoid(
		a,
		"resetAddress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		a,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

