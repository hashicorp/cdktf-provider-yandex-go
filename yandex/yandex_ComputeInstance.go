// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance yandex_compute_instance}.
type ComputeInstance interface {
	cdktf.TerraformResource
	AllowRecreate() interface{}
	SetAllowRecreate(val interface{})
	AllowRecreateInput() interface{}
	AllowStoppingForUpdate() interface{}
	SetAllowStoppingForUpdate(val interface{})
	AllowStoppingForUpdateInput() interface{}
	BootDisk() ComputeInstanceBootDiskOutputReference
	BootDiskInput() *ComputeInstanceBootDisk
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
	Fqdn() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Hostname() *string
	SetHostname(val *string)
	HostnameInput() *string
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
	LocalDisk() ComputeInstanceLocalDiskList
	LocalDiskInput() interface{}
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataInput() *map[string]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkAccelerationType() *string
	SetNetworkAccelerationType(val *string)
	NetworkAccelerationTypeInput() *string
	NetworkInterface() ComputeInstanceNetworkInterfaceList
	NetworkInterfaceInput() interface{}
	// The tree node.
	Node() constructs.Node
	PlacementPolicy() ComputeInstancePlacementPolicyOutputReference
	PlacementPolicyInput() *ComputeInstancePlacementPolicy
	PlatformId() *string
	SetPlatformId(val *string)
	PlatformIdInput() *string
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
	Resources() ComputeInstanceResourcesOutputReference
	ResourcesInput() *ComputeInstanceResources
	SchedulingPolicy() ComputeInstanceSchedulingPolicyOutputReference
	SchedulingPolicyInput() *ComputeInstanceSchedulingPolicy
	SecondaryDisk() ComputeInstanceSecondaryDiskList
	SecondaryDiskInput() interface{}
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
	Timeouts() ComputeInstanceTimeoutsOutputReference
	TimeoutsInput() interface{}
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
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
	PutBootDisk(value *ComputeInstanceBootDisk)
	PutLocalDisk(value interface{})
	PutNetworkInterface(value interface{})
	PutPlacementPolicy(value *ComputeInstancePlacementPolicy)
	PutResources(value *ComputeInstanceResources)
	PutSchedulingPolicy(value *ComputeInstanceSchedulingPolicy)
	PutSecondaryDisk(value interface{})
	PutTimeouts(value *ComputeInstanceTimeouts)
	ResetAllowRecreate()
	ResetAllowStoppingForUpdate()
	ResetDescription()
	ResetFolderId()
	ResetHostname()
	ResetId()
	ResetLabels()
	ResetLocalDisk()
	ResetMetadata()
	ResetName()
	ResetNetworkAccelerationType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlacementPolicy()
	ResetPlatformId()
	ResetSchedulingPolicy()
	ResetSecondaryDisk()
	ResetServiceAccountId()
	ResetTimeouts()
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

// The jsii proxy struct for ComputeInstance
type jsiiProxy_ComputeInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ComputeInstance) AllowRecreate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowRecreate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) AllowRecreateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowRecreateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) AllowStoppingForUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowStoppingForUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) AllowStoppingForUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowStoppingForUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) BootDisk() ComputeInstanceBootDiskOutputReference {
	var returns ComputeInstanceBootDiskOutputReference
	_jsii_.Get(
		j,
		"bootDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) BootDiskInput() *ComputeInstanceBootDisk {
	var returns *ComputeInstanceBootDisk
	_jsii_.Get(
		j,
		"bootDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) FolderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) FolderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Fqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) LocalDisk() ComputeInstanceLocalDiskList {
	var returns ComputeInstanceLocalDiskList
	_jsii_.Get(
		j,
		"localDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) LocalDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) NetworkAccelerationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkAccelerationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) NetworkAccelerationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkAccelerationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) NetworkInterface() ComputeInstanceNetworkInterfaceList {
	var returns ComputeInstanceNetworkInterfaceList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) NetworkInterfaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) PlacementPolicy() ComputeInstancePlacementPolicyOutputReference {
	var returns ComputeInstancePlacementPolicyOutputReference
	_jsii_.Get(
		j,
		"placementPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) PlacementPolicyInput() *ComputeInstancePlacementPolicy {
	var returns *ComputeInstancePlacementPolicy
	_jsii_.Get(
		j,
		"placementPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) PlatformId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) PlatformIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Resources() ComputeInstanceResourcesOutputReference {
	var returns ComputeInstanceResourcesOutputReference
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) ResourcesInput() *ComputeInstanceResources {
	var returns *ComputeInstanceResources
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) SchedulingPolicy() ComputeInstanceSchedulingPolicyOutputReference {
	var returns ComputeInstanceSchedulingPolicyOutputReference
	_jsii_.Get(
		j,
		"schedulingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) SchedulingPolicyInput() *ComputeInstanceSchedulingPolicy {
	var returns *ComputeInstanceSchedulingPolicy
	_jsii_.Get(
		j,
		"schedulingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) SecondaryDisk() ComputeInstanceSecondaryDiskList {
	var returns ComputeInstanceSecondaryDiskList
	_jsii_.Get(
		j,
		"secondaryDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) SecondaryDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) ServiceAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) ServiceAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Timeouts() ComputeInstanceTimeoutsOutputReference {
	var returns ComputeInstanceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance yandex_compute_instance} Resource.
func NewComputeInstance(scope constructs.Construct, id *string, config *ComputeInstanceConfig) ComputeInstance {
	_init_.Initialize()

	j := jsiiProxy_ComputeInstance{}

	_jsii_.Create(
		"@cdktf/provider-yandex.ComputeInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance yandex_compute_instance} Resource.
func NewComputeInstance_Override(c ComputeInstance, scope constructs.Construct, id *string, config *ComputeInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.ComputeInstance",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ComputeInstance) SetAllowRecreate(val interface{}) {
	_jsii_.Set(
		j,
		"allowRecreate",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance) SetAllowStoppingForUpdate(val interface{}) {
	_jsii_.Set(
		j,
		"allowStoppingForUpdate",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance) SetFolderId(val *string) {
	_jsii_.Set(
		j,
		"folderId",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance) SetHostname(val *string) {
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance) SetLabels(val *map[string]*string) {
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance) SetMetadata(val *map[string]*string) {
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance) SetNetworkAccelerationType(val *string) {
	_jsii_.Set(
		j,
		"networkAccelerationType",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance) SetPlatformId(val *string) {
	_jsii_.Set(
		j,
		"platformId",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance) SetServiceAccountId(val *string) {
	_jsii_.Set(
		j,
		"serviceAccountId",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance) SetZone(val *string) {
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
func ComputeInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-yandex.ComputeInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ComputeInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-yandex.ComputeInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ComputeInstance) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ComputeInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstance) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstance) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstance) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ComputeInstance) PutBootDisk(value *ComputeInstanceBootDisk) {
	_jsii_.InvokeVoid(
		c,
		"putBootDisk",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstance) PutLocalDisk(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putLocalDisk",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstance) PutNetworkInterface(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putNetworkInterface",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstance) PutPlacementPolicy(value *ComputeInstancePlacementPolicy) {
	_jsii_.InvokeVoid(
		c,
		"putPlacementPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstance) PutResources(value *ComputeInstanceResources) {
	_jsii_.InvokeVoid(
		c,
		"putResources",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstance) PutSchedulingPolicy(value *ComputeInstanceSchedulingPolicy) {
	_jsii_.InvokeVoid(
		c,
		"putSchedulingPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstance) PutSecondaryDisk(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putSecondaryDisk",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstance) PutTimeouts(value *ComputeInstanceTimeouts) {
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstance) ResetAllowRecreate() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowRecreate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetAllowStoppingForUpdate() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowStoppingForUpdate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetFolderId() {
	_jsii_.InvokeVoid(
		c,
		"resetFolderId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetHostname() {
	_jsii_.InvokeVoid(
		c,
		"resetHostname",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetLabels() {
	_jsii_.InvokeVoid(
		c,
		"resetLabels",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetLocalDisk() {
	_jsii_.InvokeVoid(
		c,
		"resetLocalDisk",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetMetadata() {
	_jsii_.InvokeVoid(
		c,
		"resetMetadata",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetName() {
	_jsii_.InvokeVoid(
		c,
		"resetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetNetworkAccelerationType() {
	_jsii_.InvokeVoid(
		c,
		"resetNetworkAccelerationType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetPlacementPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetPlacementPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetPlatformId() {
	_jsii_.InvokeVoid(
		c,
		"resetPlatformId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetSchedulingPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetSchedulingPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetSecondaryDisk() {
	_jsii_.InvokeVoid(
		c,
		"resetSecondaryDisk",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetServiceAccountId() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceAccountId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetZone() {
	_jsii_.InvokeVoid(
		c,
		"resetZone",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

