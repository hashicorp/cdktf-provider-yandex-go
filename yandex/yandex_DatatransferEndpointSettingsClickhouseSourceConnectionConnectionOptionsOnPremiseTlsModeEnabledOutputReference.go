// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference interface {
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
	InternalValue() *DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabled
	SetInternalValue(val *DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabled)
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

// The jsii proxy struct for DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference
type jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference) CaCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference) CaCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference) InternalValue() *DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabled {
	var returns *DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabled
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference_Override(d DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference) SetCaCertificate(val *string) {
	_jsii_.Set(
		j,
		"caCertificate",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference) SetInternalValue(val *DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabled) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference) ResetCaCertificate() {
	_jsii_.InvokeVoid(
		d,
		"resetCaCertificate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabledOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

