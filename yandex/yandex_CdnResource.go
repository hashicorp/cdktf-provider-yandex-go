// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource yandex_cdn_resource}.
type CdnResource interface {
	cdktf.TerraformResource
	Active() interface{}
	SetActive(val interface{})
	ActiveInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Cname() *string
	SetCname(val *string)
	CnameInput() *string
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
	Options() CdnResourceOptionsOutputReference
	OptionsInput() *CdnResourceOptions
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
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	SecondaryHostnames() *[]*string
	SetSecondaryHostnames(val *[]*string)
	SecondaryHostnamesInput() *[]*string
	SslCertificate() CdnResourceSslCertificateOutputReference
	SslCertificateInput() *CdnResourceSslCertificate
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() CdnResourceTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutOptions(value *CdnResourceOptions)
	PutSslCertificate(value *CdnResourceSslCertificate)
	PutTimeouts(value *CdnResourceTimeouts)
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
	ResetSecondaryHostnames()
	ResetSslCertificate()
	ResetTimeouts()
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

// The jsii proxy struct for CdnResource
type jsiiProxy_CdnResource struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CdnResource) Active() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"active",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) ActiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) Cname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) CnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) FolderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) FolderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) Options() CdnResourceOptionsOutputReference {
	var returns CdnResourceOptionsOutputReference
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) OptionsInput() *CdnResourceOptions {
	var returns *CdnResourceOptions
	_jsii_.Get(
		j,
		"optionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) OriginGroupId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"originGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) OriginGroupIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"originGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) OriginGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) OriginGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) OriginProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) OriginProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) SecondaryHostnames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secondaryHostnames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) SecondaryHostnamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secondaryHostnamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) SslCertificate() CdnResourceSslCertificateOutputReference {
	var returns CdnResourceSslCertificateOutputReference
	_jsii_.Get(
		j,
		"sslCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) SslCertificateInput() *CdnResourceSslCertificate {
	var returns *CdnResourceSslCertificate
	_jsii_.Get(
		j,
		"sslCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) Timeouts() CdnResourceTimeoutsOutputReference {
	var returns CdnResourceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResource) UpdatedAtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAtInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource yandex_cdn_resource} Resource.
func NewCdnResource(scope constructs.Construct, id *string, config *CdnResourceConfig) CdnResource {
	_init_.Initialize()

	j := jsiiProxy_CdnResource{}

	_jsii_.Create(
		"@cdktf/provider-yandex.CdnResource",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource yandex_cdn_resource} Resource.
func NewCdnResource_Override(c CdnResource, scope constructs.Construct, id *string, config *CdnResourceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.CdnResource",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CdnResource) SetActive(val interface{}) {
	_jsii_.Set(
		j,
		"active",
		val,
	)
}

func (j *jsiiProxy_CdnResource) SetCname(val *string) {
	_jsii_.Set(
		j,
		"cname",
		val,
	)
}

func (j *jsiiProxy_CdnResource) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CdnResource) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CdnResource) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CdnResource) SetFolderId(val *string) {
	_jsii_.Set(
		j,
		"folderId",
		val,
	)
}

func (j *jsiiProxy_CdnResource) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CdnResource) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CdnResource) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CdnResource) SetOriginGroupId(val *float64) {
	_jsii_.Set(
		j,
		"originGroupId",
		val,
	)
}

func (j *jsiiProxy_CdnResource) SetOriginGroupName(val *string) {
	_jsii_.Set(
		j,
		"originGroupName",
		val,
	)
}

func (j *jsiiProxy_CdnResource) SetOriginProtocol(val *string) {
	_jsii_.Set(
		j,
		"originProtocol",
		val,
	)
}

func (j *jsiiProxy_CdnResource) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CdnResource) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CdnResource) SetSecondaryHostnames(val *[]*string) {
	_jsii_.Set(
		j,
		"secondaryHostnames",
		val,
	)
}

func (j *jsiiProxy_CdnResource) SetUpdatedAt(val *string) {
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
func CdnResource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-yandex.CdnResource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CdnResource_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-yandex.CdnResource",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CdnResource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CdnResource) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnResource) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnResource) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnResource) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnResource) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnResource) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnResource) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnResource) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnResource) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnResource) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnResource) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CdnResource) PutOptions(value *CdnResourceOptions) {
	_jsii_.InvokeVoid(
		c,
		"putOptions",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnResource) PutSslCertificate(value *CdnResourceSslCertificate) {
	_jsii_.InvokeVoid(
		c,
		"putSslCertificate",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnResource) PutTimeouts(value *CdnResourceTimeouts) {
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnResource) ResetActive() {
	_jsii_.InvokeVoid(
		c,
		"resetActive",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResource) ResetCname() {
	_jsii_.InvokeVoid(
		c,
		"resetCname",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResource) ResetFolderId() {
	_jsii_.InvokeVoid(
		c,
		"resetFolderId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResource) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResource) ResetOptions() {
	_jsii_.InvokeVoid(
		c,
		"resetOptions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResource) ResetOriginGroupId() {
	_jsii_.InvokeVoid(
		c,
		"resetOriginGroupId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResource) ResetOriginGroupName() {
	_jsii_.InvokeVoid(
		c,
		"resetOriginGroupName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResource) ResetOriginProtocol() {
	_jsii_.InvokeVoid(
		c,
		"resetOriginProtocol",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResource) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResource) ResetSecondaryHostnames() {
	_jsii_.InvokeVoid(
		c,
		"resetSecondaryHostnames",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResource) ResetSslCertificate() {
	_jsii_.InvokeVoid(
		c,
		"resetSslCertificate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResource) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResource) ResetUpdatedAt() {
	_jsii_.InvokeVoid(
		c,
		"resetUpdatedAt",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResource) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnResource) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnResource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnResource) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

