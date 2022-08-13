// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/yandex/d/cdn_resource yandex_cdn_resource}.
type DataYandexCdnResource interface {
	cdktf.TerraformDataSource
	Active() interface{}
	SetActive(val interface{})
	ActiveInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Cname() *string
	SetCname(val *string)
	CnameInput() *string
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	Options() DataYandexCdnResourceOptionsOutputReference
	OptionsInput() *DataYandexCdnResourceOptions
	OriginGroupId() *float64
	SetOriginGroupId(val *float64)
	OriginGroupIdInput() *float64
	OriginGroupName() *string
	SetOriginGroupName(val *string)
	OriginGroupNameInput() *string
	OriginProtocol() *string
	SetOriginProtocol(val *string)
	OriginProtocolInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	ResourceId() *string
	SetResourceId(val *string)
	ResourceIdInput() *string
	SecondaryHostnames() *[]*string
	SetSecondaryHostnames(val *[]*string)
	SecondaryHostnamesInput() *[]*string
	SslCertificate() DataYandexCdnResourceSslCertificateOutputReference
	SslCertificateInput() *DataYandexCdnResourceSslCertificate
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdatedAt() *string
	SetUpdatedAt(val *string)
	UpdatedAtInput() *string
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
	PutOptions(value *DataYandexCdnResourceOptions)
	PutSslCertificate(value *DataYandexCdnResourceSslCertificate)
	ResetActive()
	ResetCname()
	ResetFolderId()
	ResetId()
	ResetOptions()
	ResetOriginGroupId()
	ResetOriginGroupName()
	ResetOriginProtocol()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetResourceId()
	ResetSecondaryHostnames()
	ResetSslCertificate()
	ResetUpdatedAt()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataYandexCdnResource
type jsiiProxy_DataYandexCdnResource struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataYandexCdnResource) Active() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"active",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) ActiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) Cname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) CnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) FolderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) FolderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) Options() DataYandexCdnResourceOptionsOutputReference {
	var returns DataYandexCdnResourceOptionsOutputReference
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) OptionsInput() *DataYandexCdnResourceOptions {
	var returns *DataYandexCdnResourceOptions
	_jsii_.Get(
		j,
		"optionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) OriginGroupId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"originGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) OriginGroupIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"originGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) OriginGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) OriginGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) OriginProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) OriginProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) ResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) SecondaryHostnames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secondaryHostnames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) SecondaryHostnamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secondaryHostnamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) SslCertificate() DataYandexCdnResourceSslCertificateOutputReference {
	var returns DataYandexCdnResourceSslCertificateOutputReference
	_jsii_.Get(
		j,
		"sslCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) SslCertificateInput() *DataYandexCdnResourceSslCertificate {
	var returns *DataYandexCdnResourceSslCertificate
	_jsii_.Get(
		j,
		"sslCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResource) UpdatedAtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAtInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/yandex/d/cdn_resource yandex_cdn_resource} Data Source.
func NewDataYandexCdnResource(scope constructs.Construct, id *string, config *DataYandexCdnResourceConfig) DataYandexCdnResource {
	_init_.Initialize()

	j := jsiiProxy_DataYandexCdnResource{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexCdnResource",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/yandex/d/cdn_resource yandex_cdn_resource} Data Source.
func NewDataYandexCdnResource_Override(d DataYandexCdnResource, scope constructs.Construct, id *string, config *DataYandexCdnResourceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexCdnResource",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataYandexCdnResource) SetActive(val interface{}) {
	_jsii_.Set(
		j,
		"active",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResource) SetCname(val *string) {
	_jsii_.Set(
		j,
		"cname",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResource) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResource) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResource) SetFolderId(val *string) {
	_jsii_.Set(
		j,
		"folderId",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResource) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResource) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResource) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResource) SetOriginGroupId(val *float64) {
	_jsii_.Set(
		j,
		"originGroupId",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResource) SetOriginGroupName(val *string) {
	_jsii_.Set(
		j,
		"originGroupName",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResource) SetOriginProtocol(val *string) {
	_jsii_.Set(
		j,
		"originProtocol",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResource) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResource) SetResourceId(val *string) {
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResource) SetSecondaryHostnames(val *[]*string) {
	_jsii_.Set(
		j,
		"secondaryHostnames",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResource) SetUpdatedAt(val *string) {
	_jsii_.Set(
		j,
		"updatedAt",
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
func DataYandexCdnResource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-yandex.DataYandexCdnResource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataYandexCdnResource_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-yandex.DataYandexCdnResource",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataYandexCdnResource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataYandexCdnResource) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexCdnResource) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexCdnResource) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexCdnResource) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexCdnResource) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexCdnResource) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexCdnResource) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexCdnResource) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexCdnResource) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexCdnResource) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexCdnResource) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataYandexCdnResource) PutOptions(value *DataYandexCdnResourceOptions) {
	_jsii_.InvokeVoid(
		d,
		"putOptions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexCdnResource) PutSslCertificate(value *DataYandexCdnResourceSslCertificate) {
	_jsii_.InvokeVoid(
		d,
		"putSslCertificate",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexCdnResource) ResetActive() {
	_jsii_.InvokeVoid(
		d,
		"resetActive",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResource) ResetCname() {
	_jsii_.InvokeVoid(
		d,
		"resetCname",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResource) ResetFolderId() {
	_jsii_.InvokeVoid(
		d,
		"resetFolderId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResource) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResource) ResetOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResource) ResetOriginGroupId() {
	_jsii_.InvokeVoid(
		d,
		"resetOriginGroupId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResource) ResetOriginGroupName() {
	_jsii_.InvokeVoid(
		d,
		"resetOriginGroupName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResource) ResetOriginProtocol() {
	_jsii_.InvokeVoid(
		d,
		"resetOriginProtocol",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResource) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResource) ResetResourceId() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResource) ResetSecondaryHostnames() {
	_jsii_.InvokeVoid(
		d,
		"resetSecondaryHostnames",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResource) ResetSslCertificate() {
	_jsii_.InvokeVoid(
		d,
		"resetSslCertificate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResource) ResetUpdatedAt() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdatedAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResource) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexCdnResource) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexCdnResource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexCdnResource) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

