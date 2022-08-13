// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DatatransferEndpointSettingsClickhouseSource struct {
	// connection block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#connection DatatransferEndpoint#connection}
	Connection *DatatransferEndpointSettingsClickhouseSourceConnection `field:"optional" json:"connection" yaml:"connection"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#exclude_tables DatatransferEndpoint#exclude_tables}.
	ExcludeTables *[]*string `field:"optional" json:"excludeTables" yaml:"excludeTables"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#include_tables DatatransferEndpoint#include_tables}.
	IncludeTables *[]*string `field:"optional" json:"includeTables" yaml:"includeTables"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#security_groups DatatransferEndpoint#security_groups}.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#subnet_id DatatransferEndpoint#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

