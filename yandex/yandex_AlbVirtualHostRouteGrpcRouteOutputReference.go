// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbVirtualHostRouteGrpcRouteOutputReference interface {
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
	GrpcMatch() AlbVirtualHostRouteGrpcRouteGrpcMatchList
	GrpcMatchInput() interface{}
	GrpcRouteAction() AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference
	GrpcRouteActionInput() *AlbVirtualHostRouteGrpcRouteGrpcRouteAction
	GrpcStatusResponseAction() AlbVirtualHostRouteGrpcRouteGrpcStatusResponseActionOutputReference
	GrpcStatusResponseActionInput() *AlbVirtualHostRouteGrpcRouteGrpcStatusResponseAction
	InternalValue() *AlbVirtualHostRouteGrpcRoute
	SetInternalValue(val *AlbVirtualHostRouteGrpcRoute)
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
	PutGrpcMatch(value interface{})
	PutGrpcRouteAction(value *AlbVirtualHostRouteGrpcRouteGrpcRouteAction)
	PutGrpcStatusResponseAction(value *AlbVirtualHostRouteGrpcRouteGrpcStatusResponseAction)
	ResetGrpcMatch()
	ResetGrpcRouteAction()
	ResetGrpcStatusResponseAction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlbVirtualHostRouteGrpcRouteOutputReference
type jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) GrpcMatch() AlbVirtualHostRouteGrpcRouteGrpcMatchList {
	var returns AlbVirtualHostRouteGrpcRouteGrpcMatchList
	_jsii_.Get(
		j,
		"grpcMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) GrpcMatchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"grpcMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) GrpcRouteAction() AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference {
	var returns AlbVirtualHostRouteGrpcRouteGrpcRouteActionOutputReference
	_jsii_.Get(
		j,
		"grpcRouteAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) GrpcRouteActionInput() *AlbVirtualHostRouteGrpcRouteGrpcRouteAction {
	var returns *AlbVirtualHostRouteGrpcRouteGrpcRouteAction
	_jsii_.Get(
		j,
		"grpcRouteActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) GrpcStatusResponseAction() AlbVirtualHostRouteGrpcRouteGrpcStatusResponseActionOutputReference {
	var returns AlbVirtualHostRouteGrpcRouteGrpcStatusResponseActionOutputReference
	_jsii_.Get(
		j,
		"grpcStatusResponseAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) GrpcStatusResponseActionInput() *AlbVirtualHostRouteGrpcRouteGrpcStatusResponseAction {
	var returns *AlbVirtualHostRouteGrpcRouteGrpcStatusResponseAction
	_jsii_.Get(
		j,
		"grpcStatusResponseActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) InternalValue() *AlbVirtualHostRouteGrpcRoute {
	var returns *AlbVirtualHostRouteGrpcRoute
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAlbVirtualHostRouteGrpcRouteOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlbVirtualHostRouteGrpcRouteOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbVirtualHostRouteGrpcRouteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlbVirtualHostRouteGrpcRouteOutputReference_Override(a AlbVirtualHostRouteGrpcRouteOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbVirtualHostRouteGrpcRouteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) SetInternalValue(val *AlbVirtualHostRouteGrpcRoute) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) PutGrpcMatch(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putGrpcMatch",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) PutGrpcRouteAction(value *AlbVirtualHostRouteGrpcRouteGrpcRouteAction) {
	_jsii_.InvokeVoid(
		a,
		"putGrpcRouteAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) PutGrpcStatusResponseAction(value *AlbVirtualHostRouteGrpcRouteGrpcStatusResponseAction) {
	_jsii_.InvokeVoid(
		a,
		"putGrpcStatusResponseAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) ResetGrpcMatch() {
	_jsii_.InvokeVoid(
		a,
		"resetGrpcMatch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) ResetGrpcRouteAction() {
	_jsii_.InvokeVoid(
		a,
		"resetGrpcRouteAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) ResetGrpcStatusResponseAction() {
	_jsii_.InvokeVoid(
		a,
		"resetGrpcStatusResponseAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

