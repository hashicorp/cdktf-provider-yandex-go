// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbVirtualHostRouteHttpRouteRedirectActionOutputReference interface {
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
	InternalValue() *AlbVirtualHostRouteHttpRouteRedirectAction
	SetInternalValue(val *AlbVirtualHostRouteHttpRouteRedirectAction)
	RemoveQuery() interface{}
	SetRemoveQuery(val interface{})
	RemoveQueryInput() interface{}
	ReplaceHost() *string
	SetReplaceHost(val *string)
	ReplaceHostInput() *string
	ReplacePath() *string
	SetReplacePath(val *string)
	ReplacePathInput() *string
	ReplacePort() *float64
	SetReplacePort(val *float64)
	ReplacePortInput() *float64
	ReplacePrefix() *string
	SetReplacePrefix(val *string)
	ReplacePrefixInput() *string
	ReplaceScheme() *string
	SetReplaceScheme(val *string)
	ReplaceSchemeInput() *string
	ResponseCode() *string
	SetResponseCode(val *string)
	ResponseCodeInput() *string
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
	ResetRemoveQuery()
	ResetReplaceHost()
	ResetReplacePath()
	ResetReplacePort()
	ResetReplacePrefix()
	ResetReplaceScheme()
	ResetResponseCode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlbVirtualHostRouteHttpRouteRedirectActionOutputReference
type jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) InternalValue() *AlbVirtualHostRouteHttpRouteRedirectAction {
	var returns *AlbVirtualHostRouteHttpRouteRedirectAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) RemoveQuery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) RemoveQueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ReplaceHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ReplaceHostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ReplacePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replacePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ReplacePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replacePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ReplacePort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replacePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ReplacePortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replacePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ReplacePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replacePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ReplacePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replacePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ReplaceScheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceScheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ReplaceSchemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceSchemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ResponseCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ResponseCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAlbVirtualHostRouteHttpRouteRedirectActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlbVirtualHostRouteHttpRouteRedirectActionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbVirtualHostRouteHttpRouteRedirectActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlbVirtualHostRouteHttpRouteRedirectActionOutputReference_Override(a AlbVirtualHostRouteHttpRouteRedirectActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbVirtualHostRouteHttpRouteRedirectActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) SetInternalValue(val *AlbVirtualHostRouteHttpRouteRedirectAction) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) SetRemoveQuery(val interface{}) {
	_jsii_.Set(
		j,
		"removeQuery",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) SetReplaceHost(val *string) {
	_jsii_.Set(
		j,
		"replaceHost",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) SetReplacePath(val *string) {
	_jsii_.Set(
		j,
		"replacePath",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) SetReplacePort(val *float64) {
	_jsii_.Set(
		j,
		"replacePort",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) SetReplacePrefix(val *string) {
	_jsii_.Set(
		j,
		"replacePrefix",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) SetReplaceScheme(val *string) {
	_jsii_.Set(
		j,
		"replaceScheme",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) SetResponseCode(val *string) {
	_jsii_.Set(
		j,
		"responseCode",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ResetRemoveQuery() {
	_jsii_.InvokeVoid(
		a,
		"resetRemoveQuery",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ResetReplaceHost() {
	_jsii_.InvokeVoid(
		a,
		"resetReplaceHost",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ResetReplacePath() {
	_jsii_.InvokeVoid(
		a,
		"resetReplacePath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ResetReplacePort() {
	_jsii_.InvokeVoid(
		a,
		"resetReplacePort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ResetReplacePrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetReplacePrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ResetReplaceScheme() {
	_jsii_.InvokeVoid(
		a,
		"resetReplaceScheme",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ResetResponseCode() {
	_jsii_.InvokeVoid(
		a,
		"resetResponseCode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

