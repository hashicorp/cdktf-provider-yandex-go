// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexMdbMongodbClusterClusterConfigMongodSecurityKmip struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#client_certificate DataYandexMdbMongodbCluster#client_certificate}.
	ClientCertificate *string `field:"optional" json:"clientCertificate" yaml:"clientCertificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#key_identifier DataYandexMdbMongodbCluster#key_identifier}.
	KeyIdentifier *string `field:"optional" json:"keyIdentifier" yaml:"keyIdentifier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#port DataYandexMdbMongodbCluster#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#server_ca DataYandexMdbMongodbCluster#server_ca}.
	ServerCa *string `field:"optional" json:"serverCa" yaml:"serverCa"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#server_name DataYandexMdbMongodbCluster#server_name}.
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
}

