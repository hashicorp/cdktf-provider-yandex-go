// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataprocClusterClusterConfigSubclusterSpecAutoscalingConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#max_hosts_count DataprocCluster#max_hosts_count}.
	MaxHostsCount *float64 `field:"required" json:"maxHostsCount" yaml:"maxHostsCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#cpu_utilization_target DataprocCluster#cpu_utilization_target}.
	CpuUtilizationTarget *float64 `field:"optional" json:"cpuUtilizationTarget" yaml:"cpuUtilizationTarget"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#decommission_timeout DataprocCluster#decommission_timeout}.
	DecommissionTimeout *float64 `field:"optional" json:"decommissionTimeout" yaml:"decommissionTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#measurement_duration DataprocCluster#measurement_duration}.
	MeasurementDuration *float64 `field:"optional" json:"measurementDuration" yaml:"measurementDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#preemptible DataprocCluster#preemptible}.
	Preemptible interface{} `field:"optional" json:"preemptible" yaml:"preemptible"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#stabilization_duration DataprocCluster#stabilization_duration}.
	StabilizationDuration *float64 `field:"optional" json:"stabilizationDuration" yaml:"stabilizationDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#warmup_duration DataprocCluster#warmup_duration}.
	WarmupDuration *float64 `field:"optional" json:"warmupDuration" yaml:"warmupDuration"`
}

