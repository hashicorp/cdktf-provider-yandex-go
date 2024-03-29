// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ContainerRegistryTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/container_registry#create ContainerRegistry#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/container_registry#delete ContainerRegistry#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/container_registry#update ContainerRegistry#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

