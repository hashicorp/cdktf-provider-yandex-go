// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/yandex yandex}.
type YandexProvider interface {
	cdktf.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudId() *string
	SetCloudId(val *string)
	CloudIdInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Endpoint() *string
	SetEndpoint(val *string)
	EndpointInput() *string
	FolderId() *string
	SetFolderId(val *string)
	FolderIdInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Insecure() interface{}
	SetInsecure(val interface{})
	InsecureInput() interface{}
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	MaxRetriesInput() *float64
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	OrganizationId() *string
	SetOrganizationId(val *string)
	OrganizationIdInput() *string
	Plaintext() interface{}
	SetPlaintext(val interface{})
	PlaintextInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	RegionId() *string
	SetRegionId(val *string)
	RegionIdInput() *string
	ServiceAccountKeyFile() *string
	SetServiceAccountKeyFile(val *string)
	ServiceAccountKeyFileInput() *string
	StorageAccessKey() *string
	SetStorageAccessKey(val *string)
	StorageAccessKeyInput() *string
	StorageEndpoint() *string
	SetStorageEndpoint(val *string)
	StorageEndpointInput() *string
	StorageSecretKey() *string
	SetStorageSecretKey(val *string)
	StorageSecretKeyInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	Token() *string
	SetToken(val *string)
	TokenInput() *string
	YmqAccessKey() *string
	SetYmqAccessKey(val *string)
	YmqAccessKeyInput() *string
	YmqEndpoint() *string
	SetYmqEndpoint(val *string)
	YmqEndpointInput() *string
	YmqSecretKey() *string
	SetYmqSecretKey(val *string)
	YmqSecretKeyInput() *string
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAlias()
	ResetCloudId()
	ResetEndpoint()
	ResetFolderId()
	ResetInsecure()
	ResetMaxRetries()
	ResetOrganizationId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlaintext()
	ResetRegionId()
	ResetServiceAccountKeyFile()
	ResetStorageAccessKey()
	ResetStorageEndpoint()
	ResetStorageSecretKey()
	ResetToken()
	ResetYmqAccessKey()
	ResetYmqEndpoint()
	ResetYmqSecretKey()
	ResetZone()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for YandexProvider
type jsiiProxy_YandexProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_YandexProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) CloudId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) CloudIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) EndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) FolderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) FolderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) Insecure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) InsecureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) OrganizationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) OrganizationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) Plaintext() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"plaintext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) PlaintextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"plaintextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) RegionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) RegionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) ServiceAccountKeyFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountKeyFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) ServiceAccountKeyFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountKeyFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) StorageAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) StorageAccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) StorageEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) StorageEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) StorageSecretKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageSecretKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) StorageSecretKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageSecretKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) YmqAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ymqAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) YmqAccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ymqAccessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) YmqEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ymqEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) YmqEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ymqEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) YmqSecretKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ymqSecretKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) YmqSecretKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ymqSecretKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YandexProvider) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/yandex yandex} Resource.
func NewYandexProvider(scope constructs.Construct, id *string, config *YandexProviderConfig) YandexProvider {
	_init_.Initialize()

	j := jsiiProxy_YandexProvider{}

	_jsii_.Create(
		"@cdktf/provider-yandex.YandexProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/yandex yandex} Resource.
func NewYandexProvider_Override(y YandexProvider, scope constructs.Construct, id *string, config *YandexProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.YandexProvider",
		[]interface{}{scope, id, config},
		y,
	)
}

func (j *jsiiProxy_YandexProvider) SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_YandexProvider) SetCloudId(val *string) {
	_jsii_.Set(
		j,
		"cloudId",
		val,
	)
}

func (j *jsiiProxy_YandexProvider) SetEndpoint(val *string) {
	_jsii_.Set(
		j,
		"endpoint",
		val,
	)
}

func (j *jsiiProxy_YandexProvider) SetFolderId(val *string) {
	_jsii_.Set(
		j,
		"folderId",
		val,
	)
}

func (j *jsiiProxy_YandexProvider) SetInsecure(val interface{}) {
	_jsii_.Set(
		j,
		"insecure",
		val,
	)
}

func (j *jsiiProxy_YandexProvider) SetMaxRetries(val *float64) {
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_YandexProvider) SetOrganizationId(val *string) {
	_jsii_.Set(
		j,
		"organizationId",
		val,
	)
}

