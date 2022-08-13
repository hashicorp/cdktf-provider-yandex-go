// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbVirtualHostRouteGrpcRouteGrpcMatchList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) AlbVirtualHostRouteGrpcRouteGrpcMatchOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlbVirtualHostRouteGrpcRouteGrpcMatchList
type jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcMatchList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcMatchList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcMatchList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcMatchList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcMatchList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcMatchList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcMatchList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewAlbVirtualHostRouteGrpcRouteGrpcMatchList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AlbVirtualHostRouteGrpcRouteGrpcMatchList {
	_init_.Initialize()

	j := jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcMatchList{}

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbVirtualHostRouteGrpcRouteGrpcMatchList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewAlbVirtualHostRouteGrpcRouteGrpcMatchList_Override(a AlbVirtualHostRouteGrpcRouteGrpcMatchList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.AlbVirtualHostRouteGrpcRouteGrpcMatchList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcMatchList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcMatchList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcMatchList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcMatchList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcMatchList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcMatchList) Get(index *float64) AlbVirtualHostRouteGrpcRouteGrpcMatchOutputReference {
	var returns AlbVirtualHostRouteGrpcRouteGrpcMatchOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcMatchList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbVirtualHostRouteGrpcRouteGrpcMatchList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

