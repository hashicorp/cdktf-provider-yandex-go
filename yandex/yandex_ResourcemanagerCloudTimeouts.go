// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ResourcemanagerCloudTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/resourcemanager_cloud#create ResourcemanagerCloud#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/resourcemanager_cloud#delete ResourcemanagerCloud#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/resourcemanager_cloud#read ResourcemanagerCloud#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/resourcemanager_cloud#update ResourcemanagerCloud#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

