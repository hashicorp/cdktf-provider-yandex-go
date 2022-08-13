// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ApiGatewayTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/api_gateway#create ApiGateway#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/api_gateway#delete ApiGateway#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/api_gateway#update ApiGateway#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

