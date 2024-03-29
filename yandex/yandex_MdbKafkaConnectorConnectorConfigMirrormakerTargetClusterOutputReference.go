// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference interface {
	cdktf.ComplexObject
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
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
	ExternalCluster() MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterList
	ExternalClusterInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *MdbKafkaConnectorConnectorConfigMirrormakerTargetCluster
	SetInternalValue(val *MdbKafkaConnectorConnectorConfigMirrormakerTargetCluster)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ThisCluster() MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterThisClusterList
	ThisClusterInput() interface{}
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
	PutExternalCluster(value interface{})
	PutThisCluster(value interface{})
	ResetAlias()
	ResetExternalCluster()
	ResetThisCluster()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference
type jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) ExternalCluster() MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterList {
	var returns MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterList
	_jsii_.Get(
		j,
		"externalCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) ExternalClusterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) InternalValue() *MdbKafkaConnectorConnectorConfigMirrormakerTargetCluster {
	var returns *MdbKafkaConnectorConnectorConfigMirrormakerTargetCluster
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) ThisCluster() MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterThisClusterList {
	var returns MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterThisClusterList
	_jsii_.Get(
		j,
		"thisCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) ThisClusterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"thisClusterInput",
		&returns,
	)
	return returns
}


func NewMdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference_Override(m MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) SetInternalValue(val *MdbKafkaConnectorConnectorConfigMirrormakerTargetCluster) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) PutExternalCluster(value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"putExternalCluster",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) PutThisCluster(value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"putThisCluster",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) ResetAlias() {
	_jsii_.InvokeVoid(
		m,
		"resetAlias",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) ResetExternalCluster() {
	_jsii_.InvokeVoid(
		m,
		"resetExternalCluster",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) ResetThisCluster() {
	_jsii_.InvokeVoid(
		m,
		"resetThisCluster",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

