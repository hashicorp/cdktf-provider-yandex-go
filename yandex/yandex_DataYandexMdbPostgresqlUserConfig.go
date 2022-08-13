// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataYandexMdbPostgresqlUserConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_postgresql_user#cluster_id DataYandexMdbPostgresqlUser#cluster_id}.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_postgresql_user#name DataYandexMdbPostgresqlUser#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_postgresql_user#conn_limit DataYandexMdbPostgresqlUser#conn_limit}.
	ConnLimit *float64 `field:"optional" json:"connLimit" yaml:"connLimit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_postgresql_user#id DataYandexMdbPostgresqlUser#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_postgresql_user#login DataYandexMdbPostgresqlUser#login}.
	Login interface{} `field:"optional" json:"login" yaml:"login"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_postgresql_user#settings DataYandexMdbPostgresqlUser#settings}.
	Settings *map[string]*string `field:"optional" json:"settings" yaml:"settings"`
}

