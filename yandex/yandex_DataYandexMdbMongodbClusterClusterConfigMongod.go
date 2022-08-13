// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type DataYandexMdbMongodbClusterClusterConfigMongod struct {
	// audit_log block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#audit_log DataYandexMdbMongodbCluster#audit_log}
	AuditLog *DataYandexMdbMongodbClusterClusterConfigMongodAuditLog `field:"optional" json:"auditLog" yaml:"auditLog"`
	// security block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#security DataYandexMdbMongodbCluster#security}
	Security *DataYandexMdbMongodbClusterClusterConfigMongodSecurity `field:"optional" json:"security" yaml:"security"`
	// set_parameter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/d/mdb_mongodb_cluster#set_parameter DataYandexMdbMongodbCluster#set_parameter}
	SetParameter *DataYandexMdbMongodbClusterClusterConfigMongodSetParameter `field:"optional" json:"setParameter" yaml:"setParameter"`
}

