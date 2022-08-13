// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster yandex_mdb_kafka_cluster}.
type DataYandexMdbKafkaCluster interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
	Config() DataYandexMdbKafkaClusterConfigAOutputReference
	ConfigInput() *DataYandexMdbKafkaClusterConfigA
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
	Health() *string
	Host() DataYandexMdbKafkaClusterHostList
	HostGroupIds() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Labels() cdktf.StringMap
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaintenanceWindow() DataYandexMdbKafkaClusterMaintenanceWindowList
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
	Status() *string
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	SubnetIdsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Topic() DataYandexMdbKafkaClusterTopicList
	TopicInput() interface{}
	User() DataYandexMdbKafkaClusterUserList
	UserInput() interface{}
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
	PutConfig(value *DataYandexMdbKafkaClusterConfigA)
	PutTopic(value interface{})
	PutUser(value interface{})
	ResetClusterId()
	ResetConfig()
	ResetDeletionProtection()
	ResetFolderId()
	ResetId()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSubnetIds()
	ResetTopic()
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

// The jsii proxy struct for DataYandexMdbKafkaCluster
type jsiiProxy_DataYandexMdbKafkaCluster struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) Config() DataYandexMdbKafkaClusterConfigAOutputReference {
	var returns DataYandexMdbKafkaClusterConfigAOutputReference
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) ConfigInput() *DataYandexMdbKafkaClusterConfigA {
	var returns *DataYandexMdbKafkaClusterConfigA
	_jsii_.Get(
		j,
		"configInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) Environment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) FolderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) FolderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) Health() *string {
	var returns *string
	_jsii_.Get(
		j,
		"health",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) Host() DataYandexMdbKafkaClusterHostList {
	var returns DataYandexMdbKafkaClusterHostList
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) HostGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) Labels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) MaintenanceWindow() DataYandexMdbKafkaClusterMaintenanceWindowList {
	var returns DataYandexMdbKafkaClusterMaintenanceWindowList
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) NetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) Topic() DataYandexMdbKafkaClusterTopicList {
	var returns DataYandexMdbKafkaClusterTopicList
	_jsii_.Get(
		j,
		"topic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) TopicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"topicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) User() DataYandexMdbKafkaClusterUserList {
	var returns DataYandexMdbKafkaClusterUserList
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) UserInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster yandex_mdb_kafka_cluster} Data Source.
func NewDataYandexMdbKafkaCluster(scope constructs.Construct, id *string, config *DataYandexMdbKafkaClusterConfig) DataYandexMdbKafkaCluster {
	_init_.Initialize()

	j := jsiiProxy_DataYandexMdbKafkaCluster{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexMdbKafkaCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/yandex/d/mdb_kafka_cluster yandex_mdb_kafka_cluster} Data Source.
func NewDataYandexMdbKafkaCluster_Override(d DataYandexMdbKafkaCluster, scope constructs.Construct, id *string, config *DataYandexMdbKafkaClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexMdbKafkaCluster",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) SetClusterId(val *string) {
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) SetDeletionProtection(val interface{}) {
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) SetFolderId(val *string) {
	_jsii_.Set(
		j,
		"folderId",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbKafkaCluster) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
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
func DataYandexMdbKafkaCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-yandex.DataYandexMdbKafkaCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataYandexMdbKafkaCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-yandex.DataYandexMdbKafkaCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaCluster) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaCluster) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaCluster) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaCluster) PutConfig(value *DataYandexMdbKafkaClusterConfigA) {
	_jsii_.InvokeVoid(
		d,
		"putConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaCluster) PutTopic(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putTopic",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaCluster) PutUser(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putUser",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaCluster) ResetClusterId() {
	_jsii_.InvokeVoid(
		d,
		"resetClusterId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaCluster) ResetConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaCluster) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		d,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaCluster) ResetFolderId() {
	_jsii_.InvokeVoid(
		d,
		"resetFolderId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaCluster) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaCluster) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaCluster) ResetSubnetIds() {
	_jsii_.InvokeVoid(
		d,
		"resetSubnetIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaCluster) ResetTopic() {
	_jsii_.InvokeVoid(
		d,
		"resetTopic",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaCluster) ResetUser() {
	_jsii_.InvokeVoid(
		d,
		"resetUser",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbKafkaCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbKafkaCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

