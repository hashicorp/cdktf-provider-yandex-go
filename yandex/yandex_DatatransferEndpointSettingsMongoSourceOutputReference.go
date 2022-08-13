// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatatransferEndpointSettingsMongoSourceOutputReference interface {
	cdktf.ComplexObject
	Collections() DatatransferEndpointSettingsMongoSourceCollectionsList
	CollectionsInput() interface{}
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
	Connection() DatatransferEndpointSettingsMongoSourceConnectionOutputReference
	ConnectionInput() *DatatransferEndpointSettingsMongoSourceConnection
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExcludedCollections() DatatransferEndpointSettingsMongoSourceExcludedCollectionsList
	ExcludedCollectionsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *DatatransferEndpointSettingsMongoSource
	SetInternalValue(val *DatatransferEndpointSettingsMongoSource)
	SecondaryPreferredMode() interface{}
	SetSecondaryPreferredMode(val interface{})
	SecondaryPreferredModeInput() interface{}
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
	PutCollections(value interface{})
	PutConnection(value *DatatransferEndpointSettingsMongoSourceConnection)
	PutExcludedCollections(value interface{})
	ResetCollections()
	ResetConnection()
	ResetExcludedCollections()
	ResetSecondaryPreferredMode()
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

// The jsii proxy struct for DatatransferEndpointSettingsMongoSourceOutputReference
type jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) Collections() DatatransferEndpointSettingsMongoSourceCollectionsList {
	var returns DatatransferEndpointSettingsMongoSourceCollectionsList
	_jsii_.Get(
		j,
		"collections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) CollectionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"collectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) Connection() DatatransferEndpointSettingsMongoSourceConnectionOutputReference {
	var returns DatatransferEndpointSettingsMongoSourceConnectionOutputReference
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) ConnectionInput() *DatatransferEndpointSettingsMongoSourceConnection {
	var returns *DatatransferEndpointSettingsMongoSourceConnection
	_jsii_.Get(
		j,
		"connectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) ExcludedCollections() DatatransferEndpointSettingsMongoSourceExcludedCollectionsList {
	var returns DatatransferEndpointSettingsMongoSourceExcludedCollectionsList
	_jsii_.Get(
		j,
		"excludedCollections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) ExcludedCollectionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludedCollectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) InternalValue() *DatatransferEndpointSettingsMongoSource {
	var returns *DatatransferEndpointSettingsMongoSource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) SecondaryPreferredMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryPreferredMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) SecondaryPreferredModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryPreferredModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatatransferEndpointSettingsMongoSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatatransferEndpointSettingsMongoSourceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsMongoSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatatransferEndpointSettingsMongoSourceOutputReference_Override(d DatatransferEndpointSettingsMongoSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsMongoSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) SetInternalValue(val *DatatransferEndpointSettingsMongoSource) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) SetSecondaryPreferredMode(val interface{}) {
	_jsii_.Set(
		j,
		"secondaryPreferredMode",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) SetSecurityGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) PutCollections(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putCollections",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) PutConnection(value *DatatransferEndpointSettingsMongoSourceConnection) {
	_jsii_.InvokeVoid(
		d,
		"putConnection",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) PutExcludedCollections(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putExcludedCollections",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) ResetCollections() {
	_jsii_.InvokeVoid(
		d,
		"resetCollections",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) ResetConnection() {
	_jsii_.InvokeVoid(
		d,
		"resetConnection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) ResetExcludedCollections() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludedCollections",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) ResetSecondaryPreferredMode() {
	_jsii_.InvokeVoid(
		d,
		"resetSecondaryPreferredMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		d,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

