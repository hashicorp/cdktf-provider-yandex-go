// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterList
type jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewMdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterList {
	_init_.Initialize()

	j := jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterList{}

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewMdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterList_Override(m MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		m,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterList) Get(index *float64) MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference {
	var returns MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterOutputReference

	_jsii_.Invoke(
		m,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MdbKafkaConnectorConnectorConfigMirrormakerSourceClusterExternalClusterList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

