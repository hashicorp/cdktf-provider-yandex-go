// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster yandex_kubernetes_cluster}.
type KubernetesCluster interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterIpv4Range() *string
	SetClusterIpv4Range(val *string)
	ClusterIpv4RangeInput() *string
	ClusterIpv6Range() *string
	SetClusterIpv6Range(val *string)
	ClusterIpv6RangeInput() *string
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
	Health() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	KmsProvider() KubernetesClusterKmsProviderOutputReference
	KmsProviderInput() *KubernetesClusterKmsProvider
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogGroupId() *string
	Master() KubernetesClusterMasterOutputReference
	MasterInput() *KubernetesClusterMaster
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkId() *string
	SetNetworkId(val *string)
	NetworkIdInput() *string
	NetworkImplementation() KubernetesClusterNetworkImplementationOutputReference
	NetworkImplementationInput() *KubernetesClusterNetworkImplementation
	NetworkPolicyProvider() *string
	SetNetworkPolicyProvider(val *string)
	NetworkPolicyProviderInput() *string
	// The tree node.
	Node() constructs.Node
	NodeIpv4CidrMaskSize() *float64
	SetNodeIpv4CidrMaskSize(val *float64)
	NodeIpv4CidrMaskSizeInput() *float64
	NodeServiceAccountId() *string
	SetNodeServiceAccountId(val *string)
	NodeServiceAccountIdInput() *string
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
	ReleaseChannel() *string
	SetReleaseChannel(val *string)
	ReleaseChannelInput() *string
	ServiceAccountId() *string
	SetServiceAccountId(val *string)
	ServiceAccountIdInput() *string
	ServiceIpv4Range() *string
	SetServiceIpv4Range(val *string)
	ServiceIpv4RangeInput() *string
	ServiceIpv6Range() *string
	SetServiceIpv6Range(val *string)
	ServiceIpv6RangeInput() *string
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() KubernetesClusterTimeoutsOutputReference
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
	PutKmsProvider(value *KubernetesClusterKmsProvider)
	PutMaster(value *KubernetesClusterMaster)
	PutNetworkImplementation(value *KubernetesClusterNetworkImplementation)
	PutTimeouts(value *KubernetesClusterTimeouts)
	ResetClusterIpv4Range()
	ResetClusterIpv6Range()
	ResetDescription()
	ResetFolderId()
	ResetId()
	ResetKmsProvider()
	ResetLabels()
	ResetName()
	ResetNetworkImplementation()
	ResetNetworkPolicyProvider()
	ResetNodeIpv4CidrMaskSize()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReleaseChannel()
	ResetServiceIpv4Range()
	ResetServiceIpv6Range()
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

