// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/yandex/r/ydb_database_dedicated yandex_ydb_database_dedicated}.
type YdbDatabaseDedicated interface {
	cdktf.TerraformResource
	AssignPublicIps() interface{}
	SetAssignPublicIps(val interface{})
	AssignPublicIpsInput() interface{}
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
	Location() YdbDatabaseDedicatedLocationOutputReference
	LocationId() *string
	SetLocationId(val *string)
	LocationIdInput() *string
	LocationInput() *YdbDatabaseDedicatedLocation
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkId() *string
	SetNetworkId(val *string)
	NetworkIdInput() *string
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
	ResourcePresetId() *string
	SetResourcePresetId(val *string)
	ResourcePresetIdInput() *string
	ScalePolicy() YdbDatabaseDedicatedScalePolicyOutputReference
	ScalePolicyInput() *YdbDatabaseDedicatedScalePolicy
	Status() *string
	StorageConfig() YdbDatabaseDedicatedStorageConfigOutputReference
	StorageConfigInput() *YdbDatabaseDedicatedStorageConfig
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	SubnetIdsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() YdbDatabaseDedicatedTimeoutsOutputReference
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
	PutLocation(value *YdbDatabaseDedicatedLocation)
	PutScalePolicy(value *YdbDatabaseDedicatedScalePolicy)
	PutStorageConfig(value *YdbDatabaseDedicatedStorageConfig)
	PutTimeouts(value *YdbDatabaseDedicatedTimeouts)
	ResetAssignPublicIps()
	ResetDescription()
	ResetFolderId()
	ResetId()
	ResetLabels()
	ResetLocation()
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

// The jsii proxy struct for YdbDatabaseDedicated
type jsiiProxy_YdbDatabaseDedicated struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_YdbDatabaseDedicated) AssignPublicIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assignPublicIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) AssignPublicIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assignPublicIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) DatabasePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databasePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) FolderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) FolderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) Location() YdbDatabaseDedicatedLocationOutputReference {
	var returns YdbDatabaseDedicatedLocationOutputReference
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) LocationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) LocationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) LocationInput() *YdbDatabaseDedicatedLocation {
	var returns *YdbDatabaseDedicatedLocation
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) NetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) NetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) ResourcePresetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourcePresetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) ResourcePresetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourcePresetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) ScalePolicy() YdbDatabaseDedicatedScalePolicyOutputReference {
	var returns YdbDatabaseDedicatedScalePolicyOutputReference
	_jsii_.Get(
		j,
		"scalePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) ScalePolicyInput() *YdbDatabaseDedicatedScalePolicy {
	var returns *YdbDatabaseDedicatedScalePolicy
	_jsii_.Get(
		j,
		"scalePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) StorageConfig() YdbDatabaseDedicatedStorageConfigOutputReference {
	var returns YdbDatabaseDedicatedStorageConfigOutputReference
	_jsii_.Get(
		j,
		"storageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) StorageConfigInput() *YdbDatabaseDedicatedStorageConfig {
	var returns *YdbDatabaseDedicatedStorageConfig
	_jsii_.Get(
		j,
		"storageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) Timeouts() YdbDatabaseDedicatedTimeoutsOutputReference {
	var returns YdbDatabaseDedicatedTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) TlsEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"tlsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) YdbApiEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ydbApiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseDedicated) YdbFullEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ydbFullEndpoint",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/ydb_database_dedicated yandex_ydb_database_dedicated} Resource.
func NewYdbDatabaseDedicated(scope constructs.Construct, id *string, config *YdbDatabaseDedicatedConfig) YdbDatabaseDedicated {
	_init_.Initialize()

	j := jsiiProxy_YdbDatabaseDedicated{}

	_jsii_.Create(
		"@cdktf/provider-yandex.YdbDatabaseDedicated",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/ydb_database_dedicated yandex_ydb_database_dedicated} Resource.
func NewYdbDatabaseDedicated_Override(y YdbDatabaseDedicated, scope constructs.Construct, id *string, config *YdbDatabaseDedicatedConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.YdbDatabaseDedicated",
		[]interface{}{scope, id, config},
		y,
	)
}

func (j *jsiiProxy_YdbDatabaseDedicated) SetAssignPublicIps(val interface{}) {
	_jsii_.Set(
		j,
		"assignPublicIps",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseDedicated) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseDedicated) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseDedicated) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseDedicated) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseDedicated) SetFolderId(val *string) {
	_jsii_.Set(
		j,
		"folderId",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseDedicated) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseDedicated) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseDedicated) SetLabels(val *map[string]*string) {
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseDedicated) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseDedicated) SetLocationId(val *string) {
	_jsii_.Set(
		j,
		"locationId",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseDedicated) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseDedicated) SetNetworkId(val *string) {
	_jsii_.Set(
		j,
		"networkId",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseDedicated) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseDedicated) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseDedicated) SetResourcePresetId(val *string) {
	_jsii_.Set(
		j,
		"resourcePresetId",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseDedicated) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
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
func YdbDatabaseDedicated_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-yandex.YdbDatabaseDedicated",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func YdbDatabaseDedicated_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-yandex.YdbDatabaseDedicated",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (y *jsiiProxy_YdbDatabaseDedicated) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		y,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (y *jsiiProxy_YdbDatabaseDedicated) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		y,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseDedicated) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		y,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseDedicated) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		y,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseDedicated) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		y,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseDedicated) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		y,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseDedicated) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		y,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseDedicated) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		y,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseDedicated) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		y,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseDedicated) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		y,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseDedicated) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		y,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseDedicated) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		y,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (y *jsiiProxy_YdbDatabaseDedicated) PutLocation(value *YdbDatabaseDedicatedLocation) {
	_jsii_.InvokeVoid(
		y,
		"putLocation",
		[]interface{}{value},
	)
}

