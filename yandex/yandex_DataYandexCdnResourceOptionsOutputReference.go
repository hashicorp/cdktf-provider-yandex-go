// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexCdnResourceOptionsOutputReference interface {
	cdktf.ComplexObject
	AllowedHttpMethods() *[]*string
	SetAllowedHttpMethods(val *[]*string)
	AllowedHttpMethodsInput() *[]*string
	BrowserCacheSettings() *float64
	SetBrowserCacheSettings(val *float64)
	BrowserCacheSettingsInput() *float64
	CacheHttpHeaders() *[]*string
	SetCacheHttpHeaders(val *[]*string)
	CacheHttpHeadersInput() *[]*string
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
	Cors() *[]*string
	SetCors(val *[]*string)
	CorsInput() *[]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomHostHeader() *string
	SetCustomHostHeader(val *string)
	CustomHostHeaderInput() *string
	CustomServerName() *string
	SetCustomServerName(val *string)
	CustomServerNameInput() *string
	DisableCache() interface{}
	SetDisableCache(val interface{})
	DisableCacheInput() interface{}
	DisableProxyForceRanges() interface{}
	SetDisableProxyForceRanges(val interface{})
	DisableProxyForceRangesInput() interface{}
	EdgeCacheSettings() *float64
	SetEdgeCacheSettings(val *float64)
	EdgeCacheSettingsInput() *float64
	FetchedCompressed() interface{}
	SetFetchedCompressed(val interface{})
	FetchedCompressedInput() interface{}
	ForwardHostHeader() interface{}
	SetForwardHostHeader(val interface{})
	ForwardHostHeaderInput() interface{}
	// Experimental.
	Fqn() *string
	GzipOn() interface{}
	SetGzipOn(val interface{})
	GzipOnInput() interface{}
	IgnoreCookie() interface{}
	SetIgnoreCookie(val interface{})
	IgnoreCookieInput() interface{}
	IgnoreQueryParams() interface{}
	SetIgnoreQueryParams(val interface{})
	IgnoreQueryParamsInput() interface{}
	InternalValue() *DataYandexCdnResourceOptions
	SetInternalValue(val *DataYandexCdnResourceOptions)
	ProxyCacheMethodsSet() interface{}
	SetProxyCacheMethodsSet(val interface{})
	ProxyCacheMethodsSetInput() interface{}
	QueryParamsBlacklist() *[]*string
	SetQueryParamsBlacklist(val *[]*string)
	QueryParamsBlacklistInput() *[]*string
	QueryParamsWhitelist() *[]*string
	SetQueryParamsWhitelist(val *[]*string)
	QueryParamsWhitelistInput() *[]*string
	RedirectHttpsToHttp() interface{}
	SetRedirectHttpsToHttp(val interface{})
	RedirectHttpsToHttpInput() interface{}
	RedirectHttpToHttps() interface{}
	SetRedirectHttpToHttps(val interface{})
	RedirectHttpToHttpsInput() interface{}
	Slice() interface{}
	SetSlice(val interface{})
	SliceInput() interface{}
	StaticRequestHeaders() *[]*string
	SetStaticRequestHeaders(val *[]*string)
	StaticRequestHeadersInput() *[]*string
	StaticResponseHeaders() *map[string]*string
	SetStaticResponseHeaders(val *map[string]*string)
	StaticResponseHeadersInput() *map[string]*string
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
	ResetAllowedHttpMethods()
	ResetBrowserCacheSettings()
	ResetCacheHttpHeaders()
	ResetCors()
	ResetCustomHostHeader()
	ResetCustomServerName()
	ResetDisableCache()
	ResetDisableProxyForceRanges()
	ResetEdgeCacheSettings()
	ResetFetchedCompressed()
	ResetForwardHostHeader()
	ResetGzipOn()
	ResetIgnoreCookie()
	ResetIgnoreQueryParams()
	ResetProxyCacheMethodsSet()
	ResetQueryParamsBlacklist()
	ResetQueryParamsWhitelist()
	ResetRedirectHttpsToHttp()
	ResetRedirectHttpToHttps()
	ResetSlice()
	ResetStaticRequestHeaders()
	ResetStaticResponseHeaders()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataYandexCdnResourceOptionsOutputReference
