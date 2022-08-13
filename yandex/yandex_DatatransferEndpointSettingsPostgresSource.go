// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DatatransferEndpointSettingsPostgresSource struct {
	// connection block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#connection DatatransferEndpoint#connection}
	Connection *DatatransferEndpointSettingsPostgresSourceConnection `field:"optional" json:"connection" yaml:"connection"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#database DatatransferEndpoint#database}.
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#exclude_tables DatatransferEndpoint#exclude_tables}.
	ExcludeTables *[]*string `field:"optional" json:"excludeTables" yaml:"excludeTables"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#include_tables DatatransferEndpoint#include_tables}.
	IncludeTables *[]*string `field:"optional" json:"includeTables" yaml:"includeTables"`
	// object_transfer_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#object_transfer_settings DatatransferEndpoint#object_transfer_settings}
	ObjectTransferSettings *DatatransferEndpointSettingsPostgresSourceObjectTransferSettings `field:"optional" json:"objectTransferSettings" yaml:"objectTransferSettings"`
	// password block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#password DatatransferEndpoint#password}
	Password *DatatransferEndpointSettingsPostgresSourcePassword `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#security_groups DatatransferEndpoint#security_groups}.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#service_schema DatatransferEndpoint#service_schema}.
	ServiceSchema *string `field:"optional" json:"serviceSchema" yaml:"serviceSchema"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#slot_gigabyte_lag_limit DatatransferEndpoint#slot_gigabyte_lag_limit}.
	SlotGigabyteLagLimit *float64 `field:"optional" json:"slotGigabyteLagLimit" yaml:"slotGigabyteLagLimit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#user DatatransferEndpoint#user}.
	User *string `field:"optional" json:"user" yaml:"user"`
}