func (y *jsiiProxy_YdbDatabaseDedicated) PutScalePolicy(value *YdbDatabaseDedicatedScalePolicy) {
	_jsii_.InvokeVoid(
		y,
		"putScalePolicy",
		[]interface{}{value},
	)
}

func (y *jsiiProxy_YdbDatabaseDedicated) PutStorageConfig(value *YdbDatabaseDedicatedStorageConfig) {
	_jsii_.InvokeVoid(
		y,
		"putStorageConfig",
		[]interface{}{value},
	)
}

func (y *jsiiProxy_YdbDatabaseDedicated) PutTimeouts(value *YdbDatabaseDedicatedTimeouts) {
	_jsii_.InvokeVoid(
		y,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (y *jsiiProxy_YdbDatabaseDedicated) ResetAssignPublicIps() {
	_jsii_.InvokeVoid(
		y,
		"resetAssignPublicIps",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YdbDatabaseDedicated) ResetDescription() {
	_jsii_.InvokeVoid(
		y,
		"resetDescription",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YdbDatabaseDedicated) ResetFolderId() {
	_jsii_.InvokeVoid(
		y,
		"resetFolderId",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YdbDatabaseDedicated) ResetId() {
	_jsii_.InvokeVoid(
		y,
		"resetId",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YdbDatabaseDedicated) ResetLabels() {
	_jsii_.InvokeVoid(
		y,
		"resetLabels",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YdbDatabaseDedicated) ResetLocation() {
	_jsii_.InvokeVoid(
		y,
		"resetLocation",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YdbDatabaseDedicated) ResetLocationId() {
	_jsii_.InvokeVoid(
		y,
		"resetLocationId",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YdbDatabaseDedicated) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		y,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YdbDatabaseDedicated) ResetTimeouts() {
	_jsii_.InvokeVoid(
		y,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YdbDatabaseDedicated) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		y,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseDedicated) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		y,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseDedicated) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		y,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseDedicated) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		y,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

