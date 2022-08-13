// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeInstanceGroupInstanceTemplateOutputReference interface {
	cdktf.ComplexObject
	BootDisk() ComputeInstanceGroupInstanceTemplateBootDiskOutputReference
	BootDiskInput() *ComputeInstanceGroupInstanceTemplateBootDisk
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	Hostname() *string
	SetHostname(val *string)
	HostnameInput() *string
	InternalValue() *ComputeInstanceGroupInstanceTemplate
	SetInternalValue(val *ComputeInstanceGroupInstanceTemplate)
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataInput() *map[string]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkInterface() ComputeInstanceGroupInstanceTemplateNetworkInterfaceList
	NetworkInterfaceInput() interface{}
	NetworkSettings() ComputeInstanceGroupInstanceTemplateNetworkSettingsList
	NetworkSettingsInput() interface{}
	PlacementPolicy() ComputeInstanceGroupInstanceTemplatePlacementPolicyOutputReference
	PlacementPolicyInput() *ComputeInstanceGroupInstanceTemplatePlacementPolicy
	PlatformId() *string
	SetPlatformId(val *string)
	PlatformIdInput() *string
	Resources() ComputeInstanceGroupInstanceTemplateResourcesOutputReference
	ResourcesInput() *ComputeInstanceGroupInstanceTemplateResources
	SchedulingPolicy() ComputeInstanceGroupInstanceTemplateSchedulingPolicyOutputReference
	SchedulingPolicyInput() *ComputeInstanceGroupInstanceTemplateSchedulingPolicy
	SecondaryDisk() ComputeInstanceGroupInstanceTemplateSecondaryDiskList
	SecondaryDiskInput() interface{}
	ServiceAccountId() *string
	SetServiceAccountId(val *string)
	ServiceAccountIdInput() *string
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
	PutBootDisk(value *ComputeInstanceGroupInstanceTemplateBootDisk)
	PutNetworkInterface(value interface{})
	PutNetworkSettings(value interface{})
	PutPlacementPolicy(value *ComputeInstanceGroupInstanceTemplatePlacementPolicy)
	PutResources(value *ComputeInstanceGroupInstanceTemplateResources)
	PutSchedulingPolicy(value *ComputeInstanceGroupInstanceTemplateSchedulingPolicy)
	PutSecondaryDisk(value interface{})
	ResetDescription()
	ResetHostname()
	ResetLabels()
	ResetMetadata()
	ResetName()
	ResetNetworkSettings()
	ResetPlacementPolicy()
	ResetPlatformId()
	ResetSchedulingPolicy()
	ResetSecondaryDisk()
	ResetServiceAccountId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeInstanceGroupInstanceTemplateOutputReference
type jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) BootDisk() ComputeInstanceGroupInstanceTemplateBootDiskOutputReference {
	var returns ComputeInstanceGroupInstanceTemplateBootDiskOutputReference
	_jsii_.Get(
		j,
		"bootDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) BootDiskInput() *ComputeInstanceGroupInstanceTemplateBootDisk {
	var returns *ComputeInstanceGroupInstanceTemplateBootDisk
	_jsii_.Get(
		j,
		"bootDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) InternalValue() *ComputeInstanceGroupInstanceTemplate {
	var returns *ComputeInstanceGroupInstanceTemplate
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) NetworkInterface() ComputeInstanceGroupInstanceTemplateNetworkInterfaceList {
	var returns ComputeInstanceGroupInstanceTemplateNetworkInterfaceList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) NetworkInterfaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) NetworkSettings() ComputeInstanceGroupInstanceTemplateNetworkSettingsList {
	var returns ComputeInstanceGroupInstanceTemplateNetworkSettingsList
	_jsii_.Get(
		j,
		"networkSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) NetworkSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) PlacementPolicy() ComputeInstanceGroupInstanceTemplatePlacementPolicyOutputReference {
	var returns ComputeInstanceGroupInstanceTemplatePlacementPolicyOutputReference
	_jsii_.Get(
		j,
		"placementPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) PlacementPolicyInput() *ComputeInstanceGroupInstanceTemplatePlacementPolicy {
	var returns *ComputeInstanceGroupInstanceTemplatePlacementPolicy
	_jsii_.Get(
		j,
		"placementPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) PlatformId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) PlatformIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) Resources() ComputeInstanceGroupInstanceTemplateResourcesOutputReference {
	var returns ComputeInstanceGroupInstanceTemplateResourcesOutputReference
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) ResourcesInput() *ComputeInstanceGroupInstanceTemplateResources {
	var returns *ComputeInstanceGroupInstanceTemplateResources
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) SchedulingPolicy() ComputeInstanceGroupInstanceTemplateSchedulingPolicyOutputReference {
	var returns ComputeInstanceGroupInstanceTemplateSchedulingPolicyOutputReference
	_jsii_.Get(
		j,
		"schedulingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) SchedulingPolicyInput() *ComputeInstanceGroupInstanceTemplateSchedulingPolicy {
	var returns *ComputeInstanceGroupInstanceTemplateSchedulingPolicy
	_jsii_.Get(
		j,
		"schedulingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) SecondaryDisk() ComputeInstanceGroupInstanceTemplateSecondaryDiskList {
	var returns ComputeInstanceGroupInstanceTemplateSecondaryDiskList
	_jsii_.Get(
		j,
		"secondaryDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) SecondaryDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) ServiceAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) ServiceAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeInstanceGroupInstanceTemplateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeInstanceGroupInstanceTemplateOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.ComputeInstanceGroupInstanceTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeInstanceGroupInstanceTemplateOutputReference_Override(c ComputeInstanceGroupInstanceTemplateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.ComputeInstanceGroupInstanceTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) SetHostname(val *string) {
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) SetInternalValue(val *ComputeInstanceGroupInstanceTemplate) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) SetLabels(val *map[string]*string) {
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) SetMetadata(val *map[string]*string) {
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) SetPlatformId(val *string) {
	_jsii_.Set(
		j,
		"platformId",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) SetServiceAccountId(val *string) {
	_jsii_.Set(
		j,
		"serviceAccountId",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) PutBootDisk(value *ComputeInstanceGroupInstanceTemplateBootDisk) {
	_jsii_.InvokeVoid(
		c,
		"putBootDisk",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) PutNetworkInterface(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putNetworkInterface",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) PutNetworkSettings(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putNetworkSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) PutPlacementPolicy(value *ComputeInstanceGroupInstanceTemplatePlacementPolicy) {
	_jsii_.InvokeVoid(
		c,
		"putPlacementPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) PutResources(value *ComputeInstanceGroupInstanceTemplateResources) {
	_jsii_.InvokeVoid(
		c,
		"putResources",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) PutSchedulingPolicy(value *ComputeInstanceGroupInstanceTemplateSchedulingPolicy) {
	_jsii_.InvokeVoid(
		c,
		"putSchedulingPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) PutSecondaryDisk(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putSecondaryDisk",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) ResetHostname() {
	_jsii_.InvokeVoid(
		c,
		"resetHostname",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		c,
		"resetLabels",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) ResetMetadata() {
	_jsii_.InvokeVoid(
		c,
		"resetMetadata",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		c,
		"resetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) ResetNetworkSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetNetworkSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) ResetPlacementPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetPlacementPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) ResetPlatformId() {
	_jsii_.InvokeVoid(
		c,
		"resetPlatformId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) ResetSchedulingPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetSchedulingPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) ResetSecondaryDisk() {
	_jsii_.InvokeVoid(
		c,
		"resetSecondaryDisk",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) ResetServiceAccountId() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceAccountId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

