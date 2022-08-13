// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type YdbDatabaseDedicatedConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/ydb_database_dedicated#name YdbDatabaseDedicated#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/ydb_database_dedicated#network_id YdbDatabaseDedicated#network_id}.
	NetworkId *string `field:"required" json:"networkId" yaml:"networkId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/ydb_database_dedicated#resource_preset_id YdbDatabaseDedicated#resource_preset_id}.
	ResourcePresetId *string `field:"required" json:"resourcePresetId" yaml:"resourcePresetId"`
	// scale_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/ydb_database_dedicated#scale_policy YdbDatabaseDedicated#scale_policy}
	ScalePolicy *YdbDatabaseDedicatedScalePolicy `field:"required" json:"scalePolicy" yaml:"scalePolicy"`
	// storage_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/ydb_database_dedicated#storage_config YdbDatabaseDedicated#storage_config}
	StorageConfig *YdbDatabaseDedicatedStorageConfig `field:"required" json:"storageConfig" yaml:"storageConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/ydb_database_dedicated#subnet_ids YdbDatabaseDedicated#subnet_ids}.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/ydb_database_dedicated#assign_public_ips YdbDatabaseDedicated#assign_public_ips}.
	AssignPublicIps interface{} `field:"optional" json:"assignPublicIps" yaml:"assignPublicIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/ydb_database_dedicated#description YdbDatabaseDedicated#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/ydb_database_dedicated#folder_id YdbDatabaseDedicated#folder_id}.
	FolderId *string `field:"optional" json:"folderId" yaml:"folderId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/ydb_database_dedicated#id YdbDatabaseDedicated#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/ydb_database_dedicated#labels YdbDatabaseDedicated#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// location block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/ydb_database_dedicated#location YdbDatabaseDedicated#location}
	Location *YdbDatabaseDedicatedLocation `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/ydb_database_dedicated#location_id YdbDatabaseDedicated#location_id}.
	LocationId *string `field:"optional" json:"locationId" yaml:"locationId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/ydb_database_dedicated#timeouts YdbDatabaseDedicated#timeouts}
	Timeouts *YdbDatabaseDedicatedTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

