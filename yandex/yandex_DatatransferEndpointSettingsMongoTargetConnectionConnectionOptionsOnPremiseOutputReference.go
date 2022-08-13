// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference interface {
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
	InternalValue() *DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremise
	SetInternalValue(val *DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremise)
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
	TlsMode() DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseTlsModeOutputReference
	TlsModeInput() *DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseTlsMode
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
	PutTlsMode(value *DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseTlsMode)
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

// The jsii proxy struct for DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference
type jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) Hosts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) HostsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) InternalValue() *DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremise {
	var returns *DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremise
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) ReplicaSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) ReplicaSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) TlsMode() DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseTlsModeOutputReference {
	var returns DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseTlsModeOutputReference
	_jsii_.Get(
		j,
		"tlsMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) TlsModeInput() *DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseTlsMode {
	var returns *DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseTlsMode
	_jsii_.Get(
		j,
		"tlsModeInput",
		&returns,
	)
	return returns
}


func NewDatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference_Override(d DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) SetHosts(val *[]*string) {
	_jsii_.Set(
		j,
		"hosts",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) SetInternalValue(val *DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremise) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) SetReplicaSet(val *string) {
	_jsii_.Set(
		j,
		"replicaSet",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) PutTlsMode(value *DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseTlsMode) {
	_jsii_.InvokeVoid(
		d,
		"putTlsMode",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) ResetHosts() {
	_jsii_.InvokeVoid(
		d,
		"resetHosts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		d,
		"resetPort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) ResetReplicaSet() {
	_jsii_.InvokeVoid(
		d,
		"resetReplicaSet",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) ResetTlsMode() {
	_jsii_.InvokeVoid(
		d,
		"resetTlsMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

