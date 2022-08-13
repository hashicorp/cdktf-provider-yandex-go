// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *MdbClickhouseClusterClickhouseConfigMergeTree
	SetInternalValue(val *MdbClickhouseClusterClickhouseConfigMergeTree)
	MaxBytesToMergeAtMinSpaceInPool() *float64
	SetMaxBytesToMergeAtMinSpaceInPool(val *float64)
	MaxBytesToMergeAtMinSpaceInPoolInput() *float64
	MaxReplicatedMergesInQueue() *float64
	SetMaxReplicatedMergesInQueue(val *float64)
	MaxReplicatedMergesInQueueInput() *float64
	NumberOfFreeEntriesInPoolToLowerMaxSizeOfMerge() *float64
	SetNumberOfFreeEntriesInPoolToLowerMaxSizeOfMerge(val *float64)
	NumberOfFreeEntriesInPoolToLowerMaxSizeOfMergeInput() *float64
	PartsToDelayInsert() *float64
	SetPartsToDelayInsert(val *float64)
	PartsToDelayInsertInput() *float64
	PartsToThrowInsert() *float64
	SetPartsToThrowInsert(val *float64)
	PartsToThrowInsertInput() *float64
	ReplicatedDeduplicationWindow() *float64
	SetReplicatedDeduplicationWindow(val *float64)
	ReplicatedDeduplicationWindowInput() *float64
	ReplicatedDeduplicationWindowSeconds() *float64
	SetReplicatedDeduplicationWindowSeconds(val *float64)
	ReplicatedDeduplicationWindowSecondsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetMaxBytesToMergeAtMinSpaceInPool()
	ResetMaxReplicatedMergesInQueue()
	ResetNumberOfFreeEntriesInPoolToLowerMaxSizeOfMerge()
	ResetPartsToDelayInsert()
	ResetPartsToThrowInsert()
	ResetReplicatedDeduplicationWindow()
	ResetReplicatedDeduplicationWindowSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference
type jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) InternalValue() *MdbClickhouseClusterClickhouseConfigMergeTree {
	var returns *MdbClickhouseClusterClickhouseConfigMergeTree
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) MaxBytesToMergeAtMinSpaceInPool() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBytesToMergeAtMinSpaceInPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) MaxBytesToMergeAtMinSpaceInPoolInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBytesToMergeAtMinSpaceInPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) MaxReplicatedMergesInQueue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxReplicatedMergesInQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) MaxReplicatedMergesInQueueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxReplicatedMergesInQueueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) NumberOfFreeEntriesInPoolToLowerMaxSizeOfMerge() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfFreeEntriesInPoolToLowerMaxSizeOfMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) NumberOfFreeEntriesInPoolToLowerMaxSizeOfMergeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfFreeEntriesInPoolToLowerMaxSizeOfMergeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) PartsToDelayInsert() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partsToDelayInsert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) PartsToDelayInsertInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partsToDelayInsertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) PartsToThrowInsert() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partsToThrowInsert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) PartsToThrowInsertInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partsToThrowInsertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) ReplicatedDeduplicationWindow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicatedDeduplicationWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) ReplicatedDeduplicationWindowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicatedDeduplicationWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) ReplicatedDeduplicationWindowSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicatedDeduplicationWindowSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) ReplicatedDeduplicationWindowSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicatedDeduplicationWindowSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMdbClickhouseClusterClickhouseConfigMergeTreeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMdbClickhouseClusterClickhouseConfigMergeTreeOutputReference_Override(m MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) SetInternalValue(val *MdbClickhouseClusterClickhouseConfigMergeTree) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) SetMaxBytesToMergeAtMinSpaceInPool(val *float64) {
	_jsii_.Set(
		j,
		"maxBytesToMergeAtMinSpaceInPool",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) SetMaxReplicatedMergesInQueue(val *float64) {
	_jsii_.Set(
		j,
		"maxReplicatedMergesInQueue",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) SetNumberOfFreeEntriesInPoolToLowerMaxSizeOfMerge(val *float64) {
	_jsii_.Set(
		j,
		"numberOfFreeEntriesInPoolToLowerMaxSizeOfMerge",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) SetPartsToDelayInsert(val *float64) {
	_jsii_.Set(
		j,
		"partsToDelayInsert",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) SetPartsToThrowInsert(val *float64) {
	_jsii_.Set(
		j,
		"partsToThrowInsert",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) SetReplicatedDeduplicationWindow(val *float64) {
	_jsii_.Set(
		j,
		"replicatedDeduplicationWindow",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) SetReplicatedDeduplicationWindowSeconds(val *float64) {
	_jsii_.Set(
		j,
		"replicatedDeduplicationWindowSeconds",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) ResetMaxBytesToMergeAtMinSpaceInPool() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxBytesToMergeAtMinSpaceInPool",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) ResetMaxReplicatedMergesInQueue() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxReplicatedMergesInQueue",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) ResetNumberOfFreeEntriesInPoolToLowerMaxSizeOfMerge() {
	_jsii_.InvokeVoid(
		m,
		"resetNumberOfFreeEntriesInPoolToLowerMaxSizeOfMerge",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) ResetPartsToDelayInsert() {
	_jsii_.InvokeVoid(
		m,
		"resetPartsToDelayInsert",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) ResetPartsToThrowInsert() {
	_jsii_.InvokeVoid(
		m,
		"resetPartsToThrowInsert",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) ResetReplicatedDeduplicationWindow() {
	_jsii_.InvokeVoid(
		m,
		"resetReplicatedDeduplicationWindow",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) ResetReplicatedDeduplicationWindowSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetReplicatedDeduplicationWindowSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbClickhouseClusterClickhouseConfigMergeTreeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

