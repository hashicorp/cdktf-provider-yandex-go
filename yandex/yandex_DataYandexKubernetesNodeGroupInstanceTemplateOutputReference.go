// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexKubernetesNodeGroupInstanceTemplateOutputReference interface {
	cdktf.ComplexObject
	BootDisk() DataYandexKubernetesNodeGroupInstanceTemplateBootDiskList
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
	ContainerRuntime() DataYandexKubernetesNodeGroupInstanceTemplateContainerRuntimeList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *DataYandexKubernetesNodeGroupInstanceTemplate
	SetInternalValue(val *DataYandexKubernetesNodeGroupInstanceTemplate)
	Labels() cdktf.StringMap
	Metadata() cdktf.StringMap
	Name() *string
	Nat() cdktf.IResolvable
	NetworkAccelerationType() *string
	NetworkInterface() DataYandexKubernetesNodeGroupInstanceTemplateNetworkInterfaceList
	PlacementPolicy() DataYandexKubernetesNodeGroupInstanceTemplatePlacementPolicyList
	PlatformId() *string
	Resources() DataYandexKubernetesNodeGroupInstanceTemplateResourcesList
	SchedulingPolicy() DataYandexKubernetesNodeGroupInstanceTemplateSchedulingPolicyList
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

// The jsii proxy struct for DataYandexKubernetesNodeGroupInstanceTemplateOutputReference
type jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) BootDisk() DataYandexKubernetesNodeGroupInstanceTemplateBootDiskList {
	var returns DataYandexKubernetesNodeGroupInstanceTemplateBootDiskList
	_jsii_.Get(
		j,
		"bootDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) ContainerRuntime() DataYandexKubernetesNodeGroupInstanceTemplateContainerRuntimeList {
	var returns DataYandexKubernetesNodeGroupInstanceTemplateContainerRuntimeList
	_jsii_.Get(
		j,
		"containerRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) InternalValue() *DataYandexKubernetesNodeGroupInstanceTemplate {
	var returns *DataYandexKubernetesNodeGroupInstanceTemplate
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) Labels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) Metadata() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) Nat() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"nat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) NetworkAccelerationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkAccelerationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) NetworkInterface() DataYandexKubernetesNodeGroupInstanceTemplateNetworkInterfaceList {
	var returns DataYandexKubernetesNodeGroupInstanceTemplateNetworkInterfaceList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) PlacementPolicy() DataYandexKubernetesNodeGroupInstanceTemplatePlacementPolicyList {
	var returns DataYandexKubernetesNodeGroupInstanceTemplatePlacementPolicyList
	_jsii_.Get(
		j,
		"placementPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) PlatformId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) Resources() DataYandexKubernetesNodeGroupInstanceTemplateResourcesList {
	var returns DataYandexKubernetesNodeGroupInstanceTemplateResourcesList
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) SchedulingPolicy() DataYandexKubernetesNodeGroupInstanceTemplateSchedulingPolicyList {
	var returns DataYandexKubernetesNodeGroupInstanceTemplateSchedulingPolicyList
	_jsii_.Get(
		j,
		"schedulingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataYandexKubernetesNodeGroupInstanceTemplateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataYandexKubernetesNodeGroupInstanceTemplateOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexKubernetesNodeGroupInstanceTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataYandexKubernetesNodeGroupInstanceTemplateOutputReference_Override(d DataYandexKubernetesNodeGroupInstanceTemplateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexKubernetesNodeGroupInstanceTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) SetInternalValue(val *DataYandexKubernetesNodeGroupInstanceTemplate) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexKubernetesNodeGroupInstanceTemplateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

