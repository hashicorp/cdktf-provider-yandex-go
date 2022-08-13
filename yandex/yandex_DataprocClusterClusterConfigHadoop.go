// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataprocClusterClusterConfigHadoop struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#properties DataprocCluster#properties}.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#services DataprocCluster#services}.
	Services *[]*string `field:"optional" json:"services" yaml:"services"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#ssh_public_keys DataprocCluster#ssh_public_keys}.
	SshPublicKeys *[]*string `field:"optional" json:"sshPublicKeys" yaml:"sshPublicKeys"`
}

