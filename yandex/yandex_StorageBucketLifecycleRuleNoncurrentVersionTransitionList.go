// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageBucketLifecycleRuleNoncurrentVersionTransitionList interface {
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
	Get(index *float64) StorageBucketLifecycleRuleNoncurrentVersionTransitionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageBucketLifecycleRuleNoncurrentVersionTransitionList
type jsiiProxy_StorageBucketLifecycleRuleNoncurrentVersionTransitionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_StorageBucketLifecycleRuleNoncurrentVersionTransitionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketLifecycleRuleNoncurrentVersionTransitionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketLifecycleRuleNoncurrentVersionTransitionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketLifecycleRuleNoncurrentVersionTransitionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketLifecycleRuleNoncurrentVersionTransitionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketLifecycleRuleNoncurrentVersionTransitionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewStorageBucketLifecycleRuleNoncurrentVersionTransitionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) StorageBucketLifecycleRuleNoncurrentVersionTransitionList {
	_init_.Initialize()

	j := jsiiProxy_StorageBucketLifecycleRuleNoncurrentVersionTransitionList{}

	_jsii_.Create(
		"@cdktf/provider-yandex.StorageBucketLifecycleRuleNoncurrentVersionTransitionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewStorageBucketLifecycleRuleNoncurrentVersionTransitionList_Override(s StorageBucketLifecycleRuleNoncurrentVersionTransitionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.StorageBucketLifecycleRuleNoncurrentVersionTransitionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_StorageBucketLifecycleRuleNoncurrentVersionTransitionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageBucketLifecycleRuleNoncurrentVersionTransitionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageBucketLifecycleRuleNoncurrentVersionTransitionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StorageBucketLifecycleRuleNoncurrentVersionTransitionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_StorageBucketLifecycleRuleNoncurrentVersionTransitionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketLifecycleRuleNoncurrentVersionTransitionList) Get(index *float64) StorageBucketLifecycleRuleNoncurrentVersionTransitionOutputReference {
	var returns StorageBucketLifecycleRuleNoncurrentVersionTransitionOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketLifecycleRuleNoncurrentVersionTransitionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketLifecycleRuleNoncurrentVersionTransitionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

