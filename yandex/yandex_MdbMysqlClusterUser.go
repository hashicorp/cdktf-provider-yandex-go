// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbMysqlClusterUser struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#name MdbMysqlCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#password MdbMysqlCluster#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#authentication_plugin MdbMysqlCluster#authentication_plugin}.
	AuthenticationPlugin *string `field:"optional" json:"authenticationPlugin" yaml:"authenticationPlugin"`
	// connection_limits block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#connection_limits MdbMysqlCluster#connection_limits}
	ConnectionLimits *MdbMysqlClusterUserConnectionLimits `field:"optional" json:"connectionLimits" yaml:"connectionLimits"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#global_permissions MdbMysqlCluster#global_permissions}.
	GlobalPermissions *[]*string `field:"optional" json:"globalPermissions" yaml:"globalPermissions"`
	// permission block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#permission MdbMysqlCluster#permission}
	Permission interface{} `field:"optional" json:"permission" yaml:"permission"`
}

