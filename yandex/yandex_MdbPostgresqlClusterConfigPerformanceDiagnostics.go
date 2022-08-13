// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbPostgresqlClusterConfigPerformanceDiagnostics struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#sessions_sampling_interval MdbPostgresqlCluster#sessions_sampling_interval}.
	SessionsSamplingInterval *float64 `field:"required" json:"sessionsSamplingInterval" yaml:"sessionsSamplingInterval"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#statements_sampling_interval MdbPostgresqlCluster#statements_sampling_interval}.
	StatementsSamplingInterval *float64 `field:"required" json:"statementsSamplingInterval" yaml:"statementsSamplingInterval"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#enabled MdbPostgresqlCluster#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

