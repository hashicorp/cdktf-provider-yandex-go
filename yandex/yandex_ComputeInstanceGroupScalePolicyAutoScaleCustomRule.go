// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ComputeInstanceGroupScalePolicyAutoScaleCustomRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#metric_name ComputeInstanceGroup#metric_name}.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#metric_type ComputeInstanceGroup#metric_type}.
	MetricType *string `field:"required" json:"metricType" yaml:"metricType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#rule_type ComputeInstanceGroup#rule_type}.
	RuleType *string `field:"required" json:"ruleType" yaml:"ruleType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#target ComputeInstanceGroup#target}.
	Target *float64 `field:"required" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#folder_id ComputeInstanceGroup#folder_id}.
	FolderId *string `field:"optional" json:"folderId" yaml:"folderId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#labels ComputeInstanceGroup#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#service ComputeInstanceGroup#service}.
	Service *string `field:"optional" json:"service" yaml:"service"`
}

