// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type FunctionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function#create Function#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function#delete Function#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function#update Function#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

