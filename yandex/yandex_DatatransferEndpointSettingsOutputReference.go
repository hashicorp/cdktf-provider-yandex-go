// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatatransferEndpointSettingsOutputReference interface {
	cdktf.ComplexObject
	ClickhouseSource() DatatransferEndpointSettingsClickhouseSourceOutputReference
	ClickhouseSourceInput() *DatatransferEndpointSettingsClickhouseSource
	ClickhouseTarget() DatatransferEndpointSettingsClickhouseTargetOutputReference
	ClickhouseTargetInput() *DatatransferEndpointSettingsClickhouseTarget
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
	InternalValue() *DatatransferEndpointSettings
	SetInternalValue(val *DatatransferEndpointSettings)
	MongoSource() DatatransferEndpointSettingsMongoSourceOutputReference
	MongoSourceInput() *DatatransferEndpointSettingsMongoSource
	MongoTarget() DatatransferEndpointSettingsMongoTargetOutputReference
	MongoTargetInput() *DatatransferEndpointSettingsMongoTarget
	MysqlSource() DatatransferEndpointSettingsMysqlSourceOutputReference
	MysqlSourceInput() *DatatransferEndpointSettingsMysqlSource
	MysqlTarget() DatatransferEndpointSettingsMysqlTargetOutputReference
	MysqlTargetInput() *DatatransferEndpointSettingsMysqlTarget
	PostgresSource() DatatransferEndpointSettingsPostgresSourceOutputReference
	PostgresSourceInput() *DatatransferEndpointSettingsPostgresSource
	PostgresTarget() DatatransferEndpointSettingsPostgresTargetOutputReference
	PostgresTargetInput() *DatatransferEndpointSettingsPostgresTarget
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
	PutClickhouseSource(value *DatatransferEndpointSettingsClickhouseSource)
	PutClickhouseTarget(value *DatatransferEndpointSettingsClickhouseTarget)
	PutMongoSource(value *DatatransferEndpointSettingsMongoSource)
	PutMongoTarget(value *DatatransferEndpointSettingsMongoTarget)
	PutMysqlSource(value *DatatransferEndpointSettingsMysqlSource)
	PutMysqlTarget(value *DatatransferEndpointSettingsMysqlTarget)
	PutPostgresSource(value *DatatransferEndpointSettingsPostgresSource)
	PutPostgresTarget(value *DatatransferEndpointSettingsPostgresTarget)
	ResetClickhouseSource()
	ResetClickhouseTarget()
	ResetMongoSource()
	ResetMongoTarget()
	ResetMysqlSource()
	ResetMysqlTarget()
	ResetPostgresSource()
	ResetPostgresTarget()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatatransferEndpointSettingsOutputReference
