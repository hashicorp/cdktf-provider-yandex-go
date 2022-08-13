// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DatatransferEndpointSettingsMongoTarget struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#cleanup_policy DatatransferEndpoint#cleanup_policy}.
	CleanupPolicy *string `field:"optional" json:"cleanupPolicy" yaml:"cleanupPolicy"`
	// connection block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#connection DatatransferEndpoint#connection}
	Connection *DatatransferEndpointSettingsMongoTargetConnection `field:"optional" json:"connection" yaml:"connection"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#database DatatransferEndpoint#database}.
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#security_groups DatatransferEndpoint#security_groups}.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#subnet_id DatatransferEndpoint#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

