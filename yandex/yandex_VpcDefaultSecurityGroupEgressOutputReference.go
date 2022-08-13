// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpcDefaultSecurityGroupEgressOutputReference interface {
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	FromPort() *float64
	SetFromPort(val *float64)
	FromPortInput() *float64
	Id() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	PredefinedTarget() *string
	SetPredefinedTarget(val *string)
	PredefinedTargetInput() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	SecurityGroupId() *string
	SetSecurityGroupId(val *string)
	SecurityGroupIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ToPort() *float64
	SetToPort(val *float64)
	ToPortInput() *float64
	V4CidrBlocks() *[]*string
	SetV4CidrBlocks(val *[]*string)
	V4CidrBlocksInput() *[]*string
	V6CidrBlocks() *[]*string
	SetV6CidrBlocks(val *[]*string)
	V6CidrBlocksInput() *[]*string
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
	ResetDescription()
	ResetFromPort()
	ResetLabels()
	ResetPort()
	ResetPredefinedTarget()
	ResetSecurityGroupId()
	ResetToPort()
	ResetV4CidrBlocks()
	ResetV6CidrBlocks()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VpcDefaultSecurityGroupEgressOutputReference
type jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) FromPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fromPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) FromPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fromPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) PredefinedTarget() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predefinedTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) PredefinedTargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predefinedTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) SecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) SecurityGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) ToPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"toPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) ToPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"toPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) V4CidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"v4CidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) V4CidrBlocksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"v4CidrBlocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) V6CidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"v6CidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) V6CidrBlocksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"v6CidrBlocksInput",
		&returns,
	)
	return returns
}


func NewVpcDefaultSecurityGroupEgressOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VpcDefaultSecurityGroupEgressOutputReference {
	_init_.Initialize()

	j := jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.VpcDefaultSecurityGroupEgressOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVpcDefaultSecurityGroupEgressOutputReference_Override(v VpcDefaultSecurityGroupEgressOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.VpcDefaultSecurityGroupEgressOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) SetFromPort(val *float64) {
	_jsii_.Set(
		j,
		"fromPort",
		val,
	)
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) SetLabels(val *map[string]*string) {
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) SetPredefinedTarget(val *string) {
	_jsii_.Set(
		j,
		"predefinedTarget",
		val,
	)
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) SetProtocol(val *string) {
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) SetSecurityGroupId(val *string) {
	_jsii_.Set(
		j,
		"securityGroupId",
		val,
	)
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) SetToPort(val *float64) {
	_jsii_.Set(
		j,
		"toPort",
		val,
	)
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) SetV4CidrBlocks(val *[]*string) {
	_jsii_.Set(
		j,
		"v4CidrBlocks",
		val,
	)
}

func (j *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) SetV6CidrBlocks(val *[]*string) {
	_jsii_.Set(
		j,
		"v6CidrBlocks",
		val,
	)
}

func (v *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		v,
		"resetDescription",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) ResetFromPort() {
	_jsii_.InvokeVoid(
		v,
		"resetFromPort",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		v,
		"resetLabels",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		v,
		"resetPort",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) ResetPredefinedTarget() {
	_jsii_.InvokeVoid(
		v,
		"resetPredefinedTarget",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) ResetSecurityGroupId() {
	_jsii_.InvokeVoid(
		v,
		"resetSecurityGroupId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) ResetToPort() {
	_jsii_.InvokeVoid(
		v,
		"resetToPort",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) ResetV4CidrBlocks() {
	_jsii_.InvokeVoid(
		v,
		"resetV4CidrBlocks",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) ResetV6CidrBlocks() {
	_jsii_.InvokeVoid(
		v,
		"resetV6CidrBlocks",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcDefaultSecurityGroupEgressOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

