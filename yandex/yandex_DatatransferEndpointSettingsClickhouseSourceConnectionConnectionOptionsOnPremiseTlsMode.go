// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsMode struct {
	// disabled block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#disabled DatatransferEndpoint#disabled}
	Disabled *DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeDisabled `field:"optional" json:"disabled" yaml:"disabled"`
	// enabled block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#enabled DatatransferEndpoint#enabled}
	Enabled *DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsModeEnabled `field:"optional" json:"enabled" yaml:"enabled"`
}

