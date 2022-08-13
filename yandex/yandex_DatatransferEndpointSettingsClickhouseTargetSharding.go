// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DatatransferEndpointSettingsClickhouseTargetSharding struct {
	// column_value_hash block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#column_value_hash DatatransferEndpoint#column_value_hash}
	ColumnValueHash *DatatransferEndpointSettingsClickhouseTargetShardingColumnValueHash `field:"optional" json:"columnValueHash" yaml:"columnValueHash"`
	// transfer_id block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#transfer_id DatatransferEndpoint#transfer_id}
	TransferId *DatatransferEndpointSettingsClickhouseTargetShardingTransferId `field:"optional" json:"transferId" yaml:"transferId"`
}

