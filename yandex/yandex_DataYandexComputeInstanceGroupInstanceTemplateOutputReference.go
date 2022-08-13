// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexComputeInstanceGroupInstanceTemplateOutputReference interface {
	cdktf.ComplexObject
	BootDisk() DataYandexComputeInstanceGroupInstanceTemplateBootDiskList
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
	// Experimental.
	Fqn() *string
	Hostname() *string
	InternalValue() *DataYandexComputeInstanceGroupInstanceTemplate
	SetInternalValue(val *DataYandexComputeInstanceGroupInstanceTemplate)
	Labels() cdktf.StringMap
	Metadata() cdktf.StringMap
	Name() *string
	NetworkInterface() DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceList
	NetworkSettings() DataYandexComputeInstanceGroupInstanceTemplateNetworkSettingsList
	PlacementPolicy() DataYandexComputeInstanceGroupInstanceTemplatePlacementPolicyList
	PlatformId() *string
	Resources() DataYandexComputeInstanceGroupInstanceTemplateResourcesList
	SchedulingPolicy() DataYandexComputeInstanceGroupInstanceTemplateSchedulingPolicyList
	SecondaryDisk() DataYandexComputeInstanceGroupInstanceTemplateSecondaryDiskList
	ServiceAccountId() *string
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

// The jsii proxy struct for DataYandexComputeInstanceGroupInstanceTemplateOutputReference
type jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) BootDisk() DataYandexComputeInstanceGroupInstanceTemplateBootDiskList {
	var returns DataYandexComputeInstanceGroupInstanceTemplateBootDiskList
	_jsii_.Get(
		j,
		"bootDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) InternalValue() *DataYandexComputeInstanceGroupInstanceTemplate {
	var returns *DataYandexComputeInstanceGroupInstanceTemplate
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) Labels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) Metadata() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) NetworkInterface() DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceList {
	var returns DataYandexComputeInstanceGroupInstanceTemplateNetworkInterfaceList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) NetworkSettings() DataYandexComputeInstanceGroupInstanceTemplateNetworkSettingsList {
	var returns DataYandexComputeInstanceGroupInstanceTemplateNetworkSettingsList
	_jsii_.Get(
		j,
		"networkSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) PlacementPolicy() DataYandexComputeInstanceGroupInstanceTemplatePlacementPolicyList {
	var returns DataYandexComputeInstanceGroupInstanceTemplatePlacementPolicyList
	_jsii_.Get(
		j,
		"placementPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) PlatformId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) Resources() DataYandexComputeInstanceGroupInstanceTemplateResourcesList {
	var returns DataYandexComputeInstanceGroupInstanceTemplateResourcesList
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) SchedulingPolicy() DataYandexComputeInstanceGroupInstanceTemplateSchedulingPolicyList {
	var returns DataYandexComputeInstanceGroupInstanceTemplateSchedulingPolicyList
	_jsii_.Get(
		j,
		"schedulingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) SecondaryDisk() DataYandexComputeInstanceGroupInstanceTemplateSecondaryDiskList {
	var returns DataYandexComputeInstanceGroupInstanceTemplateSecondaryDiskList
	_jsii_.Get(
		j,
		"secondaryDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) ServiceAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataYandexComputeInstanceGroupInstanceTemplateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataYandexComputeInstanceGroupInstanceTemplateOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexComputeInstanceGroupInstanceTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataYandexComputeInstanceGroupInstanceTemplateOutputReference_Override(d DataYandexComputeInstanceGroupInstanceTemplateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexComputeInstanceGroupInstanceTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) SetInternalValue(val *DataYandexComputeInstanceGroupInstanceTemplate) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexComputeInstanceGroupInstanceTemplateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

