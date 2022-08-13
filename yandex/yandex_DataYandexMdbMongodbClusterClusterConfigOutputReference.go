// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexMdbMongodbClusterClusterConfigOutputReference interface {
	cdktf.ComplexObject
	Access() DataYandexMdbMongodbClusterClusterConfigAccessOutputReference
	AccessInput() *DataYandexMdbMongodbClusterClusterConfigAccess
	BackupWindowStart() DataYandexMdbMongodbClusterClusterConfigBackupWindowStartOutputReference
	BackupWindowStartInput() *DataYandexMdbMongodbClusterClusterConfigBackupWindowStart
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
	InternalValue() *DataYandexMdbMongodbClusterClusterConfig
	SetInternalValue(val *DataYandexMdbMongodbClusterClusterConfig)
	Mongod() DataYandexMdbMongodbClusterClusterConfigMongodOutputReference
	MongodInput() *DataYandexMdbMongodbClusterClusterConfigMongod
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
	PutAccess(value *DataYandexMdbMongodbClusterClusterConfigAccess)
	PutBackupWindowStart(value *DataYandexMdbMongodbClusterClusterConfigBackupWindowStart)
	PutMongod(value *DataYandexMdbMongodbClusterClusterConfigMongod)
	ResetAccess()
	ResetBackupWindowStart()
	ResetFeatureCompatibilityVersion()
	ResetMongod()
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataYandexMdbMongodbClusterClusterConfigOutputReference
type jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) Access() DataYandexMdbMongodbClusterClusterConfigAccessOutputReference {
	var returns DataYandexMdbMongodbClusterClusterConfigAccessOutputReference
	_jsii_.Get(
		j,
		"access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) AccessInput() *DataYandexMdbMongodbClusterClusterConfigAccess {
	var returns *DataYandexMdbMongodbClusterClusterConfigAccess
	_jsii_.Get(
		j,
		"accessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) BackupWindowStart() DataYandexMdbMongodbClusterClusterConfigBackupWindowStartOutputReference {
	var returns DataYandexMdbMongodbClusterClusterConfigBackupWindowStartOutputReference
	_jsii_.Get(
		j,
		"backupWindowStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) BackupWindowStartInput() *DataYandexMdbMongodbClusterClusterConfigBackupWindowStart {
	var returns *DataYandexMdbMongodbClusterClusterConfigBackupWindowStart
	_jsii_.Get(
		j,
		"backupWindowStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) FeatureCompatibilityVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"featureCompatibilityVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) FeatureCompatibilityVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"featureCompatibilityVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) InternalValue() *DataYandexMdbMongodbClusterClusterConfig {
	var returns *DataYandexMdbMongodbClusterClusterConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) Mongod() DataYandexMdbMongodbClusterClusterConfigMongodOutputReference {
	var returns DataYandexMdbMongodbClusterClusterConfigMongodOutputReference
	_jsii_.Get(
		j,
		"mongod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) MongodInput() *DataYandexMdbMongodbClusterClusterConfigMongod {
	var returns *DataYandexMdbMongodbClusterClusterConfigMongod
	_jsii_.Get(
		j,
		"mongodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewDataYandexMdbMongodbClusterClusterConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataYandexMdbMongodbClusterClusterConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexMdbMongodbClusterClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataYandexMdbMongodbClusterClusterConfigOutputReference_Override(d DataYandexMdbMongodbClusterClusterConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexMdbMongodbClusterClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) SetFeatureCompatibilityVersion(val *string) {
	_jsii_.Set(
		j,
		"featureCompatibilityVersion",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) SetInternalValue(val *DataYandexMdbMongodbClusterClusterConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) PutAccess(value *DataYandexMdbMongodbClusterClusterConfigAccess) {
	_jsii_.InvokeVoid(
		d,
		"putAccess",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) PutBackupWindowStart(value *DataYandexMdbMongodbClusterClusterConfigBackupWindowStart) {
	_jsii_.InvokeVoid(
		d,
		"putBackupWindowStart",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) PutMongod(value *DataYandexMdbMongodbClusterClusterConfigMongod) {
	_jsii_.InvokeVoid(
		d,
		"putMongod",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) ResetAccess() {
	_jsii_.InvokeVoid(
		d,
		"resetAccess",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) ResetBackupWindowStart() {
	_jsii_.InvokeVoid(
		d,
		"resetBackupWindowStart",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) ResetFeatureCompatibilityVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetFeatureCompatibilityVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) ResetMongod() {
	_jsii_.InvokeVoid(
		d,
		"resetMongod",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

