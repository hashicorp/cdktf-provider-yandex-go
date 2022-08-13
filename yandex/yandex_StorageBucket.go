// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket yandex_storage_bucket}.
type StorageBucket interface {
	cdktf.TerraformResource
	AccessKey() *string
	SetAccessKey(val *string)
	AccessKeyInput() *string
	Acl() *string
	SetAcl(val *string)
	AclInput() *string
	AnonymousAccessFlags() StorageBucketAnonymousAccessFlagsOutputReference
	AnonymousAccessFlagsInput() *StorageBucketAnonymousAccessFlags
	Bucket() *string
	SetBucket(val *string)
	BucketDomainName() *string
	BucketInput() *string
	BucketPrefix() *string
	SetBucketPrefix(val *string)
	BucketPrefixInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CorsRule() StorageBucketCorsRuleList
	CorsRuleInput() interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	DefaultStorageClass() *string
	SetDefaultStorageClass(val *string)
	DefaultStorageClassInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	FolderId() *string
	SetFolderId(val *string)
	FolderIdInput() *string
	ForceDestroy() interface{}
	SetForceDestroy(val interface{})
	ForceDestroyInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Grant() StorageBucketGrantList
	GrantInput() interface{}
	Https() StorageBucketHttpsOutputReference
	HttpsInput() *StorageBucketHttps
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LifecycleRule() StorageBucketLifecycleRuleList
	LifecycleRuleInput() interface{}
	Logging() StorageBucketLoggingList
	LoggingInput() interface{}
	MaxSize() *float64
	SetMaxSize(val *float64)
	MaxSizeInput() *float64
	// The tree node.
	Node() constructs.Node
	Policy() *string
	SetPolicy(val *string)
	PolicyInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	SecretKey() *string
	SetSecretKey(val *string)
	SecretKeyInput() *string
	ServerSideEncryptionConfiguration() StorageBucketServerSideEncryptionConfigurationOutputReference
	ServerSideEncryptionConfigurationInput() *StorageBucketServerSideEncryptionConfiguration
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Versioning() StorageBucketVersioningOutputReference
	VersioningInput() *StorageBucketVersioning
	Website() StorageBucketWebsiteOutputReference
	WebsiteDomain() *string
	SetWebsiteDomain(val *string)
	WebsiteDomainInput() *string
	WebsiteEndpoint() *string
	SetWebsiteEndpoint(val *string)
	WebsiteEndpointInput() *string
	WebsiteInput() *StorageBucketWebsite
	// Experimental.
	AddOverride(path *string, value interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAnonymousAccessFlags(value *StorageBucketAnonymousAccessFlags)
	PutCorsRule(value interface{})
	PutGrant(value interface{})
	PutHttps(value *StorageBucketHttps)
	PutLifecycleRule(value interface{})
	PutLogging(value interface{})
	PutServerSideEncryptionConfiguration(value *StorageBucketServerSideEncryptionConfiguration)
	PutVersioning(value *StorageBucketVersioning)
	PutWebsite(value *StorageBucketWebsite)
	ResetAccessKey()
	ResetAcl()
	ResetAnonymousAccessFlags()
	ResetBucket()
	ResetBucketPrefix()
	ResetCorsRule()
	ResetDefaultStorageClass()
	ResetFolderId()
	ResetForceDestroy()
	ResetGrant()
	ResetHttps()
	ResetId()
	ResetLifecycleRule()
	ResetLogging()
	ResetMaxSize()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPolicy()
	ResetSecretKey()
	ResetServerSideEncryptionConfiguration()
	ResetVersioning()
	ResetWebsite()
	ResetWebsiteDomain()
	ResetWebsiteEndpoint()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for StorageBucket
type jsiiProxy_StorageBucket struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StorageBucket) AccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) AccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) Acl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) AclInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) AnonymousAccessFlags() StorageBucketAnonymousAccessFlagsOutputReference {
	var returns StorageBucketAnonymousAccessFlagsOutputReference
	_jsii_.Get(
		j,
		"anonymousAccessFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) AnonymousAccessFlagsInput() *StorageBucketAnonymousAccessFlags {
	var returns *StorageBucketAnonymousAccessFlags
	_jsii_.Get(
		j,
		"anonymousAccessFlagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) BucketDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) BucketPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) BucketPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) CorsRule() StorageBucketCorsRuleList {
	var returns StorageBucketCorsRuleList
	_jsii_.Get(
		j,
		"corsRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) CorsRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"corsRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) DefaultStorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultStorageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) DefaultStorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultStorageClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) FolderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) FolderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) ForceDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) ForceDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) Grant() StorageBucketGrantList {
	var returns StorageBucketGrantList
	_jsii_.Get(
		j,
		"grant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) GrantInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"grantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) Https() StorageBucketHttpsOutputReference {
	var returns StorageBucketHttpsOutputReference
	_jsii_.Get(
		j,
		"https",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) HttpsInput() *StorageBucketHttps {
	var returns *StorageBucketHttps
	_jsii_.Get(
		j,
		"httpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) LifecycleRule() StorageBucketLifecycleRuleList {
	var returns StorageBucketLifecycleRuleList
	_jsii_.Get(
		j,
		"lifecycleRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) LifecycleRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lifecycleRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) Logging() StorageBucketLoggingList {
	var returns StorageBucketLoggingList
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) LoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) MaxSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) MaxSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) PolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) SecretKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) SecretKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) ServerSideEncryptionConfiguration() StorageBucketServerSideEncryptionConfigurationOutputReference {
	var returns StorageBucketServerSideEncryptionConfigurationOutputReference
	_jsii_.Get(
		j,
		"serverSideEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) ServerSideEncryptionConfigurationInput() *StorageBucketServerSideEncryptionConfiguration {
	var returns *StorageBucketServerSideEncryptionConfiguration
	_jsii_.Get(
		j,
		"serverSideEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) Versioning() StorageBucketVersioningOutputReference {
	var returns StorageBucketVersioningOutputReference
	_jsii_.Get(
		j,
		"versioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) VersioningInput() *StorageBucketVersioning {
	var returns *StorageBucketVersioning
	_jsii_.Get(
		j,
		"versioningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) Website() StorageBucketWebsiteOutputReference {
	var returns StorageBucketWebsiteOutputReference
	_jsii_.Get(
		j,
		"website",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) WebsiteDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websiteDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) WebsiteDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websiteDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) WebsiteEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websiteEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) WebsiteEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websiteEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucket) WebsiteInput() *StorageBucketWebsite {
	var returns *StorageBucketWebsite
	_jsii_.Get(
		j,
		"websiteInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket yandex_storage_bucket} Resource.
func NewStorageBucket(scope constructs.Construct, id *string, config *StorageBucketConfig) StorageBucket {
	_init_.Initialize()

	j := jsiiProxy_StorageBucket{}

	_jsii_.Create(
		"@cdktf/provider-yandex.StorageBucket",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket yandex_storage_bucket} Resource.
func NewStorageBucket_Override(s StorageBucket, scope constructs.Construct, id *string, config *StorageBucketConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.StorageBucket",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StorageBucket) SetAccessKey(val *string) {
	_jsii_.Set(
		j,
		"accessKey",
		val,
	)
}

func (j *jsiiProxy_StorageBucket) SetAcl(val *string) {
	_jsii_.Set(
		j,
		"acl",
		val,
	)
}

func (j *jsiiProxy_StorageBucket) SetBucket(val *string) {
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_StorageBucket) SetBucketPrefix(val *string) {
	_jsii_.Set(
		j,
		"bucketPrefix",
		val,
	)
}

func (j *jsiiProxy_StorageBucket) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_StorageBucket) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StorageBucket) SetDefaultStorageClass(val *string) {
	_jsii_.Set(
		j,
		"defaultStorageClass",
		val,
	)
}

func (j *jsiiProxy_StorageBucket) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StorageBucket) SetFolderId(val *string) {
	_jsii_.Set(
		j,
		"folderId",
		val,
	)
}

func (j *jsiiProxy_StorageBucket) SetForceDestroy(val interface{}) {
	_jsii_.Set(
		j,
		"forceDestroy",
		val,
	)
}

func (j *jsiiProxy_StorageBucket) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_StorageBucket) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_StorageBucket) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StorageBucket) SetMaxSize(val *float64) {
	_jsii_.Set(
		j,
		"maxSize",
		val,
	)
}

