// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/yandex/r/organizationmanager_saml_federation yandex_organizationmanager_saml_federation}.
type OrganizationmanagerSamlFederation interface {
	cdktf.TerraformResource
	AutoCreateAccountOnLogin() interface{}
	SetAutoCreateAccountOnLogin(val interface{})
	AutoCreateAccountOnLoginInput() interface{}
	CaseInsensitiveNameIds() interface{}
	SetCaseInsensitiveNameIds(val interface{})
	CaseInsensitiveNameIdsInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CookieMaxAge() *string
	SetCookieMaxAge(val *string)
	CookieMaxAgeInput() *string
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
	Issuer() *string
	SetIssuer(val *string)
	IssuerInput() *string
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
	OrganizationId() *string
	SetOrganizationId(val *string)
	OrganizationIdInput() *string
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
	SecuritySettings() OrganizationmanagerSamlFederationSecuritySettingsOutputReference
	SecuritySettingsInput() *OrganizationmanagerSamlFederationSecuritySettings
	SsoBinding() *string
	SetSsoBinding(val *string)
	SsoBindingInput() *string
	SsoUrl() *string
	SetSsoUrl(val *string)
	SsoUrlInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() OrganizationmanagerSamlFederationTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutSecuritySettings(value *OrganizationmanagerSamlFederationSecuritySettings)
	PutTimeouts(value *OrganizationmanagerSamlFederationTimeouts)
	ResetAutoCreateAccountOnLogin()
	ResetCaseInsensitiveNameIds()
	ResetCookieMaxAge()
	ResetDescription()
	ResetId()
	ResetLabels()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSecuritySettings()
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

// The jsii proxy struct for OrganizationmanagerSamlFederation
type jsiiProxy_OrganizationmanagerSamlFederation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) AutoCreateAccountOnLogin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoCreateAccountOnLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) AutoCreateAccountOnLoginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoCreateAccountOnLoginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) CaseInsensitiveNameIds() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"caseInsensitiveNameIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) CaseInsensitiveNameIdsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"caseInsensitiveNameIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) CookieMaxAge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cookieMaxAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) CookieMaxAgeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cookieMaxAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) OrganizationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) OrganizationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) SecuritySettings() OrganizationmanagerSamlFederationSecuritySettingsOutputReference {
	var returns OrganizationmanagerSamlFederationSecuritySettingsOutputReference
	_jsii_.Get(
		j,
		"securitySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) SecuritySettingsInput() *OrganizationmanagerSamlFederationSecuritySettings {
	var returns *OrganizationmanagerSamlFederationSecuritySettings
	_jsii_.Get(
		j,
		"securitySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) SsoBinding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) SsoBindingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) SsoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) SsoUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) Timeouts() OrganizationmanagerSamlFederationTimeoutsOutputReference {
	var returns OrganizationmanagerSamlFederationTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/organizationmanager_saml_federation yandex_organizationmanager_saml_federation} Resource.
func NewOrganizationmanagerSamlFederation(scope constructs.Construct, id *string, config *OrganizationmanagerSamlFederationConfig) OrganizationmanagerSamlFederation {
	_init_.Initialize()

	j := jsiiProxy_OrganizationmanagerSamlFederation{}

	_jsii_.Create(
		"@cdktf/provider-yandex.OrganizationmanagerSamlFederation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/organizationmanager_saml_federation yandex_organizationmanager_saml_federation} Resource.
func NewOrganizationmanagerSamlFederation_Override(o OrganizationmanagerSamlFederation, scope constructs.Construct, id *string, config *OrganizationmanagerSamlFederationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.OrganizationmanagerSamlFederation",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) SetAutoCreateAccountOnLogin(val interface{}) {
	_jsii_.Set(
		j,
		"autoCreateAccountOnLogin",
		val,
	)
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) SetCaseInsensitiveNameIds(val interface{}) {
	_jsii_.Set(
		j,
		"caseInsensitiveNameIds",
		val,
	)
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) SetCookieMaxAge(val *string) {
	_jsii_.Set(
		j,
		"cookieMaxAge",
		val,
	)
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) SetIssuer(val *string) {
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) SetLabels(val *map[string]*string) {
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) SetOrganizationId(val *string) {
	_jsii_.Set(
		j,
		"organizationId",
		val,
	)
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) SetSsoBinding(val *string) {
	_jsii_.Set(
		j,
		"ssoBinding",
		val,
	)
}

func (j *jsiiProxy_OrganizationmanagerSamlFederation) SetSsoUrl(val *string) {
	_jsii_.Set(
		j,
		"ssoUrl",
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
func OrganizationmanagerSamlFederation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-yandex.OrganizationmanagerSamlFederation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OrganizationmanagerSamlFederation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-yandex.OrganizationmanagerSamlFederation",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OrganizationmanagerSamlFederation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OrganizationmanagerSamlFederation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationmanagerSamlFederation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationmanagerSamlFederation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationmanagerSamlFederation) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationmanagerSamlFederation) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationmanagerSamlFederation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationmanagerSamlFederation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationmanagerSamlFederation) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationmanagerSamlFederation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationmanagerSamlFederation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationmanagerSamlFederation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OrganizationmanagerSamlFederation) PutSecuritySettings(value *OrganizationmanagerSamlFederationSecuritySettings) {
	_jsii_.InvokeVoid(
		o,
		"putSecuritySettings",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrganizationmanagerSamlFederation) PutTimeouts(value *OrganizationmanagerSamlFederationTimeouts) {
	_jsii_.InvokeVoid(
		o,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrganizationmanagerSamlFederation) ResetAutoCreateAccountOnLogin() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoCreateAccountOnLogin",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationmanagerSamlFederation) ResetCaseInsensitiveNameIds() {
	_jsii_.InvokeVoid(
		o,
		"resetCaseInsensitiveNameIds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationmanagerSamlFederation) ResetCookieMaxAge() {
	_jsii_.InvokeVoid(
		o,
		"resetCookieMaxAge",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationmanagerSamlFederation) ResetDescription() {
	_jsii_.InvokeVoid(
		o,
		"resetDescription",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationmanagerSamlFederation) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationmanagerSamlFederation) ResetLabels() {
	_jsii_.InvokeVoid(
		o,
		"resetLabels",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationmanagerSamlFederation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationmanagerSamlFederation) ResetSecuritySettings() {
	_jsii_.InvokeVoid(
		o,
		"resetSecuritySettings",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationmanagerSamlFederation) ResetTimeouts() {
	_jsii_.InvokeVoid(
		o,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationmanagerSamlFederation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationmanagerSamlFederation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationmanagerSamlFederation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationmanagerSamlFederation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

