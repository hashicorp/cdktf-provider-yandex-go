// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexFunctionScalingPolicyPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/function_scaling_policy#zone_instances_limit DataYandexFunctionScalingPolicy#zone_instances_limit}.
	ZoneInstancesLimit *float64 `field:"optional" json:"zoneInstancesLimit" yaml:"zoneInstancesLimit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/function_scaling_policy#zone_requests_limit DataYandexFunctionScalingPolicy#zone_requests_limit}.
	ZoneRequestsLimit *float64 `field:"optional" json:"zoneRequestsLimit" yaml:"zoneRequestsLimit"`
}

