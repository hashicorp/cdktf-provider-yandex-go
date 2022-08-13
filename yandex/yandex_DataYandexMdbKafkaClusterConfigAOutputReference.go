// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexMdbKafkaClusterConfigAOutputReference interface {
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
	InternalValue() *DataYandexMdbKafkaClusterConfigA
	SetInternalValue(val *DataYandexMdbKafkaClusterConfigA)
	Kafka() DataYandexMdbKafkaClusterConfigKafkaOutputReference
	KafkaInput() *DataYandexMdbKafkaClusterConfigKafka
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
	Zookeeper() DataYandexMdbKafkaClusterConfigZookeeperOutputReference
	ZookeeperInput() *DataYandexMdbKafkaClusterConfigZookeeper
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
	PutKafka(value *DataYandexMdbKafkaClusterConfigKafka)
	PutZookeeper(value *DataYandexMdbKafkaClusterConfigZookeeper)
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

// The jsii proxy struct for DataYandexMdbKafkaClusterConfigAOutputReference
type jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) AssignPublicIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assignPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) AssignPublicIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assignPublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) BrokersCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"brokersCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) BrokersCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"brokersCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) InternalValue() *DataYandexMdbKafkaClusterConfigA {
	var returns *DataYandexMdbKafkaClusterConfigA
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) Kafka() DataYandexMdbKafkaClusterConfigKafkaOutputReference {
	var returns DataYandexMdbKafkaClusterConfigKafkaOutputReference
	_jsii_.Get(
		j,
		"kafka",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) KafkaInput() *DataYandexMdbKafkaClusterConfigKafka {
	var returns *DataYandexMdbKafkaClusterConfigKafka
	_jsii_.Get(
		j,
		"kafkaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) SchemaRegistry() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaRegistry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) SchemaRegistryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaRegistryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) UnmanagedTopics() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unmanagedTopics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) UnmanagedTopicsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unmanagedTopicsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) Zones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) ZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) Zookeeper() DataYandexMdbKafkaClusterConfigZookeeperOutputReference {
	var returns DataYandexMdbKafkaClusterConfigZookeeperOutputReference
	_jsii_.Get(
		j,
		"zookeeper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) ZookeeperInput() *DataYandexMdbKafkaClusterConfigZookeeper {
	var returns *DataYandexMdbKafkaClusterConfigZookeeper
	_jsii_.Get(
		j,
		"zookeeperInput",
		&returns,
	)
	return returns
}


func NewDataYandexMdbKafkaClusterConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataYandexMdbKafkaClusterConfigAOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexMdbKafkaClusterConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataYandexMdbKafkaClusterConfigAOutputReference_Override(d DataYandexMdbKafkaClusterConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexMdbKafkaClusterConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) SetAssignPublicIp(val interface{}) {
	_jsii_.Set(
		j,
		"assignPublicIp",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) SetBrokersCount(val *float64) {
	_jsii_.Set(
		j,
		"brokersCount",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) SetInternalValue(val *DataYandexMdbKafkaClusterConfigA) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) SetSchemaRegistry(val interface{}) {
	_jsii_.Set(
		j,
		"schemaRegistry",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) SetUnmanagedTopics(val interface{}) {
	_jsii_.Set(
		j,
		"unmanagedTopics",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) SetZones(val *[]*string) {
	_jsii_.Set(
		j,
		"zones",
		val,
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) PutKafka(value *DataYandexMdbKafkaClusterConfigKafka) {
	_jsii_.InvokeVoid(
		d,
		"putKafka",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) PutZookeeper(value *DataYandexMdbKafkaClusterConfigZookeeper) {
	_jsii_.InvokeVoid(
		d,
		"putZookeeper",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) ResetAssignPublicIp() {
	_jsii_.InvokeVoid(
		d,
		"resetAssignPublicIp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) ResetBrokersCount() {
	_jsii_.InvokeVoid(
		d,
		"resetBrokersCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) ResetSchemaRegistry() {
	_jsii_.InvokeVoid(
		d,
		"resetSchemaRegistry",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) ResetUnmanagedTopics() {
	_jsii_.InvokeVoid(
		d,
		"resetUnmanagedTopics",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) ResetZookeeper() {
	_jsii_.InvokeVoid(
		d,
		"resetZookeeper",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaClusterConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

