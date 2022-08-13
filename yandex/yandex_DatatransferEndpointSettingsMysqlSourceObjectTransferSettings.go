// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DatatransferEndpointSettingsMysqlSourceObjectTransferSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#routine DatatransferEndpoint#routine}.
	Routine *string `field:"optional" json:"routine" yaml:"routine"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#trigger DatatransferEndpoint#trigger}.
	Trigger *string `field:"optional" json:"trigger" yaml:"trigger"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#view DatatransferEndpoint#view}.
	View *string `field:"optional" json:"view" yaml:"view"`
}

