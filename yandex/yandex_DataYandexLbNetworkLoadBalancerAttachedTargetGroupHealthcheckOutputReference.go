// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference interface {
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
	HealthyThreshold() *float64
	HttpOptions() DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckHttpOptionsList
	InternalValue() *DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheck
	SetInternalValue(val *DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheck)
	Interval() *float64
	Name() *string
	TcpOptions() DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckTcpOptionsList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() *float64
	UnhealthyThreshold() *float64
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

// The jsii proxy struct for DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference
type jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) HealthyThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthyThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) HttpOptions() DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckHttpOptionsList {
	var returns DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckHttpOptionsList
	_jsii_.Get(
		j,
		"httpOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) InternalValue() *DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheck {
	var returns *DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheck
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) Interval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) TcpOptions() DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckTcpOptionsList {
	var returns DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckTcpOptionsList
	_jsii_.Get(
		j,
		"tcpOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) UnhealthyThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyThreshold",
		&returns,
	)
	return returns
}


func NewDataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference_Override(d DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) SetInternalValue(val *DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheck) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexLbNetworkLoadBalancerAttachedTargetGroupHealthcheckOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

