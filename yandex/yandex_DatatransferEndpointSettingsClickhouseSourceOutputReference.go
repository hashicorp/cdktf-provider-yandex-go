// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatatransferEndpointSettingsClickhouseSourceOutputReference interface {
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
	Connection() DatatransferEndpointSettingsClickhouseSourceConnectionOutputReference
	ConnectionInput() *DatatransferEndpointSettingsClickhouseSourceConnection
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExcludeTables() *[]*string
	SetExcludeTables(val *[]*string)
	ExcludeTablesInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludeTables() *[]*string
	SetIncludeTables(val *[]*string)
	IncludeTablesInput() *[]*string
	InternalValue() *DatatransferEndpointSettingsClickhouseSource
	SetInternalValue(val *DatatransferEndpointSettingsClickhouseSource)
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	SecurityGroupsInput() *[]*string
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
	PutConnection(value *DatatransferEndpointSettingsClickhouseSourceConnection)
	ResetConnection()
	ResetExcludeTables()
	ResetIncludeTables()
	ResetSecurityGroups()
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

// The jsii proxy struct for DatatransferEndpointSettingsClickhouseSourceOutputReference
type jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) Connection() DatatransferEndpointSettingsClickhouseSourceConnectionOutputReference {
	var returns DatatransferEndpointSettingsClickhouseSourceConnectionOutputReference
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) ConnectionInput() *DatatransferEndpointSettingsClickhouseSourceConnection {
	var returns *DatatransferEndpointSettingsClickhouseSourceConnection
	_jsii_.Get(
		j,
		"connectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) ExcludeTables() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeTables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) ExcludeTablesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeTablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) IncludeTables() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeTables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) IncludeTablesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeTablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) InternalValue() *DatatransferEndpointSettingsClickhouseSource {
	var returns *DatatransferEndpointSettingsClickhouseSource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatatransferEndpointSettingsClickhouseSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatatransferEndpointSettingsClickhouseSourceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsClickhouseSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatatransferEndpointSettingsClickhouseSourceOutputReference_Override(d DatatransferEndpointSettingsClickhouseSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsClickhouseSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) SetExcludeTables(val *[]*string) {
	_jsii_.Set(
		j,
		"excludeTables",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) SetIncludeTables(val *[]*string) {
	_jsii_.Set(
		j,
		"includeTables",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) SetInternalValue(val *DatatransferEndpointSettingsClickhouseSource) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) SetSecurityGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) PutConnection(value *DatatransferEndpointSettingsClickhouseSourceConnection) {
	_jsii_.InvokeVoid(
		d,
		"putConnection",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) ResetConnection() {
	_jsii_.InvokeVoid(
		d,
		"resetConnection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) ResetExcludeTables() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludeTables",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) ResetIncludeTables() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeTables",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		d,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

