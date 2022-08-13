// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ComputeInstanceGroupScalePolicyAutoScale struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#initial_size ComputeInstanceGroup#initial_size}.
	InitialSize *float64 `field:"required" json:"initialSize" yaml:"initialSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#measurement_duration ComputeInstanceGroup#measurement_duration}.
	MeasurementDuration *float64 `field:"required" json:"measurementDuration" yaml:"measurementDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#cpu_utilization_target ComputeInstanceGroup#cpu_utilization_target}.
	CpuUtilizationTarget *float64 `field:"optional" json:"cpuUtilizationTarget" yaml:"cpuUtilizationTarget"`
	// custom_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#custom_rule ComputeInstanceGroup#custom_rule}
	CustomRule interface{} `field:"optional" json:"customRule" yaml:"customRule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#max_size ComputeInstanceGroup#max_size}.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#min_zone_size ComputeInstanceGroup#min_zone_size}.
	MinZoneSize *float64 `field:"optional" json:"minZoneSize" yaml:"minZoneSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#stabilization_duration ComputeInstanceGroup#stabilization_duration}.
	StabilizationDuration *float64 `field:"optional" json:"stabilizationDuration" yaml:"stabilizationDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#warmup_duration ComputeInstanceGroup#warmup_duration}.
	WarmupDuration *float64 `field:"optional" json:"warmupDuration" yaml:"warmupDuration"`
}

