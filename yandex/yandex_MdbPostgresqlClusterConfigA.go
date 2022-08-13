// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbPostgresqlClusterConfigA struct {
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#resources MdbPostgresqlCluster#resources}
	Resources *MdbPostgresqlClusterConfigResources `field:"required" json:"resources" yaml:"resources"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#version MdbPostgresqlCluster#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// access block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#access MdbPostgresqlCluster#access}
	Access *MdbPostgresqlClusterConfigAccess `field:"optional" json:"access" yaml:"access"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#autofailover MdbPostgresqlCluster#autofailover}.
	Autofailover interface{} `field:"optional" json:"autofailover" yaml:"autofailover"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#backup_retain_period_days MdbPostgresqlCluster#backup_retain_period_days}.
	BackupRetainPeriodDays *float64 `field:"optional" json:"backupRetainPeriodDays" yaml:"backupRetainPeriodDays"`
	// backup_window_start block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#backup_window_start MdbPostgresqlCluster#backup_window_start}
	BackupWindowStart *MdbPostgresqlClusterConfigBackupWindowStart `field:"optional" json:"backupWindowStart" yaml:"backupWindowStart"`
	// performance_diagnostics block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#performance_diagnostics MdbPostgresqlCluster#performance_diagnostics}
	PerformanceDiagnostics *MdbPostgresqlClusterConfigPerformanceDiagnostics `field:"optional" json:"performanceDiagnostics" yaml:"performanceDiagnostics"`
	// pooler_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#pooler_config MdbPostgresqlCluster#pooler_config}
	PoolerConfig *MdbPostgresqlClusterConfigPoolerConfig `field:"optional" json:"poolerConfig" yaml:"poolerConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#postgresql_config MdbPostgresqlCluster#postgresql_config}.
	PostgresqlConfig *map[string]*string `field:"optional" json:"postgresqlConfig" yaml:"postgresqlConfig"`
}

