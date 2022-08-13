// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbMongodbClusterClusterConfigMongodOutputReference interface {
	cdktf.ComplexObject
	AuditLog() MdbMongodbClusterClusterConfigMongodAuditLogOutputReference
	AuditLogInput() *MdbMongodbClusterClusterConfigMongodAuditLog
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
	InternalValue() *MdbMongodbClusterClusterConfigMongod
	SetInternalValue(val *MdbMongodbClusterClusterConfigMongod)
	Security() MdbMongodbClusterClusterConfigMongodSecurityOutputReference
	SecurityInput() *MdbMongodbClusterClusterConfigMongodSecurity
	SetParameter() MdbMongodbClusterClusterConfigMongodSetParameterOutputReference
	SetParameterInput() *MdbMongodbClusterClusterConfigMongodSetParameter
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
	PutAuditLog(value *MdbMongodbClusterClusterConfigMongodAuditLog)
	PutSecurity(value *MdbMongodbClusterClusterConfigMongodSecurity)
	PutSetParameter(value *MdbMongodbClusterClusterConfigMongodSetParameter)
	ResetAuditLog()
	ResetSecurity()
	ResetSetParameter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MdbMongodbClusterClusterConfigMongodOutputReference
type jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) AuditLog() MdbMongodbClusterClusterConfigMongodAuditLogOutputReference {
	var returns MdbMongodbClusterClusterConfigMongodAuditLogOutputReference
	_jsii_.Get(
		j,
		"auditLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) AuditLogInput() *MdbMongodbClusterClusterConfigMongodAuditLog {
	var returns *MdbMongodbClusterClusterConfigMongodAuditLog
	_jsii_.Get(
		j,
		"auditLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) InternalValue() *MdbMongodbClusterClusterConfigMongod {
	var returns *MdbMongodbClusterClusterConfigMongod
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) Security() MdbMongodbClusterClusterConfigMongodSecurityOutputReference {
	var returns MdbMongodbClusterClusterConfigMongodSecurityOutputReference
	_jsii_.Get(
		j,
		"security",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) SecurityInput() *MdbMongodbClusterClusterConfigMongodSecurity {
	var returns *MdbMongodbClusterClusterConfigMongodSecurity
	_jsii_.Get(
		j,
		"securityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) SetParameter() MdbMongodbClusterClusterConfigMongodSetParameterOutputReference {
	var returns MdbMongodbClusterClusterConfigMongodSetParameterOutputReference
	_jsii_.Get(
		j,
		"setParameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) SetParameterInput() *MdbMongodbClusterClusterConfigMongodSetParameter {
	var returns *MdbMongodbClusterClusterConfigMongodSetParameter
	_jsii_.Get(
		j,
		"setParameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMdbMongodbClusterClusterConfigMongodOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MdbMongodbClusterClusterConfigMongodOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbMongodbClusterClusterConfigMongodOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMdbMongodbClusterClusterConfigMongodOutputReference_Override(m MdbMongodbClusterClusterConfigMongodOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbMongodbClusterClusterConfigMongodOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) SetInternalValue(val *MdbMongodbClusterClusterConfigMongod) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) PutAuditLog(value *MdbMongodbClusterClusterConfigMongodAuditLog) {
	_jsii_.InvokeVoid(
		m,
		"putAuditLog",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) PutSecurity(value *MdbMongodbClusterClusterConfigMongodSecurity) {
	_jsii_.InvokeVoid(
		m,
		"putSecurity",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) PutSetParameter(value *MdbMongodbClusterClusterConfigMongodSetParameter) {
	_jsii_.InvokeVoid(
		m,
		"putSetParameter",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) ResetAuditLog() {
	_jsii_.InvokeVoid(
		m,
		"resetAuditLog",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) ResetSecurity() {
	_jsii_.InvokeVoid(
		m,
		"resetSecurity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) ResetSetParameter() {
	_jsii_.InvokeVoid(
		m,
		"resetSetParameter",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

