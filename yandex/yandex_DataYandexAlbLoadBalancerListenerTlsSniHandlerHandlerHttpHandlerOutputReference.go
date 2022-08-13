// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference interface {
	cdktf.ComplexObject
	AllowHttp10() cdktf.IResolvable
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
	Http2Options() DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsList
	HttpRouterId() *string
	InternalValue() *DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandler
	SetInternalValue(val *DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandler)
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

// The jsii proxy struct for DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference
type jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference) AllowHttp10() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"allowHttp10",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference) Http2Options() DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsList {
	var returns DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsList
	_jsii_.Get(
		j,
		"http2Options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference) HttpRouterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpRouterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference) InternalValue() *DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandler {
	var returns *DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandler
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference_Override(d DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference) SetInternalValue(val *DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandler) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

