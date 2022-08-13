// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster yandex_mdb_mysql_cluster}.
type MdbMysqlCluster interface {
	cdktf.TerraformResource
	Access() MdbMysqlClusterAccessOutputReference
	AccessInput() *MdbMysqlClusterAccess
	AllowRegenerationHost() interface{}
	SetAllowRegenerationHost(val interface{})
	AllowRegenerationHostInput() interface{}
	BackupWindowStart() MdbMysqlClusterBackupWindowStartOutputReference
	BackupWindowStartInput() *MdbMysqlClusterBackupWindowStart
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
	Database() MdbMysqlClusterDatabaseList
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
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Health() *string
	Host() MdbMysqlClusterHostList
	HostGroupIds() *[]*string
	SetHostGroupIds(val *[]*string)
	HostGroupIdsInput() *[]*string
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
	MaintenanceWindow() MdbMysqlClusterMaintenanceWindowOutputReference
	MaintenanceWindowInput() *MdbMysqlClusterMaintenanceWindow
	MysqlConfig() *map[string]*string
	SetMysqlConfig(val *map[string]*string)
	MysqlConfigInput() *map[string]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkId() *string
	SetNetworkId(val *string)
	NetworkIdInput() *string
	// The tree node.
	Node() constructs.Node
	PerformanceDiagnostics() MdbMysqlClusterPerformanceDiagnosticsOutputReference
	PerformanceDiagnosticsInput() *MdbMysqlClusterPerformanceDiagnostics
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
	Resources() MdbMysqlClusterResourcesOutputReference
	ResourcesInput() *MdbMysqlClusterResources
	Restore() MdbMysqlClusterRestoreOutputReference
	RestoreInput() *MdbMysqlClusterRestore
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MdbMysqlClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	User() MdbMysqlClusterUserList
	UserInput() interface{}
	Version() *string
	SetVersion(val *string)
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
	PutAccess(value *MdbMysqlClusterAccess)
	PutBackupWindowStart(value *MdbMysqlClusterBackupWindowStart)
	PutDatabase(value interface{})
	PutHost(value interface{})
	PutMaintenanceWindow(value *MdbMysqlClusterMaintenanceWindow)
	PutPerformanceDiagnostics(value *MdbMysqlClusterPerformanceDiagnostics)
	PutResources(value *MdbMysqlClusterResources)
	PutRestore(value *MdbMysqlClusterRestore)
	PutTimeouts(value *MdbMysqlClusterTimeouts)
	PutUser(value interface{})
	ResetAccess()
	ResetAllowRegenerationHost()
	ResetBackupWindowStart()
	ResetDatabase()
	ResetDeletionProtection()
	ResetDescription()
	ResetFolderId()
	ResetHostGroupIds()
	ResetId()
	ResetLabels()
	ResetMaintenanceWindow()
	ResetMysqlConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPerformanceDiagnostics()
	ResetRestore()
	ResetSecurityGroupIds()
	ResetTimeouts()
	ResetUser()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for MdbMysqlCluster
type jsiiProxy_MdbMysqlCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MdbMysqlCluster) Access() MdbMysqlClusterAccessOutputReference {
	var returns MdbMysqlClusterAccessOutputReference
	_jsii_.Get(
		j,
		"access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) AccessInput() *MdbMysqlClusterAccess {
	var returns *MdbMysqlClusterAccess
	_jsii_.Get(
		j,
		"accessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) AllowRegenerationHost() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowRegenerationHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) AllowRegenerationHostInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowRegenerationHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) BackupWindowStart() MdbMysqlClusterBackupWindowStartOutputReference {
	var returns MdbMysqlClusterBackupWindowStartOutputReference
	_jsii_.Get(
		j,
		"backupWindowStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) BackupWindowStartInput() *MdbMysqlClusterBackupWindowStart {
	var returns *MdbMysqlClusterBackupWindowStart
	_jsii_.Get(
		j,
		"backupWindowStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) Database() MdbMysqlClusterDatabaseList {
	var returns MdbMysqlClusterDatabaseList
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) DatabaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) Environment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) EnvironmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) FolderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) FolderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) Health() *string {
	var returns *string
	_jsii_.Get(
		j,
		"health",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) Host() MdbMysqlClusterHostList {
	var returns MdbMysqlClusterHostList
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) HostGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) HostGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) HostInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) MaintenanceWindow() MdbMysqlClusterMaintenanceWindowOutputReference {
	var returns MdbMysqlClusterMaintenanceWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) MaintenanceWindowInput() *MdbMysqlClusterMaintenanceWindow {
	var returns *MdbMysqlClusterMaintenanceWindow
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) MysqlConfig() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"mysqlConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) MysqlConfigInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"mysqlConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) NetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) NetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) PerformanceDiagnostics() MdbMysqlClusterPerformanceDiagnosticsOutputReference {
	var returns MdbMysqlClusterPerformanceDiagnosticsOutputReference
	_jsii_.Get(
		j,
		"performanceDiagnostics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) PerformanceDiagnosticsInput() *MdbMysqlClusterPerformanceDiagnostics {
	var returns *MdbMysqlClusterPerformanceDiagnostics
	_jsii_.Get(
		j,
		"performanceDiagnosticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) Resources() MdbMysqlClusterResourcesOutputReference {
	var returns MdbMysqlClusterResourcesOutputReference
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) ResourcesInput() *MdbMysqlClusterResources {
	var returns *MdbMysqlClusterResources
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) Restore() MdbMysqlClusterRestoreOutputReference {
	var returns MdbMysqlClusterRestoreOutputReference
	_jsii_.Get(
		j,
		"restore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) RestoreInput() *MdbMysqlClusterRestore {
	var returns *MdbMysqlClusterRestore
	_jsii_.Get(
		j,
		"restoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) Timeouts() MdbMysqlClusterTimeoutsOutputReference {
	var returns MdbMysqlClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) User() MdbMysqlClusterUserList {
	var returns MdbMysqlClusterUserList
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) UserInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMysqlCluster) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster yandex_mdb_mysql_cluster} Resource.
func NewMdbMysqlCluster(scope constructs.Construct, id *string, config *MdbMysqlClusterConfig) MdbMysqlCluster {
	_init_.Initialize()

	j := jsiiProxy_MdbMysqlCluster{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbMysqlCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster yandex_mdb_mysql_cluster} Resource.
func NewMdbMysqlCluster_Override(m MdbMysqlCluster, scope constructs.Construct, id *string, config *MdbMysqlClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbMysqlCluster",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MdbMysqlCluster) SetAllowRegenerationHost(val interface{}) {
	_jsii_.Set(
		j,
		"allowRegenerationHost",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlCluster) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlCluster) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlCluster) SetDeletionProtection(val interface{}) {
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlCluster) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlCluster) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlCluster) SetEnvironment(val *string) {
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlCluster) SetFolderId(val *string) {
	_jsii_.Set(
		j,
		"folderId",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlCluster) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlCluster) SetHostGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"hostGroupIds",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlCluster) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlCluster) SetLabels(val *map[string]*string) {
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlCluster) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlCluster) SetMysqlConfig(val *map[string]*string) {
	_jsii_.Set(
		j,
		"mysqlConfig",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlCluster) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlCluster) SetNetworkId(val *string) {
	_jsii_.Set(
		j,
		"networkId",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlCluster) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlCluster) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlCluster) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_MdbMysqlCluster) SetVersion(val *string) {
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
func MdbMysqlCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-yandex.MdbMysqlCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MdbMysqlCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-yandex.MdbMysqlCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MdbMysqlCluster) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MdbMysqlCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlCluster) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlCluster) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MdbMysqlCluster) PutAccess(value *MdbMysqlClusterAccess) {
	_jsii_.InvokeVoid(
		m,
		"putAccess",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbMysqlCluster) PutBackupWindowStart(value *MdbMysqlClusterBackupWindowStart) {
	_jsii_.InvokeVoid(
		m,
		"putBackupWindowStart",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbMysqlCluster) PutDatabase(value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"putDatabase",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbMysqlCluster) PutHost(value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"putHost",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbMysqlCluster) PutMaintenanceWindow(value *MdbMysqlClusterMaintenanceWindow) {
	_jsii_.InvokeVoid(
		m,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbMysqlCluster) PutPerformanceDiagnostics(value *MdbMysqlClusterPerformanceDiagnostics) {
	_jsii_.InvokeVoid(
		m,
		"putPerformanceDiagnostics",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbMysqlCluster) PutResources(value *MdbMysqlClusterResources) {
	_jsii_.InvokeVoid(
		m,
		"putResources",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbMysqlCluster) PutRestore(value *MdbMysqlClusterRestore) {
	_jsii_.InvokeVoid(
		m,
		"putRestore",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbMysqlCluster) PutTimeouts(value *MdbMysqlClusterTimeouts) {
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbMysqlCluster) PutUser(value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"putUser",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbMysqlCluster) ResetAccess() {
	_jsii_.InvokeVoid(
		m,
		"resetAccess",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMysqlCluster) ResetAllowRegenerationHost() {
	_jsii_.InvokeVoid(
		m,
		"resetAllowRegenerationHost",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMysqlCluster) ResetBackupWindowStart() {
	_jsii_.InvokeVoid(
		m,
		"resetBackupWindowStart",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMysqlCluster) ResetDatabase() {
	_jsii_.InvokeVoid(
		m,
		"resetDatabase",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMysqlCluster) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		m,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMysqlCluster) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMysqlCluster) ResetFolderId() {
	_jsii_.InvokeVoid(
		m,
		"resetFolderId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMysqlCluster) ResetHostGroupIds() {
	_jsii_.InvokeVoid(
		m,
		"resetHostGroupIds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMysqlCluster) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMysqlCluster) ResetLabels() {
	_jsii_.InvokeVoid(
		m,
		"resetLabels",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMysqlCluster) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		m,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMysqlCluster) ResetMysqlConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetMysqlConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMysqlCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMysqlCluster) ResetPerformanceDiagnostics() {
	_jsii_.InvokeVoid(
		m,
		"resetPerformanceDiagnostics",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMysqlCluster) ResetRestore() {
	_jsii_.InvokeVoid(
		m,
		"resetRestore",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMysqlCluster) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		m,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMysqlCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMysqlCluster) ResetUser() {
	_jsii_.InvokeVoid(
		m,
		"resetUser",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMysqlCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMysqlCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

