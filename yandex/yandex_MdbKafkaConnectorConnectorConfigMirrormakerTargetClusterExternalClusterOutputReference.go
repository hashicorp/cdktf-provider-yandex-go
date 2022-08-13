// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference interface {
	cdktf.ComplexObject
	BootstrapServers() *string
	SetBootstrapServers(val *string)
	BootstrapServersInput() *string
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
	SaslMechanism() *string
	SetSaslMechanism(val *string)
	SaslMechanismInput() *string
	SaslPassword() *string
	SetSaslPassword(val *string)
	SaslPasswordInput() *string
	SaslUsername() *string
	SetSaslUsername(val *string)
	SaslUsernameInput() *string
	SecurityProtocol() *string
	SetSecurityProtocol(val *string)
	SecurityProtocolInput() *string
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
	ResetSaslMechanism()
	ResetSaslPassword()
	ResetSaslUsername()
	ResetSecurityProtocol()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference
type jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) BootstrapServers() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) BootstrapServersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) SaslMechanism() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslMechanism",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) SaslMechanismInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslMechanismInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) SaslPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) SaslPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) SaslUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) SaslUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) SecurityProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) SecurityProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference_Override(m MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) SetBootstrapServers(val *string) {
	_jsii_.Set(
		j,
		"bootstrapServers",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) SetSaslMechanism(val *string) {
	_jsii_.Set(
		j,
		"saslMechanism",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) SetSaslPassword(val *string) {
	_jsii_.Set(
		j,
		"saslPassword",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) SetSaslUsername(val *string) {
	_jsii_.Set(
		j,
		"saslUsername",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) SetSecurityProtocol(val *string) {
	_jsii_.Set(
		j,
		"securityProtocol",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) ResetSaslMechanism() {
	_jsii_.InvokeVoid(
		m,
		"resetSaslMechanism",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) ResetSaslPassword() {
	_jsii_.InvokeVoid(
		m,
		"resetSaslPassword",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) ResetSaslUsername() {
	_jsii_.InvokeVoid(
		m,
		"resetSaslUsername",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) ResetSecurityProtocol() {
	_jsii_.InvokeVoid(
		m,
		"resetSecurityProtocol",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalClusterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

