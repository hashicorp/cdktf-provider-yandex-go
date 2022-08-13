// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ComputeInstanceGroupApplicationLoadBalancer struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#max_opening_traffic_duration ComputeInstanceGroup#max_opening_traffic_duration}.
	MaxOpeningTrafficDuration *float64 `field:"optional" json:"maxOpeningTrafficDuration" yaml:"maxOpeningTrafficDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#target_group_description ComputeInstanceGroup#target_group_description}.
	TargetGroupDescription *string `field:"optional" json:"targetGroupDescription" yaml:"targetGroupDescription"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#target_group_labels ComputeInstanceGroup#target_group_labels}.
	TargetGroupLabels *map[string]*string `field:"optional" json:"targetGroupLabels" yaml:"targetGroupLabels"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#target_group_name ComputeInstanceGroup#target_group_name}.
	TargetGroupName *string `field:"optional" json:"targetGroupName" yaml:"targetGroupName"`
}

