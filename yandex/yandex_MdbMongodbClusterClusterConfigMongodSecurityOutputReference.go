// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbMongodbClusterClusterConfigMongodSecurityOutputReference interface {
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
	EnableEncryption() interface{}
	SetEnableEncryption(val interface{})
	EnableEncryptionInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *MdbMongodbClusterClusterConfigMongodSecurity
	SetInternalValue(val *MdbMongodbClusterClusterConfigMongodSecurity)
	Kmip() MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference
	KmipInput() *MdbMongodbClusterClusterConfigMongodSecurityKmip
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
	PutKmip(value *MdbMongodbClusterClusterConfigMongodSecurityKmip)
	ResetEnableEncryption()
	ResetKmip()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MdbMongodbClusterClusterConfigMongodSecurityOutputReference
type jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) EnableEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) EnableEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) InternalValue() *MdbMongodbClusterClusterConfigMongodSecurity {
	var returns *MdbMongodbClusterClusterConfigMongodSecurity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) Kmip() MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference {
	var returns MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference
	_jsii_.Get(
		j,
		"kmip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) KmipInput() *MdbMongodbClusterClusterConfigMongodSecurityKmip {
	var returns *MdbMongodbClusterClusterConfigMongodSecurityKmip
	_jsii_.Get(
		j,
		"kmipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMdbMongodbClusterClusterConfigMongodSecurityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MdbMongodbClusterClusterConfigMongodSecurityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbMongodbClusterClusterConfigMongodSecurityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMdbMongodbClusterClusterConfigMongodSecurityOutputReference_Override(m MdbMongodbClusterClusterConfigMongodSecurityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbMongodbClusterClusterConfigMongodSecurityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) SetEnableEncryption(val interface{}) {
	_jsii_.Set(
		j,
		"enableEncryption",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) SetInternalValue(val *MdbMongodbClusterClusterConfigMongodSecurity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) PutKmip(value *MdbMongodbClusterClusterConfigMongodSecurityKmip) {
	_jsii_.InvokeVoid(
		m,
		"putKmip",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) ResetEnableEncryption() {
	_jsii_.InvokeVoid(
		m,
		"resetEnableEncryption",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) ResetKmip() {
	_jsii_.InvokeVoid(
		m,
		"resetKmip",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

