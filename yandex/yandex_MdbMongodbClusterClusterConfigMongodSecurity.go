// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbMongodbClusterClusterConfigMongodSecurity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#enable_encryption MdbMongodbCluster#enable_encryption}.
	EnableEncryption interface{} `field:"optional" json:"enableEncryption" yaml:"enableEncryption"`
	// kmip block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#kmip MdbMongodbCluster#kmip}
	Kmip *MdbMongodbClusterClusterConfigMongodSecurityKmip `field:"optional" json:"kmip" yaml:"kmip"`
}

