// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexAlbLoadBalancerListenerEndpointAddressExternalIpv4AddressList interface {
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
	Get(index *float64) DataYandexAlbLoadBalancerListenerEndpointAddressExternalIpv4AddressOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataYandexAlbLoadBalancerListenerEndpointAddressExternalIpv4AddressList
type jsiiProxy_DataYandexAlbLoadBalancerListenerEndpointAddressExternalIpv4AddressList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerEndpointAddressExternalIpv4AddressList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerEndpointAddressExternalIpv4AddressList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerEndpointAddressExternalIpv4AddressList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerEndpointAddressExternalIpv4AddressList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerEndpointAddressExternalIpv4AddressList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataYandexAlbLoadBalancerListenerEndpointAddressExternalIpv4AddressList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataYandexAlbLoadBalancerListenerEndpointAddressExternalIpv4AddressList {
	_init_.Initialize()

	j := jsiiProxy_DataYandexAlbLoadBalancerListenerEndpointAddressExternalIpv4AddressList{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexAlbLoadBalancerListenerEndpointAddressExternalIpv4AddressList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataYandexAlbLoadBalancerListenerEndpointAddressExternalIpv4AddressList_Override(d DataYandexAlbLoadBalancerListenerEndpointAddressExternalIpv4AddressList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexAlbLoadBalancerListenerEndpointAddressExternalIpv4AddressList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerEndpointAddressExternalIpv4AddressList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerEndpointAddressExternalIpv4AddressList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerEndpointAddressExternalIpv4AddressList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataYandexAlbLoadBalancerListenerEndpointAddressExternalIpv4AddressList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbLoadBalancerListenerEndpointAddressExternalIpv4AddressList) Get(index *float64) DataYandexAlbLoadBalancerListenerEndpointAddressExternalIpv4AddressOutputReference {
	var returns DataYandexAlbLoadBalancerListenerEndpointAddressExternalIpv4AddressOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbLoadBalancerListenerEndpointAddressExternalIpv4AddressList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbLoadBalancerListenerEndpointAddressExternalIpv4AddressList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
