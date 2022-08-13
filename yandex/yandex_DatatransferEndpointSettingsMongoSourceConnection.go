// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DatatransferEndpointSettingsMongoSourceConnection struct {
	// connection_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#connection_options DatatransferEndpoint#connection_options}
	ConnectionOptions *DatatransferEndpointSettingsMongoSourceConnectionConnectionOptions `field:"optional" json:"connectionOptions" yaml:"connectionOptions"`
}

