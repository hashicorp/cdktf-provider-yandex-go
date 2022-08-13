// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexMdbMongodbClusterClusterConfig struct {
	// access block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#access DataYandexMdbMongodbCluster#access}
	Access *DataYandexMdbMongodbClusterClusterConfigAccess `field:"optional" json:"access" yaml:"access"`
	// backup_window_start block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#backup_window_start DataYandexMdbMongodbCluster#backup_window_start}
	BackupWindowStart *DataYandexMdbMongodbClusterClusterConfigBackupWindowStart `field:"optional" json:"backupWindowStart" yaml:"backupWindowStart"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#feature_compatibility_version DataYandexMdbMongodbCluster#feature_compatibility_version}.
	FeatureCompatibilityVersion *string `field:"optional" json:"featureCompatibilityVersion" yaml:"featureCompatibilityVersion"`
	// mongod block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#mongod DataYandexMdbMongodbCluster#mongod}
	Mongod *DataYandexMdbMongodbClusterClusterConfigMongod `field:"optional" json:"mongod" yaml:"mongod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#version DataYandexMdbMongodbCluster#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

