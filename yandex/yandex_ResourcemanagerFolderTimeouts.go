// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ResourcemanagerFolderTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/resourcemanager_folder#create ResourcemanagerFolder#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/resourcemanager_folder#delete ResourcemanagerFolder#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/resourcemanager_folder#read ResourcemanagerFolder#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/resourcemanager_folder#update ResourcemanagerFolder#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

