// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/yandex/r/ydb_database_serverless yandex_ydb_database_serverless}.
type YdbDatabaseServerless interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	CreatedAt() *string
	DatabasePath() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DocumentApiEndpoint() *string
	FolderId() *string
	SetFolderId(val *string)
	FolderIdInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocationId() *string
	SetLocationId(val *string)
	LocationIdInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() YdbDatabaseServerlessTimeoutsOutputReference
	TimeoutsInput() interface{}
	TlsEnabled() cdktf.IResolvable
	YdbApiEndpoint() *string
	YdbFullEndpoint() *string
	// Experimental.
	AddOverride(path *string, value interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *YdbDatabaseServerlessTimeouts)
	ResetDescription()
	ResetFolderId()
	ResetId()
	ResetLabels()
	ResetLocationId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for YdbDatabaseServerless
type jsiiProxy_YdbDatabaseServerless struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_YdbDatabaseServerless) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) DatabasePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databasePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) DocumentApiEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentApiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) FolderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) FolderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) LocationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) LocationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) Timeouts() YdbDatabaseServerlessTimeoutsOutputReference {
	var returns YdbDatabaseServerlessTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) TlsEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"tlsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) YdbApiEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ydbApiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseServerless) YdbFullEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ydbFullEndpoint",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/ydb_database_serverless yandex_ydb_database_serverless} Resource.
func NewYdbDatabaseServerless(scope constructs.Construct, id *string, config *YdbDatabaseServerlessConfig) YdbDatabaseServerless {
	_init_.Initialize()

	j := jsiiProxy_YdbDatabaseServerless{}

	_jsii_.Create(
		"@cdktf/provider-yandex.YdbDatabaseServerless",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/ydb_database_serverless yandex_ydb_database_serverless} Resource.
func NewYdbDatabaseServerless_Override(y YdbDatabaseServerless, scope constructs.Construct, id *string, config *YdbDatabaseServerlessConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.YdbDatabaseServerless",
		[]interface{}{scope, id, config},
		y,
	)
}

func (j *jsiiProxy_YdbDatabaseServerless) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseServerless) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseServerless) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseServerless) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseServerless) SetFolderId(val *string) {
	_jsii_.Set(
		j,
		"folderId",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseServerless) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseServerless) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseServerless) SetLabels(val *map[string]*string) {
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseServerless) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseServerless) SetLocationId(val *string) {
	_jsii_.Set(
		j,
		"locationId",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseServerless) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseServerless) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseServerless) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func YdbDatabaseServerless_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-yandex.YdbDatabaseServerless",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func YdbDatabaseServerless_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-yandex.YdbDatabaseServerless",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (y *jsiiProxy_YdbDatabaseServerless) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		y,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (y *jsiiProxy_YdbDatabaseServerless) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		y,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseServerless) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		y,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseServerless) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		y,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseServerless) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		y,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseServerless) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		y,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseServerless) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		y,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseServerless) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		y,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseServerless) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		y,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseServerless) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		y,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseServerless) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		y,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseServerless) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		y,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (y *jsiiProxy_YdbDatabaseServerless) PutTimeouts(value *YdbDatabaseServerlessTimeouts) {
	_jsii_.InvokeVoid(
		y,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (y *jsiiProxy_YdbDatabaseServerless) ResetDescription() {
	_jsii_.InvokeVoid(
		y,
		"resetDescription",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YdbDatabaseServerless) ResetFolderId() {
	_jsii_.InvokeVoid(
		y,
		"resetFolderId",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YdbDatabaseServerless) ResetId() {
	_jsii_.InvokeVoid(
		y,
		"resetId",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YdbDatabaseServerless) ResetLabels() {
	_jsii_.InvokeVoid(
		y,
		"resetLabels",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YdbDatabaseServerless) ResetLocationId() {
	_jsii_.InvokeVoid(
		y,
		"resetLocationId",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YdbDatabaseServerless) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		y,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YdbDatabaseServerless) ResetTimeouts() {
	_jsii_.InvokeVoid(
		y,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YdbDatabaseServerless) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		y,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseServerless) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		y,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseServerless) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		y,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseServerless) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		y,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

