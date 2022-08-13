// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference interface {
	cdktf.ComplexObject
	ClientCertificate() *string
	SetClientCertificate(val *string)
	ClientCertificateInput() *string
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
	InternalValue() *MdbMongodbClusterClusterConfigMongodSecurityKmip
	SetInternalValue(val *MdbMongodbClusterClusterConfigMongodSecurityKmip)
	KeyIdentifier() *string
	SetKeyIdentifier(val *string)
	KeyIdentifierInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	ServerCa() *string
	SetServerCa(val *string)
	ServerCaInput() *string
	ServerName() *string
	SetServerName(val *string)
	ServerNameInput() *string
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
	ResetClientCertificate()
	ResetKeyIdentifier()
	ResetPort()
	ResetServerCa()
	ResetServerName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference
type jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) ClientCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) ClientCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) InternalValue() *MdbMongodbClusterClusterConfigMongodSecurityKmip {
	var returns *MdbMongodbClusterClusterConfigMongodSecurityKmip
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) KeyIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) KeyIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) ServerCa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) ServerCaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverCaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) ServerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) ServerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference_Override(m MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) SetClientCertificate(val *string) {
	_jsii_.Set(
		j,
		"clientCertificate",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) SetInternalValue(val *MdbMongodbClusterClusterConfigMongodSecurityKmip) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) SetKeyIdentifier(val *string) {
	_jsii_.Set(
		j,
		"keyIdentifier",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) SetServerCa(val *string) {
	_jsii_.Set(
		j,
		"serverCa",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) SetServerName(val *string) {
	_jsii_.Set(
		j,
		"serverName",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) ResetClientCertificate() {
	_jsii_.InvokeVoid(
		m,
		"resetClientCertificate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) ResetKeyIdentifier() {
	_jsii_.InvokeVoid(
		m,
		"resetKeyIdentifier",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		m,
		"resetPort",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) ResetServerCa() {
	_jsii_.InvokeVoid(
		m,
		"resetServerCa",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) ResetServerName() {
	_jsii_.InvokeVoid(
		m,
		"resetServerName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

