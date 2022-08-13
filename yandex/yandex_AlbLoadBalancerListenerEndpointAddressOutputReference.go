// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbLoadBalancerListenerEndpointAddressOutputReference interface {
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
	ExternalIpv4Address() AlbLoadBalancerListenerEndpointAddressExternalIpv4AddressOutputReference
	ExternalIpv4AddressInput() *AlbLoadBalancerListenerEndpointAddressExternalIpv4Address
	ExternalIpv6Address() AlbLoadBalancerListenerEndpointAddressExternalIpv6AddressOutputReference
	ExternalIpv6AddressInput() *AlbLoadBalancerListenerEndpointAddressExternalIpv6Address
	// Experimental.
	Fqn() *string
	InternalIpv4Address() AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference
	InternalIpv4AddressInput() *AlbLoadBalancerListenerEndpointAddressInternalIpv4Address
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	PutExternalIpv4Address(value *AlbLoadBalancerListenerEndpointAddressExternalIpv4Address)
	PutExternalIpv6Address(value *AlbLoadBalancerListenerEndpointAddressExternalIpv6Address)
	PutInternalIpv4Address(value *AlbLoadBalancerListenerEndpointAddressInternalIpv4Address)
	ResetExternalIpv4Address()
	ResetExternalIpv6Address()
	ResetInternalIpv4Address()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlbLoadBalancerListenerEndpointAddressOutputReference
type jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) ExternalIpv4Address() AlbLoadBalancerListenerEndpointAddressExternalIpv4AddressOutputReference {
	var returns AlbLoadBalancerListenerEndpointAddressExternalIpv4AddressOutputReference
	_jsii_.Get(
		j,
		"externalIpv4Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) ExternalIpv4AddressInput() *AlbLoadBalancerListenerEndpointAddressExternalIpv4Address {
	var returns *AlbLoadBalancerListenerEndpointAddressExternalIpv4Address
	_jsii_.Get(
		j,
		"externalIpv4AddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) ExternalIpv6Address() AlbLoadBalancerListenerEndpointAddressExternalIpv6AddressOutputReference {
	var returns AlbLoadBalancerListenerEndpointAddressExternalIpv6AddressOutputReference
	_jsii_.Get(
		j,
		"externalIpv6Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) ExternalIpv6AddressInput() *AlbLoadBalancerListenerEndpointAddressExternalIpv6Address {
	var returns *AlbLoadBalancerListenerEndpointAddressExternalIpv6Address
	_jsii_.Get(
		j,
		"externalIpv6AddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) InternalIpv4Address() AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference {
	var returns AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressOutputReference
	_jsii_.Get(
		j,
		"internalIpv4Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) InternalIpv4AddressInput() *AlbLoadBalancerListenerEndpointAddressInternalIpv4Address {
	var returns *AlbLoadBalancerListenerEndpointAddressInternalIpv4Address
	_jsii_.Get(
		j,
		"internalIpv4AddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAlbLoadBalancerListenerEndpointAddressOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AlbLoadBalancerListenerEndpointAddressOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbLoadBalancerListenerEndpointAddressOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAlbLoadBalancerListenerEndpointAddressOutputReference_Override(a AlbLoadBalancerListenerEndpointAddressOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbLoadBalancerListenerEndpointAddressOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) PutExternalIpv4Address(value *AlbLoadBalancerListenerEndpointAddressExternalIpv4Address) {
	_jsii_.InvokeVoid(
		a,
		"putExternalIpv4Address",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) PutExternalIpv6Address(value *AlbLoadBalancerListenerEndpointAddressExternalIpv6Address) {
	_jsii_.InvokeVoid(
		a,
		"putExternalIpv6Address",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) PutInternalIpv4Address(value *AlbLoadBalancerListenerEndpointAddressInternalIpv4Address) {
	_jsii_.InvokeVoid(
		a,
		"putInternalIpv4Address",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) ResetExternalIpv4Address() {
	_jsii_.InvokeVoid(
		a,
		"resetExternalIpv4Address",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) ResetExternalIpv6Address() {
	_jsii_.InvokeVoid(
		a,
		"resetExternalIpv6Address",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) ResetInternalIpv4Address() {
	_jsii_.InvokeVoid(
		a,
		"resetInternalIpv4Address",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLoadBalancerListenerEndpointAddressOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

