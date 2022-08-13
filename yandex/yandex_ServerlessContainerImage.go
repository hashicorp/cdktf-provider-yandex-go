// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ServerlessContainerImage struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/serverless_container#url ServerlessContainer#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/serverless_container#args ServerlessContainer#args}.
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/serverless_container#command ServerlessContainer#command}.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/serverless_container#digest ServerlessContainer#digest}.
	Digest *string `field:"optional" json:"digest" yaml:"digest"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/serverless_container#environment ServerlessContainer#environment}.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/serverless_container#work_dir ServerlessContainer#work_dir}.
	WorkDir *string `field:"optional" json:"workDir" yaml:"workDir"`
}

