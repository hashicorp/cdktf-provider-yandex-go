// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexMdbMongodbClusterClusterConfigMongodOutputReference interface {
	cdktf.ComplexObject
	AuditLog() DataYandexMdbMongodbClusterClusterConfigMongodAuditLogOutputReference
	AuditLogInput() *DataYandexMdbMongodbClusterClusterConfigMongodAuditLog
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
	InternalValue() *DataYandexMdbMongodbClusterClusterConfigMongod
	SetInternalValue(val *DataYandexMdbMongodbClusterClusterConfigMongod)
	Security() DataYandexMdbMongodbClusterClusterConfigMongodSecurityOutputReference
	SecurityInput() *DataYandexMdbMongodbClusterClusterConfigMongodSecurity
	SetParameter() DataYandexMdbMongodbClusterClusterConfigMongodSetParameterOutputReference
	SetParameterInput() *DataYandexMdbMongodbClusterClusterConfigMongodSetParameter
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
	PutAuditLog(value *DataYandexMdbMongodbClusterClusterConfigMongodAuditLog)
	PutSecurity(value *DataYandexMdbMongodbClusterClusterConfigMongodSecurity)
	PutSetParameter(value *DataYandexMdbMongodbClusterClusterConfigMongodSetParameter)
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

// The jsii proxy struct for DataYandexMdbMongodbClusterClusterConfigMongodOutputReference
type jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) AuditLog() DataYandexMdbMongodbClusterClusterConfigMongodAuditLogOutputReference {
	var returns DataYandexMdbMongodbClusterClusterConfigMongodAuditLogOutputReference
	_jsii_.Get(
		j,
		"auditLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) AuditLogInput() *DataYandexMdbMongodbClusterClusterConfigMongodAuditLog {
	var returns *DataYandexMdbMongodbClusterClusterConfigMongodAuditLog
	_jsii_.Get(
		j,
		"auditLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) InternalValue() *DataYandexMdbMongodbClusterClusterConfigMongod {
	var returns *DataYandexMdbMongodbClusterClusterConfigMongod
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) Security() DataYandexMdbMongodbClusterClusterConfigMongodSecurityOutputReference {
	var returns DataYandexMdbMongodbClusterClusterConfigMongodSecurityOutputReference
	_jsii_.Get(
		j,
		"security",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) SecurityInput() *DataYandexMdbMongodbClusterClusterConfigMongodSecurity {
	var returns *DataYandexMdbMongodbClusterClusterConfigMongodSecurity
	_jsii_.Get(
		j,
		"securityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) SetParameter() DataYandexMdbMongodbClusterClusterConfigMongodSetParameterOutputReference {
	var returns DataYandexMdbMongodbClusterClusterConfigMongodSetParameterOutputReference
	_jsii_.Get(
		j,
		"setParameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) SetParameterInput() *DataYandexMdbMongodbClusterClusterConfigMongodSetParameter {
	var returns *DataYandexMdbMongodbClusterClusterConfigMongodSetParameter
	_jsii_.Get(
		j,
		"setParameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataYandexMdbMongodbClusterClusterConfigMongodOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataYandexMdbMongodbClusterClusterConfigMongodOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexMdbMongodbClusterClusterConfigMongodOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataYandexMdbMongodbClusterClusterConfigMongodOutputReference_Override(d DataYandexMdbMongodbClusterClusterConfigMongodOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexMdbMongodbClusterClusterConfigMongodOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) SetInternalValue(val *DataYandexMdbMongodbClusterClusterConfigMongod) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) PutAuditLog(value *DataYandexMdbMongodbClusterClusterConfigMongodAuditLog) {
	_jsii_.InvokeVoid(
		d,
		"putAuditLog",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) PutSecurity(value *DataYandexMdbMongodbClusterClusterConfigMongodSecurity) {
	_jsii_.InvokeVoid(
		d,
		"putSecurity",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) PutSetParameter(value *DataYandexMdbMongodbClusterClusterConfigMongodSetParameter) {
	_jsii_.InvokeVoid(
		d,
		"putSetParameter",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) ResetAuditLog() {
	_jsii_.InvokeVoid(
		d,
		"resetAuditLog",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) ResetSecurity() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) ResetSetParameter() {
	_jsii_.InvokeVoid(
		d,
		"resetSetParameter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

