// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremise struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#hosts DatatransferEndpoint#hosts}.
	Hosts *[]*string `field:"optional" json:"hosts" yaml:"hosts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#port DatatransferEndpoint#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#replica_set DatatransferEndpoint#replica_set}.
	ReplicaSet *string `field:"optional" json:"replicaSet" yaml:"replicaSet"`
	// tls_mode block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#tls_mode DatatransferEndpoint#tls_mode}
	TlsMode *DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremiseTlsMode `field:"optional" json:"tlsMode" yaml:"tlsMode"`
}

