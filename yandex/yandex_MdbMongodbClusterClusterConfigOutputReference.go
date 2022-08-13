// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbMongodbClusterClusterConfigOutputReference interface {
	cdktf.ComplexObject
	Access() MdbMongodbClusterClusterConfigAccessOutputReference
	AccessInput() *MdbMongodbClusterClusterConfigAccess
	BackupWindowStart() MdbMongodbClusterClusterConfigBackupWindowStartOutputReference
	BackupWindowStartInput() *MdbMongodbClusterClusterConfigBackupWindowStart
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
	FeatureCompatibilityVersion() *string
	SetFeatureCompatibilityVersion(val *string)
	FeatureCompatibilityVersionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MdbMongodbClusterClusterConfig
	SetInternalValue(val *MdbMongodbClusterClusterConfig)
	Mongod() MdbMongodbClusterClusterConfigMongodOutputReference
	MongodInput() *MdbMongodbClusterClusterConfigMongod
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutAccess(value *MdbMongodbClusterClusterConfigAccess)
	PutBackupWindowStart(value *MdbMongodbClusterClusterConfigBackupWindowStart)
	PutMongod(value *MdbMongodbClusterClusterConfigMongod)
	ResetAccess()
	ResetBackupWindowStart()
	ResetFeatureCompatibilityVersion()
	ResetMongod()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MdbMongodbClusterClusterConfigOutputReference
type jsiiProxy_MdbMongodbClusterClusterConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) Access() MdbMongodbClusterClusterConfigAccessOutputReference {
	var returns MdbMongodbClusterClusterConfigAccessOutputReference
	_jsii_.Get(
		j,
		"access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) AccessInput() *MdbMongodbClusterClusterConfigAccess {
	var returns *MdbMongodbClusterClusterConfigAccess
	_jsii_.Get(
		j,
		"accessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) BackupWindowStart() MdbMongodbClusterClusterConfigBackupWindowStartOutputReference {
	var returns MdbMongodbClusterClusterConfigBackupWindowStartOutputReference
	_jsii_.Get(
		j,
		"backupWindowStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) BackupWindowStartInput() *MdbMongodbClusterClusterConfigBackupWindowStart {
	var returns *MdbMongodbClusterClusterConfigBackupWindowStart
	_jsii_.Get(
		j,
		"backupWindowStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) FeatureCompatibilityVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"featureCompatibilityVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) FeatureCompatibilityVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"featureCompatibilityVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) InternalValue() *MdbMongodbClusterClusterConfig {
	var returns *MdbMongodbClusterClusterConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) Mongod() MdbMongodbClusterClusterConfigMongodOutputReference {
	var returns MdbMongodbClusterClusterConfigMongodOutputReference
	_jsii_.Get(
		j,
		"mongod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) MongodInput() *MdbMongodbClusterClusterConfigMongod {
	var returns *MdbMongodbClusterClusterConfigMongod
	_jsii_.Get(
		j,
		"mongodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewMdbMongodbClusterClusterConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MdbMongodbClusterClusterConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MdbMongodbClusterClusterConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbMongodbClusterClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMdbMongodbClusterClusterConfigOutputReference_Override(m MdbMongodbClusterClusterConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbMongodbClusterClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) SetFeatureCompatibilityVersion(val *string) {
	_jsii_.Set(
		j,
		"featureCompatibilityVersion",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) SetInternalValue(val *MdbMongodbClusterClusterConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) PutAccess(value *MdbMongodbClusterClusterConfigAccess) {
	_jsii_.InvokeVoid(
		m,
		"putAccess",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) PutBackupWindowStart(value *MdbMongodbClusterClusterConfigBackupWindowStart) {
	_jsii_.InvokeVoid(
		m,
		"putBackupWindowStart",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) PutMongod(value *MdbMongodbClusterClusterConfigMongod) {
	_jsii_.InvokeVoid(
		m,
		"putMongod",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) ResetAccess() {
	_jsii_.InvokeVoid(
		m,
		"resetAccess",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) ResetBackupWindowStart() {
	_jsii_.InvokeVoid(
		m,
		"resetBackupWindowStart",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) ResetFeatureCompatibilityVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetFeatureCompatibilityVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) ResetMongod() {
	_jsii_.InvokeVoid(
		m,
		"resetMongod",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

