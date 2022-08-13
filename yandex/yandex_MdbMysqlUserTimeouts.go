// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbMysqlUserTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_user#create MdbMysqlUser#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_user#delete MdbMysqlUser#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_user#read MdbMysqlUser#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_user#update MdbMysqlUser#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

