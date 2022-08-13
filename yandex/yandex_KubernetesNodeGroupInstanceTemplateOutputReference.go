// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesNodeGroupInstanceTemplateOutputReference interface {
	cdktf.ComplexObject
	BootDisk() KubernetesNodeGroupInstanceTemplateBootDiskOutputReference
	BootDiskInput() *KubernetesNodeGroupInstanceTemplateBootDisk
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
	ContainerRuntime() KubernetesNodeGroupInstanceTemplateContainerRuntimeOutputReference
	ContainerRuntimeInput() *KubernetesNodeGroupInstanceTemplateContainerRuntime
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *KubernetesNodeGroupInstanceTemplate
	SetInternalValue(val *KubernetesNodeGroupInstanceTemplate)
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataInput() *map[string]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Nat() interface{}
	SetNat(val interface{})
	NatInput() interface{}
	NetworkAccelerationType() *string
	SetNetworkAccelerationType(val *string)
	NetworkAccelerationTypeInput() *string
	NetworkInterface() KubernetesNodeGroupInstanceTemplateNetworkInterfaceList
	NetworkInterfaceInput() interface{}
	PlacementPolicy() KubernetesNodeGroupInstanceTemplatePlacementPolicyOutputReference
	PlacementPolicyInput() *KubernetesNodeGroupInstanceTemplatePlacementPolicy
	PlatformId() *string
	SetPlatformId(val *string)
	PlatformIdInput() *string
	Resources() KubernetesNodeGroupInstanceTemplateResourcesOutputReference
	ResourcesInput() *KubernetesNodeGroupInstanceTemplateResources
	SchedulingPolicy() KubernetesNodeGroupInstanceTemplateSchedulingPolicyOutputReference
	SchedulingPolicyInput() *KubernetesNodeGroupInstanceTemplateSchedulingPolicy
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
	PutBootDisk(value *KubernetesNodeGroupInstanceTemplateBootDisk)
	PutContainerRuntime(value *KubernetesNodeGroupInstanceTemplateContainerRuntime)
	PutNetworkInterface(value interface{})
	PutPlacementPolicy(value *KubernetesNodeGroupInstanceTemplatePlacementPolicy)
	PutResources(value *KubernetesNodeGroupInstanceTemplateResources)
	PutSchedulingPolicy(value *KubernetesNodeGroupInstanceTemplateSchedulingPolicy)
	ResetBootDisk()
	ResetContainerRuntime()
	ResetLabels()
	ResetMetadata()
	ResetName()
	ResetNat()
	ResetNetworkAccelerationType()
	ResetNetworkInterface()
	ResetPlacementPolicy()
	ResetPlatformId()
	ResetResources()
	ResetSchedulingPolicy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesNodeGroupInstanceTemplateOutputReference
type jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) BootDisk() KubernetesNodeGroupInstanceTemplateBootDiskOutputReference {
	var returns KubernetesNodeGroupInstanceTemplateBootDiskOutputReference
	_jsii_.Get(
		j,
		"bootDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) BootDiskInput() *KubernetesNodeGroupInstanceTemplateBootDisk {
	var returns *KubernetesNodeGroupInstanceTemplateBootDisk
	_jsii_.Get(
		j,
		"bootDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) ContainerRuntime() KubernetesNodeGroupInstanceTemplateContainerRuntimeOutputReference {
	var returns KubernetesNodeGroupInstanceTemplateContainerRuntimeOutputReference
	_jsii_.Get(
		j,
		"containerRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) ContainerRuntimeInput() *KubernetesNodeGroupInstanceTemplateContainerRuntime {
	var returns *KubernetesNodeGroupInstanceTemplateContainerRuntime
	_jsii_.Get(
		j,
		"containerRuntimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) InternalValue() *KubernetesNodeGroupInstanceTemplate {
	var returns *KubernetesNodeGroupInstanceTemplate
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) Nat() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) NatInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"natInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) NetworkAccelerationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkAccelerationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) NetworkAccelerationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkAccelerationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) NetworkInterface() KubernetesNodeGroupInstanceTemplateNetworkInterfaceList {
	var returns KubernetesNodeGroupInstanceTemplateNetworkInterfaceList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) NetworkInterfaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) PlacementPolicy() KubernetesNodeGroupInstanceTemplatePlacementPolicyOutputReference {
	var returns KubernetesNodeGroupInstanceTemplatePlacementPolicyOutputReference
	_jsii_.Get(
		j,
		"placementPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) PlacementPolicyInput() *KubernetesNodeGroupInstanceTemplatePlacementPolicy {
	var returns *KubernetesNodeGroupInstanceTemplatePlacementPolicy
	_jsii_.Get(
		j,
		"placementPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) PlatformId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) PlatformIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) Resources() KubernetesNodeGroupInstanceTemplateResourcesOutputReference {
	var returns KubernetesNodeGroupInstanceTemplateResourcesOutputReference
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) ResourcesInput() *KubernetesNodeGroupInstanceTemplateResources {
	var returns *KubernetesNodeGroupInstanceTemplateResources
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) SchedulingPolicy() KubernetesNodeGroupInstanceTemplateSchedulingPolicyOutputReference {
	var returns KubernetesNodeGroupInstanceTemplateSchedulingPolicyOutputReference
	_jsii_.Get(
		j,
		"schedulingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) SchedulingPolicyInput() *KubernetesNodeGroupInstanceTemplateSchedulingPolicy {
	var returns *KubernetesNodeGroupInstanceTemplateSchedulingPolicy
	_jsii_.Get(
		j,
		"schedulingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesNodeGroupInstanceTemplateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesNodeGroupInstanceTemplateOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.KubernetesNodeGroupInstanceTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesNodeGroupInstanceTemplateOutputReference_Override(k KubernetesNodeGroupInstanceTemplateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.KubernetesNodeGroupInstanceTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) SetInternalValue(val *KubernetesNodeGroupInstanceTemplate) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) SetLabels(val *map[string]*string) {
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) SetMetadata(val *map[string]*string) {
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) SetNat(val interface{}) {
	_jsii_.Set(
		j,
		"nat",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) SetNetworkAccelerationType(val *string) {
	_jsii_.Set(
		j,
		"networkAccelerationType",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) SetPlatformId(val *string) {
	_jsii_.Set(
		j,
		"platformId",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) PutBootDisk(value *KubernetesNodeGroupInstanceTemplateBootDisk) {
	_jsii_.InvokeVoid(
		k,
		"putBootDisk",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) PutContainerRuntime(value *KubernetesNodeGroupInstanceTemplateContainerRuntime) {
	_jsii_.InvokeVoid(
		k,
		"putContainerRuntime",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) PutNetworkInterface(value interface{}) {
	_jsii_.InvokeVoid(
		k,
		"putNetworkInterface",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) PutPlacementPolicy(value *KubernetesNodeGroupInstanceTemplatePlacementPolicy) {
	_jsii_.InvokeVoid(
		k,
		"putPlacementPolicy",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) PutResources(value *KubernetesNodeGroupInstanceTemplateResources) {
	_jsii_.InvokeVoid(
		k,
		"putResources",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) PutSchedulingPolicy(value *KubernetesNodeGroupInstanceTemplateSchedulingPolicy) {
	_jsii_.InvokeVoid(
		k,
		"putSchedulingPolicy",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) ResetBootDisk() {
	_jsii_.InvokeVoid(
		k,
		"resetBootDisk",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) ResetContainerRuntime() {
	_jsii_.InvokeVoid(
		k,
		"resetContainerRuntime",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		k,
		"resetLabels",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) ResetMetadata() {
	_jsii_.InvokeVoid(
		k,
		"resetMetadata",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		k,
		"resetName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) ResetNat() {
	_jsii_.InvokeVoid(
		k,
		"resetNat",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) ResetNetworkAccelerationType() {
	_jsii_.InvokeVoid(
		k,
		"resetNetworkAccelerationType",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) ResetNetworkInterface() {
	_jsii_.InvokeVoid(
		k,
		"resetNetworkInterface",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) ResetPlacementPolicy() {
	_jsii_.InvokeVoid(
		k,
		"resetPlacementPolicy",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) ResetPlatformId() {
	_jsii_.InvokeVoid(
		k,
		"resetPlatformId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) ResetResources() {
	_jsii_.InvokeVoid(
		k,
		"resetResources",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) ResetSchedulingPolicy() {
	_jsii_.InvokeVoid(
		k,
		"resetSchedulingPolicy",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupInstanceTemplateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

