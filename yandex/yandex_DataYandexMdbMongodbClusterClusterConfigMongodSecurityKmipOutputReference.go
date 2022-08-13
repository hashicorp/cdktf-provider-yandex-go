// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference interface {
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
	InternalValue() *DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmip
	SetInternalValue(val *DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmip)
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

// The jsii proxy struct for DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference
type jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) ClientCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) ClientCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) InternalValue() *DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmip {
	var returns *DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmip
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) KeyIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) KeyIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) ServerCa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) ServerCaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverCaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) ServerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) ServerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference_Override(d DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) SetClientCertificate(val *string) {
	_jsii_.Set(
		j,
		"clientCertificate",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) SetInternalValue(val *DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmip) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) SetKeyIdentifier(val *string) {
	_jsii_.Set(
		j,
		"keyIdentifier",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) SetServerCa(val *string) {
	_jsii_.Set(
		j,
		"serverCa",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) SetServerName(val *string) {
	_jsii_.Set(
		j,
		"serverName",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) ResetClientCertificate() {
	_jsii_.InvokeVoid(
		d,
		"resetClientCertificate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) ResetKeyIdentifier() {
	_jsii_.InvokeVoid(
		d,
		"resetKeyIdentifier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		d,
		"resetPort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) ResetServerCa() {
	_jsii_.InvokeVoid(
		d,
		"resetServerCa",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) ResetServerName() {
	_jsii_.InvokeVoid(
		d,
		"resetServerName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmipOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

