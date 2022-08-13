// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeInstanceNetworkInterfaceDnsRecordOutputReference interface {
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

// The jsii proxy struct for ComputeInstanceNetworkInterfaceDnsRecordOutputReference
type jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) DnsZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) DnsZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsZoneIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) Fqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) FqdnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) Ptr() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ptr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) PtrInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ptrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) Ttl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) TtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}


func NewComputeInstanceNetworkInterfaceDnsRecordOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComputeInstanceNetworkInterfaceDnsRecordOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.ComputeInstanceNetworkInterfaceDnsRecordOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComputeInstanceNetworkInterfaceDnsRecordOutputReference_Override(c ComputeInstanceNetworkInterfaceDnsRecordOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.ComputeInstanceNetworkInterfaceDnsRecordOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) SetDnsZoneId(val *string) {
	_jsii_.Set(
		j,
		"dnsZoneId",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) SetFqdn(val *string) {
	_jsii_.Set(
		j,
		"fqdn",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) SetPtr(val interface{}) {
	_jsii_.Set(
		j,
		"ptr",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) SetTtl(val *float64) {
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) ResetDnsZoneId() {
	_jsii_.InvokeVoid(
		c,
		"resetDnsZoneId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) ResetPtr() {
	_jsii_.InvokeVoid(
		c,
		"resetPtr",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) ResetTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceNetworkInterfaceDnsRecordOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

