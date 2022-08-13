// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type MdbMysqlUserConnectionLimits struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_user#max_connections_per_hour MdbMysqlUser#max_connections_per_hour}.
	MaxConnectionsPerHour *float64 `field:"optional" json:"maxConnectionsPerHour" yaml:"maxConnectionsPerHour"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_user#max_questions_per_hour MdbMysqlUser#max_questions_per_hour}.
	MaxQuestionsPerHour *float64 `field:"optional" json:"maxQuestionsPerHour" yaml:"maxQuestionsPerHour"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_user#max_updates_per_hour MdbMysqlUser#max_updates_per_hour}.
	MaxUpdatesPerHour *float64 `field:"optional" json:"maxUpdatesPerHour" yaml:"maxUpdatesPerHour"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/mdb_mysql_user#max_user_connections MdbMysqlUser#max_user_connections}.
	MaxUserConnections *float64 `field:"optional" json:"maxUserConnections" yaml:"maxUserConnections"`
}

