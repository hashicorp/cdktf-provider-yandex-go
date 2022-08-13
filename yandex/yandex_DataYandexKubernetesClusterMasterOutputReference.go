// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexKubernetesClusterMasterOutputReference interface {
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
	InternalValue() *DataYandexKubernetesClusterMaster
	SetInternalValue(val *DataYandexKubernetesClusterMaster)
	MaintenancePolicy() DataYandexKubernetesClusterMasterMaintenancePolicyList
	PublicIp() cdktf.IResolvable
	Regional() DataYandexKubernetesClusterMasterRegionalList
	SecurityGroupIds() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Version() *string
	VersionInfo() DataYandexKubernetesClusterMasterVersionInfoList
	Zonal() DataYandexKubernetesClusterMasterZonalList
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataYandexKubernetesClusterMasterOutputReference
type jsiiProxy_DataYandexKubernetesClusterMasterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) ClusterCaCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterCaCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) ExternalV4Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalV4Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) ExternalV4Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalV4Endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) InternalV4Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internalV4Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) InternalV4Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internalV4Endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) InternalValue() *DataYandexKubernetesClusterMaster {
	var returns *DataYandexKubernetesClusterMaster
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) MaintenancePolicy() DataYandexKubernetesClusterMasterMaintenancePolicyList {
	var returns DataYandexKubernetesClusterMasterMaintenancePolicyList
	_jsii_.Get(
		j,
		"maintenancePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) PublicIp() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"publicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) Regional() DataYandexKubernetesClusterMasterRegionalList {
	var returns DataYandexKubernetesClusterMasterRegionalList
	_jsii_.Get(
		j,
		"regional",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) VersionInfo() DataYandexKubernetesClusterMasterVersionInfoList {
	var returns DataYandexKubernetesClusterMasterVersionInfoList
	_jsii_.Get(
		j,
		"versionInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) Zonal() DataYandexKubernetesClusterMasterZonalList {
	var returns DataYandexKubernetesClusterMasterZonalList
	_jsii_.Get(
		j,
		"zonal",
		&returns,
	)
	return returns
}


func NewDataYandexKubernetesClusterMasterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataYandexKubernetesClusterMasterOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataYandexKubernetesClusterMasterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexKubernetesClusterMasterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataYandexKubernetesClusterMasterOutputReference_Override(d DataYandexKubernetesClusterMasterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexKubernetesClusterMasterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) SetInternalValue(val *DataYandexKubernetesClusterMaster) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexKubernetesClusterMasterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

