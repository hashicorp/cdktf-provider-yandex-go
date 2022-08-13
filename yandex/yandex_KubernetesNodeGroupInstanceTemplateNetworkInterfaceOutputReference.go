// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ipv4() interface{}
	SetIpv4(val interface{})
	Ipv4DnsRecords() KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsList
	Ipv4DnsRecordsInput() interface{}
	Ipv4Input() interface{}
	Ipv6() interface{}
	SetIpv6(val interface{})
	Ipv6DnsRecords() KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordsList
	Ipv6DnsRecordsInput() interface{}
	Ipv6Input() interface{}
	Nat() interface{}
	SetNat(val interface{})
	NatInput() interface{}
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
	PutIpv4DnsRecords(value interface{})
	PutIpv6DnsRecords(value interface{})
	ResetIpv4()
	ResetIpv4DnsRecords()
	ResetIpv6()
	ResetIpv6DnsRecords()
	ResetNat()
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

// The jsii proxy struct for KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference
type jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) Ipv4() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) Ipv4DnsRecords() KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsList {
	var returns KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsList
	_jsii_.Get(
		j,
		"ipv4DnsRecords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) Ipv4DnsRecordsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv4DnsRecordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) Ipv4Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) Ipv6() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) Ipv6DnsRecords() KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordsList {
	var returns KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordsList
	_jsii_.Get(
		j,
		"ipv6DnsRecords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) Ipv6DnsRecordsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6DnsRecordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) Ipv6Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) Nat() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) NatInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"natInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewKubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference_Override(k KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		k,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) SetIpv4(val interface{}) {
	_jsii_.Set(
		j,
		"ipv4",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) SetIpv6(val interface{}) {
	_jsii_.Set(
		j,
		"ipv6",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) SetNat(val interface{}) {
	_jsii_.Set(
		j,
		"nat",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) PutIpv4DnsRecords(value interface{}) {
	_jsii_.InvokeVoid(
		k,
		"putIpv4DnsRecords",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) PutIpv6DnsRecords(value interface{}) {
	_jsii_.InvokeVoid(
		k,
		"putIpv6DnsRecords",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) ResetIpv4() {
	_jsii_.InvokeVoid(
		k,
		"resetIpv4",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) ResetIpv4DnsRecords() {
	_jsii_.InvokeVoid(
		k,
		"resetIpv4DnsRecords",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) ResetIpv6() {
	_jsii_.InvokeVoid(
		k,
		"resetIpv6",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) ResetIpv6DnsRecords() {
	_jsii_.InvokeVoid(
		k,
		"resetIpv6DnsRecords",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) ResetNat() {
	_jsii_.InvokeVoid(
		k,
		"resetNat",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		k,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

