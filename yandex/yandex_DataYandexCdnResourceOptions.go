// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexCdnResourceOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/cdn_resource#allowed_http_methods DataYandexCdnResource#allowed_http_methods}.
	AllowedHttpMethods *[]*string `field:"optional" json:"allowedHttpMethods" yaml:"allowedHttpMethods"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/cdn_resource#browser_cache_settings DataYandexCdnResource#browser_cache_settings}.
	BrowserCacheSettings *float64 `field:"optional" json:"browserCacheSettings" yaml:"browserCacheSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/cdn_resource#cache_http_headers DataYandexCdnResource#cache_http_headers}.
	CacheHttpHeaders *[]*string `field:"optional" json:"cacheHttpHeaders" yaml:"cacheHttpHeaders"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/cdn_resource#cors DataYandexCdnResource#cors}.
	Cors *[]*string `field:"optional" json:"cors" yaml:"cors"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/cdn_resource#custom_host_header DataYandexCdnResource#custom_host_header}.
	CustomHostHeader *string `field:"optional" json:"customHostHeader" yaml:"customHostHeader"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/cdn_resource#custom_server_name DataYandexCdnResource#custom_server_name}.
	CustomServerName *string `field:"optional" json:"customServerName" yaml:"customServerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/cdn_resource#disable_cache DataYandexCdnResource#disable_cache}.
	DisableCache interface{} `field:"optional" json:"disableCache" yaml:"disableCache"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/cdn_resource#disable_proxy_force_ranges DataYandexCdnResource#disable_proxy_force_ranges}.
	DisableProxyForceRanges interface{} `field:"optional" json:"disableProxyForceRanges" yaml:"disableProxyForceRanges"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/cdn_resource#edge_cache_settings DataYandexCdnResource#edge_cache_settings}.
	EdgeCacheSettings *float64 `field:"optional" json:"edgeCacheSettings" yaml:"edgeCacheSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/cdn_resource#fetched_compressed DataYandexCdnResource#fetched_compressed}.
	FetchedCompressed interface{} `field:"optional" json:"fetchedCompressed" yaml:"fetchedCompressed"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/cdn_resource#forward_host_header DataYandexCdnResource#forward_host_header}.
	ForwardHostHeader interface{} `field:"optional" json:"forwardHostHeader" yaml:"forwardHostHeader"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/cdn_resource#gzip_on DataYandexCdnResource#gzip_on}.
	GzipOn interface{} `field:"optional" json:"gzipOn" yaml:"gzipOn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/cdn_resource#ignore_cookie DataYandexCdnResource#ignore_cookie}.
	IgnoreCookie interface{} `field:"optional" json:"ignoreCookie" yaml:"ignoreCookie"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/cdn_resource#ignore_query_params DataYandexCdnResource#ignore_query_params}.
	IgnoreQueryParams interface{} `field:"optional" json:"ignoreQueryParams" yaml:"ignoreQueryParams"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/cdn_resource#proxy_cache_methods_set DataYandexCdnResource#proxy_cache_methods_set}.
	ProxyCacheMethodsSet interface{} `field:"optional" json:"proxyCacheMethodsSet" yaml:"proxyCacheMethodsSet"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/cdn_resource#query_params_blacklist DataYandexCdnResource#query_params_blacklist}.
	QueryParamsBlacklist *[]*string `field:"optional" json:"queryParamsBlacklist" yaml:"queryParamsBlacklist"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/cdn_resource#query_params_whitelist DataYandexCdnResource#query_params_whitelist}.
	QueryParamsWhitelist *[]*string `field:"optional" json:"queryParamsWhitelist" yaml:"queryParamsWhitelist"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/cdn_resource#redirect_https_to_http DataYandexCdnResource#redirect_https_to_http}.
	RedirectHttpsToHttp interface{} `field:"optional" json:"redirectHttpsToHttp" yaml:"redirectHttpsToHttp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/cdn_resource#redirect_http_to_https DataYandexCdnResource#redirect_http_to_https}.
	RedirectHttpToHttps interface{} `field:"optional" json:"redirectHttpToHttps" yaml:"redirectHttpToHttps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/cdn_resource#slice DataYandexCdnResource#slice}.
	Slice interface{} `field:"optional" json:"slice" yaml:"slice"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/cdn_resource#static_request_headers DataYandexCdnResource#static_request_headers}.
	StaticRequestHeaders *[]*string `field:"optional" json:"staticRequestHeaders" yaml:"staticRequestHeaders"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/cdn_resource#static_response_headers DataYandexCdnResource#static_response_headers}.
	StaticResponseHeaders *map[string]*string `field:"optional" json:"staticResponseHeaders" yaml:"staticResponseHeaders"`
}