type jsiiProxy_DataYandexCdnResourceOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) AllowedHttpMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedHttpMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) AllowedHttpMethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedHttpMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) BrowserCacheSettings() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"browserCacheSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) BrowserCacheSettingsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"browserCacheSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) CacheHttpHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheHttpHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) CacheHttpHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheHttpHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) Cors() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) CorsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"corsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) CustomHostHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customHostHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) CustomHostHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customHostHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) CustomServerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customServerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) CustomServerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customServerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) DisableCache() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) DisableCacheInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableCacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) DisableProxyForceRanges() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableProxyForceRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) DisableProxyForceRangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableProxyForceRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) EdgeCacheSettings() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"edgeCacheSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) EdgeCacheSettingsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"edgeCacheSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) FetchedCompressed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchedCompressed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) FetchedCompressedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchedCompressedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) ForwardHostHeader() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardHostHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) ForwardHostHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardHostHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) GzipOn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gzipOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) GzipOnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gzipOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) IgnoreCookie() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreCookie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) IgnoreCookieInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreCookieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) IgnoreQueryParams() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreQueryParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) IgnoreQueryParamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreQueryParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) InternalValue() *DataYandexCdnResourceOptions {
	var returns *DataYandexCdnResourceOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) ProxyCacheMethodsSet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proxyCacheMethodsSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) ProxyCacheMethodsSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proxyCacheMethodsSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) QueryParamsBlacklist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryParamsBlacklist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) QueryParamsBlacklistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryParamsBlacklistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) QueryParamsWhitelist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryParamsWhitelist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) QueryParamsWhitelistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryParamsWhitelistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) RedirectHttpsToHttp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redirectHttpsToHttp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) RedirectHttpsToHttpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redirectHttpsToHttpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) RedirectHttpToHttps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redirectHttpToHttps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) RedirectHttpToHttpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redirectHttpToHttpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) Slice() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) SliceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sliceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) StaticRequestHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"staticRequestHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) StaticRequestHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"staticRequestHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) StaticResponseHeaders() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"staticResponseHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) StaticResponseHeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"staticResponseHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataYandexCdnResourceOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataYandexCdnResourceOptionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataYandexCdnResourceOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexCdnResourceOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataYandexCdnResourceOptionsOutputReference_Override(d DataYandexCdnResourceOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.DataYandexCdnResourceOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) SetAllowedHttpMethods(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedHttpMethods",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) SetBrowserCacheSettings(val *float64) {
	_jsii_.Set(
		j,
		"browserCacheSettings",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) SetCacheHttpHeaders(val *[]*string) {
	_jsii_.Set(
		j,
		"cacheHttpHeaders",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) SetCors(val *[]*string) {
	_jsii_.Set(
		j,
		"cors",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) SetCustomHostHeader(val *string) {
	_jsii_.Set(
		j,
		"customHostHeader",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) SetCustomServerName(val *string) {
	_jsii_.Set(
		j,
		"customServerName",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) SetDisableCache(val interface{}) {
	_jsii_.Set(
		j,
		"disableCache",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) SetDisableProxyForceRanges(val interface{}) {
	_jsii_.Set(
		j,
		"disableProxyForceRanges",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) SetEdgeCacheSettings(val *float64) {
	_jsii_.Set(
		j,
		"edgeCacheSettings",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) SetFetchedCompressed(val interface{}) {
	_jsii_.Set(
		j,
		"fetchedCompressed",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) SetForwardHostHeader(val interface{}) {
	_jsii_.Set(
		j,
		"forwardHostHeader",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) SetGzipOn(val interface{}) {
	_jsii_.Set(
		j,
		"gzipOn",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) SetIgnoreCookie(val interface{}) {
	_jsii_.Set(
		j,
		"ignoreCookie",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) SetIgnoreQueryParams(val interface{}) {
	_jsii_.Set(
		j,
		"ignoreQueryParams",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) SetInternalValue(val *DataYandexCdnResourceOptions) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) SetProxyCacheMethodsSet(val interface{}) {
	_jsii_.Set(
		j,
		"proxyCacheMethodsSet",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) SetQueryParamsBlacklist(val *[]*string) {
	_jsii_.Set(
		j,
		"queryParamsBlacklist",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) SetQueryParamsWhitelist(val *[]*string) {
	_jsii_.Set(
		j,
		"queryParamsWhitelist",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) SetRedirectHttpsToHttp(val interface{}) {
	_jsii_.Set(
		j,
		"redirectHttpsToHttp",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) SetRedirectHttpToHttps(val interface{}) {
	_jsii_.Set(
		j,
		"redirectHttpToHttps",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) SetSlice(val interface{}) {
	_jsii_.Set(
		j,
		"slice",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) SetStaticRequestHeaders(val *[]*string) {
	_jsii_.Set(
		j,
		"staticRequestHeaders",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) SetStaticResponseHeaders(val *map[string]*string) {
	_jsii_.Set(
		j,
		"staticResponseHeaders",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) ResetAllowedHttpMethods() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowedHttpMethods",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) ResetBrowserCacheSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetBrowserCacheSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) ResetCacheHttpHeaders() {
	_jsii_.InvokeVoid(
		d,
		"resetCacheHttpHeaders",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) ResetCors() {
	_jsii_.InvokeVoid(
		d,
		"resetCors",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) ResetCustomHostHeader() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomHostHeader",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) ResetCustomServerName() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomServerName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) ResetDisableCache() {
	_jsii_.InvokeVoid(
		d,
		"resetDisableCache",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) ResetDisableProxyForceRanges() {
	_jsii_.InvokeVoid(
		d,
		"resetDisableProxyForceRanges",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) ResetEdgeCacheSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetEdgeCacheSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) ResetFetchedCompressed() {
	_jsii_.InvokeVoid(
		d,
		"resetFetchedCompressed",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) ResetForwardHostHeader() {
	_jsii_.InvokeVoid(
		d,
		"resetForwardHostHeader",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) ResetGzipOn() {
	_jsii_.InvokeVoid(
		d,
		"resetGzipOn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) ResetIgnoreCookie() {
	_jsii_.InvokeVoid(
		d,
		"resetIgnoreCookie",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) ResetIgnoreQueryParams() {
	_jsii_.InvokeVoid(
		d,
		"resetIgnoreQueryParams",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) ResetProxyCacheMethodsSet() {
	_jsii_.InvokeVoid(
		d,
		"resetProxyCacheMethodsSet",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) ResetQueryParamsBlacklist() {
	_jsii_.InvokeVoid(
		d,
		"resetQueryParamsBlacklist",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) ResetQueryParamsWhitelist() {
	_jsii_.InvokeVoid(
		d,
		"resetQueryParamsWhitelist",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) ResetRedirectHttpsToHttp() {
	_jsii_.InvokeVoid(
		d,
		"resetRedirectHttpsToHttp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) ResetRedirectHttpToHttps() {
	_jsii_.InvokeVoid(
		d,
		"resetRedirectHttpToHttps",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) ResetSlice() {
	_jsii_.InvokeVoid(
		d,
		"resetSlice",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) ResetStaticRequestHeaders() {
	_jsii_.InvokeVoid(
		d,
		"resetStaticRequestHeaders",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) ResetStaticResponseHeaders() {
	_jsii_.InvokeVoid(
		d,
		"resetStaticResponseHeaders",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataYandexCdnResourceOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

