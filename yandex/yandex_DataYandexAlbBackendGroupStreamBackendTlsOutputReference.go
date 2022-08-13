// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexAlbBackendGroupStreamBackendTlsOutputReference interface {
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
	InternalValue() *DataYandexAlbBackendGroupStreamBackendTls
	SetInternalValue(val *DataYandexAlbBackendGroupStreamBackendTls)
	Sni() *string
	SetSni(val *string)
	SniInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ValidationContext() DataYandexAlbBackendGroupStreamBackendTlsValidationContextOutputReference
	ValidationContextInput() *DataYandexAlbBackendGroupStreamBackendTlsValidationContext
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
	PutValidationContext(value *DataYandexAlbBackendGroupStreamBackendTlsValidationContext)
	ResetSni()
	ResetValidationContext()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataYandexAlbBackendGroupStreamBackendTlsOutputReference
type jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) InternalValue() *DataYandexAlbBackendGroupStreamBackendTls {
	var returns *DataYandexAlbBackendGroupStreamBackendTls
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) Sni() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sni",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) SniInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sniInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) ValidationContext() DataYandexAlbBackendGroupStreamBackendTlsValidationContextOutputReference {
	var returns DataYandexAlbBackendGroupStreamBackendTlsValidationContextOutputReference
	_jsii_.Get(
		j,
		"validationContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) ValidationContextInput() *DataYandexAlbBackendGroupStreamBackendTlsValidationContext {
	var returns *DataYandexAlbBackendGroupStreamBackendTlsValidationContext
	_jsii_.Get(
		j,
		"validationContextInput",
		&returns,
	)
	return returns
}


func NewDataYandexAlbBackendGroupStreamBackendTlsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataYandexAlbBackendGroupStreamBackendTlsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexAlbBackendGroupStreamBackendTlsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataYandexAlbBackendGroupStreamBackendTlsOutputReference_Override(d DataYandexAlbBackendGroupStreamBackendTlsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexAlbBackendGroupStreamBackendTlsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) SetInternalValue(val *DataYandexAlbBackendGroupStreamBackendTls) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) SetSni(val *string) {
	_jsii_.Set(
		j,
		"sni",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) PutValidationContext(value *DataYandexAlbBackendGroupStreamBackendTlsValidationContext) {
	_jsii_.InvokeVoid(
		d,
		"putValidationContext",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) ResetSni() {
	_jsii_.InvokeVoid(
		d,
		"resetSni",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) ResetValidationContext() {
	_jsii_.InvokeVoid(
		d,
		"resetValidationContext",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbBackendGroupStreamBackendTlsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

