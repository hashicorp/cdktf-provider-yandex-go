// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DatatransferEndpointSettingsMongoSourceExcludedCollections struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#collection_name DatatransferEndpoint#collection_name}.
	CollectionName *string `field:"optional" json:"collectionName" yaml:"collectionName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#database_name DatatransferEndpoint#database_name}.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
}

