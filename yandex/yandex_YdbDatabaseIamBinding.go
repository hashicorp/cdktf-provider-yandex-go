// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/yandex/r/ydb_database_iam_binding yandex_ydb_database_iam_binding}.
type YdbDatabaseIamBinding interface {
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
	DatabaseId() *string
	SetDatabaseId(val *string)
	DatabaseIdInput() *string
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Members() *[]*string
	SetMembers(val *[]*string)
	MembersInput() *[]*string
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
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	SleepAfter() *float64
	SetSleepAfter(val *float64)
	SleepAfterInput() *float64
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
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSleepAfter()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for YdbDatabaseIamBinding
type jsiiProxy_YdbDatabaseIamBinding struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_YdbDatabaseIamBinding) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseIamBinding) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseIamBinding) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseIamBinding) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseIamBinding) DatabaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseIamBinding) DatabaseIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseIamBinding) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseIamBinding) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseIamBinding) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseIamBinding) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseIamBinding) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseIamBinding) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseIamBinding) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseIamBinding) Members() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"members",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseIamBinding) MembersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"membersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseIamBinding) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseIamBinding) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseIamBinding) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseIamBinding) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseIamBinding) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseIamBinding) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseIamBinding) SleepAfter() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sleepAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseIamBinding) SleepAfterInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sleepAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseIamBinding) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseIamBinding) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YdbDatabaseIamBinding) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/ydb_database_iam_binding yandex_ydb_database_iam_binding} Resource.
func NewYdbDatabaseIamBinding(scope constructs.Construct, id *string, config *YdbDatabaseIamBindingConfig) YdbDatabaseIamBinding {
	_init_.Initialize()

	j := jsiiProxy_YdbDatabaseIamBinding{}

	_jsii_.Create(
		"@cdktf/provider-yandex.YdbDatabaseIamBinding",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/ydb_database_iam_binding yandex_ydb_database_iam_binding} Resource.
func NewYdbDatabaseIamBinding_Override(y YdbDatabaseIamBinding, scope constructs.Construct, id *string, config *YdbDatabaseIamBindingConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.YdbDatabaseIamBinding",
		[]interface{}{scope, id, config},
		y,
	)
}

func (j *jsiiProxy_YdbDatabaseIamBinding) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseIamBinding) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseIamBinding) SetDatabaseId(val *string) {
	_jsii_.Set(
		j,
		"databaseId",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseIamBinding) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseIamBinding) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseIamBinding) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseIamBinding) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseIamBinding) SetMembers(val *[]*string) {
	_jsii_.Set(
		j,
		"members",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseIamBinding) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseIamBinding) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseIamBinding) SetRole(val *string) {
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_YdbDatabaseIamBinding) SetSleepAfter(val *float64) {
	_jsii_.Set(
		j,
		"sleepAfter",
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
func YdbDatabaseIamBinding_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-yandex.YdbDatabaseIamBinding",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func YdbDatabaseIamBinding_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-yandex.YdbDatabaseIamBinding",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (y *jsiiProxy_YdbDatabaseIamBinding) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		y,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (y *jsiiProxy_YdbDatabaseIamBinding) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		y,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseIamBinding) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		y,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseIamBinding) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		y,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseIamBinding) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		y,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseIamBinding) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		y,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseIamBinding) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		y,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseIamBinding) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		y,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseIamBinding) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		y,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseIamBinding) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		y,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseIamBinding) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		y,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseIamBinding) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		y,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (y *jsiiProxy_YdbDatabaseIamBinding) ResetId() {
	_jsii_.InvokeVoid(
		y,
		"resetId",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YdbDatabaseIamBinding) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		y,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YdbDatabaseIamBinding) ResetSleepAfter() {
	_jsii_.InvokeVoid(
		y,
		"resetSleepAfter",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YdbDatabaseIamBinding) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		y,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseIamBinding) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		y,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseIamBinding) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		y,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YdbDatabaseIamBinding) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		y,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

