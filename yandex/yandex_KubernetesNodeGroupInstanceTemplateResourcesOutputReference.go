// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesNodeGroupInstanceTemplateResourcesOutputReference interface {
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
	CoreFraction() *float64
	SetCoreFraction(val *float64)
	CoreFractionInput() *float64
	Cores() *float64
	SetCores(val *float64)
	CoresInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Gpus() *float64
	SetGpus(val *float64)
	GpusInput() *float64
	InternalValue() *KubernetesNodeGroupInstanceTemplateResources
	SetInternalValue(val *KubernetesNodeGroupInstanceTemplateResources)
	Memory() *float64
	SetMemory(val *float64)
	MemoryInput() *float64
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
	ResetCoreFraction()
	ResetCores()
	ResetGpus()
	ResetMemory()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesNodeGroupInstanceTemplateResourcesOutputReference
type jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) CoreFraction() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coreFraction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) CoreFractionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coreFractionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) Cores() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) CoresInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) Gpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) GpusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gpusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) InternalValue() *KubernetesNodeGroupInstanceTemplateResources {
	var returns *KubernetesNodeGroupInstanceTemplateResources
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) Memory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) MemoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesNodeGroupInstanceTemplateResourcesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesNodeGroupInstanceTemplateResourcesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.KubernetesNodeGroupInstanceTemplateResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesNodeGroupInstanceTemplateResourcesOutputReference_Override(k KubernetesNodeGroupInstanceTemplateResourcesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.KubernetesNodeGroupInstanceTemplateResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) SetCoreFraction(val *float64) {
	_jsii_.Set(
		j,
		"coreFraction",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) SetCores(val *float64) {
	_jsii_.Set(
		j,
		"cores",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) SetGpus(val *float64) {
	_jsii_.Set(
		j,
		"gpus",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) SetInternalValue(val *KubernetesNodeGroupInstanceTemplateResources) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) SetMemory(val *float64) {
	_jsii_.Set(
		j,
		"memory",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) ResetCoreFraction() {
	_jsii_.InvokeVoid(
		k,
		"resetCoreFraction",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) ResetCores() {
	_jsii_.InvokeVoid(
		k,
		"resetCores",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) ResetGpus() {
	_jsii_.InvokeVoid(
		k,
		"resetGpus",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) ResetMemory() {
	_jsii_.InvokeVoid(
		k,
		"resetMemory",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateResourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

