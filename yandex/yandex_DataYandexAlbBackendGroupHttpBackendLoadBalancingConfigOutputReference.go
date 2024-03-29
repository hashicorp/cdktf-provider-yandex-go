// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference interface {
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
	InternalValue() *DataYandexAlbBackendGroupHttpBackendLoadBalancingConfig
	SetInternalValue(val *DataYandexAlbBackendGroupHttpBackendLoadBalancingConfig)
	LocalityAwareRoutingPercent() *float64
	SetLocalityAwareRoutingPercent(val *float64)
	LocalityAwareRoutingPercentInput() *float64
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	PanicThreshold() *float64
	SetPanicThreshold(val *float64)
	PanicThresholdInput() *float64
	StrictLocality() interface{}
	SetStrictLocality(val interface{})
	StrictLocalityInput() interface{}
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
	ResetLocalityAwareRoutingPercent()
	ResetMode()
	ResetPanicThreshold()
	ResetStrictLocality()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference
type jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) InternalValue() *DataYandexAlbBackendGroupHttpBackendLoadBalancingConfig {
	var returns *DataYandexAlbBackendGroupHttpBackendLoadBalancingConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) LocalityAwareRoutingPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localityAwareRoutingPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) LocalityAwareRoutingPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localityAwareRoutingPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) PanicThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"panicThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) PanicThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"panicThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) StrictLocality() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictLocality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) StrictLocalityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictLocalityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference_Override(d DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) SetInternalValue(val *DataYandexAlbBackendGroupHttpBackendLoadBalancingConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) SetLocalityAwareRoutingPercent(val *float64) {
	_jsii_.Set(
		j,
		"localityAwareRoutingPercent",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) SetMode(val *string) {
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) SetPanicThreshold(val *float64) {
	_jsii_.Set(
		j,
		"panicThreshold",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) SetStrictLocality(val interface{}) {
	_jsii_.Set(
		j,
		"strictLocality",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) ResetLocalityAwareRoutingPercent() {
	_jsii_.InvokeVoid(
		d,
		"resetLocalityAwareRoutingPercent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		d,
		"resetMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) ResetPanicThreshold() {
	_jsii_.InvokeVoid(
		d,
		"resetPanicThreshold",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) ResetStrictLocality() {
	_jsii_.InvokeVoid(
		d,
		"resetStrictLocality",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupHttpBackendLoadBalancingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

