// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type StorageBucketWebsite struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#error_document StorageBucket#error_document}.
	ErrorDocument *string `field:"optional" json:"errorDocument" yaml:"errorDocument"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#index_document StorageBucket#index_document}.
	IndexDocument *string `field:"optional" json:"indexDocument" yaml:"indexDocument"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#redirect_all_requests_to StorageBucket#redirect_all_requests_to}.
	RedirectAllRequestsTo *string `field:"optional" json:"redirectAllRequestsTo" yaml:"redirectAllRequestsTo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/storage_bucket#routing_rules StorageBucket#routing_rules}.
	RoutingRules *string `field:"optional" json:"routingRules" yaml:"routingRules"`
}

