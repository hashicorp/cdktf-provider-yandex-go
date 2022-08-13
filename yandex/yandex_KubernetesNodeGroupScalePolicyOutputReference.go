// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesNodeGroupScalePolicyOutputReference interface {
	cdktf.ComplexObject
	AutoScale() KubernetesNodeGroupScalePolicyAutoScaleOutputReference
	AutoScaleInput() *KubernetesNodeGroupScalePolicyAutoScale
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
	FixedScale() KubernetesNodeGroupScalePolicyFixedScaleOutputReference
	FixedScaleInput() *KubernetesNodeGroupScalePolicyFixedScale
	// Experimental.
	Fqn() *string
	InternalValue() *KubernetesNodeGroupScalePolicy
	SetInternalValue(val *KubernetesNodeGroupScalePolicy)
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
	PutAutoScale(value *KubernetesNodeGroupScalePolicyAutoScale)
	PutFixedScale(value *KubernetesNodeGroupScalePolicyFixedScale)
	ResetAutoScale()
	ResetFixedScale()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesNodeGroupScalePolicyOutputReference
type jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) AutoScale() KubernetesNodeGroupScalePolicyAutoScaleOutputReference {
	var returns KubernetesNodeGroupScalePolicyAutoScaleOutputReference
	_jsii_.Get(
		j,
		"autoScale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) AutoScaleInput() *KubernetesNodeGroupScalePolicyAutoScale {
	var returns *KubernetesNodeGroupScalePolicyAutoScale
	_jsii_.Get(
		j,
		"autoScaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) FixedScale() KubernetesNodeGroupScalePolicyFixedScaleOutputReference {
	var returns KubernetesNodeGroupScalePolicyFixedScaleOutputReference
	_jsii_.Get(
		j,
		"fixedScale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) FixedScaleInput() *KubernetesNodeGroupScalePolicyFixedScale {
	var returns *KubernetesNodeGroupScalePolicyFixedScale
	_jsii_.Get(
		j,
		"fixedScaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) InternalValue() *KubernetesNodeGroupScalePolicy {
	var returns *KubernetesNodeGroupScalePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesNodeGroupScalePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesNodeGroupScalePolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.KubernetesNodeGroupScalePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesNodeGroupScalePolicyOutputReference_Override(k KubernetesNodeGroupScalePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.KubernetesNodeGroupScalePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) SetInternalValue(val *KubernetesNodeGroupScalePolicy) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) PutAutoScale(value *KubernetesNodeGroupScalePolicyAutoScale) {
	_jsii_.InvokeVoid(
		k,
		"putAutoScale",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) PutFixedScale(value *KubernetesNodeGroupScalePolicyFixedScale) {
	_jsii_.InvokeVoid(
		k,
		"putFixedScale",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) ResetAutoScale() {
	_jsii_.InvokeVoid(
		k,
		"resetAutoScale",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) ResetFixedScale() {
	_jsii_.InvokeVoid(
		k,
		"resetFixedScale",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupScalePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

