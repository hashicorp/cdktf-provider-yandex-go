// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference interface {
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
	InternalValue() *DataYandexMdbKafkaClusterConfigKafkaResources
	SetInternalValue(val *DataYandexMdbKafkaClusterConfigKafkaResources)
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

// The jsii proxy struct for DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference
type jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) DiskSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) DiskSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) DiskTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) DiskTypeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) InternalValue() *DataYandexMdbKafkaClusterConfigKafkaResources {
	var returns *DataYandexMdbKafkaClusterConfigKafkaResources
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) ResourcePresetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourcePresetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) ResourcePresetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourcePresetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference_Override(d DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) SetDiskSize(val *float64) {
	_jsii_.Set(
		j,
		"diskSize",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) SetDiskTypeId(val *string) {
	_jsii_.Set(
		j,
		"diskTypeId",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) SetInternalValue(val *DataYandexMdbKafkaClusterConfigKafkaResources) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) SetResourcePresetId(val *string) {
	_jsii_.Set(
		j,
		"resourcePresetId",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigKafkaResourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

