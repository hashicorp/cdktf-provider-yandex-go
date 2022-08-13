// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference interface {
	cdktf.ComplexObject
	CaCertificate() *string
	SetCaCertificate(val *string)
	CaCertificateInput() *string
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
	InternalValue() *DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabled
	SetInternalValue(val *DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabled)
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
	ResetCaCertificate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference
type jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference) CaCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference) CaCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference) InternalValue() *DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabled {
	var returns *DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabled
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference_Override(d DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference) SetCaCertificate(val *string) {
	_jsii_.Set(
		j,
		"caCertificate",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference) SetInternalValue(val *DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabled) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference) ResetCaCertificate() {
	_jsii_.InvokeVoid(
		d,
		"resetCaCertificate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeEnabledOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

