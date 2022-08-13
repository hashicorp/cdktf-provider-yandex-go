// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbBackendGroupGrpcBackendTlsValidationContextOutputReference interface {
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
	InternalValue() *AlbBackendGroupGrpcBackendTlsValidationContext
	SetInternalValue(val *AlbBackendGroupGrpcBackendTlsValidationContext)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrustedCaBytes() *string
	SetTrustedCaBytes(val *string)
	TrustedCaBytesInput() *string
	TrustedCaId() *string
	SetTrustedCaId(val *string)
	TrustedCaIdInput() *string
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
	ResetTrustedCaBytes()
	ResetTrustedCaId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlbBackendGroupGrpcBackendTlsValidationContextOutputReference
type jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) InternalValue() *AlbBackendGroupGrpcBackendTlsValidationContext {
	var returns *AlbBackendGroupGrpcBackendTlsValidationContext
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) TrustedCaBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustedCaBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) TrustedCaBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustedCaBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) TrustedCaId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustedCaId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) TrustedCaIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustedCaIdInput",
		&returns,
	)
	return returns
}


func NewAlbBackendGroupGrpcBackendTlsValidationContextOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlbBackendGroupGrpcBackendTlsValidationContextOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbBackendGroupGrpcBackendTlsValidationContextOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlbBackendGroupGrpcBackendTlsValidationContextOutputReference_Override(a AlbBackendGroupGrpcBackendTlsValidationContextOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbBackendGroupGrpcBackendTlsValidationContextOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) SetInternalValue(val *AlbBackendGroupGrpcBackendTlsValidationContext) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) SetTrustedCaBytes(val *string) {
	_jsii_.Set(
		j,
		"trustedCaBytes",
		val,
	)
}

func (j *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) SetTrustedCaId(val *string) {
	_jsii_.Set(
		j,
		"trustedCaId",
		val,
	)
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) ResetTrustedCaBytes() {
	_jsii_.InvokeVoid(
		a,
		"resetTrustedCaBytes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) ResetTrustedCaId() {
	_jsii_.InvokeVoid(
		a,
		"resetTrustedCaId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbBackendGroupGrpcBackendTlsValidationContextOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

