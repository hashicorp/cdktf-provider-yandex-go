// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DatatransferEndpointSettingsMongoSource struct {
	// collections block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#collections DatatransferEndpoint#collections}
	Collections interface{} `field:"optional" json:"collections" yaml:"collections"`
	// connection block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#connection DatatransferEndpoint#connection}
	Connection *DatatransferEndpointSettingsMongoSourceConnection `field:"optional" json:"connection" yaml:"connection"`
	// excluded_collections block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#excluded_collections DatatransferEndpoint#excluded_collections}
	ExcludedCollections interface{} `field:"optional" json:"excludedCollections" yaml:"excludedCollections"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#secondary_preferred_mode DatatransferEndpoint#secondary_preferred_mode}.
	SecondaryPreferredMode interface{} `field:"optional" json:"secondaryPreferredMode" yaml:"secondaryPreferredMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#security_groups DatatransferEndpoint#security_groups}.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#subnet_id DatatransferEndpoint#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

