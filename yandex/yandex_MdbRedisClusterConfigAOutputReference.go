// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbRedisClusterConfigAOutputReference interface {
	cdktf.ComplexObject
	ClientOutputBufferLimitNormal() *string
	SetClientOutputBufferLimitNormal(val *string)
	ClientOutputBufferLimitNormalInput() *string
	ClientOutputBufferLimitPubsub() *string
	SetClientOutputBufferLimitPubsub(val *string)
	ClientOutputBufferLimitPubsubInput() *string
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
	Databases() *float64
	SetDatabases(val *float64)
	DatabasesInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *MdbRedisClusterConfigA
	SetInternalValue(val *MdbRedisClusterConfigA)
	MaxmemoryPolicy() *string
	SetMaxmemoryPolicy(val *string)
	MaxmemoryPolicyInput() *string
	NotifyKeyspaceEvents() *string
	SetNotifyKeyspaceEvents(val *string)
	NotifyKeyspaceEventsInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	SlowlogLogSlowerThan() *float64
	SetSlowlogLogSlowerThan(val *float64)
	SlowlogLogSlowerThanInput() *float64
	SlowlogMaxLen() *float64
	SetSlowlogMaxLen(val *float64)
	SlowlogMaxLenInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() *float64
	SetTimeout(val *float64)
	TimeoutInput() *float64
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	ResetClientOutputBufferLimitNormal()
	ResetClientOutputBufferLimitPubsub()
	ResetDatabases()
	ResetMaxmemoryPolicy()
	ResetNotifyKeyspaceEvents()
	ResetSlowlogLogSlowerThan()
	ResetSlowlogMaxLen()
	ResetTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MdbRedisClusterConfigAOutputReference
type jsiiProxy_MdbRedisClusterConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) ClientOutputBufferLimitNormal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientOutputBufferLimitNormal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) ClientOutputBufferLimitNormalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientOutputBufferLimitNormalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) ClientOutputBufferLimitPubsub() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientOutputBufferLimitPubsub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) ClientOutputBufferLimitPubsubInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientOutputBufferLimitPubsubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) Databases() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"databases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) DatabasesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"databasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) InternalValue() *MdbRedisClusterConfigA {
	var returns *MdbRedisClusterConfigA
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) MaxmemoryPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxmemoryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) MaxmemoryPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxmemoryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) NotifyKeyspaceEvents() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notifyKeyspaceEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) NotifyKeyspaceEventsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notifyKeyspaceEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) SlowlogLogSlowerThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"slowlogLogSlowerThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) SlowlogLogSlowerThanInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"slowlogLogSlowerThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) SlowlogMaxLen() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"slowlogMaxLen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) SlowlogMaxLenInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"slowlogMaxLenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewMdbRedisClusterConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MdbRedisClusterConfigAOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MdbRedisClusterConfigAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbRedisClusterConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMdbRedisClusterConfigAOutputReference_Override(m MdbRedisClusterConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbRedisClusterConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) SetClientOutputBufferLimitNormal(val *string) {
	_jsii_.Set(
		j,
		"clientOutputBufferLimitNormal",
		val,
	)
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) SetClientOutputBufferLimitPubsub(val *string) {
	_jsii_.Set(
		j,
		"clientOutputBufferLimitPubsub",
		val,
	)
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) SetDatabases(val *float64) {
	_jsii_.Set(
		j,
		"databases",
		val,
	)
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) SetInternalValue(val *MdbRedisClusterConfigA) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) SetMaxmemoryPolicy(val *string) {
	_jsii_.Set(
		j,
		"maxmemoryPolicy",
		val,
	)
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) SetNotifyKeyspaceEvents(val *string) {
	_jsii_.Set(
		j,
		"notifyKeyspaceEvents",
		val,
	)
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) SetSlowlogLogSlowerThan(val *float64) {
	_jsii_.Set(
		j,
		"slowlogLogSlowerThan",
		val,
	)
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) SetSlowlogMaxLen(val *float64) {
	_jsii_.Set(
		j,
		"slowlogMaxLen",
		val,
	)
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) SetTimeout(val *float64) {
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_MdbRedisClusterConfigAOutputReference) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (m *jsiiProxy_MdbRedisClusterConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbRedisClusterConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbRedisClusterConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbRedisClusterConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbRedisClusterConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbRedisClusterConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbRedisClusterConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbRedisClusterConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbRedisClusterConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbRedisClusterConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbRedisClusterConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbRedisClusterConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbRedisClusterConfigAOutputReference) ResetClientOutputBufferLimitNormal() {
	_jsii_.InvokeVoid(
		m,
		"resetClientOutputBufferLimitNormal",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbRedisClusterConfigAOutputReference) ResetClientOutputBufferLimitPubsub() {
	_jsii_.InvokeVoid(
		m,
		"resetClientOutputBufferLimitPubsub",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbRedisClusterConfigAOutputReference) ResetDatabases() {
	_jsii_.InvokeVoid(
		m,
		"resetDatabases",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbRedisClusterConfigAOutputReference) ResetMaxmemoryPolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxmemoryPolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbRedisClusterConfigAOutputReference) ResetNotifyKeyspaceEvents() {
	_jsii_.InvokeVoid(
		m,
		"resetNotifyKeyspaceEvents",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbRedisClusterConfigAOutputReference) ResetSlowlogLogSlowerThan() {
	_jsii_.InvokeVoid(
		m,
		"resetSlowlogLogSlowerThan",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbRedisClusterConfigAOutputReference) ResetSlowlogMaxLen() {
	_jsii_.InvokeVoid(
		m,
		"resetSlowlogMaxLen",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbRedisClusterConfigAOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeout",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbRedisClusterConfigAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbRedisClusterConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

