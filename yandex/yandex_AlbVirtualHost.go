// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host yandex_alb_virtual_host}.
type AlbVirtualHost interface {
	cdktf.TerraformResource
	Authority() *[]*string
	SetAuthority(val *[]*string)
	AuthorityInput() *[]*string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpRouterId() *string
	SetHttpRouterId(val *string)
	HttpRouterIdInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ModifyRequestHeaders() AlbVirtualHostModifyRequestHeadersList
	ModifyRequestHeadersInput() interface{}
	ModifyResponseHeaders() AlbVirtualHostModifyResponseHeadersList
	ModifyResponseHeadersInput() interface{}
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
	Route() AlbVirtualHostRouteList
	RouteInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() AlbVirtualHostTimeoutsOutputReference
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
	PutModifyRequestHeaders(value interface{})
	PutModifyResponseHeaders(value interface{})
	PutRoute(value interface{})
	PutTimeouts(value *AlbVirtualHostTimeouts)
	ResetAuthority()
	ResetId()
	ResetModifyRequestHeaders()
	ResetModifyResponseHeaders()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRoute()
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

// The jsii proxy struct for AlbVirtualHost
type jsiiProxy_AlbVirtualHost struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AlbVirtualHost) Authority() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) AuthorityInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) HttpRouterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpRouterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) HttpRouterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpRouterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) ModifyRequestHeaders() AlbVirtualHostModifyRequestHeadersList {
	var returns AlbVirtualHostModifyRequestHeadersList
	_jsii_.Get(
		j,
		"modifyRequestHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) ModifyRequestHeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modifyRequestHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) ModifyResponseHeaders() AlbVirtualHostModifyResponseHeadersList {
	var returns AlbVirtualHostModifyResponseHeadersList
	_jsii_.Get(
		j,
		"modifyResponseHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) ModifyResponseHeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modifyResponseHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) Route() AlbVirtualHostRouteList {
	var returns AlbVirtualHostRouteList
	_jsii_.Get(
		j,
		"route",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) RouteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) Timeouts() AlbVirtualHostTimeoutsOutputReference {
	var returns AlbVirtualHostTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHost) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host yandex_alb_virtual_host} Resource.
func NewAlbVirtualHost(scope constructs.Construct, id *string, config *AlbVirtualHostConfig) AlbVirtualHost {
	_init_.Initialize()

	j := jsiiProxy_AlbVirtualHost{}

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbVirtualHost",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/alb_virtual_host yandex_alb_virtual_host} Resource.
func NewAlbVirtualHost_Override(a AlbVirtualHost, scope constructs.Construct, id *string, config *AlbVirtualHostConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbVirtualHost",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AlbVirtualHost) SetAuthority(val *[]*string) {
	_jsii_.Set(
		j,
		"authority",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHost) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHost) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHost) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHost) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHost) SetHttpRouterId(val *string) {
	_jsii_.Set(
		j,
		"httpRouterId",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHost) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHost) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHost) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHost) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHost) SetProvisioners(val *[]interface{}) {
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
func AlbVirtualHost_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-yandex.AlbVirtualHost",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AlbVirtualHost_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-yandex.AlbVirtualHost",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AlbVirtualHost) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AlbVirtualHost) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHost) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHost) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHost) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHost) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHost) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHost) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHost) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHost) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHost) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHost) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AlbVirtualHost) PutModifyRequestHeaders(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putModifyRequestHeaders",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbVirtualHost) PutModifyResponseHeaders(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putModifyResponseHeaders",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbVirtualHost) PutRoute(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putRoute",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbVirtualHost) PutTimeouts(value *AlbVirtualHostTimeouts) {
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbVirtualHost) ResetAuthority() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthority",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHost) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHost) ResetModifyRequestHeaders() {
	_jsii_.InvokeVoid(
		a,
		"resetModifyRequestHeaders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHost) ResetModifyResponseHeaders() {
	_jsii_.InvokeVoid(
		a,
		"resetModifyResponseHeaders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHost) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHost) ResetRoute() {
	_jsii_.InvokeVoid(
		a,
		"resetRoute",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHost) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHost) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHost) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHost) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHost) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

