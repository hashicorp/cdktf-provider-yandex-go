// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpcAddressExternalIpv4AddressOutputReference interface {
	cdktf.ComplexObject
	Address() *string
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
	DdosProtectionProvider() *string
	SetDdosProtectionProvider(val *string)
	DdosProtectionProviderInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *VpcAddressExternalIpv4Address
	SetInternalValue(val *VpcAddressExternalIpv4Address)
	OutgoingSmtpCapability() *string
	SetOutgoingSmtpCapability(val *string)
	OutgoingSmtpCapabilityInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ZoneId() *string
	SetZoneId(val *string)
	ZoneIdInput() *string
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
	ResetDdosProtectionProvider()
	ResetOutgoingSmtpCapability()
	ResetZoneId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VpcAddressExternalIpv4AddressOutputReference
type jsiiProxy_VpcAddressExternalIpv4AddressOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) DdosProtectionProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ddosProtectionProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) DdosProtectionProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ddosProtectionProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) InternalValue() *VpcAddressExternalIpv4Address {
	var returns *VpcAddressExternalIpv4Address
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) OutgoingSmtpCapability() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outgoingSmtpCapability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) OutgoingSmtpCapabilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outgoingSmtpCapabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) ZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneIdInput",
		&returns,
	)
	return returns
}


func NewVpcAddressExternalIpv4AddressOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VpcAddressExternalIpv4AddressOutputReference {
	_init_.Initialize()

	j := jsiiProxy_VpcAddressExternalIpv4AddressOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.VpcAddressExternalIpv4AddressOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVpcAddressExternalIpv4AddressOutputReference_Override(v VpcAddressExternalIpv4AddressOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.VpcAddressExternalIpv4AddressOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) SetDdosProtectionProvider(val *string) {
	_jsii_.Set(
		j,
		"ddosProtectionProvider",
		val,
	)
}

func (j *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) SetInternalValue(val *VpcAddressExternalIpv4Address) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) SetOutgoingSmtpCapability(val *string) {
	_jsii_.Set(
		j,
		"outgoingSmtpCapability",
		val,
	)
}

func (j *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) SetZoneId(val *string) {
	_jsii_.Set(
		j,
		"zoneId",
		val,
	)
}

func (v *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) ResetDdosProtectionProvider() {
	_jsii_.InvokeVoid(
		v,
		"resetDdosProtectionProvider",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) ResetOutgoingSmtpCapability() {
	_jsii_.InvokeVoid(
		v,
		"resetOutgoingSmtpCapability",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) ResetZoneId() {
	_jsii_.InvokeVoid(
		v,
		"resetZoneId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcAddressExternalIpv4AddressOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

