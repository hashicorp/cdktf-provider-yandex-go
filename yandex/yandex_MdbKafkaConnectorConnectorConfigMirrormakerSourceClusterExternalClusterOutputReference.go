// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference interface {
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

// The jsii proxy struct for MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference
type jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) BootstrapServers() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) BootstrapServersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) SaslMechanism() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslMechanism",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) SaslMechanismInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslMechanismInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) SaslPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) SaslPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) SaslUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) SaslUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) SecurityProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) SecurityProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference_Override(m MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) SetBootstrapServers(val *string) {
	_jsii_.Set(
		j,
		"bootstrapServers",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) SetSaslMechanism(val *string) {
	_jsii_.Set(
		j,
		"saslMechanism",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) SetSaslPassword(val *string) {
	_jsii_.Set(
		j,
		"saslPassword",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) SetSaslUsername(val *string) {
	_jsii_.Set(
		j,
		"saslUsername",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) SetSecurityProtocol(val *string) {
	_jsii_.Set(
		j,
		"securityProtocol",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) ResetSaslMechanism() {
	_jsii_.InvokeVoid(
		m,
		"resetSaslMechanism",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) ResetSaslPassword() {
	_jsii_.InvokeVoid(
		m,
		"resetSaslPassword",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) ResetSaslUsername() {
	_jsii_.InvokeVoid(
		m,
		"resetSaslUsername",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) ResetSecurityProtocol() {
	_jsii_.InvokeVoid(
		m,
		"resetSecurityProtocol",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

