// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference interface {
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
	InternalValue() *DataYandexAlbVirtualHostRouteHttpRouteRedirectAction
	SetInternalValue(val *DataYandexAlbVirtualHostRouteHttpRouteRedirectAction)
	RemoveQuery() cdktf.IResolvable
	ReplaceHost() *string
	ReplacePath() *string
	ReplacePort() *float64
	ReplacePrefix() *string
	ReplaceScheme() *string
	ResponseCode() *string
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

// The jsii proxy struct for DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference
type jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) InternalValue() *DataYandexAlbVirtualHostRouteHttpRouteRedirectAction {
	var returns *DataYandexAlbVirtualHostRouteHttpRouteRedirectAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) RemoveQuery() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"removeQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ReplaceHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ReplacePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replacePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ReplacePort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replacePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ReplacePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replacePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ReplaceScheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceScheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ResponseCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference_Override(d DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) SetInternalValue(val *DataYandexAlbVirtualHostRouteHttpRouteRedirectAction) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexAlbVirtualHostRouteHttpRouteRedirectActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

