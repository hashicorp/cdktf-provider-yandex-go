// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbBackendGroupHttpBackendTlsOutputReference interface {
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
	InternalValue() *AlbBackendGroupHttpBackendTls
	SetInternalValue(val *AlbBackendGroupHttpBackendTls)
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
	ValidationContext() AlbBackendGroupHttpBackendTlsValidationContextOutputReference
	ValidationContextInput() *AlbBackendGroupHttpBackendTlsValidationContext
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
	PutValidationContext(value *AlbBackendGroupHttpBackendTlsValidationContext)
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

// The jsii proxy struct for AlbBackendGroupHttpBackendTlsOutputReference
type jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) InternalValue() *AlbBackendGroupHttpBackendTls {
	var returns *AlbBackendGroupHttpBackendTls
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) Sni() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sni",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) SniInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sniInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) ValidationContext() AlbBackendGroupHttpBackendTlsValidationContextOutputReference {
	var returns AlbBackendGroupHttpBackendTlsValidationContextOutputReference
	_jsii_.Get(
		j,
		"validationContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) ValidationContextInput() *AlbBackendGroupHttpBackendTlsValidationContext {
	var returns *AlbBackendGroupHttpBackendTlsValidationContext
	_jsii_.Get(
		j,
		"validationContextInput",
		&returns,
	)
	return returns
}


func NewAlbBackendGroupHttpBackendTlsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlbBackendGroupHttpBackendTlsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbBackendGroupHttpBackendTlsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlbBackendGroupHttpBackendTlsOutputReference_Override(a AlbBackendGroupHttpBackendTlsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbBackendGroupHttpBackendTlsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) SetInternalValue(val *AlbBackendGroupHttpBackendTls) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) SetSni(val *string) {
	_jsii_.Set(
		j,
		"sni",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) PutValidationContext(value *AlbBackendGroupHttpBackendTlsValidationContext) {
	_jsii_.InvokeVoid(
		a,
		"putValidationContext",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) ResetSni() {
	_jsii_.InvokeVoid(
		a,
		"resetSni",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) ResetValidationContext() {
	_jsii_.InvokeVoid(
		a,
		"resetValidationContext",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupHttpBackendTlsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

