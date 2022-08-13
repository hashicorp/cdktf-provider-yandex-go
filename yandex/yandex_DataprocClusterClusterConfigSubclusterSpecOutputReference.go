// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocClusterClusterConfigSubclusterSpecOutputReference interface {
	cdktf.ComplexObject
	AssignPublicIp() interface{}
	SetAssignPublicIp(val interface{})
	AssignPublicIpInput() interface{}
	AutoscalingConfig() DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference
	AutoscalingConfigInput() *DataprocClusterClusterConfigSubclusterSpecAutoscalingConfig
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
	HostsCount() *float64
	SetHostsCount(val *float64)
	HostsCountInput() *float64
	Id() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Resources() DataprocClusterClusterConfigSubclusterSpecResourcesOutputReference
	ResourcesInput() *DataprocClusterClusterConfigSubclusterSpecResources
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
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
	PutAutoscalingConfig(value *DataprocClusterClusterConfigSubclusterSpecAutoscalingConfig)
	PutResources(value *DataprocClusterClusterConfigSubclusterSpecResources)
	ResetAssignPublicIp()
	ResetAutoscalingConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataprocClusterClusterConfigSubclusterSpecOutputReference
type jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) AssignPublicIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assignPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) AssignPublicIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assignPublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) AutoscalingConfig() DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference {
	var returns DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigOutputReference
	_jsii_.Get(
		j,
		"autoscalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) AutoscalingConfigInput() *DataprocClusterClusterConfigSubclusterSpecAutoscalingConfig {
	var returns *DataprocClusterClusterConfigSubclusterSpecAutoscalingConfig
	_jsii_.Get(
		j,
		"autoscalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) HostsCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hostsCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) HostsCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hostsCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) Resources() DataprocClusterClusterConfigSubclusterSpecResourcesOutputReference {
	var returns DataprocClusterClusterConfigSubclusterSpecResourcesOutputReference
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) ResourcesInput() *DataprocClusterClusterConfigSubclusterSpecResources {
	var returns *DataprocClusterClusterConfigSubclusterSpecResources
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataprocClusterClusterConfigSubclusterSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataprocClusterClusterConfigSubclusterSpecOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataprocClusterClusterConfigSubclusterSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataprocClusterClusterConfigSubclusterSpecOutputReference_Override(d DataprocClusterClusterConfigSubclusterSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataprocClusterClusterConfigSubclusterSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) SetAssignPublicIp(val interface{}) {
	_jsii_.Set(
		j,
		"assignPublicIp",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) SetHostsCount(val *float64) {
	_jsii_.Set(
		j,
		"hostsCount",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) SetRole(val *string) {
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) PutAutoscalingConfig(value *DataprocClusterClusterConfigSubclusterSpecAutoscalingConfig) {
	_jsii_.InvokeVoid(
		d,
		"putAutoscalingConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) PutResources(value *DataprocClusterClusterConfigSubclusterSpecResources) {
	_jsii_.InvokeVoid(
		d,
		"putResources",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) ResetAssignPublicIp() {
	_jsii_.InvokeVoid(
		d,
		"resetAssignPublicIp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) ResetAutoscalingConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoscalingConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigSubclusterSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

