// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group yandex_compute_instance_group}.
type ComputeInstanceGroup interface {
	cdktf.TerraformResource
	AllocationPolicy() ComputeInstanceGroupAllocationPolicyOutputReference
	AllocationPolicyInput() *ComputeInstanceGroupAllocationPolicy
	ApplicationLoadBalancer() ComputeInstanceGroupApplicationLoadBalancerOutputReference
	ApplicationLoadBalancerInput() *ComputeInstanceGroupApplicationLoadBalancer
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
	DeletionProtection() interface{}
	SetDeletionProtection(val interface{})
	DeletionProtectionInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeployPolicy() ComputeInstanceGroupDeployPolicyOutputReference
	DeployPolicyInput() *ComputeInstanceGroupDeployPolicy
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
	HealthCheck() ComputeInstanceGroupHealthCheckList
	HealthCheckInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	Instances() ComputeInstanceGroupInstancesList
	InstanceTemplate() ComputeInstanceGroupInstanceTemplateOutputReference
	InstanceTemplateInput() *ComputeInstanceGroupInstanceTemplate
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancer() ComputeInstanceGroupLoadBalancerOutputReference
	LoadBalancerInput() *ComputeInstanceGroupLoadBalancer
	MaxCheckingHealthDuration() *float64
	SetMaxCheckingHealthDuration(val *float64)
	MaxCheckingHealthDurationInput() *float64
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
	ScalePolicy() ComputeInstanceGroupScalePolicyOutputReference
	ScalePolicyInput() *ComputeInstanceGroupScalePolicy
	ServiceAccountId() *string
	SetServiceAccountId(val *string)
	ServiceAccountIdInput() *string
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ComputeInstanceGroupTimeoutsOutputReference
	TimeoutsInput() interface{}
	Variables() *map[string]*string
	SetVariables(val *map[string]*string)
	VariablesInput() *map[string]*string
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
	PutAllocationPolicy(value *ComputeInstanceGroupAllocationPolicy)
	PutApplicationLoadBalancer(value *ComputeInstanceGroupApplicationLoadBalancer)
	PutDeployPolicy(value *ComputeInstanceGroupDeployPolicy)
	PutHealthCheck(value interface{})
	PutInstanceTemplate(value *ComputeInstanceGroupInstanceTemplate)
	PutLoadBalancer(value *ComputeInstanceGroupLoadBalancer)
	PutScalePolicy(value *ComputeInstanceGroupScalePolicy)
	PutTimeouts(value *ComputeInstanceGroupTimeouts)
	ResetApplicationLoadBalancer()
	ResetDeletionProtection()
	ResetDescription()
	ResetFolderId()
	ResetHealthCheck()
	ResetId()
	ResetLabels()
	ResetLoadBalancer()
	ResetMaxCheckingHealthDuration()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	ResetVariables()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ComputeInstanceGroup
type jsiiProxy_ComputeInstanceGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ComputeInstanceGroup) AllocationPolicy() ComputeInstanceGroupAllocationPolicyOutputReference {
	var returns ComputeInstanceGroupAllocationPolicyOutputReference
	_jsii_.Get(
		j,
		"allocationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) AllocationPolicyInput() *ComputeInstanceGroupAllocationPolicy {
	var returns *ComputeInstanceGroupAllocationPolicy
	_jsii_.Get(
		j,
		"allocationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) ApplicationLoadBalancer() ComputeInstanceGroupApplicationLoadBalancerOutputReference {
	var returns ComputeInstanceGroupApplicationLoadBalancerOutputReference
	_jsii_.Get(
		j,
		"applicationLoadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) ApplicationLoadBalancerInput() *ComputeInstanceGroupApplicationLoadBalancer {
	var returns *ComputeInstanceGroupApplicationLoadBalancer
	_jsii_.Get(
		j,
		"applicationLoadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) DeployPolicy() ComputeInstanceGroupDeployPolicyOutputReference {
	var returns ComputeInstanceGroupDeployPolicyOutputReference
	_jsii_.Get(
		j,
		"deployPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) DeployPolicyInput() *ComputeInstanceGroupDeployPolicy {
	var returns *ComputeInstanceGroupDeployPolicy
	_jsii_.Get(
		j,
		"deployPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) FolderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) FolderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) HealthCheck() ComputeInstanceGroupHealthCheckList {
	var returns ComputeInstanceGroupHealthCheckList
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) HealthCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) Instances() ComputeInstanceGroupInstancesList {
	var returns ComputeInstanceGroupInstancesList
	_jsii_.Get(
		j,
		"instances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) InstanceTemplate() ComputeInstanceGroupInstanceTemplateOutputReference {
	var returns ComputeInstanceGroupInstanceTemplateOutputReference
	_jsii_.Get(
		j,
		"instanceTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) InstanceTemplateInput() *ComputeInstanceGroupInstanceTemplate {
	var returns *ComputeInstanceGroupInstanceTemplate
	_jsii_.Get(
		j,
		"instanceTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) LoadBalancer() ComputeInstanceGroupLoadBalancerOutputReference {
	var returns ComputeInstanceGroupLoadBalancerOutputReference
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) LoadBalancerInput() *ComputeInstanceGroupLoadBalancer {
	var returns *ComputeInstanceGroupLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) MaxCheckingHealthDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCheckingHealthDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) MaxCheckingHealthDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCheckingHealthDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) ScalePolicy() ComputeInstanceGroupScalePolicyOutputReference {
	var returns ComputeInstanceGroupScalePolicyOutputReference
	_jsii_.Get(
		j,
		"scalePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) ScalePolicyInput() *ComputeInstanceGroupScalePolicy {
	var returns *ComputeInstanceGroupScalePolicy
	_jsii_.Get(
		j,
		"scalePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) ServiceAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) ServiceAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) Timeouts() ComputeInstanceGroupTimeoutsOutputReference {
	var returns ComputeInstanceGroupTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) Variables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroup) VariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"variablesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group yandex_compute_instance_group} Resource.
func NewComputeInstanceGroup(scope constructs.Construct, id *string, config *ComputeInstanceGroupConfig) ComputeInstanceGroup {
	_init_.Initialize()

	j := jsiiProxy_ComputeInstanceGroup{}

	_jsii_.Create(
		"@cdktf/provider-yandex.ComputeInstanceGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group yandex_compute_instance_group} Resource.
func NewComputeInstanceGroup_Override(c ComputeInstanceGroup, scope constructs.Construct, id *string, config *ComputeInstanceGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.ComputeInstanceGroup",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ComputeInstanceGroup) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroup) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroup) SetDeletionProtection(val interface{}) {
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroup) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroup) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroup) SetFolderId(val *string) {
	_jsii_.Set(
		j,
		"folderId",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroup) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroup) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroup) SetLabels(val *map[string]*string) {
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroup) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroup) SetMaxCheckingHealthDuration(val *float64) {
	_jsii_.Set(
		j,
		"maxCheckingHealthDuration",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroup) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroup) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroup) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroup) SetServiceAccountId(val *string) {
	_jsii_.Set(
		j,
		"serviceAccountId",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroup) SetVariables(val *map[string]*string) {
	_jsii_.Set(
		j,
		"variables",
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
func ComputeInstanceGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-yandex.ComputeInstanceGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ComputeInstanceGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-yandex.ComputeInstanceGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ComputeInstanceGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroup) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroup) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ComputeInstanceGroup) PutAllocationPolicy(value *ComputeInstanceGroupAllocationPolicy) {
	_jsii_.InvokeVoid(
		c,
		"putAllocationPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroup) PutApplicationLoadBalancer(value *ComputeInstanceGroupApplicationLoadBalancer) {
	_jsii_.InvokeVoid(
		c,
		"putApplicationLoadBalancer",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroup) PutDeployPolicy(value *ComputeInstanceGroupDeployPolicy) {
	_jsii_.InvokeVoid(
		c,
		"putDeployPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroup) PutHealthCheck(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putHealthCheck",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroup) PutInstanceTemplate(value *ComputeInstanceGroupInstanceTemplate) {
	_jsii_.InvokeVoid(
		c,
		"putInstanceTemplate",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroup) PutLoadBalancer(value *ComputeInstanceGroupLoadBalancer) {
	_jsii_.InvokeVoid(
		c,
		"putLoadBalancer",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroup) PutScalePolicy(value *ComputeInstanceGroupScalePolicy) {
	_jsii_.InvokeVoid(
		c,
		"putScalePolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroup) PutTimeouts(value *ComputeInstanceGroupTimeouts) {
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroup) ResetApplicationLoadBalancer() {
	_jsii_.InvokeVoid(
		c,
		"resetApplicationLoadBalancer",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroup) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		c,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroup) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroup) ResetFolderId() {
	_jsii_.InvokeVoid(
		c,
		"resetFolderId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroup) ResetHealthCheck() {
	_jsii_.InvokeVoid(
		c,
		"resetHealthCheck",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroup) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroup) ResetLabels() {
	_jsii_.InvokeVoid(
		c,
		"resetLabels",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroup) ResetLoadBalancer() {
	_jsii_.InvokeVoid(
		c,
		"resetLoadBalancer",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroup) ResetMaxCheckingHealthDuration() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxCheckingHealthDuration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroup) ResetName() {
	_jsii_.InvokeVoid(
		c,
		"resetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroup) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroup) ResetVariables() {
	_jsii_.InvokeVoid(
		c,
		"resetVariables",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

