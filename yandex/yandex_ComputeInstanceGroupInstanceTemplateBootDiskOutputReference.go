// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeInstanceGroupInstanceTemplateBootDiskOutputReference interface {
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
	DeviceName() *string
	SetDeviceName(val *string)
	DeviceNameInput() *string
	DiskId() *string
	SetDiskId(val *string)
	DiskIdInput() *string
	// Experimental.
	Fqn() *string
	InitializeParams() ComputeInstanceGroupInstanceTemplateBootDiskInitializeParamsOutputReference
	InitializeParamsInput() *ComputeInstanceGroupInstanceTemplateBootDiskInitializeParams
	InternalValue() *ComputeInstanceGroupInstanceTemplateBootDisk
	SetInternalValue(val *ComputeInstanceGroupInstanceTemplateBootDisk)
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
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
	PutInitializeParams(value *ComputeInstanceGroupInstanceTemplateBootDiskInitializeParams)
	ResetDeviceName()
	ResetDiskId()
	ResetInitializeParams()
	ResetMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeInstanceGroupInstanceTemplateBootDiskOutputReference
type jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) DeviceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) DeviceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) DiskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) DiskIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) InitializeParams() ComputeInstanceGroupInstanceTemplateBootDiskInitializeParamsOutputReference {
	var returns ComputeInstanceGroupInstanceTemplateBootDiskInitializeParamsOutputReference
	_jsii_.Get(
		j,
		"initializeParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) InitializeParamsInput() *ComputeInstanceGroupInstanceTemplateBootDiskInitializeParams {
	var returns *ComputeInstanceGroupInstanceTemplateBootDiskInitializeParams
	_jsii_.Get(
		j,
		"initializeParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) InternalValue() *ComputeInstanceGroupInstanceTemplateBootDisk {
	var returns *ComputeInstanceGroupInstanceTemplateBootDisk
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeInstanceGroupInstanceTemplateBootDiskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeInstanceGroupInstanceTemplateBootDiskOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.ComputeInstanceGroupInstanceTemplateBootDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeInstanceGroupInstanceTemplateBootDiskOutputReference_Override(c ComputeInstanceGroupInstanceTemplateBootDiskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.ComputeInstanceGroupInstanceTemplateBootDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) SetDeviceName(val *string) {
	_jsii_.Set(
		j,
		"deviceName",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) SetDiskId(val *string) {
	_jsii_.Set(
		j,
		"diskId",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) SetInternalValue(val *ComputeInstanceGroupInstanceTemplateBootDisk) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) SetMode(val *string) {
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) PutInitializeParams(value *ComputeInstanceGroupInstanceTemplateBootDiskInitializeParams) {
	_jsii_.InvokeVoid(
		c,
		"putInitializeParams",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) ResetDeviceName() {
	_jsii_.InvokeVoid(
		c,
		"resetDeviceName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) ResetDiskId() {
	_jsii_.InvokeVoid(
		c,
		"resetDiskId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) ResetInitializeParams() {
	_jsii_.InvokeVoid(
		c,
		"resetInitializeParams",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		c,
		"resetMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupInstanceTemplateBootDiskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

