// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster yandex_mdb_clickhouse_cluster}.
type MdbClickhouseCluster interface {
	cdktf.TerraformResource
	Access() MdbClickhouseClusterAccessOutputReference
	AccessInput() *MdbClickhouseClusterAccess
	AdminPassword() *string
	SetAdminPassword(val *string)
	AdminPasswordInput() *string
	BackupWindowStart() MdbClickhouseClusterBackupWindowStartOutputReference
	BackupWindowStartInput() *MdbClickhouseClusterBackupWindowStart
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Clickhouse() MdbClickhouseClusterClickhouseOutputReference
	ClickhouseInput() *MdbClickhouseClusterClickhouse
	CloudStorage() MdbClickhouseClusterCloudStorageOutputReference
	CloudStorageInput() *MdbClickhouseClusterCloudStorage
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CopySchemaOnNewHosts() interface{}
	SetCopySchemaOnNewHosts(val interface{})
	CopySchemaOnNewHostsInput() interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	CreatedAt() *string
	Database() MdbClickhouseClusterDatabaseList
	DatabaseInput() interface{}
	DeletionProtection() interface{}
	SetDeletionProtection(val interface{})
	DeletionProtectionInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EmbeddedKeeper() interface{}
	SetEmbeddedKeeper(val interface{})
	EmbeddedKeeperInput() interface{}
	Environment() *string
	SetEnvironment(val *string)
	EnvironmentInput() *string
	FolderId() *string
	SetFolderId(val *string)
	FolderIdInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	FormatSchema() MdbClickhouseClusterFormatSchemaList
	FormatSchemaInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Health() *string
	Host() MdbClickhouseClusterHostList
	HostInput() interface{}
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
	MaintenanceWindow() MdbClickhouseClusterMaintenanceWindowOutputReference
	MaintenanceWindowInput() *MdbClickhouseClusterMaintenanceWindow
	MlModel() MdbClickhouseClusterMlModelList
	MlModelInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkId() *string
	SetNetworkId(val *string)
	NetworkIdInput() *string
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
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	ServiceAccountId() *string
	SetServiceAccountId(val *string)
	ServiceAccountIdInput() *string
	ShardGroup() MdbClickhouseClusterShardGroupList
	ShardGroupInput() interface{}
	SqlDatabaseManagement() interface{}
	SetSqlDatabaseManagement(val interface{})
	SqlDatabaseManagementInput() interface{}
	SqlUserManagement() interface{}
	SetSqlUserManagement(val interface{})
	SqlUserManagementInput() interface{}
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MdbClickhouseClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	User() MdbClickhouseClusterUserList
	UserInput() interface{}
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
	Zookeeper() MdbClickhouseClusterZookeeperOutputReference
	ZookeeperInput() *MdbClickhouseClusterZookeeper
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
	PutAccess(value *MdbClickhouseClusterAccess)
	PutBackupWindowStart(value *MdbClickhouseClusterBackupWindowStart)
	PutClickhouse(value *MdbClickhouseClusterClickhouse)
	PutCloudStorage(value *MdbClickhouseClusterCloudStorage)
	PutDatabase(value interface{})
	PutFormatSchema(value interface{})
	PutHost(value interface{})
	PutMaintenanceWindow(value *MdbClickhouseClusterMaintenanceWindow)
	PutMlModel(value interface{})
	PutShardGroup(value interface{})
	PutTimeouts(value *MdbClickhouseClusterTimeouts)
	PutUser(value interface{})
	PutZookeeper(value *MdbClickhouseClusterZookeeper)
	ResetAccess()
	ResetAdminPassword()
	ResetBackupWindowStart()
	ResetCloudStorage()
	ResetCopySchemaOnNewHosts()
	ResetDatabase()
	ResetDeletionProtection()
	ResetDescription()
	ResetEmbeddedKeeper()
	ResetFolderId()
	ResetFormatSchema()
	ResetId()
	ResetLabels()
	ResetMaintenanceWindow()
	ResetMlModel()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSecurityGroupIds()
	ResetServiceAccountId()
	ResetShardGroup()
	ResetSqlDatabaseManagement()
	ResetSqlUserManagement()
	ResetTimeouts()
	ResetUser()
	ResetVersion()
	ResetZookeeper()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for MdbClickhouseCluster
type jsiiProxy_MdbClickhouseCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MdbClickhouseCluster) Access() MdbClickhouseClusterAccessOutputReference {
	var returns MdbClickhouseClusterAccessOutputReference
	_jsii_.Get(
		j,
		"access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) AccessInput() *MdbClickhouseClusterAccess {
	var returns *MdbClickhouseClusterAccess
	_jsii_.Get(
		j,
		"accessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) AdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) AdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) BackupWindowStart() MdbClickhouseClusterBackupWindowStartOutputReference {
	var returns MdbClickhouseClusterBackupWindowStartOutputReference
	_jsii_.Get(
		j,
		"backupWindowStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) BackupWindowStartInput() *MdbClickhouseClusterBackupWindowStart {
	var returns *MdbClickhouseClusterBackupWindowStart
	_jsii_.Get(
		j,
		"backupWindowStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) Clickhouse() MdbClickhouseClusterClickhouseOutputReference {
	var returns MdbClickhouseClusterClickhouseOutputReference
	_jsii_.Get(
		j,
		"clickhouse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) ClickhouseInput() *MdbClickhouseClusterClickhouse {
	var returns *MdbClickhouseClusterClickhouse
	_jsii_.Get(
		j,
		"clickhouseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) CloudStorage() MdbClickhouseClusterCloudStorageOutputReference {
	var returns MdbClickhouseClusterCloudStorageOutputReference
	_jsii_.Get(
		j,
		"cloudStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) CloudStorageInput() *MdbClickhouseClusterCloudStorage {
	var returns *MdbClickhouseClusterCloudStorage
	_jsii_.Get(
		j,
		"cloudStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) CopySchemaOnNewHosts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copySchemaOnNewHosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) CopySchemaOnNewHostsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copySchemaOnNewHostsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) Database() MdbClickhouseClusterDatabaseList {
	var returns MdbClickhouseClusterDatabaseList
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) DatabaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) EmbeddedKeeper() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"embeddedKeeper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) EmbeddedKeeperInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"embeddedKeeperInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) Environment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) EnvironmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) FolderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) FolderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) FormatSchema() MdbClickhouseClusterFormatSchemaList {
	var returns MdbClickhouseClusterFormatSchemaList
	_jsii_.Get(
		j,
		"formatSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) FormatSchemaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"formatSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) Health() *string {
	var returns *string
	_jsii_.Get(
		j,
		"health",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) Host() MdbClickhouseClusterHostList {
	var returns MdbClickhouseClusterHostList
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) HostInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) MaintenanceWindow() MdbClickhouseClusterMaintenanceWindowOutputReference {
	var returns MdbClickhouseClusterMaintenanceWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) MaintenanceWindowInput() *MdbClickhouseClusterMaintenanceWindow {
	var returns *MdbClickhouseClusterMaintenanceWindow
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) MlModel() MdbClickhouseClusterMlModelList {
	var returns MdbClickhouseClusterMlModelList
	_jsii_.Get(
		j,
		"mlModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) MlModelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mlModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) NetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) NetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) ServiceAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) ServiceAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) ShardGroup() MdbClickhouseClusterShardGroupList {
	var returns MdbClickhouseClusterShardGroupList
	_jsii_.Get(
		j,
		"shardGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) ShardGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shardGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) SqlDatabaseManagement() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlDatabaseManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) SqlDatabaseManagementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlDatabaseManagementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) SqlUserManagement() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlUserManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) SqlUserManagementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlUserManagementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) Timeouts() MdbClickhouseClusterTimeoutsOutputReference {
	var returns MdbClickhouseClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) User() MdbClickhouseClusterUserList {
	var returns MdbClickhouseClusterUserList
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) UserInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) Zookeeper() MdbClickhouseClusterZookeeperOutputReference {
	var returns MdbClickhouseClusterZookeeperOutputReference
	_jsii_.Get(
		j,
		"zookeeper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseCluster) ZookeeperInput() *MdbClickhouseClusterZookeeper {
	var returns *MdbClickhouseClusterZookeeper
	_jsii_.Get(
		j,
		"zookeeperInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster yandex_mdb_clickhouse_cluster} Resource.
func NewMdbClickhouseCluster(scope constructs.Construct, id *string, config *MdbClickhouseClusterConfig) MdbClickhouseCluster {
	_init_.Initialize()

	j := jsiiProxy_MdbClickhouseCluster{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbClickhouseCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster yandex_mdb_clickhouse_cluster} Resource.
func NewMdbClickhouseCluster_Override(m MdbClickhouseCluster, scope constructs.Construct, id *string, config *MdbClickhouseClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbClickhouseCluster",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MdbClickhouseCluster) SetAdminPassword(val *string) {
	_jsii_.Set(
		j,
		"adminPassword",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseCluster) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseCluster) SetCopySchemaOnNewHosts(val interface{}) {
	_jsii_.Set(
		j,
		"copySchemaOnNewHosts",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseCluster) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseCluster) SetDeletionProtection(val interface{}) {
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseCluster) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseCluster) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseCluster) SetEmbeddedKeeper(val interface{}) {
	_jsii_.Set(
		j,
		"embeddedKeeper",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseCluster) SetEnvironment(val *string) {
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseCluster) SetFolderId(val *string) {
	_jsii_.Set(
		j,
		"folderId",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseCluster) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseCluster) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseCluster) SetLabels(val *map[string]*string) {
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseCluster) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseCluster) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseCluster) SetNetworkId(val *string) {
	_jsii_.Set(
		j,
		"networkId",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseCluster) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseCluster) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseCluster) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseCluster) SetServiceAccountId(val *string) {
	_jsii_.Set(
		j,
		"serviceAccountId",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseCluster) SetSqlDatabaseManagement(val interface{}) {
	_jsii_.Set(
		j,
		"sqlDatabaseManagement",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseCluster) SetSqlUserManagement(val interface{}) {
	_jsii_.Set(
		j,
		"sqlUserManagement",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseCluster) SetVersion(val *string) {
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
func MdbClickhouseCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-yandex.MdbClickhouseCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MdbClickhouseCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-yandex.MdbClickhouseCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MdbClickhouseCluster) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseCluster) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseCluster) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) PutAccess(value *MdbClickhouseClusterAccess) {
	_jsii_.InvokeVoid(
		m,
		"putAccess",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) PutBackupWindowStart(value *MdbClickhouseClusterBackupWindowStart) {
	_jsii_.InvokeVoid(
		m,
		"putBackupWindowStart",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) PutClickhouse(value *MdbClickhouseClusterClickhouse) {
	_jsii_.InvokeVoid(
		m,
		"putClickhouse",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) PutCloudStorage(value *MdbClickhouseClusterCloudStorage) {
	_jsii_.InvokeVoid(
		m,
		"putCloudStorage",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) PutDatabase(value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"putDatabase",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) PutFormatSchema(value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"putFormatSchema",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) PutHost(value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"putHost",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) PutMaintenanceWindow(value *MdbClickhouseClusterMaintenanceWindow) {
	_jsii_.InvokeVoid(
		m,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) PutMlModel(value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"putMlModel",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) PutShardGroup(value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"putShardGroup",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) PutTimeouts(value *MdbClickhouseClusterTimeouts) {
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) PutUser(value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"putUser",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) PutZookeeper(value *MdbClickhouseClusterZookeeper) {
	_jsii_.InvokeVoid(
		m,
		"putZookeeper",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) ResetAccess() {
	_jsii_.InvokeVoid(
		m,
		"resetAccess",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) ResetAdminPassword() {
	_jsii_.InvokeVoid(
		m,
		"resetAdminPassword",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) ResetBackupWindowStart() {
	_jsii_.InvokeVoid(
		m,
		"resetBackupWindowStart",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) ResetCloudStorage() {
	_jsii_.InvokeVoid(
		m,
		"resetCloudStorage",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) ResetCopySchemaOnNewHosts() {
	_jsii_.InvokeVoid(
		m,
		"resetCopySchemaOnNewHosts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) ResetDatabase() {
	_jsii_.InvokeVoid(
		m,
		"resetDatabase",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		m,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) ResetEmbeddedKeeper() {
	_jsii_.InvokeVoid(
		m,
		"resetEmbeddedKeeper",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) ResetFolderId() {
	_jsii_.InvokeVoid(
		m,
		"resetFolderId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) ResetFormatSchema() {
	_jsii_.InvokeVoid(
		m,
		"resetFormatSchema",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) ResetLabels() {
	_jsii_.InvokeVoid(
		m,
		"resetLabels",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		m,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) ResetMlModel() {
	_jsii_.InvokeVoid(
		m,
		"resetMlModel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		m,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) ResetServiceAccountId() {
	_jsii_.InvokeVoid(
		m,
		"resetServiceAccountId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) ResetShardGroup() {
	_jsii_.InvokeVoid(
		m,
		"resetShardGroup",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) ResetSqlDatabaseManagement() {
	_jsii_.InvokeVoid(
		m,
		"resetSqlDatabaseManagement",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) ResetSqlUserManagement() {
	_jsii_.InvokeVoid(
		m,
		"resetSqlUserManagement",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) ResetUser() {
	_jsii_.InvokeVoid(
		m,
		"resetUser",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) ResetVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) ResetZookeeper() {
	_jsii_.InvokeVoid(
		m,
		"resetZookeeper",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

