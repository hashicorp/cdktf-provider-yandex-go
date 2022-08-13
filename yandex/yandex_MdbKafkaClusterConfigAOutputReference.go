// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbKafkaClusterConfigAOutputReference interface {
	cdktf.ComplexObject
	AssignPublicIp() interface{}
	SetAssignPublicIp(val interface{})
	AssignPublicIpInput() interface{}
	BrokersCount() *float64
	SetBrokersCount(val *float64)
	BrokersCountInput() *float64
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
	InternalValue() *MdbKafkaClusterConfigA
	SetInternalValue(val *MdbKafkaClusterConfigA)
	Kafka() MdbKafkaClusterConfigKafkaOutputReference
	KafkaInput() *MdbKafkaClusterConfigKafka
	SchemaRegistry() interface{}
	SetSchemaRegistry(val interface{})
	SchemaRegistryInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UnmanagedTopics() interface{}
	SetUnmanagedTopics(val interface{})
	UnmanagedTopicsInput() interface{}
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
	Zones() *[]*string
	SetZones(val *[]*string)
	ZonesInput() *[]*string
	Zookeeper() MdbKafkaClusterConfigZookeeperOutputReference
	ZookeeperInput() *MdbKafkaClusterConfigZookeeper
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
	PutKafka(value *MdbKafkaClusterConfigKafka)
	PutZookeeper(value *MdbKafkaClusterConfigZookeeper)
	ResetAssignPublicIp()
	ResetBrokersCount()
	ResetSchemaRegistry()
	ResetUnmanagedTopics()
	ResetZookeeper()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MdbKafkaClusterConfigAOutputReference
type jsiiProxy_MdbKafkaClusterConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) AssignPublicIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assignPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) AssignPublicIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assignPublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) BrokersCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"brokersCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) BrokersCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"brokersCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) InternalValue() *MdbKafkaClusterConfigA {
	var returns *MdbKafkaClusterConfigA
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) Kafka() MdbKafkaClusterConfigKafkaOutputReference {
	var returns MdbKafkaClusterConfigKafkaOutputReference
	_jsii_.Get(
		j,
		"kafka",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) KafkaInput() *MdbKafkaClusterConfigKafka {
	var returns *MdbKafkaClusterConfigKafka
	_jsii_.Get(
		j,
		"kafkaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) SchemaRegistry() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaRegistry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) SchemaRegistryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaRegistryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) UnmanagedTopics() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unmanagedTopics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) UnmanagedTopicsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unmanagedTopicsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) Zones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) ZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) Zookeeper() MdbKafkaClusterConfigZookeeperOutputReference {
	var returns MdbKafkaClusterConfigZookeeperOutputReference
	_jsii_.Get(
		j,
		"zookeeper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) ZookeeperInput() *MdbKafkaClusterConfigZookeeper {
	var returns *MdbKafkaClusterConfigZookeeper
	_jsii_.Get(
		j,
		"zookeeperInput",
		&returns,
	)
	return returns
}


func NewMdbKafkaClusterConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MdbKafkaClusterConfigAOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MdbKafkaClusterConfigAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbKafkaClusterConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMdbKafkaClusterConfigAOutputReference_Override(m MdbKafkaClusterConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbKafkaClusterConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) SetAssignPublicIp(val interface{}) {
	_jsii_.Set(
		j,
		"assignPublicIp",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) SetBrokersCount(val *float64) {
	_jsii_.Set(
		j,
		"brokersCount",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) SetInternalValue(val *MdbKafkaClusterConfigA) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) SetSchemaRegistry(val interface{}) {
	_jsii_.Set(
		j,
		"schemaRegistry",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) SetUnmanagedTopics(val interface{}) {
	_jsii_.Set(
		j,
		"unmanagedTopics",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaClusterConfigAOutputReference) SetZones(val *[]*string) {
	_jsii_.Set(
		j,
		"zones",
		val,
	)
}

func (m *jsiiProxy_MdbKafkaClusterConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigAOutputReference) PutKafka(value *MdbKafkaClusterConfigKafka) {
	_jsii_.InvokeVoid(
		m,
		"putKafka",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbKafkaClusterConfigAOutputReference) PutZookeeper(value *MdbKafkaClusterConfigZookeeper) {
	_jsii_.InvokeVoid(
		m,
		"putZookeeper",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbKafkaClusterConfigAOutputReference) ResetAssignPublicIp() {
	_jsii_.InvokeVoid(
		m,
		"resetAssignPublicIp",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterConfigAOutputReference) ResetBrokersCount() {
	_jsii_.InvokeVoid(
		m,
		"resetBrokersCount",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterConfigAOutputReference) ResetSchemaRegistry() {
	_jsii_.InvokeVoid(
		m,
		"resetSchemaRegistry",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterConfigAOutputReference) ResetUnmanagedTopics() {
	_jsii_.InvokeVoid(
		m,
		"resetUnmanagedTopics",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterConfigAOutputReference) ResetZookeeper() {
	_jsii_.InvokeVoid(
		m,
		"resetZookeeper",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaClusterConfigAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaClusterConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

