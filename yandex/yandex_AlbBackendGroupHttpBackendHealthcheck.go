// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbBackendGroupHttpBackendHealthcheck struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#interval AlbBackendGroup#interval}.
	Interval *string `field:"required" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#timeout AlbBackendGroup#timeout}.
	Timeout *string `field:"required" json:"timeout" yaml:"timeout"`
	// grpc_healthcheck block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#grpc_healthcheck AlbBackendGroup#grpc_healthcheck}
	GrpcHealthcheck *AlbBackendGroupHttpBackendHealthcheckGrpcHealthcheck `field:"optional" json:"grpcHealthcheck" yaml:"grpcHealthcheck"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#healthcheck_port AlbBackendGroup#healthcheck_port}.
	HealthcheckPort *float64 `field:"optional" json:"healthcheckPort" yaml:"healthcheckPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#healthy_threshold AlbBackendGroup#healthy_threshold}.
	HealthyThreshold *float64 `field:"optional" json:"healthyThreshold" yaml:"healthyThreshold"`
	// http_healthcheck block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#http_healthcheck AlbBackendGroup#http_healthcheck}
	HttpHealthcheck *AlbBackendGroupHttpBackendHealthcheckHttpHealthcheck `field:"optional" json:"httpHealthcheck" yaml:"httpHealthcheck"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#interval_jitter_percent AlbBackendGroup#interval_jitter_percent}.
	IntervalJitterPercent *float64 `field:"optional" json:"intervalJitterPercent" yaml:"intervalJitterPercent"`
	// stream_healthcheck block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#stream_healthcheck AlbBackendGroup#stream_healthcheck}
	StreamHealthcheck *AlbBackendGroupHttpBackendHealthcheckStreamHealthcheck `field:"optional" json:"streamHealthcheck" yaml:"streamHealthcheck"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_backend_group#unhealthy_threshold AlbBackendGroup#unhealthy_threshold}.
	UnhealthyThreshold *float64 `field:"optional" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

