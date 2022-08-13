// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesClusterMasterOutputReference interface {
	cdktf.ComplexObject
	ClusterCaCertificate() *string
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
	ExternalV4Address() *string
	ExternalV4Endpoint() *string
	// Experimental.
	Fqn() *string
	InternalV4Address() *string
	InternalV4Endpoint() *string
	InternalValue() *KubernetesClusterMaster
	SetInternalValue(val *KubernetesClusterMaster)
	MaintenancePolicy() KubernetesClusterMasterMaintenancePolicyOutputReference
	MaintenancePolicyInput() *KubernetesClusterMasterMaintenancePolicy
	PublicIp() interface{}
	SetPublicIp(val interface{})
	PublicIpInput() interface{}
	Regional() KubernetesClusterMasterRegionalOutputReference
	RegionalInput() *KubernetesClusterMasterRegional
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
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
	VersionInfo() KubernetesClusterMasterVersionInfoList
	VersionInput() *string
	Zonal() KubernetesClusterMasterZonalOutputReference
	ZonalInput() *KubernetesClusterMasterZonal
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
	PutMaintenancePolicy(value *KubernetesClusterMasterMaintenancePolicy)
	PutRegional(value *KubernetesClusterMasterRegional)
	PutZonal(value *KubernetesClusterMasterZonal)
	ResetMaintenancePolicy()
	ResetPublicIp()
	ResetRegional()
	ResetSecurityGroupIds()
	ResetVersion()
	ResetZonal()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterMasterOutputReference
type jsiiProxy_KubernetesClusterMasterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) ClusterCaCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterCaCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) ExternalV4Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalV4Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) ExternalV4Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalV4Endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) InternalV4Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internalV4Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) InternalV4Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internalV4Endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) InternalValue() *KubernetesClusterMaster {
	var returns *KubernetesClusterMaster
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) MaintenancePolicy() KubernetesClusterMasterMaintenancePolicyOutputReference {
	var returns KubernetesClusterMasterMaintenancePolicyOutputReference
	_jsii_.Get(
		j,
		"maintenancePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) MaintenancePolicyInput() *KubernetesClusterMasterMaintenancePolicy {
	var returns *KubernetesClusterMasterMaintenancePolicy
	_jsii_.Get(
		j,
		"maintenancePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) PublicIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) PublicIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) Regional() KubernetesClusterMasterRegionalOutputReference {
	var returns KubernetesClusterMasterRegionalOutputReference
	_jsii_.Get(
		j,
		"regional",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) RegionalInput() *KubernetesClusterMasterRegional {
	var returns *KubernetesClusterMasterRegional
	_jsii_.Get(
		j,
		"regionalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) VersionInfo() KubernetesClusterMasterVersionInfoList {
	var returns KubernetesClusterMasterVersionInfoList
	_jsii_.Get(
		j,
		"versionInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) Zonal() KubernetesClusterMasterZonalOutputReference {
	var returns KubernetesClusterMasterZonalOutputReference
	_jsii_.Get(
		j,
		"zonal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) ZonalInput() *KubernetesClusterMasterZonal {
	var returns *KubernetesClusterMasterZonal
	_jsii_.Get(
		j,
		"zonalInput",
		&returns,
	)
	return returns
}


func NewKubernetesClusterMasterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterMasterOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterMasterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.KubernetesClusterMasterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterMasterOutputReference_Override(k KubernetesClusterMasterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.KubernetesClusterMasterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) SetInternalValue(val *KubernetesClusterMaster) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) SetPublicIp(val interface{}) {
	_jsii_.Set(
		j,
		"publicIp",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMasterOutputReference) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterMasterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMasterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMasterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMasterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMasterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMasterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMasterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMasterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMasterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMasterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMasterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMasterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMasterOutputReference) PutMaintenancePolicy(value *KubernetesClusterMasterMaintenancePolicy) {
	_jsii_.InvokeVoid(
		k,
		"putMaintenancePolicy",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterMasterOutputReference) PutRegional(value *KubernetesClusterMasterRegional) {
	_jsii_.InvokeVoid(
		k,
		"putRegional",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterMasterOutputReference) PutZonal(value *KubernetesClusterMasterZonal) {
	_jsii_.InvokeVoid(
		k,
		"putZonal",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterMasterOutputReference) ResetMaintenancePolicy() {
	_jsii_.InvokeVoid(
		k,
		"resetMaintenancePolicy",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterMasterOutputReference) ResetPublicIp() {
	_jsii_.InvokeVoid(
		k,
		"resetPublicIp",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterMasterOutputReference) ResetRegional() {
	_jsii_.InvokeVoid(
		k,
		"resetRegional",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterMasterOutputReference) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		k,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterMasterOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		k,
		"resetVersion",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterMasterOutputReference) ResetZonal() {
	_jsii_.InvokeVoid(
		k,
		"resetZonal",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterMasterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMasterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

