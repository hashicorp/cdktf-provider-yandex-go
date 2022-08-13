// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster yandex_mdb_greenplum_cluster}.
type MdbGreenplumCluster interface {
	cdktf.TerraformResource
	Access() MdbGreenplumClusterAccessOutputReference
	AccessInput() *MdbGreenplumClusterAccess
	AssignPublicIp() interface{}
	SetAssignPublicIp(val interface{})
	AssignPublicIpInput() interface{}
	BackupWindowStart() MdbGreenplumClusterBackupWindowStartOutputReference
	BackupWindowStartInput() *MdbGreenplumClusterBackupWindowStart
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
	GreenplumConfig() *map[string]*string
	SetGreenplumConfig(val *map[string]*string)
	GreenplumConfigInput() *map[string]*string
	Health() *string
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
	MasterHostCount() *float64
	SetMasterHostCount(val *float64)
	MasterHostCountInput() *float64
	MasterHosts() MdbGreenplumClusterMasterHostsList
	MasterSubcluster() MdbGreenplumClusterMasterSubclusterOutputReference
	MasterSubclusterInput() *MdbGreenplumClusterMasterSubcluster
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkId() *string
	SetNetworkId(val *string)
	NetworkIdInput() *string
	// The tree node.
	Node() constructs.Node
	PoolerConfig() MdbGreenplumClusterPoolerConfigOutputReference
	PoolerConfigInput() *MdbGreenplumClusterPoolerConfig
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
	SegmentHostCount() *float64
	SetSegmentHostCount(val *float64)
	SegmentHostCountInput() *float64
	SegmentHosts() MdbGreenplumClusterSegmentHostsList
	SegmentInHost() *float64
	SetSegmentInHost(val *float64)
	SegmentInHostInput() *float64
	SegmentSubcluster() MdbGreenplumClusterSegmentSubclusterOutputReference
	SegmentSubclusterInput() *MdbGreenplumClusterSegmentSubcluster
	Status() *string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MdbGreenplumClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	UserName() *string
	SetUserName(val *string)
	UserNameInput() *string
	UserPassword() *string
	SetUserPassword(val *string)
	UserPasswordInput() *string
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutAccess(value *MdbGreenplumClusterAccess)
	PutBackupWindowStart(value *MdbGreenplumClusterBackupWindowStart)
	PutMasterSubcluster(value *MdbGreenplumClusterMasterSubcluster)
	PutPoolerConfig(value *MdbGreenplumClusterPoolerConfig)
	PutSegmentSubcluster(value *MdbGreenplumClusterSegmentSubcluster)
	PutTimeouts(value *MdbGreenplumClusterTimeouts)
	ResetAccess()
	ResetBackupWindowStart()
	ResetDeletionProtection()
	ResetDescription()
	ResetFolderId()
	ResetGreenplumConfig()
	ResetId()
	ResetLabels()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPoolerConfig()
	ResetSecurityGroupIds()
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

// The jsii proxy struct for MdbGreenplumCluster
type jsiiProxy_MdbGreenplumCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MdbGreenplumCluster) Access() MdbGreenplumClusterAccessOutputReference {
	var returns MdbGreenplumClusterAccessOutputReference
	_jsii_.Get(
		j,
		"access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) AccessInput() *MdbGreenplumClusterAccess {
	var returns *MdbGreenplumClusterAccess
	_jsii_.Get(
		j,
		"accessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) AssignPublicIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assignPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) AssignPublicIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assignPublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) BackupWindowStart() MdbGreenplumClusterBackupWindowStartOutputReference {
	var returns MdbGreenplumClusterBackupWindowStartOutputReference
	_jsii_.Get(
		j,
		"backupWindowStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) BackupWindowStartInput() *MdbGreenplumClusterBackupWindowStart {
	var returns *MdbGreenplumClusterBackupWindowStart
	_jsii_.Get(
		j,
		"backupWindowStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) Environment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) EnvironmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) FolderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) FolderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) GreenplumConfig() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"greenplumConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) GreenplumConfigInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"greenplumConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) Health() *string {
	var returns *string
	_jsii_.Get(
		j,
		"health",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) MasterHostCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"masterHostCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) MasterHostCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"masterHostCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) MasterHosts() MdbGreenplumClusterMasterHostsList {
	var returns MdbGreenplumClusterMasterHostsList
	_jsii_.Get(
		j,
		"masterHosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) MasterSubcluster() MdbGreenplumClusterMasterSubclusterOutputReference {
	var returns MdbGreenplumClusterMasterSubclusterOutputReference
	_jsii_.Get(
		j,
		"masterSubcluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) MasterSubclusterInput() *MdbGreenplumClusterMasterSubcluster {
	var returns *MdbGreenplumClusterMasterSubcluster
	_jsii_.Get(
		j,
		"masterSubclusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) NetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) NetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) PoolerConfig() MdbGreenplumClusterPoolerConfigOutputReference {
	var returns MdbGreenplumClusterPoolerConfigOutputReference
	_jsii_.Get(
		j,
		"poolerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) PoolerConfigInput() *MdbGreenplumClusterPoolerConfig {
	var returns *MdbGreenplumClusterPoolerConfig
	_jsii_.Get(
		j,
		"poolerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) SegmentHostCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentHostCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) SegmentHostCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentHostCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) SegmentHosts() MdbGreenplumClusterSegmentHostsList {
	var returns MdbGreenplumClusterSegmentHostsList
	_jsii_.Get(
		j,
		"segmentHosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) SegmentInHost() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentInHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) SegmentInHostInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentInHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) SegmentSubcluster() MdbGreenplumClusterSegmentSubclusterOutputReference {
	var returns MdbGreenplumClusterSegmentSubclusterOutputReference
	_jsii_.Get(
		j,
		"segmentSubcluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) SegmentSubclusterInput() *MdbGreenplumClusterSegmentSubcluster {
	var returns *MdbGreenplumClusterSegmentSubcluster
	_jsii_.Get(
		j,
		"segmentSubclusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) Timeouts() MdbGreenplumClusterTimeoutsOutputReference {
	var returns MdbGreenplumClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) UserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) UserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) UserPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) UserPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbGreenplumCluster) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster yandex_mdb_greenplum_cluster} Resource.
func NewMdbGreenplumCluster(scope constructs.Construct, id *string, config *MdbGreenplumClusterConfig) MdbGreenplumCluster {
	_init_.Initialize()

	j := jsiiProxy_MdbGreenplumCluster{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbGreenplumCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/mdb_greenplum_cluster yandex_mdb_greenplum_cluster} Resource.
func NewMdbGreenplumCluster_Override(m MdbGreenplumCluster, scope constructs.Construct, id *string, config *MdbGreenplumClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbGreenplumCluster",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MdbGreenplumCluster) SetAssignPublicIp(val interface{}) {
	_jsii_.Set(
		j,
		"assignPublicIp",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumCluster) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumCluster) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumCluster) SetDeletionProtection(val interface{}) {
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumCluster) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumCluster) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumCluster) SetEnvironment(val *string) {
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumCluster) SetFolderId(val *string) {
	_jsii_.Set(
		j,
		"folderId",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumCluster) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumCluster) SetGreenplumConfig(val *map[string]*string) {
	_jsii_.Set(
		j,
		"greenplumConfig",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumCluster) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumCluster) SetLabels(val *map[string]*string) {
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumCluster) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumCluster) SetMasterHostCount(val *float64) {
	_jsii_.Set(
		j,
		"masterHostCount",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumCluster) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumCluster) SetNetworkId(val *string) {
	_jsii_.Set(
		j,
		"networkId",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumCluster) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumCluster) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumCluster) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumCluster) SetSegmentHostCount(val *float64) {
	_jsii_.Set(
		j,
		"segmentHostCount",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumCluster) SetSegmentInHost(val *float64) {
	_jsii_.Set(
		j,
		"segmentInHost",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumCluster) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumCluster) SetUserName(val *string) {
	_jsii_.Set(
		j,
		"userName",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumCluster) SetUserPassword(val *string) {
	_jsii_.Set(
		j,
		"userPassword",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumCluster) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_MdbGreenplumCluster) SetZone(val *string) {
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
func MdbGreenplumCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-yandex.MdbGreenplumCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MdbGreenplumCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-yandex.MdbGreenplumCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MdbGreenplumCluster) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MdbGreenplumCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbGreenplumCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbGreenplumCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbGreenplumCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbGreenplumCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbGreenplumCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbGreenplumCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbGreenplumCluster) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbGreenplumCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbGreenplumCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbGreenplumCluster) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MdbGreenplumCluster) PutAccess(value *MdbGreenplumClusterAccess) {
	_jsii_.InvokeVoid(
		m,
		"putAccess",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbGreenplumCluster) PutBackupWindowStart(value *MdbGreenplumClusterBackupWindowStart) {
	_jsii_.InvokeVoid(
		m,
		"putBackupWindowStart",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbGreenplumCluster) PutMasterSubcluster(value *MdbGreenplumClusterMasterSubcluster) {
	_jsii_.InvokeVoid(
		m,
		"putMasterSubcluster",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbGreenplumCluster) PutPoolerConfig(value *MdbGreenplumClusterPoolerConfig) {
	_jsii_.InvokeVoid(
		m,
		"putPoolerConfig",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbGreenplumCluster) PutSegmentSubcluster(value *MdbGreenplumClusterSegmentSubcluster) {
	_jsii_.InvokeVoid(
		m,
		"putSegmentSubcluster",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbGreenplumCluster) PutTimeouts(value *MdbGreenplumClusterTimeouts) {
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbGreenplumCluster) ResetAccess() {
	_jsii_.InvokeVoid(
		m,
		"resetAccess",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbGreenplumCluster) ResetBackupWindowStart() {
	_jsii_.InvokeVoid(
		m,
		"resetBackupWindowStart",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbGreenplumCluster) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		m,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbGreenplumCluster) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbGreenplumCluster) ResetFolderId() {
	_jsii_.InvokeVoid(
		m,
		"resetFolderId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbGreenplumCluster) ResetGreenplumConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetGreenplumConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbGreenplumCluster) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbGreenplumCluster) ResetLabels() {
	_jsii_.InvokeVoid(
		m,
		"resetLabels",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbGreenplumCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbGreenplumCluster) ResetPoolerConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetPoolerConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbGreenplumCluster) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		m,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbGreenplumCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbGreenplumCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbGreenplumCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbGreenplumCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbGreenplumCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

