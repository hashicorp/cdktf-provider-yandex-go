// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group yandex_kubernetes_node_group}.
type KubernetesNodeGroup interface {
	cdktf.TerraformResource
	AllocationPolicy() KubernetesNodeGroupAllocationPolicyOutputReference
	AllocationPolicyInput() *KubernetesNodeGroupAllocationPolicy
	AllowedUnsafeSysctls() *[]*string
	SetAllowedUnsafeSysctls(val *[]*string)
	AllowedUnsafeSysctlsInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
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
	DeployPolicy() KubernetesNodeGroupDeployPolicyOutputReference
	DeployPolicyInput() *KubernetesNodeGroupDeployPolicy
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
	InstanceGroupId() *string
	InstanceTemplate() KubernetesNodeGroupInstanceTemplateOutputReference
	InstanceTemplateInput() *KubernetesNodeGroupInstanceTemplate
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaintenancePolicy() KubernetesNodeGroupMaintenancePolicyOutputReference
	MaintenancePolicyInput() *KubernetesNodeGroupMaintenancePolicy
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NodeLabels() *map[string]*string
	SetNodeLabels(val *map[string]*string)
	NodeLabelsInput() *map[string]*string
	NodeTaints() *[]*string
	SetNodeTaints(val *[]*string)
	NodeTaintsInput() *[]*string
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
	ScalePolicy() KubernetesNodeGroupScalePolicyOutputReference
	ScalePolicyInput() *KubernetesNodeGroupScalePolicy
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() KubernetesNodeGroupTimeoutsOutputReference
	TimeoutsInput() interface{}
	Version() *string
	SetVersion(val *string)
	VersionInfo() KubernetesNodeGroupVersionInfoList
	VersionInput() *string
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
	PutAllocationPolicy(value *KubernetesNodeGroupAllocationPolicy)
	PutDeployPolicy(value *KubernetesNodeGroupDeployPolicy)
	PutInstanceTemplate(value *KubernetesNodeGroupInstanceTemplate)
	PutMaintenancePolicy(value *KubernetesNodeGroupMaintenancePolicy)
	PutScalePolicy(value *KubernetesNodeGroupScalePolicy)
	PutTimeouts(value *KubernetesNodeGroupTimeouts)
	ResetAllocationPolicy()
	ResetAllowedUnsafeSysctls()
	ResetDeployPolicy()
	ResetDescription()
	ResetId()
	ResetLabels()
	ResetMaintenancePolicy()
	ResetName()
	ResetNodeLabels()
	ResetNodeTaints()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	ResetVersion()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for KubernetesNodeGroup
type jsiiProxy_KubernetesNodeGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_KubernetesNodeGroup) AllocationPolicy() KubernetesNodeGroupAllocationPolicyOutputReference {
	var returns KubernetesNodeGroupAllocationPolicyOutputReference
	_jsii_.Get(
		j,
		"allocationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) AllocationPolicyInput() *KubernetesNodeGroupAllocationPolicy {
	var returns *KubernetesNodeGroupAllocationPolicy
	_jsii_.Get(
		j,
		"allocationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) AllowedUnsafeSysctls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUnsafeSysctls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) AllowedUnsafeSysctlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUnsafeSysctlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) DeployPolicy() KubernetesNodeGroupDeployPolicyOutputReference {
	var returns KubernetesNodeGroupDeployPolicyOutputReference
	_jsii_.Get(
		j,
		"deployPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) DeployPolicyInput() *KubernetesNodeGroupDeployPolicy {
	var returns *KubernetesNodeGroupDeployPolicy
	_jsii_.Get(
		j,
		"deployPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) InstanceGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) InstanceTemplate() KubernetesNodeGroupInstanceTemplateOutputReference {
	var returns KubernetesNodeGroupInstanceTemplateOutputReference
	_jsii_.Get(
		j,
		"instanceTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) InstanceTemplateInput() *KubernetesNodeGroupInstanceTemplate {
	var returns *KubernetesNodeGroupInstanceTemplate
	_jsii_.Get(
		j,
		"instanceTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) MaintenancePolicy() KubernetesNodeGroupMaintenancePolicyOutputReference {
	var returns KubernetesNodeGroupMaintenancePolicyOutputReference
	_jsii_.Get(
		j,
		"maintenancePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) MaintenancePolicyInput() *KubernetesNodeGroupMaintenancePolicy {
	var returns *KubernetesNodeGroupMaintenancePolicy
	_jsii_.Get(
		j,
		"maintenancePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) NodeLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodeLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) NodeLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodeLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) NodeTaints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nodeTaints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) NodeTaintsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nodeTaintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) ScalePolicy() KubernetesNodeGroupScalePolicyOutputReference {
	var returns KubernetesNodeGroupScalePolicyOutputReference
	_jsii_.Get(
		j,
		"scalePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) ScalePolicyInput() *KubernetesNodeGroupScalePolicy {
	var returns *KubernetesNodeGroupScalePolicy
	_jsii_.Get(
		j,
		"scalePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) Timeouts() KubernetesNodeGroupTimeoutsOutputReference {
	var returns KubernetesNodeGroupTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) VersionInfo() KubernetesNodeGroupVersionInfoList {
	var returns KubernetesNodeGroupVersionInfoList
	_jsii_.Get(
		j,
		"versionInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroup) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group yandex_kubernetes_node_group} Resource.
func NewKubernetesNodeGroup(scope constructs.Construct, id *string, config *KubernetesNodeGroupConfig) KubernetesNodeGroup {
	_init_.Initialize()

	j := jsiiProxy_KubernetesNodeGroup{}

	_jsii_.Create(
		"@cdktf/provider-yandex.KubernetesNodeGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_node_group yandex_kubernetes_node_group} Resource.
func NewKubernetesNodeGroup_Override(k KubernetesNodeGroup, scope constructs.Construct, id *string, config *KubernetesNodeGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.KubernetesNodeGroup",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KubernetesNodeGroup) SetAllowedUnsafeSysctls(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedUnsafeSysctls",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroup) SetClusterId(val *string) {
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroup) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroup) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroup) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroup) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroup) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroup) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroup) SetLabels(val *map[string]*string) {
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroup) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroup) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroup) SetNodeLabels(val *map[string]*string) {
	_jsii_.Set(
		j,
		"nodeLabels",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroup) SetNodeTaints(val *[]*string) {
	_jsii_.Set(
		j,
		"nodeTaints",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroup) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroup) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroup) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
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
func KubernetesNodeGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-yandex.KubernetesNodeGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KubernetesNodeGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-yandex.KubernetesNodeGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KubernetesNodeGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KubernetesNodeGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroup) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroup) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KubernetesNodeGroup) PutAllocationPolicy(value *KubernetesNodeGroupAllocationPolicy) {
	_jsii_.InvokeVoid(
		k,
		"putAllocationPolicy",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesNodeGroup) PutDeployPolicy(value *KubernetesNodeGroupDeployPolicy) {
	_jsii_.InvokeVoid(
		k,
		"putDeployPolicy",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesNodeGroup) PutInstanceTemplate(value *KubernetesNodeGroupInstanceTemplate) {
	_jsii_.InvokeVoid(
		k,
		"putInstanceTemplate",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesNodeGroup) PutMaintenancePolicy(value *KubernetesNodeGroupMaintenancePolicy) {
	_jsii_.InvokeVoid(
		k,
		"putMaintenancePolicy",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesNodeGroup) PutScalePolicy(value *KubernetesNodeGroupScalePolicy) {
	_jsii_.InvokeVoid(
		k,
		"putScalePolicy",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesNodeGroup) PutTimeouts(value *KubernetesNodeGroupTimeouts) {
	_jsii_.InvokeVoid(
		k,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesNodeGroup) ResetAllocationPolicy() {
	_jsii_.InvokeVoid(
		k,
		"resetAllocationPolicy",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroup) ResetAllowedUnsafeSysctls() {
	_jsii_.InvokeVoid(
		k,
		"resetAllowedUnsafeSysctls",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroup) ResetDeployPolicy() {
	_jsii_.InvokeVoid(
		k,
		"resetDeployPolicy",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroup) ResetDescription() {
	_jsii_.InvokeVoid(
		k,
		"resetDescription",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroup) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroup) ResetLabels() {
	_jsii_.InvokeVoid(
		k,
		"resetLabels",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroup) ResetMaintenancePolicy() {
	_jsii_.InvokeVoid(
		k,
		"resetMaintenancePolicy",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroup) ResetName() {
	_jsii_.InvokeVoid(
		k,
		"resetName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroup) ResetNodeLabels() {
	_jsii_.InvokeVoid(
		k,
		"resetNodeLabels",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroup) ResetNodeTaints() {
	_jsii_.InvokeVoid(
		k,
		"resetNodeTaints",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroup) ResetTimeouts() {
	_jsii_.InvokeVoid(
		k,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroup) ResetVersion() {
	_jsii_.InvokeVoid(
		k,
		"resetVersion",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

