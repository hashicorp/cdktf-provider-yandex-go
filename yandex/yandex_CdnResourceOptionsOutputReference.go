// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/jsii"

	"github.com/hashicorp/cdktf-provider-yandex-go/yandex/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CdnResourceOptionsOutputReference interface {
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
	InternalValue() *CdnResourceOptions
	SetInternalValue(val *CdnResourceOptions)
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

// The jsii proxy struct for CdnResourceOptionsOutputReference
type jsiiProxy_CdnResourceOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) AllowedHttpMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedHttpMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) AllowedHttpMethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedHttpMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) BrowserCacheSettings() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"browserCacheSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) BrowserCacheSettingsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"browserCacheSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) CacheHttpHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheHttpHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) CacheHttpHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheHttpHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) Cors() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) CorsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"corsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) CustomHostHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customHostHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) CustomHostHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customHostHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) CustomServerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customServerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) CustomServerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customServerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) DisableCache() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) DisableCacheInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableCacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) DisableProxyForceRanges() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableProxyForceRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) DisableProxyForceRangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableProxyForceRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) EdgeCacheSettings() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"edgeCacheSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) EdgeCacheSettingsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"edgeCacheSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) FetchedCompressed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchedCompressed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) FetchedCompressedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchedCompressedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) ForwardHostHeader() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardHostHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) ForwardHostHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardHostHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) GzipOn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gzipOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) GzipOnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gzipOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) IgnoreCookie() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreCookie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) IgnoreCookieInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreCookieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) IgnoreQueryParams() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreQueryParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) IgnoreQueryParamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreQueryParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) InternalValue() *CdnResourceOptions {
	var returns *CdnResourceOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) ProxyCacheMethodsSet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proxyCacheMethodsSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) ProxyCacheMethodsSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proxyCacheMethodsSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) QueryParamsBlacklist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryParamsBlacklist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) QueryParamsBlacklistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryParamsBlacklistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) QueryParamsWhitelist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryParamsWhitelist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) QueryParamsWhitelistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryParamsWhitelistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) RedirectHttpsToHttp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redirectHttpsToHttp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) RedirectHttpsToHttpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redirectHttpsToHttpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) RedirectHttpToHttps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redirectHttpToHttps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) RedirectHttpToHttpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redirectHttpToHttpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) Slice() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) SliceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sliceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) StaticRequestHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"staticRequestHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) StaticRequestHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"staticRequestHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) StaticResponseHeaders() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"staticResponseHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) StaticResponseHeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"staticResponseHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnResourceOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CdnResourceOptionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnResourceOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-yandex.CdnResourceOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnResourceOptionsOutputReference_Override(c CdnResourceOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-yandex.CdnResourceOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) SetAllowedHttpMethods(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedHttpMethods",
		val,
	)
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) SetBrowserCacheSettings(val *float64) {
	_jsii_.Set(
		j,
		"browserCacheSettings",
		val,
	)
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) SetCacheHttpHeaders(val *[]*string) {
	_jsii_.Set(
		j,
		"cacheHttpHeaders",
		val,
	)
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) SetCors(val *[]*string) {
	_jsii_.Set(
		j,
		"cors",
		val,
	)
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) SetCustomHostHeader(val *string) {
	_jsii_.Set(
		j,
		"customHostHeader",
		val,
	)
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) SetCustomServerName(val *string) {
	_jsii_.Set(
		j,
		"customServerName",
		val,
	)
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) SetDisableCache(val interface{}) {
	_jsii_.Set(
		j,
		"disableCache",
		val,
	)
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) SetDisableProxyForceRanges(val interface{}) {
	_jsii_.Set(
		j,
		"disableProxyForceRanges",
		val,
	)
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) SetEdgeCacheSettings(val *float64) {
	_jsii_.Set(
		j,
		"edgeCacheSettings",
		val,
	)
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) SetFetchedCompressed(val interface{}) {
	_jsii_.Set(
		j,
		"fetchedCompressed",
		val,
	)
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) SetForwardHostHeader(val interface{}) {
	_jsii_.Set(
		j,
		"forwardHostHeader",
		val,
	)
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) SetGzipOn(val interface{}) {
	_jsii_.Set(
		j,
		"gzipOn",
		val,
	)
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) SetIgnoreCookie(val interface{}) {
	_jsii_.Set(
		j,
		"ignoreCookie",
		val,
	)
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) SetIgnoreQueryParams(val interface{}) {
	_jsii_.Set(
		j,
		"ignoreQueryParams",
		val,
	)
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) SetInternalValue(val *CdnResourceOptions) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) SetProxyCacheMethodsSet(val interface{}) {
	_jsii_.Set(
		j,
		"proxyCacheMethodsSet",
		val,
	)
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) SetQueryParamsBlacklist(val *[]*string) {
	_jsii_.Set(
		j,
		"queryParamsBlacklist",
		val,
	)
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) SetQueryParamsWhitelist(val *[]*string) {
	_jsii_.Set(
		j,
		"queryParamsWhitelist",
		val,
	)
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) SetRedirectHttpsToHttp(val interface{}) {
	_jsii_.Set(
		j,
		"redirectHttpsToHttp",
		val,
	)
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) SetRedirectHttpToHttps(val interface{}) {
	_jsii_.Set(
		j,
		"redirectHttpToHttps",
		val,
	)
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) SetSlice(val interface{}) {
	_jsii_.Set(
		j,
		"slice",
		val,
	)
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) SetStaticRequestHeaders(val *[]*string) {
	_jsii_.Set(
		j,
		"staticRequestHeaders",
		val,
	)
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) SetStaticResponseHeaders(val *map[string]*string) {
	_jsii_.Set(
		j,
		"staticResponseHeaders",
		val,
	)
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnResourceOptionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) ResetAllowedHttpMethods() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedHttpMethods",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) ResetBrowserCacheSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetBrowserCacheSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) ResetCacheHttpHeaders() {
	_jsii_.InvokeVoid(
		c,
		"resetCacheHttpHeaders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) ResetCors() {
	_jsii_.InvokeVoid(
		c,
		"resetCors",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) ResetCustomHostHeader() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomHostHeader",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) ResetCustomServerName() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomServerName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) ResetDisableCache() {
	_jsii_.InvokeVoid(
		c,
		"resetDisableCache",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) ResetDisableProxyForceRanges() {
	_jsii_.InvokeVoid(
		c,
		"resetDisableProxyForceRanges",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) ResetEdgeCacheSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetEdgeCacheSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) ResetFetchedCompressed() {
	_jsii_.InvokeVoid(
		c,
		"resetFetchedCompressed",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) ResetForwardHostHeader() {
	_jsii_.InvokeVoid(
		c,
		"resetForwardHostHeader",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) ResetGzipOn() {
	_jsii_.InvokeVoid(
		c,
		"resetGzipOn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) ResetIgnoreCookie() {
	_jsii_.InvokeVoid(
		c,
		"resetIgnoreCookie",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) ResetIgnoreQueryParams() {
	_jsii_.InvokeVoid(
		c,
		"resetIgnoreQueryParams",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) ResetProxyCacheMethodsSet() {
	_jsii_.InvokeVoid(
		c,
		"resetProxyCacheMethodsSet",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) ResetQueryParamsBlacklist() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryParamsBlacklist",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) ResetQueryParamsWhitelist() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryParamsWhitelist",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) ResetRedirectHttpsToHttp() {
	_jsii_.InvokeVoid(
		c,
		"resetRedirectHttpsToHttp",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) ResetRedirectHttpToHttps() {
	_jsii_.InvokeVoid(
		c,
		"resetRedirectHttpToHttps",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) ResetSlice() {
	_jsii_.InvokeVoid(
		c,
		"resetSlice",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) ResetStaticRequestHeaders() {
	_jsii_.InvokeVoid(
		c,
		"resetStaticRequestHeaders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) ResetStaticResponseHeaders() {
	_jsii_.InvokeVoid(
		c,
		"resetStaticResponseHeaders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnResourceOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

