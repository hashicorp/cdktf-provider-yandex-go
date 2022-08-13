// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbLoadBalancerAllocationPolicy struct {
	// location block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_load_balancer#location AlbLoadBalancer#location}
	Location interface{} `field:"required" json:"location" yaml:"location"`
}

