// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataprocClusterClusterConfigSubclusterSpec struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#hosts_count DataprocCluster#hosts_count}.
	HostsCount *float64 `field:"required" json:"hostsCount" yaml:"hostsCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#name DataprocCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#resources DataprocCluster#resources}
	Resources *DataprocClusterClusterConfigSubclusterSpecResources `field:"required" json:"resources" yaml:"resources"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#role DataprocCluster#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#subnet_id DataprocCluster#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#assign_public_ip DataprocCluster#assign_public_ip}.
	AssignPublicIp interface{} `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// autoscaling_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/dataproc_cluster#autoscaling_config DataprocCluster#autoscaling_config}
	AutoscalingConfig *DataprocClusterClusterConfigSubclusterSpecAutoscalingConfig `field:"optional" json:"autoscalingConfig" yaml:"autoscalingConfig"`
}

