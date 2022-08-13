// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbClickhouseClusterClickhouseConfigKafkaOutputReference interface {
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
	InternalValue() *MdbClickhouseClusterClickhouseConfigKafka
	SetInternalValue(val *MdbClickhouseClusterClickhouseConfigKafka)
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

// The jsii proxy struct for MdbClickhouseClusterClickhouseConfigKafkaOutputReference
type jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) InternalValue() *MdbClickhouseClusterClickhouseConfigKafka {
	var returns *MdbClickhouseClusterClickhouseConfigKafka
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) SaslMechanism() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslMechanism",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) SaslMechanismInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslMechanismInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) SaslPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) SaslPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) SaslUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) SaslUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) SecurityProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) SecurityProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMdbClickhouseClusterClickhouseConfigKafkaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MdbClickhouseClusterClickhouseConfigKafkaOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbClickhouseClusterClickhouseConfigKafkaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMdbClickhouseClusterClickhouseConfigKafkaOutputReference_Override(m MdbClickhouseClusterClickhouseConfigKafkaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbClickhouseClusterClickhouseConfigKafkaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) SetInternalValue(val *MdbClickhouseClusterClickhouseConfigKafka) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) SetSaslMechanism(val *string) {
	_jsii_.Set(
		j,
		"saslMechanism",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) SetSaslPassword(val *string) {
	_jsii_.Set(
		j,
		"saslPassword",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) SetSaslUsername(val *string) {
	_jsii_.Set(
		j,
		"saslUsername",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) SetSecurityProtocol(val *string) {
	_jsii_.Set(
		j,
		"securityProtocol",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) ResetSaslMechanism() {
	_jsii_.InvokeVoid(
		m,
		"resetSaslMechanism",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) ResetSaslPassword() {
	_jsii_.InvokeVoid(
		m,
		"resetSaslPassword",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) ResetSaslUsername() {
	_jsii_.InvokeVoid(
		m,
		"resetSaslUsername",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) ResetSecurityProtocol() {
	_jsii_.InvokeVoid(
		m,
		"resetSecurityProtocol",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigKafkaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

