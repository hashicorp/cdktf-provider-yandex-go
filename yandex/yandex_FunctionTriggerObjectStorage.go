// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type FunctionTriggerObjectStorage struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#bucket_id FunctionTrigger#bucket_id}.
	BucketId *string `field:"required" json:"bucketId" yaml:"bucketId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#create FunctionTrigger#create}.
	Create interface{} `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#delete FunctionTrigger#delete}.
	Delete interface{} `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#prefix FunctionTrigger#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#suffix FunctionTrigger#suffix}.
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#update FunctionTrigger#update}.
	Update interface{} `field:"optional" json:"update" yaml:"update"`
}

