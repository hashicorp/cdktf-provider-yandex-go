// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference interface {
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
	DnsRecord() DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceDnsRecordList
	// Experimental.
	Fqn() *string
	InternalValue() *DataYandexComputeInstanceGroupInstanceTemplateNetworkInterface
	SetInternalValue(val *DataYandexComputeInstanceGroupInstanceTemplateNetworkInterface)
	IpAddress() *string
	Ipv4() cdktf.IResolvable
	Ipv6() cdktf.IResolvable
	Ipv6Address() *string
	Ipv6DnsRecord() DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordList
	Nat() cdktf.IResolvable
	NatDnsRecord() DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceNatDnsRecordList
	NatIpAddress() *string
	NetworkId() *string
	SecurityGroupIds() *[]*string
	SubnetIds() *[]*string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference
type jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) DnsRecord() DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceDnsRecordList {
	var returns DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceDnsRecordList
	_jsii_.Get(
		j,
		"dnsRecord",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) InternalValue() *DataYandexComputeInstanceGroupInstanceTemplateNetworkInterface {
	var returns *DataYandexComputeInstanceGroupInstanceTemplateNetworkInterface
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) Ipv4() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"ipv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) Ipv6() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"ipv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) Ipv6Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) Ipv6DnsRecord() DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordList {
	var returns DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordList
	_jsii_.Get(
		j,
		"ipv6DnsRecord",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) Nat() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"nat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) NatDnsRecord() DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceNatDnsRecordList {
	var returns DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceNatDnsRecordList
	_jsii_.Get(
		j,
		"natDnsRecord",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) NatIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) NetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference_Override(d DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) SetInternalValue(val *DataYandexComputeInstanceGroupInstanceTemplateNetworkInterface) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

