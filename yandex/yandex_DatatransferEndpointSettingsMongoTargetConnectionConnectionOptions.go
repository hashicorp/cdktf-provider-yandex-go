// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DatatransferEndpointSettingsMongoTargetConnectionConnectionOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#auth_source DatatransferEndpoint#auth_source}.
	AuthSource *string `field:"optional" json:"authSource" yaml:"authSource"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#mdb_cluster_id DatatransferEndpoint#mdb_cluster_id}.
	MdbClusterId *string `field:"optional" json:"mdbClusterId" yaml:"mdbClusterId"`
	// on_premise block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#on_premise DatatransferEndpoint#on_premise}
	OnPremise *DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsOnPremise `field:"optional" json:"onPremise" yaml:"onPremise"`
	// password block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#password DatatransferEndpoint#password}
	Password *DatatransferEndpointSettingsMongoTargetConnectionConnectionOptionsPassword `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#user DatatransferEndpoint#user}.
	User *string `field:"optional" json:"user" yaml:"user"`
}

