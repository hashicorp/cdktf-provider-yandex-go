// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type FunctionScalingPolicyPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_scaling_policy#tag FunctionScalingPolicy#tag}.
	Tag *string `field:"required" json:"tag" yaml:"tag"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_scaling_policy#zone_instances_limit FunctionScalingPolicy#zone_instances_limit}.
	ZoneInstancesLimit *float64 `field:"optional" json:"zoneInstancesLimit" yaml:"zoneInstancesLimit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_scaling_policy#zone_requests_limit FunctionScalingPolicy#zone_requests_limit}.
	ZoneRequestsLimit *float64 `field:"optional" json:"zoneRequestsLimit" yaml:"zoneRequestsLimit"`
}

