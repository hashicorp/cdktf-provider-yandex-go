// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexMdbMongodbClusterUser struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#name DataYandexMdbMongodbCluster#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#password DataYandexMdbMongodbCluster#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// permission block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#permission DataYandexMdbMongodbCluster#permission}
	Permission interface{} `field:"optional" json:"permission" yaml:"permission"`
}

