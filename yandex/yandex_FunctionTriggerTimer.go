// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type FunctionTriggerTimer struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function_trigger#cron_expression FunctionTrigger#cron_expression}.
	CronExpression *string `field:"required" json:"cronExpression" yaml:"cronExpression"`
}

