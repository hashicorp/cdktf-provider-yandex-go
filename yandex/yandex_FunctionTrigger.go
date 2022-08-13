// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger yandex_function_trigger}.
type FunctionTrigger interface {
	cdktf.TerraformResource
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
	Dlq() FunctionTriggerDlqOutputReference
	DlqInput() *FunctionTriggerDlq
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
	Function() FunctionTriggerFunctionOutputReference
	FunctionInput() *FunctionTriggerFunction
	Id() *string
	SetId(val *string)
	IdInput() *string
	Iot() FunctionTriggerIotOutputReference
	IotInput() *FunctionTriggerIot
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Logging() FunctionTriggerLoggingOutputReference
	LoggingInput() *FunctionTriggerLogging
	LogGroup() FunctionTriggerLogGroupOutputReference
	LogGroupInput() *FunctionTriggerLogGroup
	MessageQueue() FunctionTriggerMessageQueueOutputReference
	MessageQueueInput() *FunctionTriggerMessageQueue
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	ObjectStorage() FunctionTriggerObjectStorageOutputReference
	ObjectStorageInput() *FunctionTriggerObjectStorage
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() FunctionTriggerTimeoutsOutputReference
	TimeoutsInput() interface{}
	Timer() FunctionTriggerTimerOutputReference
	TimerInput() *FunctionTriggerTimer
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
	PutDlq(value *FunctionTriggerDlq)
	PutFunction(value *FunctionTriggerFunction)
	PutIot(value *FunctionTriggerIot)
	PutLogging(value *FunctionTriggerLogging)
	PutLogGroup(value *FunctionTriggerLogGroup)
	PutMessageQueue(value *FunctionTriggerMessageQueue)
	PutObjectStorage(value *FunctionTriggerObjectStorage)
	PutTimeouts(value *FunctionTriggerTimeouts)
	PutTimer(value *FunctionTriggerTimer)
	ResetDescription()
	ResetDlq()
	ResetFolderId()
	ResetId()
	ResetIot()
	ResetLabels()
	ResetLogging()
	ResetLogGroup()
	ResetMessageQueue()
	ResetObjectStorage()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	ResetTimer()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for FunctionTrigger
type jsiiProxy_FunctionTrigger struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FunctionTrigger) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) Dlq() FunctionTriggerDlqOutputReference {
	var returns FunctionTriggerDlqOutputReference
	_jsii_.Get(
		j,
		"dlq",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) DlqInput() *FunctionTriggerDlq {
	var returns *FunctionTriggerDlq
	_jsii_.Get(
		j,
		"dlqInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) FolderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) FolderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) Function() FunctionTriggerFunctionOutputReference {
	var returns FunctionTriggerFunctionOutputReference
	_jsii_.Get(
		j,
		"function",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) FunctionInput() *FunctionTriggerFunction {
	var returns *FunctionTriggerFunction
	_jsii_.Get(
		j,
		"functionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) Iot() FunctionTriggerIotOutputReference {
	var returns FunctionTriggerIotOutputReference
	_jsii_.Get(
		j,
		"iot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) IotInput() *FunctionTriggerIot {
	var returns *FunctionTriggerIot
	_jsii_.Get(
		j,
		"iotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) Logging() FunctionTriggerLoggingOutputReference {
	var returns FunctionTriggerLoggingOutputReference
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) LoggingInput() *FunctionTriggerLogging {
	var returns *FunctionTriggerLogging
	_jsii_.Get(
		j,
		"loggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) LogGroup() FunctionTriggerLogGroupOutputReference {
	var returns FunctionTriggerLogGroupOutputReference
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) LogGroupInput() *FunctionTriggerLogGroup {
	var returns *FunctionTriggerLogGroup
	_jsii_.Get(
		j,
		"logGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) MessageQueue() FunctionTriggerMessageQueueOutputReference {
	var returns FunctionTriggerMessageQueueOutputReference
	_jsii_.Get(
		j,
		"messageQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) MessageQueueInput() *FunctionTriggerMessageQueue {
	var returns *FunctionTriggerMessageQueue
	_jsii_.Get(
		j,
		"messageQueueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) ObjectStorage() FunctionTriggerObjectStorageOutputReference {
	var returns FunctionTriggerObjectStorageOutputReference
	_jsii_.Get(
		j,
		"objectStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) ObjectStorageInput() *FunctionTriggerObjectStorage {
	var returns *FunctionTriggerObjectStorage
	_jsii_.Get(
		j,
		"objectStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) Timeouts() FunctionTriggerTimeoutsOutputReference {
	var returns FunctionTriggerTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) Timer() FunctionTriggerTimerOutputReference {
	var returns FunctionTriggerTimerOutputReference
	_jsii_.Get(
		j,
		"timer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionTrigger) TimerInput() *FunctionTriggerTimer {
	var returns *FunctionTriggerTimer
	_jsii_.Get(
		j,
		"timerInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger yandex_function_trigger} Resource.
func NewFunctionTrigger(scope constructs.Construct, id *string, config *FunctionTriggerConfig) FunctionTrigger {
	_init_.Initialize()

	j := jsiiProxy_FunctionTrigger{}

	_jsii_.Create(
		"@cdktf/provider-yandex.FunctionTrigger",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger yandex_function_trigger} Resource.
func NewFunctionTrigger_Override(f FunctionTrigger, scope constructs.Construct, id *string, config *FunctionTriggerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.FunctionTrigger",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FunctionTrigger) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_FunctionTrigger) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FunctionTrigger) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FunctionTrigger) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_FunctionTrigger) SetFolderId(val *string) {
	_jsii_.Set(
		j,
		"folderId",
		val,
	)
}

func (j *jsiiProxy_FunctionTrigger) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_FunctionTrigger) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_FunctionTrigger) SetLabels(val *map[string]*string) {
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_FunctionTrigger) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FunctionTrigger) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FunctionTrigger) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FunctionTrigger) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
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
func FunctionTrigger_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-yandex.FunctionTrigger",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FunctionTrigger_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-yandex.FunctionTrigger",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FunctionTrigger) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_FunctionTrigger) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTrigger) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTrigger) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTrigger) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTrigger) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTrigger) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTrigger) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTrigger) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTrigger) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTrigger) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTrigger) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FunctionTrigger) PutDlq(value *FunctionTriggerDlq) {
	_jsii_.InvokeVoid(
		f,
		"putDlq",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionTrigger) PutFunction(value *FunctionTriggerFunction) {
	_jsii_.InvokeVoid(
		f,
		"putFunction",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionTrigger) PutIot(value *FunctionTriggerIot) {
	_jsii_.InvokeVoid(
		f,
		"putIot",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionTrigger) PutLogging(value *FunctionTriggerLogging) {
	_jsii_.InvokeVoid(
		f,
		"putLogging",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionTrigger) PutLogGroup(value *FunctionTriggerLogGroup) {
	_jsii_.InvokeVoid(
		f,
		"putLogGroup",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionTrigger) PutMessageQueue(value *FunctionTriggerMessageQueue) {
	_jsii_.InvokeVoid(
		f,
		"putMessageQueue",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionTrigger) PutObjectStorage(value *FunctionTriggerObjectStorage) {
	_jsii_.InvokeVoid(
		f,
		"putObjectStorage",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionTrigger) PutTimeouts(value *FunctionTriggerTimeouts) {
	_jsii_.InvokeVoid(
		f,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionTrigger) PutTimer(value *FunctionTriggerTimer) {
	_jsii_.InvokeVoid(
		f,
		"putTimer",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionTrigger) ResetDescription() {
	_jsii_.InvokeVoid(
		f,
		"resetDescription",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionTrigger) ResetDlq() {
	_jsii_.InvokeVoid(
		f,
		"resetDlq",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionTrigger) ResetFolderId() {
	_jsii_.InvokeVoid(
		f,
		"resetFolderId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionTrigger) ResetId() {
	_jsii_.InvokeVoid(
		f,
		"resetId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionTrigger) ResetIot() {
	_jsii_.InvokeVoid(
		f,
		"resetIot",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionTrigger) ResetLabels() {
	_jsii_.InvokeVoid(
		f,
		"resetLabels",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionTrigger) ResetLogging() {
	_jsii_.InvokeVoid(
		f,
		"resetLogging",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionTrigger) ResetLogGroup() {
	_jsii_.InvokeVoid(
		f,
		"resetLogGroup",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionTrigger) ResetMessageQueue() {
	_jsii_.InvokeVoid(
		f,
		"resetMessageQueue",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionTrigger) ResetObjectStorage() {
	_jsii_.InvokeVoid(
		f,
		"resetObjectStorage",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionTrigger) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionTrigger) ResetTimeouts() {
	_jsii_.InvokeVoid(
		f,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionTrigger) ResetTimer() {
	_jsii_.InvokeVoid(
		f,
		"resetTimer",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionTrigger) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTrigger) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTrigger) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionTrigger) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

