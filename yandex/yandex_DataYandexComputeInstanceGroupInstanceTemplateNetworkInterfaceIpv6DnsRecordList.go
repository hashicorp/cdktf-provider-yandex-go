// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordList
type jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordList {
	_init_.Initialize()

	j := jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordList{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordList_Override(d DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordList) Get(index *float64) DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordOutputReference {
	var returns DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

