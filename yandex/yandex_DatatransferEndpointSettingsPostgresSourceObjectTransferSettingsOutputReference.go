// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference interface {
	cdktf.ComplexObject
	Cast() *string
	SetCast(val *string)
	CastInput() *string
	Collation() *string
	SetCollation(val *string)
	CollationInput() *string
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
	Constraint() *string
	SetConstraint(val *string)
	ConstraintInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultValues() *string
	SetDefaultValues(val *string)
	DefaultValuesInput() *string
	FkConstraint() *string
	SetFkConstraint(val *string)
	FkConstraintInput() *string
	// Experimental.
	Fqn() *string
	Function() *string
	SetFunction(val *string)
	FunctionInput() *string
	Index() *string
	SetIndex(val *string)
	IndexInput() *string
	InternalValue() *DatatransferEndpointSettingsPostgresSourceObjectTransferSettings
	SetInternalValue(val *DatatransferEndpointSettingsPostgresSourceObjectTransferSettings)
	MaterializedView() *string
	SetMaterializedView(val *string)
	MaterializedViewInput() *string
	Policy() *string
	SetPolicy(val *string)
	PolicyInput() *string
	PrimaryKey() *string
	SetPrimaryKey(val *string)
	PrimaryKeyInput() *string
	Rule() *string
	SetRule(val *string)
	RuleInput() *string
	Sequence() *string
	SetSequence(val *string)
	SequenceInput() *string
	SequenceOwnedBy() *string
	SetSequenceOwnedBy(val *string)
	SequenceOwnedByInput() *string
	Table() *string
	SetTable(val *string)
	TableInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Trigger() *string
	SetTrigger(val *string)
	TriggerInput() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	View() *string
	SetView(val *string)
	ViewInput() *string
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
	ResetCast()
	ResetCollation()
	ResetConstraint()
	ResetDefaultValues()
	ResetFkConstraint()
	ResetFunction()
	ResetIndex()
	ResetMaterializedView()
	ResetPolicy()
	ResetPrimaryKey()
	ResetRule()
	ResetSequence()
	ResetSequenceOwnedBy()
	ResetTable()
	ResetTrigger()
	ResetType()
	ResetView()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference
type jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) Cast() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cast",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) CastInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"castInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) Collation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) CollationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) Constraint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) ConstraintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constraintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) DefaultValues() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) DefaultValuesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) FkConstraint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fkConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) FkConstraintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fkConstraintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) Function() *string {
	var returns *string
	_jsii_.Get(
		j,
		"function",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) FunctionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) Index() *string {
	var returns *string
	_jsii_.Get(
		j,
		"index",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) IndexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) InternalValue() *DatatransferEndpointSettingsPostgresSourceObjectTransferSettings {
	var returns *DatatransferEndpointSettingsPostgresSourceObjectTransferSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) MaterializedView() *string {
	var returns *string
	_jsii_.Get(
		j,
		"materializedView",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) MaterializedViewInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"materializedViewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) PolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) PrimaryKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) PrimaryKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) Rule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) RuleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) Sequence() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sequence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) SequenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sequenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) SequenceOwnedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sequenceOwnedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) SequenceOwnedByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sequenceOwnedByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) Table() *string {
	var returns *string
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) TableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) Trigger() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) TriggerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) View() *string {
	var returns *string
	_jsii_.Get(
		j,
		"view",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) ViewInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewInput",
		&returns,
	)
	return returns
}


func NewDatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference_Override(d DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) SetCast(val *string) {
	_jsii_.Set(
		j,
		"cast",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) SetCollation(val *string) {
	_jsii_.Set(
		j,
		"collation",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) SetConstraint(val *string) {
	_jsii_.Set(
		j,
		"constraint",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) SetDefaultValues(val *string) {
	_jsii_.Set(
		j,
		"defaultValues",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) SetFkConstraint(val *string) {
	_jsii_.Set(
		j,
		"fkConstraint",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) SetFunction(val *string) {
	_jsii_.Set(
		j,
		"function",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) SetIndex(val *string) {
	_jsii_.Set(
		j,
		"index",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) SetInternalValue(val *DatatransferEndpointSettingsPostgresSourceObjectTransferSettings) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) SetMaterializedView(val *string) {
	_jsii_.Set(
		j,
		"materializedView",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) SetPolicy(val *string) {
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) SetPrimaryKey(val *string) {
	_jsii_.Set(
		j,
		"primaryKey",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) SetRule(val *string) {
	_jsii_.Set(
		j,
		"rule",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) SetSequence(val *string) {
	_jsii_.Set(
		j,
		"sequence",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) SetSequenceOwnedBy(val *string) {
	_jsii_.Set(
		j,
		"sequenceOwnedBy",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) SetTable(val *string) {
	_jsii_.Set(
		j,
		"table",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) SetTrigger(val *string) {
	_jsii_.Set(
		j,
		"trigger",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) SetView(val *string) {
	_jsii_.Set(
		j,
		"view",
		val,
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) ResetCast() {
	_jsii_.InvokeVoid(
		d,
		"resetCast",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) ResetCollation() {
	_jsii_.InvokeVoid(
		d,
		"resetCollation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) ResetConstraint() {
	_jsii_.InvokeVoid(
		d,
		"resetConstraint",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) ResetDefaultValues() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultValues",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) ResetFkConstraint() {
	_jsii_.InvokeVoid(
		d,
		"resetFkConstraint",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) ResetFunction() {
	_jsii_.InvokeVoid(
		d,
		"resetFunction",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) ResetIndex() {
	_jsii_.InvokeVoid(
		d,
		"resetIndex",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) ResetMaterializedView() {
	_jsii_.InvokeVoid(
		d,
		"resetMaterializedView",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) ResetPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) ResetPrimaryKey() {
	_jsii_.InvokeVoid(
		d,
		"resetPrimaryKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) ResetRule() {
	_jsii_.InvokeVoid(
		d,
		"resetRule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) ResetSequence() {
	_jsii_.InvokeVoid(
		d,
		"resetSequence",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) ResetSequenceOwnedBy() {
	_jsii_.InvokeVoid(
		d,
		"resetSequenceOwnedBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) ResetTable() {
	_jsii_.InvokeVoid(
		d,
		"resetTable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) ResetTrigger() {
	_jsii_.InvokeVoid(
		d,
		"resetTrigger",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		d,
		"resetType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) ResetView() {
	_jsii_.InvokeVoid(
		d,
		"resetView",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

