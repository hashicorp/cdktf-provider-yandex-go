// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DatatransferEndpointSettingsMysqlSource struct {
	// connection block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#connection DatatransferEndpoint#connection}
	Connection *DatatransferEndpointSettingsMysqlSourceConnection `field:"optional" json:"connection" yaml:"connection"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#database DatatransferEndpoint#database}.
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#exclude_tables_regex DatatransferEndpoint#exclude_tables_regex}.
	ExcludeTablesRegex *[]*string `field:"optional" json:"excludeTablesRegex" yaml:"excludeTablesRegex"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#include_tables_regex DatatransferEndpoint#include_tables_regex}.
	IncludeTablesRegex *[]*string `field:"optional" json:"includeTablesRegex" yaml:"includeTablesRegex"`
	// object_transfer_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#object_transfer_settings DatatransferEndpoint#object_transfer_settings}
	ObjectTransferSettings *DatatransferEndpointSettingsMysqlSourceObjectTransferSettings `field:"optional" json:"objectTransferSettings" yaml:"objectTransferSettings"`
	// password block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#password DatatransferEndpoint#password}
	Password *DatatransferEndpointSettingsMysqlSourcePassword `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#security_groups DatatransferEndpoint#security_groups}.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#service_database DatatransferEndpoint#service_database}.
	ServiceDatabase *string `field:"optional" json:"serviceDatabase" yaml:"serviceDatabase"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#timezone DatatransferEndpoint#timezone}.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#user DatatransferEndpoint#user}.
	User *string `field:"optional" json:"user" yaml:"user"`
}

