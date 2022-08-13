// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/yandex/d/mdb_greenplum_cluster yandex_mdb_greenplum_cluster}.
type DataYandexMdbGreenplumCluster interface {
	cdktf.TerraformDataSource
	Access() DataYandexMdbGreenplumClusterAccessList
	AssignPublicIp() cdktf.IResolvable
	BackupWindowStart() DataYandexMdbGreenplumClusterBackupWindowStartList
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	CreatedAt() *string
	DeletionProtection() cdktf.IResolvable
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	Environment() *string
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
	GreenplumConfig() *map[string]*string
	SetGreenplumConfig(val *map[string]*string)
	GreenplumConfigInput() *map[string]*string
	Health() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Labels() cdktf.StringMap
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MasterHostCount() *float64
	MasterHosts() DataYandexMdbGreenplumClusterMasterHostsList
	MasterSubcluster() DataYandexMdbGreenplumClusterMasterSubclusterList
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkId() *string
	// The tree node.
	Node() constructs.Node
	PoolerConfig() DataYandexMdbGreenplumClusterPoolerConfigOutputReference
	PoolerConfigInput() *DataYandexMdbGreenplumClusterPoolerConfig
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	SecurityGroupIds() *[]*string
	SegmentHostCount() *float64
	SegmentHosts() DataYandexMdbGreenplumClusterSegmentHostsList
	SegmentInHost() *float64
	SegmentSubcluster() DataYandexMdbGreenplumClusterSegmentSubclusterList
	Status() *string
	SubnetId() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UserName() *string
	Version() *string
	Zone() *string
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
	PutPoolerConfig(value *DataYandexMdbGreenplumClusterPoolerConfig)
	ResetClusterId()
	ResetFolderId()
	ResetGreenplumConfig()
	ResetId()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPoolerConfig()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataYandexMdbGreenplumCluster
type jsiiProxy_DataYandexMdbGreenplumCluster struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) Access() DataYandexMdbGreenplumClusterAccessList {
	var returns DataYandexMdbGreenplumClusterAccessList
	_jsii_.Get(
		j,
		"access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) AssignPublicIp() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"assignPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) BackupWindowStart() DataYandexMdbGreenplumClusterBackupWindowStartList {
	var returns DataYandexMdbGreenplumClusterBackupWindowStartList
	_jsii_.Get(
		j,
		"backupWindowStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) DeletionProtection() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) Environment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) FolderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) FolderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) GreenplumConfig() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"greenplumConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) GreenplumConfigInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"greenplumConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) Health() *string {
	var returns *string
	_jsii_.Get(
		j,
		"health",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) Labels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) MasterHostCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"masterHostCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) MasterHosts() DataYandexMdbGreenplumClusterMasterHostsList {
	var returns DataYandexMdbGreenplumClusterMasterHostsList
	_jsii_.Get(
		j,
		"masterHosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) MasterSubcluster() DataYandexMdbGreenplumClusterMasterSubclusterList {
	var returns DataYandexMdbGreenplumClusterMasterSubclusterList
	_jsii_.Get(
		j,
		"masterSubcluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) NetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) PoolerConfig() DataYandexMdbGreenplumClusterPoolerConfigOutputReference {
	var returns DataYandexMdbGreenplumClusterPoolerConfigOutputReference
	_jsii_.Get(
		j,
		"poolerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) PoolerConfigInput() *DataYandexMdbGreenplumClusterPoolerConfig {
	var returns *DataYandexMdbGreenplumClusterPoolerConfig
	_jsii_.Get(
		j,
		"poolerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) SegmentHostCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentHostCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) SegmentHosts() DataYandexMdbGreenplumClusterSegmentHostsList {
	var returns DataYandexMdbGreenplumClusterSegmentHostsList
	_jsii_.Get(
		j,
		"segmentHosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) SegmentInHost() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentInHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) SegmentSubcluster() DataYandexMdbGreenplumClusterSegmentSubclusterList {
	var returns DataYandexMdbGreenplumClusterSegmentSubclusterList
	_jsii_.Get(
		j,
		"segmentSubcluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) UserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/yandex/d/mdb_greenplum_cluster yandex_mdb_greenplum_cluster} Data Source.
func NewDataYandexMdbGreenplumCluster(scope constructs.Construct, id *string, config *DataYandexMdbGreenplumClusterConfig) DataYandexMdbGreenplumCluster {
	_init_.Initialize()

	j := jsiiProxy_DataYandexMdbGreenplumCluster{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexMdbGreenplumCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/yandex/d/mdb_greenplum_cluster yandex_mdb_greenplum_cluster} Data Source.
func NewDataYandexMdbGreenplumCluster_Override(d DataYandexMdbGreenplumCluster, scope constructs.Construct, id *string, config *DataYandexMdbGreenplumClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexMdbGreenplumCluster",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) SetClusterId(val *string) {
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) SetFolderId(val *string) {
	_jsii_.Set(
		j,
		"folderId",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) SetGreenplumConfig(val *map[string]*string) {
	_jsii_.Set(
		j,
		"greenplumConfig",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbGreenplumCluster) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
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
func DataYandexMdbGreenplumCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-yandex.DataYandexMdbGreenplumCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataYandexMdbGreenplumCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-yandex.DataYandexMdbGreenplumCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataYandexMdbGreenplumCluster) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataYandexMdbGreenplumCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbGreenplumCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbGreenplumCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbGreenplumCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbGreenplumCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbGreenplumCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbGreenplumCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbGreenplumCluster) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbGreenplumCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbGreenplumCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbGreenplumCluster) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataYandexMdbGreenplumCluster) PutPoolerConfig(value *DataYandexMdbGreenplumClusterPoolerConfig) {
	_jsii_.InvokeVoid(
		d,
		"putPoolerConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexMdbGreenplumCluster) ResetClusterId() {
	_jsii_.InvokeVoid(
		d,
		"resetClusterId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbGreenplumCluster) ResetFolderId() {
	_jsii_.InvokeVoid(
		d,
		"resetFolderId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbGreenplumCluster) ResetGreenplumConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetGreenplumConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbGreenplumCluster) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbGreenplumCluster) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbGreenplumCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbGreenplumCluster) ResetPoolerConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetPoolerConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbGreenplumCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbGreenplumCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbGreenplumCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbGreenplumCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

