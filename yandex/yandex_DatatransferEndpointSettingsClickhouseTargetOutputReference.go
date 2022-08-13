// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatatransferEndpointSettingsClickhouseTargetOutputReference interface {
	cdktf.ComplexObject
	AltNames() DatatransferEndpointSettingsClickhouseTargetAltNamesList
	AltNamesInput() interface{}
	CleanupPolicy() *string
	SetCleanupPolicy(val *string)
	CleanupPolicyInput() *string
	ClickhouseClusterName() *string
	SetClickhouseClusterName(val *string)
	ClickhouseClusterNameInput() *string
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
	Connection() DatatransferEndpointSettingsClickhouseTargetConnectionOutputReference
	ConnectionInput() *DatatransferEndpointSettingsClickhouseTargetConnection
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *DatatransferEndpointSettingsClickhouseTarget
	SetInternalValue(val *DatatransferEndpointSettingsClickhouseTarget)
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	SecurityGroupsInput() *[]*string
	Sharding() DatatransferEndpointSettingsClickhouseTargetShardingOutputReference
	ShardingInput() *DatatransferEndpointSettingsClickhouseTargetSharding
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
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
	PutAltNames(value interface{})
	PutConnection(value *DatatransferEndpointSettingsClickhouseTargetConnection)
	PutSharding(value *DatatransferEndpointSettingsClickhouseTargetSharding)
	ResetAltNames()
	ResetCleanupPolicy()
	ResetClickhouseClusterName()
	ResetConnection()
	ResetSecurityGroups()
	ResetSharding()
	ResetSubnetId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatatransferEndpointSettingsClickhouseTargetOutputReference
type jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) AltNames() DatatransferEndpointSettingsClickhouseTargetAltNamesList {
	var returns DatatransferEndpointSettingsClickhouseTargetAltNamesList
	_jsii_.Get(
		j,
		"altNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) AltNamesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"altNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) CleanupPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cleanupPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) CleanupPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cleanupPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) ClickhouseClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clickhouseClusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) ClickhouseClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clickhouseClusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) Connection() DatatransferEndpointSettingsClickhouseTargetConnectionOutputReference {
	var returns DatatransferEndpointSettingsClickhouseTargetConnectionOutputReference
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) ConnectionInput() *DatatransferEndpointSettingsClickhouseTargetConnection {
	var returns *DatatransferEndpointSettingsClickhouseTargetConnection
	_jsii_.Get(
		j,
		"connectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) InternalValue() *DatatransferEndpointSettingsClickhouseTarget {
	var returns *DatatransferEndpointSettingsClickhouseTarget
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) Sharding() DatatransferEndpointSettingsClickhouseTargetShardingOutputReference {
	var returns DatatransferEndpointSettingsClickhouseTargetShardingOutputReference
	_jsii_.Get(
		j,
		"sharding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) ShardingInput() *DatatransferEndpointSettingsClickhouseTargetSharding {
	var returns *DatatransferEndpointSettingsClickhouseTargetSharding
	_jsii_.Get(
		j,
		"shardingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatatransferEndpointSettingsClickhouseTargetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatatransferEndpointSettingsClickhouseTargetOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsClickhouseTargetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatatransferEndpointSettingsClickhouseTargetOutputReference_Override(d DatatransferEndpointSettingsClickhouseTargetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsClickhouseTargetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) SetCleanupPolicy(val *string) {
	_jsii_.Set(
		j,
		"cleanupPolicy",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) SetClickhouseClusterName(val *string) {
	_jsii_.Set(
		j,
		"clickhouseClusterName",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) SetInternalValue(val *DatatransferEndpointSettingsClickhouseTarget) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) SetSecurityGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) PutAltNames(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putAltNames",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) PutConnection(value *DatatransferEndpointSettingsClickhouseTargetConnection) {
	_jsii_.InvokeVoid(
		d,
		"putConnection",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) PutSharding(value *DatatransferEndpointSettingsClickhouseTargetSharding) {
	_jsii_.InvokeVoid(
		d,
		"putSharding",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) ResetAltNames() {
	_jsii_.InvokeVoid(
		d,
		"resetAltNames",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) ResetCleanupPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetCleanupPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) ResetClickhouseClusterName() {
	_jsii_.InvokeVoid(
		d,
		"resetClickhouseClusterName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) ResetConnection() {
	_jsii_.InvokeVoid(
		d,
		"resetConnection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) ResetSharding() {
	_jsii_.InvokeVoid(
		d,
		"resetSharding",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		d,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseTargetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

