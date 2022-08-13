// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference interface {
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
	DnsZoneId() *string
	SetDnsZoneId(val *string)
	DnsZoneIdInput() *string
	Fqdn() *string
	SetFqdn(val *string)
	FqdnInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ptr() interface{}
	SetPtr(val interface{})
	PtrInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Ttl() *float64
	SetTtl(val *float64)
	TtlInput() *float64
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
	ResetDnsZoneId()
	ResetPtr()
	ResetTtl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference
type jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) DnsZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) DnsZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsZoneIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) Fqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) FqdnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) Ptr() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ptr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) PtrInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ptrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) Ttl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) TtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}


func NewKubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewKubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference_Override(k KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		k,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) SetDnsZoneId(val *string) {
	_jsii_.Set(
		j,
		"dnsZoneId",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) SetFqdn(val *string) {
	_jsii_.Set(
		j,
		"fqdn",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) SetPtr(val interface{}) {
	_jsii_.Set(
		j,
		"ptr",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) SetTtl(val *float64) {
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) ResetDnsZoneId() {
	_jsii_.InvokeVoid(
		k,
		"resetDnsZoneId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) ResetPtr() {
	_jsii_.InvokeVoid(
		k,
		"resetPtr",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) ResetTtl() {
	_jsii_.InvokeVoid(
		k,
		"resetTtl",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecordsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

