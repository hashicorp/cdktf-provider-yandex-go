// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DatatransferEndpointSettingsPostgresSourceConnectionOnPremise struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#hosts DatatransferEndpoint#hosts}.
	Hosts *[]*string `field:"optional" json:"hosts" yaml:"hosts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#port DatatransferEndpoint#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#subnet_id DatatransferEndpoint#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// tls_mode block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#tls_mode DatatransferEndpoint#tls_mode}
	TlsMode *DatatransferEndpointSettingsPostgresSourceConnectionOnPremiseTlsMode `field:"optional" json:"tlsMode" yaml:"tlsMode"`
}

