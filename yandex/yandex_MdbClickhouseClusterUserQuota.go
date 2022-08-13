// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbClickhouseClusterUserQuota struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#interval_duration MdbClickhouseCluster#interval_duration}.
	IntervalDuration *float64 `field:"required" json:"intervalDuration" yaml:"intervalDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#errors MdbClickhouseCluster#errors}.
	Errors *float64 `field:"optional" json:"errors" yaml:"errors"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#execution_time MdbClickhouseCluster#execution_time}.
	ExecutionTime *float64 `field:"optional" json:"executionTime" yaml:"executionTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#queries MdbClickhouseCluster#queries}.
	Queries *float64 `field:"optional" json:"queries" yaml:"queries"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#read_rows MdbClickhouseCluster#read_rows}.
	ReadRows *float64 `field:"optional" json:"readRows" yaml:"readRows"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_clickhouse_cluster#result_rows MdbClickhouseCluster#result_rows}.
	ResultRows *float64 `field:"optional" json:"resultRows" yaml:"resultRows"`
}

