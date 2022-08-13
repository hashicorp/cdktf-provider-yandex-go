// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DatatransferEndpointSettingsMysqlTarget struct {
	// connection block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#connection DatatransferEndpoint#connection}
	Connection *DatatransferEndpointSettingsMysqlTargetConnection `field:"optional" json:"connection" yaml:"connection"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#database DatatransferEndpoint#database}.
	Database *string `field:"optional" json:"database" yaml:"database"`
	// password block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#password DatatransferEndpoint#password}
	Password *DatatransferEndpointSettingsMysqlTargetPassword `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#security_groups DatatransferEndpoint#security_groups}.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#skip_constraint_checks DatatransferEndpoint#skip_constraint_checks}.
	SkipConstraintChecks interface{} `field:"optional" json:"skipConstraintChecks" yaml:"skipConstraintChecks"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#sql_mode DatatransferEndpoint#sql_mode}.
	SqlMode *string `field:"optional" json:"sqlMode" yaml:"sqlMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#timezone DatatransferEndpoint#timezone}.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#user DatatransferEndpoint#user}.
	User *string `field:"optional" json:"user" yaml:"user"`
}

