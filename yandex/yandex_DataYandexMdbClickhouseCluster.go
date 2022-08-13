// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/yandex/d/mdb_clickhouse_cluster yandex_mdb_clickhouse_cluster}.
type DataYandexMdbClickhouseCluster interface {
	cdktf.TerraformDataSource
	Access() DataYandexMdbClickhouseClusterAccessList
	BackupWindowStart() DataYandexMdbClickhouseClusterBackupWindowStartList
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Clickhouse() DataYandexMdbClickhouseClusterClickhouseList
	CloudStorage() DataYandexMdbClickhouseClusterCloudStorageOutputReference
	CloudStorageInput() *DataYandexMdbClickhouseClusterCloudStorage
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
	Database() DataYandexMdbClickhouseClusterDatabaseList
	DeletionProtection() interface{}
	SetDeletionProtection(val interface{})
	DeletionProtectionInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	EmbeddedKeeper() cdktf.IResolvable
	Environment() *string
	FolderId() *string
	SetFolderId(val *string)
	FolderIdInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	FormatSchema() DataYandexMdbClickhouseClusterFormatSchemaList
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Health() *string
	Host() DataYandexMdbClickhouseClusterHostList
	Id() *string
	SetId(val *string)
	IdInput() *string
	Labels() cdktf.StringMap
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaintenanceWindow() DataYandexMdbClickhouseClusterMaintenanceWindowList
	MlModel() DataYandexMdbClickhouseClusterMlModelList
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkId() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	SecurityGroupIds() *[]*string
	ServiceAccountId() *string
	SetServiceAccountId(val *string)
	ServiceAccountIdInput() *string
	ShardGroup() DataYandexMdbClickhouseClusterShardGroupList
	SqlDatabaseManagement() cdktf.IResolvable
	SqlUserManagement() cdktf.IResolvable
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	User() DataYandexMdbClickhouseClusterUserList
	Version() *string
	Zookeeper() DataYandexMdbClickhouseClusterZookeeperList
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
	PutCloudStorage(value *DataYandexMdbClickhouseClusterCloudStorage)
	ResetCloudStorage()
	ResetClusterId()
	ResetDeletionProtection()
	ResetFolderId()
	ResetId()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetServiceAccountId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataYandexMdbClickhouseCluster
type jsiiProxy_DataYandexMdbClickhouseCluster struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) Access() DataYandexMdbClickhouseClusterAccessList {
	var returns DataYandexMdbClickhouseClusterAccessList
	_jsii_.Get(
		j,
		"access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) BackupWindowStart() DataYandexMdbClickhouseClusterBackupWindowStartList {
	var returns DataYandexMdbClickhouseClusterBackupWindowStartList
	_jsii_.Get(
		j,
		"backupWindowStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) Clickhouse() DataYandexMdbClickhouseClusterClickhouseList {
	var returns DataYandexMdbClickhouseClusterClickhouseList
	_jsii_.Get(
		j,
		"clickhouse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) CloudStorage() DataYandexMdbClickhouseClusterCloudStorageOutputReference {
	var returns DataYandexMdbClickhouseClusterCloudStorageOutputReference
	_jsii_.Get(
		j,
		"cloudStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) CloudStorageInput() *DataYandexMdbClickhouseClusterCloudStorage {
	var returns *DataYandexMdbClickhouseClusterCloudStorage
	_jsii_.Get(
		j,
		"cloudStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) Database() DataYandexMdbClickhouseClusterDatabaseList {
	var returns DataYandexMdbClickhouseClusterDatabaseList
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) EmbeddedKeeper() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"embeddedKeeper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) Environment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) FolderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) FolderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) FormatSchema() DataYandexMdbClickhouseClusterFormatSchemaList {
	var returns DataYandexMdbClickhouseClusterFormatSchemaList
	_jsii_.Get(
		j,
		"formatSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) Health() *string {
	var returns *string
	_jsii_.Get(
		j,
		"health",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) Host() DataYandexMdbClickhouseClusterHostList {
	var returns DataYandexMdbClickhouseClusterHostList
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) Labels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) MaintenanceWindow() DataYandexMdbClickhouseClusterMaintenanceWindowList {
	var returns DataYandexMdbClickhouseClusterMaintenanceWindowList
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) MlModel() DataYandexMdbClickhouseClusterMlModelList {
	var returns DataYandexMdbClickhouseClusterMlModelList
	_jsii_.Get(
		j,
		"mlModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) NetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) ServiceAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) ServiceAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) ShardGroup() DataYandexMdbClickhouseClusterShardGroupList {
	var returns DataYandexMdbClickhouseClusterShardGroupList
	_jsii_.Get(
		j,
		"shardGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) SqlDatabaseManagement() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"sqlDatabaseManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) SqlUserManagement() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"sqlUserManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) User() DataYandexMdbClickhouseClusterUserList {
	var returns DataYandexMdbClickhouseClusterUserList
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) Zookeeper() DataYandexMdbClickhouseClusterZookeeperList {
	var returns DataYandexMdbClickhouseClusterZookeeperList
	_jsii_.Get(
		j,
		"zookeeper",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/yandex/d/mdb_clickhouse_cluster yandex_mdb_clickhouse_cluster} Data Source.
func NewDataYandexMdbClickhouseCluster(scope constructs.Construct, id *string, config *DataYandexMdbClickhouseClusterConfig) DataYandexMdbClickhouseCluster {
	_init_.Initialize()

	j := jsiiProxy_DataYandexMdbClickhouseCluster{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexMdbClickhouseCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/yandex/d/mdb_clickhouse_cluster yandex_mdb_clickhouse_cluster} Data Source.
func NewDataYandexMdbClickhouseCluster_Override(d DataYandexMdbClickhouseCluster, scope constructs.Construct, id *string, config *DataYandexMdbClickhouseClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexMdbClickhouseCluster",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) SetClusterId(val *string) {
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) SetDeletionProtection(val interface{}) {
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) SetFolderId(val *string) {
	_jsii_.Set(
		j,
		"folderId",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbClickhouseCluster) SetServiceAccountId(val *string) {
	_jsii_.Set(
		j,
		"serviceAccountId",
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
func DataYandexMdbClickhouseCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-yandex.DataYandexMdbClickhouseCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataYandexMdbClickhouseCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-yandex.DataYandexMdbClickhouseCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseCluster) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataYandexMdbClickhouseCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseCluster) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseCluster) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataYandexMdbClickhouseCluster) PutCloudStorage(value *DataYandexMdbClickhouseClusterCloudStorage) {
	_jsii_.InvokeVoid(
		d,
		"putCloudStorage",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexMdbClickhouseCluster) ResetCloudStorage() {
	_jsii_.InvokeVoid(
		d,
		"resetCloudStorage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbClickhouseCluster) ResetClusterId() {
	_jsii_.InvokeVoid(
		d,
		"resetClusterId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbClickhouseCluster) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		d,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbClickhouseCluster) ResetFolderId() {
	_jsii_.InvokeVoid(
		d,
		"resetFolderId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbClickhouseCluster) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbClickhouseCluster) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbClickhouseCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbClickhouseCluster) ResetServiceAccountId() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceAccountId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbClickhouseCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbClickhouseCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

