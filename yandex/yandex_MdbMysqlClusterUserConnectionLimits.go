// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbMysqlClusterUserConnectionLimits struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#max_connections_per_hour MdbMysqlCluster#max_connections_per_hour}.
	MaxConnectionsPerHour *float64 `field:"optional" json:"maxConnectionsPerHour" yaml:"maxConnectionsPerHour"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#max_questions_per_hour MdbMysqlCluster#max_questions_per_hour}.
	MaxQuestionsPerHour *float64 `field:"optional" json:"maxQuestionsPerHour" yaml:"maxQuestionsPerHour"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#max_updates_per_hour MdbMysqlCluster#max_updates_per_hour}.
	MaxUpdatesPerHour *float64 `field:"optional" json:"maxUpdatesPerHour" yaml:"maxUpdatesPerHour"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_cluster#max_user_connections MdbMysqlCluster#max_user_connections}.
	MaxUserConnections *float64 `field:"optional" json:"maxUserConnections" yaml:"maxUserConnections"`
}

