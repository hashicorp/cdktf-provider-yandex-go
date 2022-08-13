// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexAlbBackendGroupStreamBackendHealthcheck struct {
	// grpc_healthcheck block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#grpc_healthcheck DataYandexAlbBackendGroup#grpc_healthcheck}
	GrpcHealthcheck *DataYandexAlbBackendGroupStreamBackendHealthcheckGrpcHealthcheck `field:"optional" json:"grpcHealthcheck" yaml:"grpcHealthcheck"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#healthcheck_port DataYandexAlbBackendGroup#healthcheck_port}.
	HealthcheckPort *float64 `field:"optional" json:"healthcheckPort" yaml:"healthcheckPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#healthy_threshold DataYandexAlbBackendGroup#healthy_threshold}.
	HealthyThreshold *float64 `field:"optional" json:"healthyThreshold" yaml:"healthyThreshold"`
	// http_healthcheck block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#http_healthcheck DataYandexAlbBackendGroup#http_healthcheck}
	HttpHealthcheck *DataYandexAlbBackendGroupStreamBackendHealthcheckHttpHealthcheck `field:"optional" json:"httpHealthcheck" yaml:"httpHealthcheck"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#interval_jitter_percent DataYandexAlbBackendGroup#interval_jitter_percent}.
	IntervalJitterPercent *float64 `field:"optional" json:"intervalJitterPercent" yaml:"intervalJitterPercent"`
	// stream_healthcheck block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#stream_healthcheck DataYandexAlbBackendGroup#stream_healthcheck}
	StreamHealthcheck *DataYandexAlbBackendGroupStreamBackendHealthcheckStreamHealthcheck `field:"optional" json:"streamHealthcheck" yaml:"streamHealthcheck"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/alb_backend_group#unhealthy_threshold DataYandexAlbBackendGroup#unhealthy_threshold}.
	UnhealthyThreshold *float64 `field:"optional" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

