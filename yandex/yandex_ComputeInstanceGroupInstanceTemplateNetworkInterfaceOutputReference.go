// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference interface {
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
	DnsRecord() ComputeInstanceGroupInstanceTemplateNetworkInterfaceDnsRecordList
	DnsRecordInput() interface{}
	// Experimental.
	Fqn() *string
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
	Ipv6DnsRecord() ComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordList
	Ipv6DnsRecordInput() interface{}
	Ipv6Input() interface{}
	Nat() interface{}
	SetNat(val interface{})
	NatDnsRecord() ComputeInstanceGroupInstanceTemplateNetworkInterfaceNatDnsRecordList
	NatDnsRecordInput() interface{}
	NatInput() interface{}
	NatIpAddress() *string
	SetNatIpAddress(val *string)
	NatIpAddressInput() *string
	NetworkId() *string
	SetNetworkId(val *string)
	NetworkIdInput() *string
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	SubnetIdsInput() *[]*string
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
	ResetNetworkId()
	ResetSecurityGroupIds()
	ResetSubnetIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference
type jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) DnsRecord() ComputeInstanceGroupInstanceTemplateNetworkInterfaceDnsRecordList {
	var returns ComputeInstanceGroupInstanceTemplateNetworkInterfaceDnsRecordList
	_jsii_.Get(
		j,
		"dnsRecord",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) DnsRecordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dnsRecordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) IpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) Ipv4() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) Ipv4Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) Ipv6() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) Ipv6Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) Ipv6AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6AddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) Ipv6DnsRecord() ComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordList {
	var returns ComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordList
	_jsii_.Get(
		j,
		"ipv6DnsRecord",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) Ipv6DnsRecordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6DnsRecordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) Ipv6Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) Nat() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) NatDnsRecord() ComputeInstanceGroupInstanceTemplateNetworkInterfaceNatDnsRecordList {
	var returns ComputeInstanceGroupInstanceTemplateNetworkInterfaceNatDnsRecordList
	_jsii_.Get(
		j,
		"natDnsRecord",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) NatDnsRecordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"natDnsRecordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) NatInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"natInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) NatIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) NatIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) NetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) NetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference_Override(c ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) SetIpAddress(val *string) {
	_jsii_.Set(
		j,
		"ipAddress",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) SetIpv4(val interface{}) {
	_jsii_.Set(
		j,
		"ipv4",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) SetIpv6(val interface{}) {
	_jsii_.Set(
		j,
		"ipv6",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) SetIpv6Address(val *string) {
	_jsii_.Set(
		j,
		"ipv6Address",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) SetNat(val interface{}) {
	_jsii_.Set(
		j,
		"nat",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) SetNatIpAddress(val *string) {
	_jsii_.Set(
		j,
		"natIpAddress",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) SetNetworkId(val *string) {
	_jsii_.Set(
		j,
		"networkId",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) PutDnsRecord(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putDnsRecord",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) PutIpv6DnsRecord(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putIpv6DnsRecord",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) PutNatDnsRecord(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putNatDnsRecord",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) ResetDnsRecord() {
	_jsii_.InvokeVoid(
		c,
		"resetDnsRecord",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) ResetIpAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetIpAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) ResetIpv4() {
	_jsii_.InvokeVoid(
		c,
		"resetIpv4",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) ResetIpv6() {
	_jsii_.InvokeVoid(
		c,
		"resetIpv6",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) ResetIpv6Address() {
	_jsii_.InvokeVoid(
		c,
		"resetIpv6Address",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) ResetIpv6DnsRecord() {
	_jsii_.InvokeVoid(
		c,
		"resetIpv6DnsRecord",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) ResetNat() {
	_jsii_.InvokeVoid(
		c,
		"resetNat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) ResetNatDnsRecord() {
	_jsii_.InvokeVoid(
		c,
		"resetNatDnsRecord",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) ResetNatIpAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetNatIpAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) ResetNetworkId() {
	_jsii_.InvokeVoid(
		c,
		"resetNetworkId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		c,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) ResetSubnetIds() {
	_jsii_.InvokeVoid(
		c,
		"resetSubnetIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

