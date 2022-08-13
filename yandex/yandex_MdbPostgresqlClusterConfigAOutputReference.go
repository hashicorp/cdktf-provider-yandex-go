// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbPostgresqlClusterConfigAOutputReference interface {
	cdktf.ComplexObject
	Access() MdbPostgresqlClusterConfigAccessOutputReference
	AccessInput() *MdbPostgresqlClusterConfigAccess
	Autofailover() interface{}
	SetAutofailover(val interface{})
	AutofailoverInput() interface{}
	BackupRetainPeriodDays() *float64
	SetBackupRetainPeriodDays(val *float64)
	BackupRetainPeriodDaysInput() *float64
	BackupWindowStart() MdbPostgresqlClusterConfigBackupWindowStartOutputReference
	BackupWindowStartInput() *MdbPostgresqlClusterConfigBackupWindowStart
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
	InternalValue() *MdbPostgresqlClusterConfigA
	SetInternalValue(val *MdbPostgresqlClusterConfigA)
	PerformanceDiagnostics() MdbPostgresqlClusterConfigPerformanceDiagnosticsOutputReference
	PerformanceDiagnosticsInput() *MdbPostgresqlClusterConfigPerformanceDiagnostics
	PoolerConfig() MdbPostgresqlClusterConfigPoolerConfigOutputReference
	PoolerConfigInput() *MdbPostgresqlClusterConfigPoolerConfig
	PostgresqlConfig() *map[string]*string
	SetPostgresqlConfig(val *map[string]*string)
	PostgresqlConfigInput() *map[string]*string
	Resources() MdbPostgresqlClusterConfigResourcesOutputReference
	ResourcesInput() *MdbPostgresqlClusterConfigResources
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
	PutAccess(value *MdbPostgresqlClusterConfigAccess)
	PutBackupWindowStart(value *MdbPostgresqlClusterConfigBackupWindowStart)
	PutPerformanceDiagnostics(value *MdbPostgresqlClusterConfigPerformanceDiagnostics)
	PutPoolerConfig(value *MdbPostgresqlClusterConfigPoolerConfig)
	PutResources(value *MdbPostgresqlClusterConfigResources)
	ResetAccess()
	ResetAutofailover()
	ResetBackupRetainPeriodDays()
	ResetBackupWindowStart()
	ResetPerformanceDiagnostics()
	ResetPoolerConfig()
	ResetPostgresqlConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MdbPostgresqlClusterConfigAOutputReference
type jsiiProxy_MdbPostgresqlClusterConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) Access() MdbPostgresqlClusterConfigAccessOutputReference {
	var returns MdbPostgresqlClusterConfigAccessOutputReference
	_jsii_.Get(
		j,
		"access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) AccessInput() *MdbPostgresqlClusterConfigAccess {
	var returns *MdbPostgresqlClusterConfigAccess
	_jsii_.Get(
		j,
		"accessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) Autofailover() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autofailover",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) AutofailoverInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autofailoverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) BackupRetainPeriodDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetainPeriodDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) BackupRetainPeriodDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetainPeriodDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) BackupWindowStart() MdbPostgresqlClusterConfigBackupWindowStartOutputReference {
	var returns MdbPostgresqlClusterConfigBackupWindowStartOutputReference
	_jsii_.Get(
		j,
		"backupWindowStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) BackupWindowStartInput() *MdbPostgresqlClusterConfigBackupWindowStart {
	var returns *MdbPostgresqlClusterConfigBackupWindowStart
	_jsii_.Get(
		j,
		"backupWindowStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) InternalValue() *MdbPostgresqlClusterConfigA {
	var returns *MdbPostgresqlClusterConfigA
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) PerformanceDiagnostics() MdbPostgresqlClusterConfigPerformanceDiagnosticsOutputReference {
	var returns MdbPostgresqlClusterConfigPerformanceDiagnosticsOutputReference
	_jsii_.Get(
		j,
		"performanceDiagnostics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) PerformanceDiagnosticsInput() *MdbPostgresqlClusterConfigPerformanceDiagnostics {
	var returns *MdbPostgresqlClusterConfigPerformanceDiagnostics
	_jsii_.Get(
		j,
		"performanceDiagnosticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) PoolerConfig() MdbPostgresqlClusterConfigPoolerConfigOutputReference {
	var returns MdbPostgresqlClusterConfigPoolerConfigOutputReference
	_jsii_.Get(
		j,
		"poolerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) PoolerConfigInput() *MdbPostgresqlClusterConfigPoolerConfig {
	var returns *MdbPostgresqlClusterConfigPoolerConfig
	_jsii_.Get(
		j,
		"poolerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) PostgresqlConfig() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"postgresqlConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) PostgresqlConfigInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"postgresqlConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) Resources() MdbPostgresqlClusterConfigResourcesOutputReference {
	var returns MdbPostgresqlClusterConfigResourcesOutputReference
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) ResourcesInput() *MdbPostgresqlClusterConfigResources {
	var returns *MdbPostgresqlClusterConfigResources
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewMdbPostgresqlClusterConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MdbPostgresqlClusterConfigAOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MdbPostgresqlClusterConfigAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbPostgresqlClusterConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMdbPostgresqlClusterConfigAOutputReference_Override(m MdbPostgresqlClusterConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbPostgresqlClusterConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) SetAutofailover(val interface{}) {
	_jsii_.Set(
		j,
		"autofailover",
		val,
	)
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) SetBackupRetainPeriodDays(val *float64) {
	_jsii_.Set(
		j,
		"backupRetainPeriodDays",
		val,
	)
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) SetInternalValue(val *MdbPostgresqlClusterConfigA) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) SetPostgresqlConfig(val *map[string]*string) {
	_jsii_.Set(
		j,
		"postgresqlConfig",
		val,
	)
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) PutAccess(value *MdbPostgresqlClusterConfigAccess) {
	_jsii_.InvokeVoid(
		m,
		"putAccess",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) PutBackupWindowStart(value *MdbPostgresqlClusterConfigBackupWindowStart) {
	_jsii_.InvokeVoid(
		m,
		"putBackupWindowStart",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) PutPerformanceDiagnostics(value *MdbPostgresqlClusterConfigPerformanceDiagnostics) {
	_jsii_.InvokeVoid(
		m,
		"putPerformanceDiagnostics",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) PutPoolerConfig(value *MdbPostgresqlClusterConfigPoolerConfig) {
	_jsii_.InvokeVoid(
		m,
		"putPoolerConfig",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) PutResources(value *MdbPostgresqlClusterConfigResources) {
	_jsii_.InvokeVoid(
		m,
		"putResources",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) ResetAccess() {
	_jsii_.InvokeVoid(
		m,
		"resetAccess",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) ResetAutofailover() {
	_jsii_.InvokeVoid(
		m,
		"resetAutofailover",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) ResetBackupRetainPeriodDays() {
	_jsii_.InvokeVoid(
		m,
		"resetBackupRetainPeriodDays",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) ResetBackupWindowStart() {
	_jsii_.InvokeVoid(
		m,
		"resetBackupWindowStart",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) ResetPerformanceDiagnostics() {
	_jsii_.InvokeVoid(
		m,
		"resetPerformanceDiagnostics",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) ResetPoolerConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetPoolerConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) ResetPostgresqlConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetPostgresqlConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbPostgresqlClusterConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

