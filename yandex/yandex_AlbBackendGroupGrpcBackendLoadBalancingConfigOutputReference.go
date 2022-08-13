// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference interface {
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
	InternalValue() *AlbBackendGroupGrpcBackendLoadBalancingConfig
	SetInternalValue(val *AlbBackendGroupGrpcBackendLoadBalancingConfig)
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

// The jsii proxy struct for AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference
type jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) InternalValue() *AlbBackendGroupGrpcBackendLoadBalancingConfig {
	var returns *AlbBackendGroupGrpcBackendLoadBalancingConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) LocalityAwareRoutingPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localityAwareRoutingPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) LocalityAwareRoutingPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localityAwareRoutingPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) PanicThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"panicThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) PanicThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"panicThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) StrictLocality() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictLocality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) StrictLocalityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictLocalityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference_Override(a AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) SetInternalValue(val *AlbBackendGroupGrpcBackendLoadBalancingConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) SetLocalityAwareRoutingPercent(val *float64) {
	_jsii_.Set(
		j,
		"localityAwareRoutingPercent",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) SetMode(val *string) {
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) SetPanicThreshold(val *float64) {
	_jsii_.Set(
		j,
		"panicThreshold",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) SetStrictLocality(val interface{}) {
	_jsii_.Set(
		j,
		"strictLocality",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) ResetLocalityAwareRoutingPercent() {
	_jsii_.InvokeVoid(
		a,
		"resetLocalityAwareRoutingPercent",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		a,
		"resetMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) ResetPanicThreshold() {
	_jsii_.InvokeVoid(
		a,
		"resetPanicThreshold",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) ResetStrictLocality() {
	_jsii_.InvokeVoid(
		a,
		"resetStrictLocality",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendLoadBalancingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

