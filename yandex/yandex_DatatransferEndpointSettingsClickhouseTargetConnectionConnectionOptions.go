// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#database DatatransferEndpoint#database}.
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#mdb_cluster_id DatatransferEndpoint#mdb_cluster_id}.
	MdbClusterId *string `field:"optional" json:"mdbClusterId" yaml:"mdbClusterId"`
	// on_premise block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#on_premise DatatransferEndpoint#on_premise}
	OnPremise *DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremise `field:"optional" json:"onPremise" yaml:"onPremise"`
	// password block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#password DatatransferEndpoint#password}
	Password *DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsPassword `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#user DatatransferEndpoint#user}.
	User *string `field:"optional" json:"user" yaml:"user"`
}