type jsiiProxy_DatatransferEndpointSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatatransferEndpointSettingsOutputReference) ClickhouseSource() DatatransferEndpointSettingsClickhouseSourceOutputReference {
	var returns DatatransferEndpointSettingsClickhouseSourceOutputReference
	_jsii_.Get(
		j,
		"clickhouseSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsOutputReference) ClickhouseSourceInput() *DatatransferEndpointSettingsClickhouseSource {
	var returns *DatatransferEndpointSettingsClickhouseSource
	_jsii_.Get(
		j,
		"clickhouseSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsOutputReference) ClickhouseTarget() DatatransferEndpointSettingsClickhouseTargetOutputReference {
	var returns DatatransferEndpointSettingsClickhouseTargetOutputReference
	_jsii_.Get(
		j,
		"clickhouseTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsOutputReference) ClickhouseTargetInput() *DatatransferEndpointSettingsClickhouseTarget {
	var returns *DatatransferEndpointSettingsClickhouseTarget
	_jsii_.Get(
		j,
		"clickhouseTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsOutputReference) InternalValue() *DatatransferEndpointSettings {
	var returns *DatatransferEndpointSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsOutputReference) MongoSource() DatatransferEndpointSettingsMongoSourceOutputReference {
	var returns DatatransferEndpointSettingsMongoSourceOutputReference
	_jsii_.Get(
		j,
		"mongoSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsOutputReference) MongoSourceInput() *DatatransferEndpointSettingsMongoSource {
	var returns *DatatransferEndpointSettingsMongoSource
	_jsii_.Get(
		j,
		"mongoSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsOutputReference) MongoTarget() DatatransferEndpointSettingsMongoTargetOutputReference {
	var returns DatatransferEndpointSettingsMongoTargetOutputReference
	_jsii_.Get(
		j,
		"mongoTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsOutputReference) MongoTargetInput() *DatatransferEndpointSettingsMongoTarget {
	var returns *DatatransferEndpointSettingsMongoTarget
	_jsii_.Get(
		j,
		"mongoTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsOutputReference) MysqlSource() DatatransferEndpointSettingsMysqlSourceOutputReference {
	var returns DatatransferEndpointSettingsMysqlSourceOutputReference
	_jsii_.Get(
		j,
		"mysqlSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsOutputReference) MysqlSourceInput() *DatatransferEndpointSettingsMysqlSource {
	var returns *DatatransferEndpointSettingsMysqlSource
	_jsii_.Get(
		j,
		"mysqlSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsOutputReference) MysqlTarget() DatatransferEndpointSettingsMysqlTargetOutputReference {
	var returns DatatransferEndpointSettingsMysqlTargetOutputReference
	_jsii_.Get(
		j,
		"mysqlTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsOutputReference) MysqlTargetInput() *DatatransferEndpointSettingsMysqlTarget {
	var returns *DatatransferEndpointSettingsMysqlTarget
	_jsii_.Get(
		j,
		"mysqlTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsOutputReference) PostgresSource() DatatransferEndpointSettingsPostgresSourceOutputReference {
	var returns DatatransferEndpointSettingsPostgresSourceOutputReference
	_jsii_.Get(
		j,
		"postgresSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsOutputReference) PostgresSourceInput() *DatatransferEndpointSettingsPostgresSource {
	var returns *DatatransferEndpointSettingsPostgresSource
	_jsii_.Get(
		j,
		"postgresSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsOutputReference) PostgresTarget() DatatransferEndpointSettingsPostgresTargetOutputReference {
	var returns DatatransferEndpointSettingsPostgresTargetOutputReference
	_jsii_.Get(
		j,
		"postgresTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsOutputReference) PostgresTargetInput() *DatatransferEndpointSettingsPostgresTarget {
	var returns *DatatransferEndpointSettingsPostgresTarget
	_jsii_.Get(
		j,
		"postgresTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatatransferEndpointSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatatransferEndpointSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DatatransferEndpointSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatatransferEndpointSettingsOutputReference_Override(d DatatransferEndpointSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsOutputReference) SetInternalValue(val *DatatransferEndpointSettings) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsOutputReference) PutClickhouseSource(value *DatatransferEndpointSettingsClickhouseSource) {
	_jsii_.InvokeVoid(
		d,
		"putClickhouseSource",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsOutputReference) PutClickhouseTarget(value *DatatransferEndpointSettingsClickhouseTarget) {
	_jsii_.InvokeVoid(
		d,
		"putClickhouseTarget",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsOutputReference) PutMongoSource(value *DatatransferEndpointSettingsMongoSource) {
	_jsii_.InvokeVoid(
		d,
		"putMongoSource",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsOutputReference) PutMongoTarget(value *DatatransferEndpointSettingsMongoTarget) {
	_jsii_.InvokeVoid(
		d,
		"putMongoTarget",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsOutputReference) PutMysqlSource(value *DatatransferEndpointSettingsMysqlSource) {
	_jsii_.InvokeVoid(
		d,
		"putMysqlSource",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsOutputReference) PutMysqlTarget(value *DatatransferEndpointSettingsMysqlTarget) {
	_jsii_.InvokeVoid(
		d,
		"putMysqlTarget",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsOutputReference) PutPostgresSource(value *DatatransferEndpointSettingsPostgresSource) {
	_jsii_.InvokeVoid(
		d,
		"putPostgresSource",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsOutputReference) PutPostgresTarget(value *DatatransferEndpointSettingsPostgresTarget) {
	_jsii_.InvokeVoid(
		d,
		"putPostgresTarget",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsOutputReference) ResetClickhouseSource() {
	_jsii_.InvokeVoid(
		d,
		"resetClickhouseSource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsOutputReference) ResetClickhouseTarget() {
	_jsii_.InvokeVoid(
		d,
		"resetClickhouseTarget",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsOutputReference) ResetMongoSource() {
	_jsii_.InvokeVoid(
		d,
		"resetMongoSource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsOutputReference) ResetMongoTarget() {
	_jsii_.InvokeVoid(
		d,
		"resetMongoTarget",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsOutputReference) ResetMysqlSource() {
	_jsii_.InvokeVoid(
		d,
		"resetMysqlSource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsOutputReference) ResetMysqlTarget() {
	_jsii_.InvokeVoid(
		d,
		"resetMysqlTarget",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsOutputReference) ResetPostgresSource() {
	_jsii_.InvokeVoid(
		d,
		"resetPostgresSource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsOutputReference) ResetPostgresTarget() {
	_jsii_.InvokeVoid(
		d,
		"resetPostgresTarget",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

