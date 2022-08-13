// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexMdbMongodbClusterClusterConfigMongodSecurity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#enable_encryption DataYandexMdbMongodbCluster#enable_encryption}.
	EnableEncryption interface{} `field:"optional" json:"enableEncryption" yaml:"enableEncryption"`
	// kmip block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#kmip DataYandexMdbMongodbCluster#kmip}
	Kmip *DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmip `field:"optional" json:"kmip" yaml:"kmip"`
}

