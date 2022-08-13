// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsList interface {
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
	Get(index *float64) DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsList
type jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsList {
	_init_.Initialize()

	j := jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsList{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsList_Override(d DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsList) Get(index *float64) DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference {
	var returns DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

