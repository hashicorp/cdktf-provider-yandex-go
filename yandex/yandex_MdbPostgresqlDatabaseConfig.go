// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MdbPostgresqlDatabaseConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_database#cluster_id MdbPostgresqlDatabase#cluster_id}.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_database#name MdbPostgresqlDatabase#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_database#owner MdbPostgresqlDatabase#owner}.
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// extension block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_database#extension MdbPostgresqlDatabase#extension}
	Extension interface{} `field:"optional" json:"extension" yaml:"extension"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_database#id MdbPostgresqlDatabase#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_database#lc_collate MdbPostgresqlDatabase#lc_collate}.
	LcCollate *string `field:"optional" json:"lcCollate" yaml:"lcCollate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_database#lc_type MdbPostgresqlDatabase#lc_type}.
	LcType *string `field:"optional" json:"lcType" yaml:"lcType"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_postgresql_database#timeouts MdbPostgresqlDatabase#timeouts}
	Timeouts *MdbPostgresqlDatabaseTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

