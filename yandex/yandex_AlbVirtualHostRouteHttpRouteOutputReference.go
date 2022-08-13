// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbVirtualHostRouteHttpRouteOutputReference interface {
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
	DirectResponseAction() AlbVirtualHostRouteHttpRouteDirectResponseActionOutputReference
	DirectResponseActionInput() *AlbVirtualHostRouteHttpRouteDirectResponseAction
	// Experimental.
	Fqn() *string
	HttpMatch() AlbVirtualHostRouteHttpRouteHttpMatchList
	HttpMatchInput() interface{}
	HttpRouteAction() AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference
	HttpRouteActionInput() *AlbVirtualHostRouteHttpRouteHttpRouteAction
	InternalValue() *AlbVirtualHostRouteHttpRoute
	SetInternalValue(val *AlbVirtualHostRouteHttpRoute)
	RedirectAction() AlbVirtualHostRouteHttpRouteRedirectActionOutputReference
	RedirectActionInput() *AlbVirtualHostRouteHttpRouteRedirectAction
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
	PutDirectResponseAction(value *AlbVirtualHostRouteHttpRouteDirectResponseAction)
	PutHttpMatch(value interface{})
	PutHttpRouteAction(value *AlbVirtualHostRouteHttpRouteHttpRouteAction)
	PutRedirectAction(value *AlbVirtualHostRouteHttpRouteRedirectAction)
	ResetDirectResponseAction()
	ResetHttpMatch()
	ResetHttpRouteAction()
	ResetRedirectAction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlbVirtualHostRouteHttpRouteOutputReference
type jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) DirectResponseAction() AlbVirtualHostRouteHttpRouteDirectResponseActionOutputReference {
	var returns AlbVirtualHostRouteHttpRouteDirectResponseActionOutputReference
	_jsii_.Get(
		j,
		"directResponseAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) DirectResponseActionInput() *AlbVirtualHostRouteHttpRouteDirectResponseAction {
	var returns *AlbVirtualHostRouteHttpRouteDirectResponseAction
	_jsii_.Get(
		j,
		"directResponseActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) HttpMatch() AlbVirtualHostRouteHttpRouteHttpMatchList {
	var returns AlbVirtualHostRouteHttpRouteHttpMatchList
	_jsii_.Get(
		j,
		"httpMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) HttpMatchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) HttpRouteAction() AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference {
	var returns AlbVirtualHostRouteHttpRouteHttpRouteActionOutputReference
	_jsii_.Get(
		j,
		"httpRouteAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) HttpRouteActionInput() *AlbVirtualHostRouteHttpRouteHttpRouteAction {
	var returns *AlbVirtualHostRouteHttpRouteHttpRouteAction
	_jsii_.Get(
		j,
		"httpRouteActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) InternalValue() *AlbVirtualHostRouteHttpRoute {
	var returns *AlbVirtualHostRouteHttpRoute
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) RedirectAction() AlbVirtualHostRouteHttpRouteRedirectActionOutputReference {
	var returns AlbVirtualHostRouteHttpRouteRedirectActionOutputReference
	_jsii_.Get(
		j,
		"redirectAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) RedirectActionInput() *AlbVirtualHostRouteHttpRouteRedirectAction {
	var returns *AlbVirtualHostRouteHttpRouteRedirectAction
	_jsii_.Get(
		j,
		"redirectActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAlbVirtualHostRouteHttpRouteOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlbVirtualHostRouteHttpRouteOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbVirtualHostRouteHttpRouteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlbVirtualHostRouteHttpRouteOutputReference_Override(a AlbVirtualHostRouteHttpRouteOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbVirtualHostRouteHttpRouteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) SetInternalValue(val *AlbVirtualHostRouteHttpRoute) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) PutDirectResponseAction(value *AlbVirtualHostRouteHttpRouteDirectResponseAction) {
	_jsii_.InvokeVoid(
		a,
		"putDirectResponseAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) PutHttpMatch(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putHttpMatch",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) PutHttpRouteAction(value *AlbVirtualHostRouteHttpRouteHttpRouteAction) {
	_jsii_.InvokeVoid(
		a,
		"putHttpRouteAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) PutRedirectAction(value *AlbVirtualHostRouteHttpRouteRedirectAction) {
	_jsii_.InvokeVoid(
		a,
		"putRedirectAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) ResetDirectResponseAction() {
	_jsii_.InvokeVoid(
		a,
		"resetDirectResponseAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) ResetHttpMatch() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpMatch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) ResetHttpRouteAction() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpRouteAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) ResetRedirectAction() {
	_jsii_.InvokeVoid(
		a,
		"resetRedirectAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

