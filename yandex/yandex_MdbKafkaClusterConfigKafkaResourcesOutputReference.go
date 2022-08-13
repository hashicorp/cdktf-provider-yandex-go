// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbKafkaClusterConfigKafkaResourcesOutputReference interface {
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
	DiskSize() *float64
	SetDiskSize(val *float64)
	DiskSizeInput() *float64
	DiskTypeId() *string
	SetDiskTypeId(val *string)
	DiskTypeIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MdbKafkaClusterConfigKafkaResources
	SetInternalValue(val *MdbKafkaClusterConfigKafkaResources)
	ResourcePresetId() *string
	SetResourcePresetId(val *string)
	ResourcePresetIdInput() *string
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

// The jsii proxy struct for MdbKafkaClusterConfigKafkaResourcesOutputReference
type jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) DiskSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) DiskSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) DiskTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) DiskTypeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) InternalValue() *MdbKafkaClusterConfigKafkaResources {
	var returns *MdbKafkaClusterConfigKafkaResources
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) ResourcePresetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourcePresetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) ResourcePresetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourcePresetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMdbKafkaClusterConfigKafkaResourcesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MdbKafkaClusterConfigKafkaResourcesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbKafkaClusterConfigKafkaResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMdbKafkaClusterConfigKafkaResourcesOutputReference_Override(m MdbKafkaClusterConfigKafkaResourcesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbKafkaClusterConfigKafkaResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) SetDiskSize(val *float64) {
	_jsii_.Set(
		j,
		"diskSize",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) SetDiskTypeId(val *string) {
	_jsii_.Set(
		j,
		"diskTypeId",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) SetInternalValue(val *MdbKafkaClusterConfigKafkaResources) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) SetResourcePresetId(val *string) {
	_jsii_.Set(
		j,
		"resourcePresetId",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigKafkaResourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

