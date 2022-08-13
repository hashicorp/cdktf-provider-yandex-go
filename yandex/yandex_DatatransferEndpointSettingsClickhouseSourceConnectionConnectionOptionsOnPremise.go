// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremise struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#http_port DatatransferEndpoint#http_port}.
	HttpPort *float64 `field:"optional" json:"httpPort" yaml:"httpPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#native_port DatatransferEndpoint#native_port}.
	NativePort *float64 `field:"optional" json:"nativePort" yaml:"nativePort"`
	// shards block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#shards DatatransferEndpoint#shards}
	Shards interface{} `field:"optional" json:"shards" yaml:"shards"`
	// tls_mode block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#tls_mode DatatransferEndpoint#tls_mode}
	TlsMode *DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsMode `field:"optional" json:"tlsMode" yaml:"tlsMode"`
}

