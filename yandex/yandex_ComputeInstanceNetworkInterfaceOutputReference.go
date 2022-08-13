// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeInstanceNetworkInterfaceOutputReference interface {
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
	DnsRecord() ComputeInstanceNetworkInterfaceDnsRecordList
	DnsRecordInput() interface{}
	// Experimental.
	Fqn() *string
	Index() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpAddress() *string
	SetIpAddress(val *string)
	IpAddressInput() *string
	Ipv4() interface{}
	SetIpv4(val interface{})
	Ipv4Input() interface{}
	Ipv6() interface{}
	SetIpv6(val interface{})
	Ipv6Address() *string
	SetIpv6Address(val *string)
	Ipv6AddressInput() *string
	Ipv6DnsRecord() ComputeInstanceNetworkInterfaceIpv6DnsRecordList
	Ipv6DnsRecordInput() interface{}
	Ipv6Input() interface{}
	MacAddress() *string
	Nat() interface{}
	SetNat(val interface{})
	NatDnsRecord() ComputeInstanceNetworkInterfaceNatDnsRecordList
	NatDnsRecordInput() interface{}
	NatInput() interface{}
	NatIpAddress() *string
	SetNatIpAddress(val *string)
	NatIpAddressInput() *string
	NatIpVersion() *string
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
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
	PutDnsRecord(value interface{})
	PutIpv6DnsRecord(value interface{})
	PutNatDnsRecord(value interface{})
	ResetDnsRecord()
	ResetIpAddress()
	ResetIpv4()
	ResetIpv6()
	ResetIpv6Address()
	ResetIpv6DnsRecord()
	ResetNat()
	ResetNatDnsRecord()
	ResetNatIpAddress()
	ResetSecurityGroupIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeInstanceNetworkInterfaceOutputReference
type jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) DnsRecord() ComputeInstanceNetworkInterfaceDnsRecordList {
	var returns ComputeInstanceNetworkInterfaceDnsRecordList
	_jsii_.Get(
		j,
		"dnsRecord",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) DnsRecordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dnsRecordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) Index() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"index",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) IpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) Ipv4() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) Ipv4Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) Ipv6() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) Ipv6Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) Ipv6AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6AddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) Ipv6DnsRecord() ComputeInstanceNetworkInterfaceIpv6DnsRecordList {
	var returns ComputeInstanceNetworkInterfaceIpv6DnsRecordList
	_jsii_.Get(
		j,
		"ipv6DnsRecord",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) Ipv6DnsRecordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6DnsRecordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) Ipv6Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) MacAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"macAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) Nat() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) NatDnsRecord() ComputeInstanceNetworkInterfaceNatDnsRecordList {
	var returns ComputeInstanceNetworkInterfaceNatDnsRecordList
	_jsii_.Get(
		j,
		"natDnsRecord",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) NatDnsRecordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"natDnsRecordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) NatInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"natInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) NatIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) NatIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) NatIpVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natIpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeInstanceNetworkInterfaceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComputeInstanceNetworkInterfaceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.ComputeInstanceNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComputeInstanceNetworkInterfaceOutputReference_Override(c ComputeInstanceNetworkInterfaceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.ComputeInstanceNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) SetIpAddress(val *string) {
	_jsii_.Set(
		j,
		"ipAddress",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) SetIpv4(val interface{}) {
	_jsii_.Set(
		j,
		"ipv4",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) SetIpv6(val interface{}) {
	_jsii_.Set(
		j,
		"ipv6",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) SetIpv6Address(val *string) {
	_jsii_.Set(
		j,
		"ipv6Address",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) SetNat(val interface{}) {
	_jsii_.Set(
		j,
		"nat",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) SetNatIpAddress(val *string) {
	_jsii_.Set(
		j,
		"natIpAddress",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) PutDnsRecord(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putDnsRecord",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) PutIpv6DnsRecord(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putIpv6DnsRecord",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) PutNatDnsRecord(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putNatDnsRecord",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) ResetDnsRecord() {
	_jsii_.InvokeVoid(
		c,
		"resetDnsRecord",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) ResetIpAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetIpAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) ResetIpv4() {
	_jsii_.InvokeVoid(
		c,
		"resetIpv4",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) ResetIpv6() {
	_jsii_.InvokeVoid(
		c,
		"resetIpv6",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) ResetIpv6Address() {
	_jsii_.InvokeVoid(
		c,
		"resetIpv6Address",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) ResetIpv6DnsRecord() {
	_jsii_.InvokeVoid(
		c,
		"resetIpv6DnsRecord",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) ResetNat() {
	_jsii_.InvokeVoid(
		c,
		"resetNat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) ResetNatDnsRecord() {
	_jsii_.InvokeVoid(
		c,
		"resetNatDnsRecord",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) ResetNatIpAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetNatIpAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		c,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

