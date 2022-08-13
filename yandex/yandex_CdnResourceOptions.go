// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type CdnResourceOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#allowed_http_methods CdnResource#allowed_http_methods}.
	AllowedHttpMethods *[]*string `field:"optional" json:"allowedHttpMethods" yaml:"allowedHttpMethods"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#browser_cache_settings CdnResource#browser_cache_settings}.
	BrowserCacheSettings *float64 `field:"optional" json:"browserCacheSettings" yaml:"browserCacheSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#cache_http_headers CdnResource#cache_http_headers}.
	CacheHttpHeaders *[]*string `field:"optional" json:"cacheHttpHeaders" yaml:"cacheHttpHeaders"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#cors CdnResource#cors}.
	Cors *[]*string `field:"optional" json:"cors" yaml:"cors"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#custom_host_header CdnResource#custom_host_header}.
	CustomHostHeader *string `field:"optional" json:"customHostHeader" yaml:"customHostHeader"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#custom_server_name CdnResource#custom_server_name}.
	CustomServerName *string `field:"optional" json:"customServerName" yaml:"customServerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#disable_cache CdnResource#disable_cache}.
	DisableCache interface{} `field:"optional" json:"disableCache" yaml:"disableCache"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#disable_proxy_force_ranges CdnResource#disable_proxy_force_ranges}.
	DisableProxyForceRanges interface{} `field:"optional" json:"disableProxyForceRanges" yaml:"disableProxyForceRanges"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#edge_cache_settings CdnResource#edge_cache_settings}.
	EdgeCacheSettings *float64 `field:"optional" json:"edgeCacheSettings" yaml:"edgeCacheSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#fetched_compressed CdnResource#fetched_compressed}.
	FetchedCompressed interface{} `field:"optional" json:"fetchedCompressed" yaml:"fetchedCompressed"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#forward_host_header CdnResource#forward_host_header}.
	ForwardHostHeader interface{} `field:"optional" json:"forwardHostHeader" yaml:"forwardHostHeader"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#gzip_on CdnResource#gzip_on}.
	GzipOn interface{} `field:"optional" json:"gzipOn" yaml:"gzipOn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#ignore_cookie CdnResource#ignore_cookie}.
	IgnoreCookie interface{} `field:"optional" json:"ignoreCookie" yaml:"ignoreCookie"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#ignore_query_params CdnResource#ignore_query_params}.
	IgnoreQueryParams interface{} `field:"optional" json:"ignoreQueryParams" yaml:"ignoreQueryParams"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#proxy_cache_methods_set CdnResource#proxy_cache_methods_set}.
	ProxyCacheMethodsSet interface{} `field:"optional" json:"proxyCacheMethodsSet" yaml:"proxyCacheMethodsSet"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#query_params_blacklist CdnResource#query_params_blacklist}.
	QueryParamsBlacklist *[]*string `field:"optional" json:"queryParamsBlacklist" yaml:"queryParamsBlacklist"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#query_params_whitelist CdnResource#query_params_whitelist}.
	QueryParamsWhitelist *[]*string `field:"optional" json:"queryParamsWhitelist" yaml:"queryParamsWhitelist"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#redirect_https_to_http CdnResource#redirect_https_to_http}.
	RedirectHttpsToHttp interface{} `field:"optional" json:"redirectHttpsToHttp" yaml:"redirectHttpsToHttp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#redirect_http_to_https CdnResource#redirect_http_to_https}.
	RedirectHttpToHttps interface{} `field:"optional" json:"redirectHttpToHttps" yaml:"redirectHttpToHttps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#slice CdnResource#slice}.
	Slice interface{} `field:"optional" json:"slice" yaml:"slice"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#static_request_headers CdnResource#static_request_headers}.
	StaticRequestHeaders *[]*string `field:"optional" json:"staticRequestHeaders" yaml:"staticRequestHeaders"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/cdn_resource#static_response_headers CdnResource#static_response_headers}.
	StaticResponseHeaders *map[string]*string `field:"optional" json:"staticResponseHeaders" yaml:"staticResponseHeaders"`
}