func (j *jsiiProxy_StorageBucket) SetPolicy(val *string) {
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_StorageBucket) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StorageBucket) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_StorageBucket) SetSecretKey(val *string) {
	_jsii_.Set(
		j,
		"secretKey",
		val,
	)
}

func (j *jsiiProxy_StorageBucket) SetWebsiteDomain(val *string) {
	_jsii_.Set(
		j,
		"websiteDomain",
		val,
	)
}

func (j *jsiiProxy_StorageBucket) SetWebsiteEndpoint(val *string) {
	_jsii_.Set(
		j,
		"websiteEndpoint",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func StorageBucket_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-yandex.StorageBucket",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StorageBucket_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-yandex.StorageBucket",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_StorageBucket) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_StorageBucket) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucket) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucket) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucket) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucket) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucket) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucket) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucket) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucket) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucket) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucket) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StorageBucket) PutAnonymousAccessFlags(value *StorageBucketAnonymousAccessFlags) {
	_jsii_.InvokeVoid(
		s,
		"putAnonymousAccessFlags",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageBucket) PutCorsRule(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putCorsRule",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageBucket) PutGrant(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putGrant",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageBucket) PutHttps(value *StorageBucketHttps) {
	_jsii_.InvokeVoid(
		s,
		"putHttps",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageBucket) PutLifecycleRule(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putLifecycleRule",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageBucket) PutLogging(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putLogging",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageBucket) PutServerSideEncryptionConfiguration(value *StorageBucketServerSideEncryptionConfiguration) {
	_jsii_.InvokeVoid(
		s,
		"putServerSideEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageBucket) PutVersioning(value *StorageBucketVersioning) {
	_jsii_.InvokeVoid(
		s,
		"putVersioning",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageBucket) PutWebsite(value *StorageBucketWebsite) {
	_jsii_.InvokeVoid(
		s,
		"putWebsite",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageBucket) ResetAccessKey() {
	_jsii_.InvokeVoid(
		s,
		"resetAccessKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucket) ResetAcl() {
	_jsii_.InvokeVoid(
		s,
		"resetAcl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucket) ResetAnonymousAccessFlags() {
	_jsii_.InvokeVoid(
		s,
		"resetAnonymousAccessFlags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucket) ResetBucket() {
	_jsii_.InvokeVoid(
		s,
		"resetBucket",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucket) ResetBucketPrefix() {
	_jsii_.InvokeVoid(
		s,
		"resetBucketPrefix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucket) ResetCorsRule() {
	_jsii_.InvokeVoid(
		s,
		"resetCorsRule",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucket) ResetDefaultStorageClass() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultStorageClass",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucket) ResetFolderId() {
	_jsii_.InvokeVoid(
		s,
		"resetFolderId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucket) ResetForceDestroy() {
	_jsii_.InvokeVoid(
		s,
		"resetForceDestroy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucket) ResetGrant() {
	_jsii_.InvokeVoid(
		s,
		"resetGrant",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucket) ResetHttps() {
	_jsii_.InvokeVoid(
		s,
		"resetHttps",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucket) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucket) ResetLifecycleRule() {
	_jsii_.InvokeVoid(
		s,
		"resetLifecycleRule",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucket) ResetLogging() {
	_jsii_.InvokeVoid(
		s,
		"resetLogging",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucket) ResetMaxSize() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxSize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucket) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucket) ResetPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucket) ResetSecretKey() {
	_jsii_.InvokeVoid(
		s,
		"resetSecretKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucket) ResetServerSideEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		s,
		"resetServerSideEncryptionConfiguration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucket) ResetVersioning() {
	_jsii_.InvokeVoid(
		s,
		"resetVersioning",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucket) ResetWebsite() {
	_jsii_.InvokeVoid(
		s,
		"resetWebsite",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucket) ResetWebsiteDomain() {
	_jsii_.InvokeVoid(
		s,
		"resetWebsiteDomain",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucket) ResetWebsiteEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetWebsiteEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucket) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucket) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucket) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucket) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

