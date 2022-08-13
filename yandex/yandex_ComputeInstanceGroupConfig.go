// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeInstanceGroupConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// allocation_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#allocation_policy ComputeInstanceGroup#allocation_policy}
	AllocationPolicy *ComputeInstanceGroupAllocationPolicy `field:"required" json:"allocationPolicy" yaml:"allocationPolicy"`
	// deploy_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#deploy_policy ComputeInstanceGroup#deploy_policy}
	DeployPolicy *ComputeInstanceGroupDeployPolicy `field:"required" json:"deployPolicy" yaml:"deployPolicy"`
	// instance_template block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#instance_template ComputeInstanceGroup#instance_template}
	InstanceTemplate *ComputeInstanceGroupInstanceTemplate `field:"required" json:"instanceTemplate" yaml:"instanceTemplate"`
	// scale_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#scale_policy ComputeInstanceGroup#scale_policy}
	ScalePolicy *ComputeInstanceGroupScalePolicy `field:"required" json:"scalePolicy" yaml:"scalePolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#service_account_id ComputeInstanceGroup#service_account_id}.
	ServiceAccountId *string `field:"required" json:"serviceAccountId" yaml:"serviceAccountId"`
	// application_load_balancer block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#application_load_balancer ComputeInstanceGroup#application_load_balancer}
	ApplicationLoadBalancer *ComputeInstanceGroupApplicationLoadBalancer `field:"optional" json:"applicationLoadBalancer" yaml:"applicationLoadBalancer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#deletion_protection ComputeInstanceGroup#deletion_protection}.
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#description ComputeInstanceGroup#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#folder_id ComputeInstanceGroup#folder_id}.
	FolderId *string `field:"optional" json:"folderId" yaml:"folderId"`
	// health_check block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#health_check ComputeInstanceGroup#health_check}
	HealthCheck interface{} `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#id ComputeInstanceGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#labels ComputeInstanceGroup#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// load_balancer block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#load_balancer ComputeInstanceGroup#load_balancer}
	LoadBalancer *ComputeInstanceGroupLoadBalancer `field:"optional" json:"loadBalancer" yaml:"loadBalancer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#max_checking_health_duration ComputeInstanceGroup#max_checking_health_duration}.
	MaxCheckingHealthDuration *float64 `field:"optional" json:"maxCheckingHealthDuration" yaml:"maxCheckingHealthDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#name ComputeInstanceGroup#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#timeouts ComputeInstanceGroup#timeouts}
	Timeouts *ComputeInstanceGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#variables ComputeInstanceGroup#variables}.
	Variables *map[string]*string `field:"optional" json:"variables" yaml:"variables"`
}

