// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DatatransferEndpointSettingsClickhouseTarget struct {
	// alt_names block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#alt_names DatatransferEndpoint#alt_names}
	AltNames interface{} `field:"optional" json:"altNames" yaml:"altNames"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#cleanup_policy DatatransferEndpoint#cleanup_policy}.
	CleanupPolicy *string `field:"optional" json:"cleanupPolicy" yaml:"cleanupPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#clickhouse_cluster_name DatatransferEndpoint#clickhouse_cluster_name}.
	ClickhouseClusterName *string `field:"optional" json:"clickhouseClusterName" yaml:"clickhouseClusterName"`
	// connection block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#connection DatatransferEndpoint#connection}
	Connection *DatatransferEndpointSettingsClickhouseTargetConnection `field:"optional" json:"connection" yaml:"connection"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#security_groups DatatransferEndpoint#security_groups}.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// sharding block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#sharding DatatransferEndpoint#sharding}
	Sharding *DatatransferEndpointSettingsClickhouseTargetSharding `field:"optional" json:"sharding" yaml:"sharding"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#subnet_id DatatransferEndpoint#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

