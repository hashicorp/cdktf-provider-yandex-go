// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ContainerRepositoryTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/container_repository#create ContainerRepository#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/container_repository#delete ContainerRepository#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

