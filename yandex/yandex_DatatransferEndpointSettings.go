// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DatatransferEndpointSettings struct {
	// clickhouse_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#clickhouse_source DatatransferEndpoint#clickhouse_source}
	ClickhouseSource *DatatransferEndpointSettingsClickhouseSource `field:"optional" json:"clickhouseSource" yaml:"clickhouseSource"`
	// clickhouse_target block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#clickhouse_target DatatransferEndpoint#clickhouse_target}
	ClickhouseTarget *DatatransferEndpointSettingsClickhouseTarget `field:"optional" json:"clickhouseTarget" yaml:"clickhouseTarget"`
	// mongo_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#mongo_source DatatransferEndpoint#mongo_source}
	MongoSource *DatatransferEndpointSettingsMongoSource `field:"optional" json:"mongoSource" yaml:"mongoSource"`
	// mongo_target block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#mongo_target DatatransferEndpoint#mongo_target}
	MongoTarget *DatatransferEndpointSettingsMongoTarget `field:"optional" json:"mongoTarget" yaml:"mongoTarget"`
	// mysql_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#mysql_source DatatransferEndpoint#mysql_source}
	MysqlSource *DatatransferEndpointSettingsMysqlSource `field:"optional" json:"mysqlSource" yaml:"mysqlSource"`
	// mysql_target block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#mysql_target DatatransferEndpoint#mysql_target}
	MysqlTarget *DatatransferEndpointSettingsMysqlTarget `field:"optional" json:"mysqlTarget" yaml:"mysqlTarget"`
	// postgres_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#postgres_source DatatransferEndpoint#postgres_source}
	PostgresSource *DatatransferEndpointSettingsPostgresSource `field:"optional" json:"postgresSource" yaml:"postgresSource"`
	// postgres_target block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/datatransfer_endpoint#postgres_target DatatransferEndpoint#postgres_target}
	PostgresTarget *DatatransferEndpointSettingsPostgresTarget `field:"optional" json:"postgresTarget" yaml:"postgresTarget"`
}

