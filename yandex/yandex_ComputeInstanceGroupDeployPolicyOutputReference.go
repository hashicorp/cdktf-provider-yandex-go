// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeInstanceGroupDeployPolicyOutputReference interface {
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
	InternalValue() *ComputeInstanceGroupDeployPolicy
	SetInternalValue(val *ComputeInstanceGroupDeployPolicy)
	MaxCreating() *float64
	SetMaxCreating(val *float64)
	MaxCreatingInput() *float64
	MaxDeleting() *float64
	SetMaxDeleting(val *float64)
	MaxDeletingInput() *float64
	MaxExpansion() *float64
	SetMaxExpansion(val *float64)
	MaxExpansionInput() *float64
	MaxUnavailable() *float64
	SetMaxUnavailable(val *float64)
	MaxUnavailableInput() *float64
	StartupDuration() *float64
	SetStartupDuration(val *float64)
	StartupDurationInput() *float64
	Strategy() *string
	SetStrategy(val *string)
	StrategyInput() *string
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
	ResetMaxCreating()
	ResetMaxDeleting()
	ResetStartupDuration()
	ResetStrategy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeInstanceGroupDeployPolicyOutputReference
type jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) InternalValue() *ComputeInstanceGroupDeployPolicy {
	var returns *ComputeInstanceGroupDeployPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) MaxCreating() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCreating",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) MaxCreatingInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCreatingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) MaxDeleting() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDeleting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) MaxDeletingInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDeletingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) MaxExpansion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxExpansion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) MaxExpansionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxExpansionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) MaxUnavailable() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnavailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) MaxUnavailableInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnavailableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) StartupDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startupDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) StartupDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startupDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) Strategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) StrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeInstanceGroupDeployPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeInstanceGroupDeployPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.ComputeInstanceGroupDeployPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeInstanceGroupDeployPolicyOutputReference_Override(c ComputeInstanceGroupDeployPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.ComputeInstanceGroupDeployPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) SetInternalValue(val *ComputeInstanceGroupDeployPolicy) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) SetMaxCreating(val *float64) {
	_jsii_.Set(
		j,
		"maxCreating",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) SetMaxDeleting(val *float64) {
	_jsii_.Set(
		j,
		"maxDeleting",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) SetMaxExpansion(val *float64) {
	_jsii_.Set(
		j,
		"maxExpansion",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) SetMaxUnavailable(val *float64) {
	_jsii_.Set(
		j,
		"maxUnavailable",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) SetStartupDuration(val *float64) {
	_jsii_.Set(
		j,
		"startupDuration",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) SetStrategy(val *string) {
	_jsii_.Set(
		j,
		"strategy",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) ResetMaxCreating() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxCreating",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) ResetMaxDeleting() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxDeleting",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) ResetStartupDuration() {
	_jsii_.InvokeVoid(
		c,
		"resetStartupDuration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) ResetStrategy() {
	_jsii_.InvokeVoid(
		c,
		"resetStrategy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupDeployPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

