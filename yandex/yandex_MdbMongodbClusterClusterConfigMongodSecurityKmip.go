// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbMongodbClusterClusterConfigMongodSecurityKmip struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#client_certificate MdbMongodbCluster#client_certificate}.
	ClientCertificate *string `field:"optional" json:"clientCertificate" yaml:"clientCertificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#key_identifier MdbMongodbCluster#key_identifier}.
	KeyIdentifier *string `field:"optional" json:"keyIdentifier" yaml:"keyIdentifier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#port MdbMongodbCluster#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#server_ca MdbMongodbCluster#server_ca}.
	ServerCa *string `field:"optional" json:"serverCa" yaml:"serverCa"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#server_name MdbMongodbCluster#server_name}.
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
}

