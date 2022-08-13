// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexMdbGreenplumClusterPoolerConfigOutputReference interface {
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
	InternalValue() *DataYandexMdbGreenplumClusterPoolerConfig
	SetInternalValue(val *DataYandexMdbGreenplumClusterPoolerConfig)
	PoolClientIdleTimeout() *float64
	SetPoolClientIdleTimeout(val *float64)
	PoolClientIdleTimeoutInput() *float64
	PoolingMode() *string
	SetPoolingMode(val *string)
	PoolingModeInput() *string
	PoolSize() *float64
	SetPoolSize(val *float64)
	PoolSizeInput() *float64
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
	ResetPoolClientIdleTimeout()
	ResetPoolingMode()
	ResetPoolSize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataYandexMdbGreenplumClusterPoolerConfigOutputReference
type jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) InternalValue() *DataYandexMdbGreenplumClusterPoolerConfig {
	var returns *DataYandexMdbGreenplumClusterPoolerConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) PoolClientIdleTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"poolClientIdleTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) PoolClientIdleTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"poolClientIdleTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) PoolingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) PoolingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) PoolSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"poolSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) PoolSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"poolSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataYandexMdbGreenplumClusterPoolerConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataYandexMdbGreenplumClusterPoolerConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexMdbGreenplumClusterPoolerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataYandexMdbGreenplumClusterPoolerConfigOutputReference_Override(d DataYandexMdbGreenplumClusterPoolerConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexMdbGreenplumClusterPoolerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) SetInternalValue(val *DataYandexMdbGreenplumClusterPoolerConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) SetPoolClientIdleTimeout(val *float64) {
	_jsii_.Set(
		j,
		"poolClientIdleTimeout",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) SetPoolingMode(val *string) {
	_jsii_.Set(
		j,
		"poolingMode",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) SetPoolSize(val *float64) {
	_jsii_.Set(
		j,
		"poolSize",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) ResetPoolClientIdleTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetPoolClientIdleTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) ResetPoolingMode() {
	_jsii_.InvokeVoid(
		d,
		"resetPoolingMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) ResetPoolSize() {
	_jsii_.InvokeVoid(
		d,
		"resetPoolSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbGreenplumClusterPoolerConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

