// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference interface {
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
	InternalValue() *StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault
	SetInternalValue(val *StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault)
	KmsMasterKeyId() *string
	SetKmsMasterKeyId(val *string)
	KmsMasterKeyIdInput() *string
	SseAlgorithm() *string
	SetSseAlgorithm(val *string)
	SseAlgorithmInput() *string
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

// The jsii proxy struct for StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference
type jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) InternalValue() *StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault {
	var returns *StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) KmsMasterKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsMasterKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) KmsMasterKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsMasterKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) SseAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sseAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) SseAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sseAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference_Override(s StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) SetInternalValue(val *StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) SetKmsMasterKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsMasterKeyId",
		val,
	)
}

func (j *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) SetSseAlgorithm(val *string) {
	_jsii_.Set(
		j,
		"sseAlgorithm",
		val,
	)
}

func (j *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

