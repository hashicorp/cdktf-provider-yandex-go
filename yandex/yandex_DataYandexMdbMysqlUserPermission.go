// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexMdbMysqlUserPermission struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mysql_user#roles DataYandexMdbMysqlUser#roles}.
	Roles *[]*string `field:"optional" json:"roles" yaml:"roles"`
}

