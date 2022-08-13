// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference interface {
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
	InternalValue() *AlbBackendGroupStreamBackendLoadBalancingConfig
	SetInternalValue(val *AlbBackendGroupStreamBackendLoadBalancingConfig)
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

// The jsii proxy struct for AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference
type jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) InternalValue() *AlbBackendGroupStreamBackendLoadBalancingConfig {
	var returns *AlbBackendGroupStreamBackendLoadBalancingConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) LocalityAwareRoutingPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localityAwareRoutingPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) LocalityAwareRoutingPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localityAwareRoutingPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) PanicThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"panicThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) PanicThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"panicThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) StrictLocality() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictLocality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) StrictLocalityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictLocalityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAlbBackendGroupStreamBackendLoadBalancingConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlbBackendGroupStreamBackendLoadBalancingConfigOutputReference_Override(a AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) SetInternalValue(val *AlbBackendGroupStreamBackendLoadBalancingConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) SetLocalityAwareRoutingPercent(val *float64) {
	_jsii_.Set(
		j,
		"localityAwareRoutingPercent",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) SetMode(val *string) {
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) SetPanicThreshold(val *float64) {
	_jsii_.Set(
		j,
		"panicThreshold",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) SetStrictLocality(val interface{}) {
	_jsii_.Set(
		j,
		"strictLocality",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) ResetLocalityAwareRoutingPercent() {
	_jsii_.InvokeVoid(
		a,
		"resetLocalityAwareRoutingPercent",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		a,
		"resetMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) ResetPanicThreshold() {
	_jsii_.InvokeVoid(
		a,
		"resetPanicThreshold",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) ResetStrictLocality() {
	_jsii_.InvokeVoid(
		a,
		"resetStrictLocality",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupStreamBackendLoadBalancingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