// The jsii proxy struct for KubernetesCluster
type jsiiProxy_KubernetesCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_KubernetesCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ClusterIpv4Range() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIpv4Range",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ClusterIpv4RangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIpv4RangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ClusterIpv6Range() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIpv6Range",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ClusterIpv6RangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIpv6RangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) FolderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) FolderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Health() *string {
	var returns *string
	_jsii_.Get(
		j,
		"health",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) KmsProvider() KubernetesClusterKmsProviderOutputReference {
	var returns KubernetesClusterKmsProviderOutputReference
	_jsii_.Get(
		j,
		"kmsProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) KmsProviderInput() *KubernetesClusterKmsProvider {
	var returns *KubernetesClusterKmsProvider
	_jsii_.Get(
		j,
		"kmsProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) LogGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Master() KubernetesClusterMasterOutputReference {
	var returns KubernetesClusterMasterOutputReference
	_jsii_.Get(
		j,
		"master",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) MasterInput() *KubernetesClusterMaster {
	var returns *KubernetesClusterMaster
	_jsii_.Get(
		j,
		"masterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) NetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) NetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) NetworkImplementation() KubernetesClusterNetworkImplementationOutputReference {
	var returns KubernetesClusterNetworkImplementationOutputReference
	_jsii_.Get(
		j,
		"networkImplementation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) NetworkImplementationInput() *KubernetesClusterNetworkImplementation {
	var returns *KubernetesClusterNetworkImplementation
	_jsii_.Get(
		j,
		"networkImplementationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) NetworkPolicyProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPolicyProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) NetworkPolicyProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPolicyProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) NodeIpv4CidrMaskSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeIpv4CidrMaskSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) NodeIpv4CidrMaskSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeIpv4CidrMaskSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) NodeServiceAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeServiceAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) NodeServiceAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeServiceAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ReleaseChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ReleaseChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ServiceAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ServiceAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ServiceIpv4Range() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceIpv4Range",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ServiceIpv4RangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceIpv4RangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ServiceIpv6Range() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceIpv6Range",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ServiceIpv6RangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceIpv6RangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Timeouts() KubernetesClusterTimeoutsOutputReference {
	var returns KubernetesClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster yandex_kubernetes_cluster} Resource.
func NewKubernetesCluster(scope constructs.Construct, id *string, config *KubernetesClusterConfig) KubernetesCluster {
	_init_.Initialize()

	j := jsiiProxy_KubernetesCluster{}

	_jsii_.Create(
		"@cdktf/provider-yandex.KubernetesCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/kubernetes_cluster yandex_kubernetes_cluster} Resource.
func NewKubernetesCluster_Override(k KubernetesCluster, scope constructs.Construct, id *string, config *KubernetesClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.KubernetesCluster",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetClusterIpv4Range(val *string) {
	_jsii_.Set(
		j,
		"clusterIpv4Range",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetClusterIpv6Range(val *string) {
	_jsii_.Set(
		j,
		"clusterIpv6Range",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetFolderId(val *string) {
	_jsii_.Set(
		j,
		"folderId",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetLabels(val *map[string]*string) {
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetNetworkId(val *string) {
	_jsii_.Set(
		j,
		"networkId",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetNetworkPolicyProvider(val *string) {
	_jsii_.Set(
		j,
		"networkPolicyProvider",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetNodeIpv4CidrMaskSize(val *float64) {
	_jsii_.Set(
		j,
		"nodeIpv4CidrMaskSize",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetNodeServiceAccountId(val *string) {
	_jsii_.Set(
		j,
		"nodeServiceAccountId",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetReleaseChannel(val *string) {
	_jsii_.Set(
		j,
		"releaseChannel",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetServiceAccountId(val *string) {
	_jsii_.Set(
		j,
		"serviceAccountId",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetServiceIpv4Range(val *string) {
	_jsii_.Set(
		j,
		"serviceIpv4Range",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetServiceIpv6Range(val *string) {
	_jsii_.Set(
		j,
		"serviceIpv6Range",
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
func KubernetesCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-yandex.KubernetesCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KubernetesCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-yandex.KubernetesCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KubernetesCluster) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KubernetesCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutKmsProvider(value *KubernetesClusterKmsProvider) {
	_jsii_.InvokeVoid(
		k,
		"putKmsProvider",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutMaster(value *KubernetesClusterMaster) {
	_jsii_.InvokeVoid(
		k,
		"putMaster",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutNetworkImplementation(value *KubernetesClusterNetworkImplementation) {
	_jsii_.InvokeVoid(
		k,
		"putNetworkImplementation",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutTimeouts(value *KubernetesClusterTimeouts) {
	_jsii_.InvokeVoid(
		k,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetClusterIpv4Range() {
	_jsii_.InvokeVoid(
		k,
		"resetClusterIpv4Range",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetClusterIpv6Range() {
	_jsii_.InvokeVoid(
		k,
		"resetClusterIpv6Range",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetDescription() {
	_jsii_.InvokeVoid(
		k,
		"resetDescription",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetFolderId() {
	_jsii_.InvokeVoid(
		k,
		"resetFolderId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetKmsProvider() {
	_jsii_.InvokeVoid(
		k,
		"resetKmsProvider",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetLabels() {
	_jsii_.InvokeVoid(
		k,
		"resetLabels",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetName() {
	_jsii_.InvokeVoid(
		k,
		"resetName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetNetworkImplementation() {
	_jsii_.InvokeVoid(
		k,
		"resetNetworkImplementation",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetNetworkPolicyProvider() {
	_jsii_.InvokeVoid(
		k,
		"resetNetworkPolicyProvider",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetNodeIpv4CidrMaskSize() {
	_jsii_.InvokeVoid(
		k,
		"resetNodeIpv4CidrMaskSize",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetReleaseChannel() {
	_jsii_.InvokeVoid(
		k,
		"resetReleaseChannel",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetServiceIpv4Range() {
	_jsii_.InvokeVoid(
		k,
		"resetServiceIpv4Range",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetServiceIpv6Range() {
	_jsii_.InvokeVoid(
		k,
		"resetServiceIpv6Range",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		k,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

