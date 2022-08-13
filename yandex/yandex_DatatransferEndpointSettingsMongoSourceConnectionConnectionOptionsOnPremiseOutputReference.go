// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference interface {
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
	Hosts() *[]*string
	SetHosts(val *[]*string)
	HostsInput() *[]*string
	InternalValue() *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremise
	SetInternalValue(val *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremise)
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	ReplicaSet() *string
	SetReplicaSet(val *string)
	ReplicaSetInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TlsMode() DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeOutputReference
	TlsModeInput() *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsMode
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
	PutTlsMode(value *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsMode)
	ResetHosts()
	ResetPort()
	ResetReplicaSet()
	ResetTlsMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference
type jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) Hosts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) HostsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) InternalValue() *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremise {
	var returns *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremise
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) ReplicaSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) ReplicaSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) TlsMode() DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeOutputReference {
	var returns DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeOutputReference
	_jsii_.Get(
		j,
		"tlsMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) TlsModeInput() *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsMode {
	var returns *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsMode
	_jsii_.Get(
		j,
		"tlsModeInput",
		&returns,
	)
	return returns
}


func NewDatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference_Override(d DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) SetHosts(val *[]*string) {
	_jsii_.Set(
		j,
		"hosts",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) SetInternalValue(val *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremise) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) SetReplicaSet(val *string) {
	_jsii_.Set(
		j,
		"replicaSet",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) PutTlsMode(value *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsMode) {
	_jsii_.InvokeVoid(
		d,
		"putTlsMode",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) ResetHosts() {
	_jsii_.InvokeVoid(
		d,
		"resetHosts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		d,
		"resetPort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) ResetReplicaSet() {
	_jsii_.InvokeVoid(
		d,
		"resetReplicaSet",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) ResetTlsMode() {
	_jsii_.InvokeVoid(
		d,
		"resetTlsMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

