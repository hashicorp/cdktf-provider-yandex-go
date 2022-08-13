// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbMongodbClusterClusterConfigMongod struct {
	// audit_log block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#audit_log MdbMongodbCluster#audit_log}
	AuditLog *MdbMongodbClusterClusterConfigMongodAuditLog `field:"optional" json:"auditLog" yaml:"auditLog"`
	// security block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#security MdbMongodbCluster#security}
	Security *MdbMongodbClusterClusterConfigMongodSecurity `field:"optional" json:"security" yaml:"security"`
	// set_parameter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mongodb_cluster#set_parameter MdbMongodbCluster#set_parameter}
	SetParameter *MdbMongodbClusterClusterConfigMongodSetParameter `field:"optional" json:"setParameter" yaml:"setParameter"`
}

