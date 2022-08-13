// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbKafkaConnectorConnectorConfigMirrormakerOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ReplicationFactor() *float64
	SetReplicationFactor(val *float64)
	ReplicationFactorInput() *float64
	SourceCluster() MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterOutputReference
	SourceClusterInput() *MdbKafkaConnectorConnectorConfigMirrormakerSourceCluster
	TargetCluster() MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference
	TargetClusterInput() *MdbKafkaConnectorConnectorConfigMirrormakerTargetCluster
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Topics() *string
	SetTopics(val *string)
	TopicsInput() *string
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
	PutSourceCluster(value *MdbKafkaConnectorConnectorConfigMirrormakerSourceCluster)
	PutTargetCluster(value *MdbKafkaConnectorConnectorConfigMirrormakerTargetCluster)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MdbKafkaConnectorConnectorConfigMirrormakerOutputReference
type jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) ReplicationFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicationFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) ReplicationFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicationFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) SourceCluster() MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterOutputReference {
	var returns MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterOutputReference
	_jsii_.Get(
		j,
		"sourceCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) SourceClusterInput() *MdbKafkaConnectorConnectorConfigMirrormakerSourceCluster {
	var returns *MdbKafkaConnectorConnectorConfigMirrormakerSourceCluster
	_jsii_.Get(
		j,
		"sourceClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) TargetCluster() MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference {
	var returns MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference
	_jsii_.Get(
		j,
		"targetCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) TargetClusterInput() *MdbKafkaConnectorConnectorConfigMirrormakerTargetCluster {
	var returns *MdbKafkaConnectorConnectorConfigMirrormakerTargetCluster
	_jsii_.Get(
		j,
		"targetClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) Topics() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) TopicsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicsInput",
		&returns,
	)
	return returns
}


func NewMdbKafkaConnectorConnectorConfigMirrormakerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MdbKafkaConnectorConnectorConfigMirrormakerOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbKafkaConnectorConnectorConfigMirrormakerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMdbKafkaConnectorConnectorConfigMirrormakerOutputReference_Override(m MdbKafkaConnectorConnectorConfigMirrormakerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbKafkaConnectorConnectorConfigMirrormakerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) SetReplicationFactor(val *float64) {
	_jsii_.Set(
		j,
		"replicationFactor",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) SetTopics(val *string) {
	_jsii_.Set(
		j,
		"topics",
		val,
	)
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) PutSourceCluster(value *MdbKafkaConnectorConnectorConfigMirrormakerSourceCluster) {
	_jsii_.InvokeVoid(
		m,
		"putSourceCluster",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) PutTargetCluster(value *MdbKafkaConnectorConnectorConfigMirrormakerTargetCluster) {
	_jsii_.InvokeVoid(
		m,
		"putTargetCluster",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

