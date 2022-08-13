// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageBucketLifecycleRuleOutputReference interface {
	cdktf.ComplexObject
	AbortIncompleteMultipartUploadDays() *float64
	SetAbortIncompleteMultipartUploadDays(val *float64)
	AbortIncompleteMultipartUploadDaysInput() *float64
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Expiration() StorageBucketLifecycleRuleExpirationOutputReference
	ExpirationInput() *StorageBucketLifecycleRuleExpiration
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NoncurrentVersionExpiration() StorageBucketLifecycleRuleNoncurrentVersionExpirationOutputReference
	NoncurrentVersionExpirationInput() *StorageBucketLifecycleRuleNoncurrentVersionExpiration
	NoncurrentVersionTransition() StorageBucketLifecycleRuleNoncurrentVersionTransitionList
	NoncurrentVersionTransitionInput() interface{}
	Prefix() *string
	SetPrefix(val *string)
	PrefixInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Transition() StorageBucketLifecycleRuleTransitionList
	TransitionInput() interface{}
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
	PutExpiration(value *StorageBucketLifecycleRuleExpiration)
	PutNoncurrentVersionExpiration(value *StorageBucketLifecycleRuleNoncurrentVersionExpiration)
	PutNoncurrentVersionTransition(value interface{})
	PutTransition(value interface{})
	ResetAbortIncompleteMultipartUploadDays()
	ResetExpiration()
	ResetId()
	ResetNoncurrentVersionExpiration()
	ResetNoncurrentVersionTransition()
	ResetPrefix()
	ResetTransition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageBucketLifecycleRuleOutputReference
type jsiiProxy_StorageBucketLifecycleRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) AbortIncompleteMultipartUploadDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"abortIncompleteMultipartUploadDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) AbortIncompleteMultipartUploadDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"abortIncompleteMultipartUploadDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) Expiration() StorageBucketLifecycleRuleExpirationOutputReference {
	var returns StorageBucketLifecycleRuleExpirationOutputReference
	_jsii_.Get(
		j,
		"expiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) ExpirationInput() *StorageBucketLifecycleRuleExpiration {
	var returns *StorageBucketLifecycleRuleExpiration
	_jsii_.Get(
		j,
		"expirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) NoncurrentVersionExpiration() StorageBucketLifecycleRuleNoncurrentVersionExpirationOutputReference {
	var returns StorageBucketLifecycleRuleNoncurrentVersionExpirationOutputReference
	_jsii_.Get(
		j,
		"noncurrentVersionExpiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) NoncurrentVersionExpirationInput() *StorageBucketLifecycleRuleNoncurrentVersionExpiration {
	var returns *StorageBucketLifecycleRuleNoncurrentVersionExpiration
	_jsii_.Get(
		j,
		"noncurrentVersionExpirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) NoncurrentVersionTransition() StorageBucketLifecycleRuleNoncurrentVersionTransitionList {
	var returns StorageBucketLifecycleRuleNoncurrentVersionTransitionList
	_jsii_.Get(
		j,
		"noncurrentVersionTransition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) NoncurrentVersionTransitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noncurrentVersionTransitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) Transition() StorageBucketLifecycleRuleTransitionList {
	var returns StorageBucketLifecycleRuleTransitionList
	_jsii_.Get(
		j,
		"transition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) TransitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transitionInput",
		&returns,
	)
	return returns
}


func NewStorageBucketLifecycleRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) StorageBucketLifecycleRuleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageBucketLifecycleRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.StorageBucketLifecycleRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewStorageBucketLifecycleRuleOutputReference_Override(s StorageBucketLifecycleRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.StorageBucketLifecycleRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) SetAbortIncompleteMultipartUploadDays(val *float64) {
	_jsii_.Set(
		j,
		"abortIncompleteMultipartUploadDays",
		val,
	)
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) SetPrefix(val *string) {
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageBucketLifecycleRuleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageBucketLifecycleRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketLifecycleRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketLifecycleRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketLifecycleRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketLifecycleRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketLifecycleRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketLifecycleRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketLifecycleRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketLifecycleRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketLifecycleRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketLifecycleRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketLifecycleRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketLifecycleRuleOutputReference) PutExpiration(value *StorageBucketLifecycleRuleExpiration) {
	_jsii_.InvokeVoid(
		s,
		"putExpiration",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageBucketLifecycleRuleOutputReference) PutNoncurrentVersionExpiration(value *StorageBucketLifecycleRuleNoncurrentVersionExpiration) {
	_jsii_.InvokeVoid(
		s,
		"putNoncurrentVersionExpiration",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageBucketLifecycleRuleOutputReference) PutNoncurrentVersionTransition(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putNoncurrentVersionTransition",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageBucketLifecycleRuleOutputReference) PutTransition(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putTransition",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageBucketLifecycleRuleOutputReference) ResetAbortIncompleteMultipartUploadDays() {
	_jsii_.InvokeVoid(
		s,
		"resetAbortIncompleteMultipartUploadDays",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucketLifecycleRuleOutputReference) ResetExpiration() {
	_jsii_.InvokeVoid(
		s,
		"resetExpiration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucketLifecycleRuleOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucketLifecycleRuleOutputReference) ResetNoncurrentVersionExpiration() {
	_jsii_.InvokeVoid(
		s,
		"resetNoncurrentVersionExpiration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucketLifecycleRuleOutputReference) ResetNoncurrentVersionTransition() {
	_jsii_.InvokeVoid(
		s,
		"resetNoncurrentVersionTransition",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucketLifecycleRuleOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		s,
		"resetPrefix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucketLifecycleRuleOutputReference) ResetTransition() {
	_jsii_.InvokeVoid(
		s,
		"resetTransition",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucketLifecycleRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketLifecycleRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

