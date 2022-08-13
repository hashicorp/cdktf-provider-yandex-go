// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group yandex_alb_backend_group}.
type DataYandexAlbBackendGroup interface {
	cdktf.TerraformDataSource
	BackendGroupId() *string
	SetBackendGroupId(val *string)
	BackendGroupIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	CreatedAt() *string
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
	GrpcBackend() DataYandexAlbBackendGroupGrpcBackendList
	GrpcBackendInput() interface{}
	HttpBackend() DataYandexAlbBackendGroupHttpBackendList
	HttpBackendInput() interface{}
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
	RawOverrides() interface{}
	SessionAffinity() DataYandexAlbBackendGroupSessionAffinityOutputReference
	SessionAffinityInput() *DataYandexAlbBackendGroupSessionAffinity
	StreamBackend() DataYandexAlbBackendGroupStreamBackendList
	StreamBackendInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutGrpcBackend(value interface{})
	PutHttpBackend(value interface{})
	PutSessionAffinity(value *DataYandexAlbBackendGroupSessionAffinity)
	PutStreamBackend(value interface{})
	ResetBackendGroupId()
	ResetDescription()
	ResetFolderId()
	ResetGrpcBackend()
	ResetHttpBackend()
	ResetId()
	ResetLabels()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSessionAffinity()
	ResetStreamBackend()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataYandexAlbBackendGroup
type jsiiProxy_DataYandexAlbBackendGroup struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) BackendGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) BackendGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) FolderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) FolderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) GrpcBackend() DataYandexAlbBackendGroupGrpcBackendList {
	var returns DataYandexAlbBackendGroupGrpcBackendList
	_jsii_.Get(
		j,
		"grpcBackend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) GrpcBackendInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"grpcBackendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) HttpBackend() DataYandexAlbBackendGroupHttpBackendList {
	var returns DataYandexAlbBackendGroupHttpBackendList
	_jsii_.Get(
		j,
		"httpBackend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) HttpBackendInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpBackendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) SessionAffinity() DataYandexAlbBackendGroupSessionAffinityOutputReference {
	var returns DataYandexAlbBackendGroupSessionAffinityOutputReference
	_jsii_.Get(
		j,
		"sessionAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) SessionAffinityInput() *DataYandexAlbBackendGroupSessionAffinity {
	var returns *DataYandexAlbBackendGroupSessionAffinity
	_jsii_.Get(
		j,
		"sessionAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) StreamBackend() DataYandexAlbBackendGroupStreamBackendList {
	var returns DataYandexAlbBackendGroupStreamBackendList
	_jsii_.Get(
		j,
		"streamBackend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) StreamBackendInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamBackendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group yandex_alb_backend_group} Data Source.
func NewDataYandexAlbBackendGroup(scope constructs.Construct, id *string, config *DataYandexAlbBackendGroupConfig) DataYandexAlbBackendGroup {
	_init_.Initialize()

	j := jsiiProxy_DataYandexAlbBackendGroup{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexAlbBackendGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group yandex_alb_backend_group} Data Source.
func NewDataYandexAlbBackendGroup_Override(d DataYandexAlbBackendGroup, scope constructs.Construct, id *string, config *DataYandexAlbBackendGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexAlbBackendGroup",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) SetBackendGroupId(val *string) {
	_jsii_.Set(
		j,
		"backendGroupId",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) SetFolderId(val *string) {
	_jsii_.Set(
		j,
		"folderId",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) SetLabels(val *map[string]*string) {
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroup) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
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
func DataYandexAlbBackendGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-yandex.DataYandexAlbBackendGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataYandexAlbBackendGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-yandex.DataYandexAlbBackendGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) PutGrpcBackend(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putGrpcBackend",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) PutHttpBackend(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putHttpBackend",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) PutSessionAffinity(value *DataYandexAlbBackendGroupSessionAffinity) {
	_jsii_.InvokeVoid(
		d,
		"putSessionAffinity",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) PutStreamBackend(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putStreamBackend",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) ResetBackendGroupId() {
	_jsii_.InvokeVoid(
		d,
		"resetBackendGroupId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) ResetFolderId() {
	_jsii_.InvokeVoid(
		d,
		"resetFolderId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) ResetGrpcBackend() {
	_jsii_.InvokeVoid(
		d,
		"resetGrpcBackend",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) ResetHttpBackend() {
	_jsii_.InvokeVoid(
		d,
		"resetHttpBackend",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) ResetLabels() {
	_jsii_.InvokeVoid(
		d,
		"resetLabels",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) ResetSessionAffinity() {
	_jsii_.InvokeVoid(
		d,
		"resetSessionAffinity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) ResetStreamBackend() {
	_jsii_.InvokeVoid(
		d,
		"resetStreamBackend",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

