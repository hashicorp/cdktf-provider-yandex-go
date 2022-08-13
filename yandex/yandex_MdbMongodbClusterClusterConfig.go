// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbMongodbClusterClusterConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#version MdbMongodbCluster#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// access block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#access MdbMongodbCluster#access}
	Access *MdbMongodbClusterClusterConfigAccess `field:"optional" json:"access" yaml:"access"`
	// backup_window_start block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#backup_window_start MdbMongodbCluster#backup_window_start}
	BackupWindowStart *MdbMongodbClusterClusterConfigBackupWindowStart `field:"optional" json:"backupWindowStart" yaml:"backupWindowStart"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#feature_compatibility_version MdbMongodbCluster#feature_compatibility_version}.
	FeatureCompatibilityVersion *string `field:"optional" json:"featureCompatibilityVersion" yaml:"featureCompatibilityVersion"`
	// mongod block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#mongod MdbMongodbCluster#mongod}
	Mongod *MdbMongodbClusterClusterConfigMongod `field:"optional" json:"mongod" yaml:"mongod"`
}

