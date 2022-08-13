// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/yandex/r/serverless_container yandex_serverless_container}.
type ServerlessContainer interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Concurrency() *float64
	SetConcurrency(val *float64)
	ConcurrencyInput() *float64
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CoreFraction() *float64
	SetCoreFraction(val *float64)
	CoreFractionInput() *float64
	Cores() *float64
	SetCores(val *float64)
	CoresInput() *float64
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
	ExecutionTimeout() *string
	SetExecutionTimeout(val *string)
	ExecutionTimeoutInput() *string
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
	Image() ServerlessContainerImageOutputReference
	ImageInput() *ServerlessContainerImage
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Memory() *float64
	SetMemory(val *float64)
	MemoryInput() *float64
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
	RevisionId() *string
	ServiceAccountId() *string
	SetServiceAccountId(val *string)
	ServiceAccountIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ServerlessContainerTimeoutsOutputReference
	TimeoutsInput() interface{}
	Url() *string
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
	PutImage(value *ServerlessContainerImage)
	PutTimeouts(value *ServerlessContainerTimeouts)
	ResetConcurrency()
	ResetCoreFraction()
	ResetCores()
	ResetDescription()
	ResetExecutionTimeout()
	ResetFolderId()
	ResetId()
	ResetLabels()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetServiceAccountId()
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

// The jsii proxy struct for ServerlessContainer
type jsiiProxy_ServerlessContainer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServerlessContainer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) Concurrency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"concurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) ConcurrencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"concurrencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) CoreFraction() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coreFraction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) CoreFractionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coreFractionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) Cores() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) CoresInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) ExecutionTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) ExecutionTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) FolderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) FolderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) Image() ServerlessContainerImageOutputReference {
	var returns ServerlessContainerImageOutputReference
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) ImageInput() *ServerlessContainerImage {
	var returns *ServerlessContainerImage
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) Memory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) MemoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) RevisionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revisionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) ServiceAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) ServiceAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) Timeouts() ServerlessContainerTimeoutsOutputReference {
	var returns ServerlessContainerTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessContainer) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/serverless_container yandex_serverless_container} Resource.
func NewServerlessContainer(scope constructs.Construct, id *string, config *ServerlessContainerConfig) ServerlessContainer {
	_init_.Initialize()

	j := jsiiProxy_ServerlessContainer{}

	_jsii_.Create(
		"@cdktf/provider-yandex.ServerlessContainer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/serverless_container yandex_serverless_container} Resource.
func NewServerlessContainer_Override(s ServerlessContainer, scope constructs.Construct, id *string, config *ServerlessContainerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.ServerlessContainer",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServerlessContainer) SetConcurrency(val *float64) {
	_jsii_.Set(
		j,
		"concurrency",
		val,
	)
}

func (j *jsiiProxy_ServerlessContainer) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ServerlessContainer) SetCoreFraction(val *float64) {
	_jsii_.Set(
		j,
		"coreFraction",
		val,
	)
}

func (j *jsiiProxy_ServerlessContainer) SetCores(val *float64) {
	_jsii_.Set(
		j,
		"cores",
		val,
	)
}

func (j *jsiiProxy_ServerlessContainer) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServerlessContainer) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServerlessContainer) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ServerlessContainer) SetExecutionTimeout(val *string) {
	_jsii_.Set(
		j,
		"executionTimeout",
		val,
	)
}

func (j *jsiiProxy_ServerlessContainer) SetFolderId(val *string) {
	_jsii_.Set(
		j,
		"folderId",
		val,
	)
}

func (j *jsiiProxy_ServerlessContainer) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ServerlessContainer) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ServerlessContainer) SetLabels(val *map[string]*string) {
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_ServerlessContainer) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServerlessContainer) SetMemory(val *float64) {
	_jsii_.Set(
		j,
		"memory",
		val,
	)
}

func (j *jsiiProxy_ServerlessContainer) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServerlessContainer) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServerlessContainer) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ServerlessContainer) SetServiceAccountId(val *string) {
	_jsii_.Set(
		j,
		"serviceAccountId",
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
func ServerlessContainer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-yandex.ServerlessContainer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServerlessContainer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-yandex.ServerlessContainer",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_ServerlessContainer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_ServerlessContainer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessContainer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessContainer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessContainer) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessContainer) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessContainer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessContainer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessContainer) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessContainer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessContainer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessContainer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServerlessContainer) PutImage(value *ServerlessContainerImage) {
	_jsii_.InvokeVoid(
		s,
		"putImage",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServerlessContainer) PutTimeouts(value *ServerlessContainerTimeouts) {
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServerlessContainer) ResetConcurrency() {
	_jsii_.InvokeVoid(
		s,
		"resetConcurrency",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerlessContainer) ResetCoreFraction() {
	_jsii_.InvokeVoid(
		s,
		"resetCoreFraction",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerlessContainer) ResetCores() {
	_jsii_.InvokeVoid(
		s,
		"resetCores",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerlessContainer) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerlessContainer) ResetExecutionTimeout() {
	_jsii_.InvokeVoid(
		s,
		"resetExecutionTimeout",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerlessContainer) ResetFolderId() {
	_jsii_.InvokeVoid(
		s,
		"resetFolderId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerlessContainer) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerlessContainer) ResetLabels() {
	_jsii_.InvokeVoid(
		s,
		"resetLabels",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerlessContainer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerlessContainer) ResetServiceAccountId() {
	_jsii_.InvokeVoid(
		s,
		"resetServiceAccountId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerlessContainer) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerlessContainer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessContainer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessContainer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessContainer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

