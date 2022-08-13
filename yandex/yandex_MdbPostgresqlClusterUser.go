// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbPostgresqlClusterUser struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#name MdbPostgresqlCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#password MdbPostgresqlCluster#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#conn_limit MdbPostgresqlCluster#conn_limit}.
	ConnLimit *float64 `field:"optional" json:"connLimit" yaml:"connLimit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#grants MdbPostgresqlCluster#grants}.
	Grants *[]*string `field:"optional" json:"grants" yaml:"grants"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#login MdbPostgresqlCluster#login}.
	Login interface{} `field:"optional" json:"login" yaml:"login"`
	// permission block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#permission MdbPostgresqlCluster#permission}
	Permission interface{} `field:"optional" json:"permission" yaml:"permission"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_cluster#settings MdbPostgresqlCluster#settings}.
	Settings *map[string]*string `field:"optional" json:"settings" yaml:"settings"`
}

