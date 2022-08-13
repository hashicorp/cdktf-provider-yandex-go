// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/yandex/r/vpc_security_group_rule yandex_vpc_security_group_rule}.
type VpcSecurityGroupRule interface {
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Direction() *string
	SetDirection(val *string)
	DirectionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FromPort() *float64
	SetFromPort(val *float64)
	FromPortInput() *float64
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
	// The tree node.
	Node() constructs.Node
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	PredefinedTarget() *string
	SetPredefinedTarget(val *string)
	PredefinedTargetInput() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
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
	SecurityGroupBinding() *string
	SetSecurityGroupBinding(val *string)
	SecurityGroupBindingInput() *string
	SecurityGroupId() *string
	SetSecurityGroupId(val *string)
	SecurityGroupIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() VpcSecurityGroupRuleTimeoutsOutputReference
	TimeoutsInput() interface{}
	ToPort() *float64
	SetToPort(val *float64)
	ToPortInput() *float64
	V4CidrBlocks() *[]*string
	SetV4CidrBlocks(val *[]*string)
	V4CidrBlocksInput() *[]*string
	V6CidrBlocks() *[]*string
	SetV6CidrBlocks(val *[]*string)
	V6CidrBlocksInput() *[]*string
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
	PutTimeouts(value *VpcSecurityGroupRuleTimeouts)
	ResetDescription()
	ResetFromPort()
	ResetId()
	ResetLabels()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPort()
	ResetPredefinedTarget()
	ResetProtocol()
	ResetSecurityGroupId()
	ResetTimeouts()
	ResetToPort()
	ResetV4CidrBlocks()
	ResetV6CidrBlocks()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for VpcSecurityGroupRule
type jsiiProxy_VpcSecurityGroupRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VpcSecurityGroupRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) DirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) FromPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fromPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) FromPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fromPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) PredefinedTarget() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predefinedTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) PredefinedTargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predefinedTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) SecurityGroupBinding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) SecurityGroupBindingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) SecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) SecurityGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) Timeouts() VpcSecurityGroupRuleTimeoutsOutputReference {
	var returns VpcSecurityGroupRuleTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) ToPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"toPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) ToPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"toPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) V4CidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"v4CidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) V4CidrBlocksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"v4CidrBlocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) V6CidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"v6CidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcSecurityGroupRule) V6CidrBlocksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"v6CidrBlocksInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/vpc_security_group_rule yandex_vpc_security_group_rule} Resource.
func NewVpcSecurityGroupRule(scope constructs.Construct, id *string, config *VpcSecurityGroupRuleConfig) VpcSecurityGroupRule {
	_init_.Initialize()

	j := jsiiProxy_VpcSecurityGroupRule{}

	_jsii_.Create(
		"@cdktf/provider-yandex.VpcSecurityGroupRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/vpc_security_group_rule yandex_vpc_security_group_rule} Resource.
func NewVpcSecurityGroupRule_Override(v VpcSecurityGroupRule, scope constructs.Construct, id *string, config *VpcSecurityGroupRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.VpcSecurityGroupRule",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VpcSecurityGroupRule) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VpcSecurityGroupRule) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VpcSecurityGroupRule) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VpcSecurityGroupRule) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_VpcSecurityGroupRule) SetDirection(val *string) {
	_jsii_.Set(
		j,
		"direction",
		val,
	)
}

func (j *jsiiProxy_VpcSecurityGroupRule) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VpcSecurityGroupRule) SetFromPort(val *float64) {
	_jsii_.Set(
		j,
		"fromPort",
		val,
	)
}

func (j *jsiiProxy_VpcSecurityGroupRule) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VpcSecurityGroupRule) SetLabels(val *map[string]*string) {
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_VpcSecurityGroupRule) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VpcSecurityGroupRule) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_VpcSecurityGroupRule) SetPredefinedTarget(val *string) {
	_jsii_.Set(
		j,
		"predefinedTarget",
		val,
	)
}

func (j *jsiiProxy_VpcSecurityGroupRule) SetProtocol(val *string) {
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_VpcSecurityGroupRule) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VpcSecurityGroupRule) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VpcSecurityGroupRule) SetSecurityGroupBinding(val *string) {
	_jsii_.Set(
		j,
		"securityGroupBinding",
		val,
	)
}

func (j *jsiiProxy_VpcSecurityGroupRule) SetSecurityGroupId(val *string) {
	_jsii_.Set(
		j,
		"securityGroupId",
		val,
	)
}

func (j *jsiiProxy_VpcSecurityGroupRule) SetToPort(val *float64) {
	_jsii_.Set(
		j,
		"toPort",
		val,
	)
}

func (j *jsiiProxy_VpcSecurityGroupRule) SetV4CidrBlocks(val *[]*string) {
	_jsii_.Set(
		j,
		"v4CidrBlocks",
		val,
	)
}

func (j *jsiiProxy_VpcSecurityGroupRule) SetV6CidrBlocks(val *[]*string) {
	_jsii_.Set(
		j,
		"v6CidrBlocks",
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
func VpcSecurityGroupRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-yandex.VpcSecurityGroupRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VpcSecurityGroupRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-yandex.VpcSecurityGroupRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VpcSecurityGroupRule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VpcSecurityGroupRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcSecurityGroupRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcSecurityGroupRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcSecurityGroupRule) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcSecurityGroupRule) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcSecurityGroupRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcSecurityGroupRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcSecurityGroupRule) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcSecurityGroupRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcSecurityGroupRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcSecurityGroupRule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VpcSecurityGroupRule) PutTimeouts(value *VpcSecurityGroupRuleTimeouts) {
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpcSecurityGroupRule) ResetDescription() {
	_jsii_.InvokeVoid(
		v,
		"resetDescription",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcSecurityGroupRule) ResetFromPort() {
	_jsii_.InvokeVoid(
		v,
		"resetFromPort",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcSecurityGroupRule) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcSecurityGroupRule) ResetLabels() {
	_jsii_.InvokeVoid(
		v,
		"resetLabels",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcSecurityGroupRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcSecurityGroupRule) ResetPort() {
	_jsii_.InvokeVoid(
		v,
		"resetPort",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcSecurityGroupRule) ResetPredefinedTarget() {
	_jsii_.InvokeVoid(
		v,
		"resetPredefinedTarget",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcSecurityGroupRule) ResetProtocol() {
	_jsii_.InvokeVoid(
		v,
		"resetProtocol",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcSecurityGroupRule) ResetSecurityGroupId() {
	_jsii_.InvokeVoid(
		v,
		"resetSecurityGroupId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcSecurityGroupRule) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcSecurityGroupRule) ResetToPort() {
	_jsii_.InvokeVoid(
		v,
		"resetToPort",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcSecurityGroupRule) ResetV4CidrBlocks() {
	_jsii_.InvokeVoid(
		v,
		"resetV4CidrBlocks",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcSecurityGroupRule) ResetV6CidrBlocks() {
	_jsii_.InvokeVoid(
		v,
		"resetV6CidrBlocks",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcSecurityGroupRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcSecurityGroupRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcSecurityGroupRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcSecurityGroupRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