func (j *jsiiProxy_YandexProvider) SetPlaintext(val interface{}) {
	_jsii_.Set(
		j,
		"plaintext",
		val,
	)
}

func (j *jsiiProxy_YandexProvider) SetRegionId(val *string) {
	_jsii_.Set(
		j,
		"regionId",
		val,
	)
}

func (j *jsiiProxy_YandexProvider) SetServiceAccountKeyFile(val *string) {
	_jsii_.Set(
		j,
		"serviceAccountKeyFile",
		val,
	)
}

func (j *jsiiProxy_YandexProvider) SetStorageAccessKey(val *string) {
	_jsii_.Set(
		j,
		"storageAccessKey",
		val,
	)
}

func (j *jsiiProxy_YandexProvider) SetStorageEndpoint(val *string) {
	_jsii_.Set(
		j,
		"storageEndpoint",
		val,
	)
}

func (j *jsiiProxy_YandexProvider) SetStorageSecretKey(val *string) {
	_jsii_.Set(
		j,
		"storageSecretKey",
		val,
	)
}

func (j *jsiiProxy_YandexProvider) SetToken(val *string) {
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

func (j *jsiiProxy_YandexProvider) SetYmqAccessKey(val *string) {
	_jsii_.Set(
		j,
		"ymqAccessKey",
		val,
	)
}

func (j *jsiiProxy_YandexProvider) SetYmqEndpoint(val *string) {
	_jsii_.Set(
		j,
		"ymqEndpoint",
		val,
	)
}

func (j *jsiiProxy_YandexProvider) SetYmqSecretKey(val *string) {
	_jsii_.Set(
		j,
		"ymqSecretKey",
		val,
	)
}

func (j *jsiiProxy_YandexProvider) SetZone(val *string) {
	_jsii_.Set(
		j,
		"zone",
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
func YandexProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-yandex.YandexProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func YandexProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-yandex.YandexProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (y *jsiiProxy_YandexProvider) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		y,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (y *jsiiProxy_YandexProvider) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		y,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (y *jsiiProxy_YandexProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		y,
		"resetAlias",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YandexProvider) ResetCloudId() {
	_jsii_.InvokeVoid(
		y,
		"resetCloudId",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YandexProvider) ResetEndpoint() {
	_jsii_.InvokeVoid(
		y,
		"resetEndpoint",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YandexProvider) ResetFolderId() {
	_jsii_.InvokeVoid(
		y,
		"resetFolderId",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YandexProvider) ResetInsecure() {
	_jsii_.InvokeVoid(
		y,
		"resetInsecure",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YandexProvider) ResetMaxRetries() {
	_jsii_.InvokeVoid(
		y,
		"resetMaxRetries",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YandexProvider) ResetOrganizationId() {
	_jsii_.InvokeVoid(
		y,
		"resetOrganizationId",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YandexProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		y,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YandexProvider) ResetPlaintext() {
	_jsii_.InvokeVoid(
		y,
		"resetPlaintext",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YandexProvider) ResetRegionId() {
	_jsii_.InvokeVoid(
		y,
		"resetRegionId",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YandexProvider) ResetServiceAccountKeyFile() {
	_jsii_.InvokeVoid(
		y,
		"resetServiceAccountKeyFile",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YandexProvider) ResetStorageAccessKey() {
	_jsii_.InvokeVoid(
		y,
		"resetStorageAccessKey",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YandexProvider) ResetStorageEndpoint() {
	_jsii_.InvokeVoid(
		y,
		"resetStorageEndpoint",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YandexProvider) ResetStorageSecretKey() {
	_jsii_.InvokeVoid(
		y,
		"resetStorageSecretKey",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YandexProvider) ResetToken() {
	_jsii_.InvokeVoid(
		y,
		"resetToken",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YandexProvider) ResetYmqAccessKey() {
	_jsii_.InvokeVoid(
		y,
		"resetYmqAccessKey",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YandexProvider) ResetYmqEndpoint() {
	_jsii_.InvokeVoid(
		y,
		"resetYmqEndpoint",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YandexProvider) ResetYmqSecretKey() {
	_jsii_.InvokeVoid(
		y,
		"resetYmqSecretKey",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YandexProvider) ResetZone() {
	_jsii_.InvokeVoid(
		y,
		"resetZone",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YandexProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		y,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YandexProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		y,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YandexProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		y,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YandexProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		y,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

