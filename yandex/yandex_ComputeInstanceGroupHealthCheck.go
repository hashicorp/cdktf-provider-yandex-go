// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ComputeInstanceGroupHealthCheck struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#healthy_threshold ComputeInstanceGroup#healthy_threshold}.
	HealthyThreshold *float64 `field:"optional" json:"healthyThreshold" yaml:"healthyThreshold"`
	// http_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#http_options ComputeInstanceGroup#http_options}
	HttpOptions interface{} `field:"optional" json:"httpOptions" yaml:"httpOptions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#interval ComputeInstanceGroup#interval}.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// tcp_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#tcp_options ComputeInstanceGroup#tcp_options}
	TcpOptions *ComputeInstanceGroupHealthCheckTcpOptions `field:"optional" json:"tcpOptions" yaml:"tcpOptions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#timeout ComputeInstanceGroup#timeout}.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#unhealthy_threshold ComputeInstanceGroup#unhealthy_threshold}.
	UnhealthyThreshold *float64 `field:"optional" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

