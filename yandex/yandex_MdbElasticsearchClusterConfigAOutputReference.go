// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbElasticsearchClusterConfigAOutputReference interface {
	cdktf.ComplexObject
	AdminPassword() *string
	SetAdminPassword(val *string)
	AdminPasswordInput() *string
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
	DataNode() MdbElasticsearchClusterConfigDataNodeOutputReference
	DataNodeInput() *MdbElasticsearchClusterConfigDataNode
	Edition() *string
	SetEdition(val *string)
	EditionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MdbElasticsearchClusterConfigA
	SetInternalValue(val *MdbElasticsearchClusterConfigA)
	MasterNode() MdbElasticsearchClusterConfigMasterNodeOutputReference
	MasterNodeInput() *MdbElasticsearchClusterConfigMasterNode
	Plugins() *[]*string
	SetPlugins(val *[]*string)
	PluginsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutDataNode(value *MdbElasticsearchClusterConfigDataNode)
	PutMasterNode(value *MdbElasticsearchClusterConfigMasterNode)
	ResetEdition()
	ResetMasterNode()
	ResetPlugins()
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MdbElasticsearchClusterConfigAOutputReference
type jsiiProxy_MdbElasticsearchClusterConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) AdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) AdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) DataNode() MdbElasticsearchClusterConfigDataNodeOutputReference {
	var returns MdbElasticsearchClusterConfigDataNodeOutputReference
	_jsii_.Get(
		j,
		"dataNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) DataNodeInput() *MdbElasticsearchClusterConfigDataNode {
	var returns *MdbElasticsearchClusterConfigDataNode
	_jsii_.Get(
		j,
		"dataNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) Edition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) EditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"editionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) InternalValue() *MdbElasticsearchClusterConfigA {
	var returns *MdbElasticsearchClusterConfigA
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) MasterNode() MdbElasticsearchClusterConfigMasterNodeOutputReference {
	var returns MdbElasticsearchClusterConfigMasterNodeOutputReference
	_jsii_.Get(
		j,
		"masterNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) MasterNodeInput() *MdbElasticsearchClusterConfigMasterNode {
	var returns *MdbElasticsearchClusterConfigMasterNode
	_jsii_.Get(
		j,
		"masterNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) Plugins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"plugins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) PluginsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pluginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewMdbElasticsearchClusterConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MdbElasticsearchClusterConfigAOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MdbElasticsearchClusterConfigAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbElasticsearchClusterConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMdbElasticsearchClusterConfigAOutputReference_Override(m MdbElasticsearchClusterConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbElasticsearchClusterConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) SetAdminPassword(val *string) {
	_jsii_.Set(
		j,
		"adminPassword",
		val,
	)
}

func (j *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) SetEdition(val *string) {
	_jsii_.Set(
		j,
		"edition",
		val,
	)
}

func (j *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) SetInternalValue(val *MdbElasticsearchClusterConfigA) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) SetPlugins(val *[]*string) {
	_jsii_.Set(
		j,
		"plugins",
		val,
	)
}

func (j *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (m *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) PutDataNode(value *MdbElasticsearchClusterConfigDataNode) {
	_jsii_.InvokeVoid(
		m,
		"putDataNode",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) PutMasterNode(value *MdbElasticsearchClusterConfigMasterNode) {
	_jsii_.InvokeVoid(
		m,
		"putMasterNode",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) ResetEdition() {
	_jsii_.InvokeVoid(
		m,
		"resetEdition",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) ResetMasterNode() {
	_jsii_.InvokeVoid(
		m,
		"resetMasterNode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) ResetPlugins() {
	_jsii_.InvokeVoid(
		m,
		"resetPlugins",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbElasticsearchClusterConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

