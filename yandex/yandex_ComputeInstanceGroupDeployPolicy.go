// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type ComputeInstanceGroupDeployPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#max_expansion ComputeInstanceGroup#max_expansion}.
	MaxExpansion *float64 `field:"required" json:"maxExpansion" yaml:"maxExpansion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#max_unavailable ComputeInstanceGroup#max_unavailable}.
	MaxUnavailable *float64 `field:"required" json:"maxUnavailable" yaml:"maxUnavailable"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#max_creating ComputeInstanceGroup#max_creating}.
	MaxCreating *float64 `field:"optional" json:"maxCreating" yaml:"maxCreating"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#max_deleting ComputeInstanceGroup#max_deleting}.
	MaxDeleting *float64 `field:"optional" json:"maxDeleting" yaml:"maxDeleting"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#startup_duration ComputeInstanceGroup#startup_duration}.
	StartupDuration *float64 `field:"optional" json:"startupDuration" yaml:"startupDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/compute_instance_group#strategy ComputeInstanceGroup#strategy}.
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
}

