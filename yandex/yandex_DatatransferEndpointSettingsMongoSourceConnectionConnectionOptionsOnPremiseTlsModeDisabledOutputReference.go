// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference interface {
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
	InternalValue() *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabled
	SetInternalValue(val *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabled)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference
type jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference) InternalValue() *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabled {
	var returns *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabled
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference_Override(d DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference) SetInternalValue(val *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabled) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsMongoSourceConnectionConnectionOptionsOnPremiseTlsModeDisabledOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

