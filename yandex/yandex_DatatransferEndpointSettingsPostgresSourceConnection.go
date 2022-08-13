// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DatatransferEndpointSettingsPostgresSourceConnection struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#mdb_cluster_id DatatransferEndpoint#mdb_cluster_id}.
	MdbClusterId *string `field:"optional" json:"mdbClusterId" yaml:"mdbClusterId"`
	// on_premise block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#on_premise DatatransferEndpoint#on_premise}
	OnPremise *DatatransferEndpointSettingsPostgresSourceConnectionOnPremise `field:"optional" json:"onPremise" yaml:"onPremise"`
}

