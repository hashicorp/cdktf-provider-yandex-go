// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type AlbLoadBalancerListenerEndpoint struct {
	// address block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_load_balancer#address AlbLoadBalancer#address}
	Address interface{} `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/alb_load_balancer#ports AlbLoadBalancer#ports}.
	Ports *[]*float64 `field:"required" json:"ports" yaml:"ports"`
}

