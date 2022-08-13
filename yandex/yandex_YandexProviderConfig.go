// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type YandexProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex#alias YandexProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// ID of Yandex.Cloud tenant.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex#cloud_id YandexProvider#cloud_id}
	CloudId *string `field:"optional" json:"cloudId" yaml:"cloudId"`
	// The API endpoint for Yandex.Cloud SDK client.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex#endpoint YandexProvider#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// The default folder ID where resources will be placed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex#folder_id YandexProvider#folder_id}
	FolderId *string `field:"optional" json:"folderId" yaml:"folderId"`
	// Explicitly allow the provider to perform "insecure" SSL requests. If omitted,default value is `false`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex#insecure YandexProvider#insecure}
	Insecure interface{} `field:"optional" json:"insecure" yaml:"insecure"`
	// The maximum number of times an API request is being executed.
	//
	// If the API request still fails, an error is thrown.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex#max_retries YandexProvider#max_retries}
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex#organization_id YandexProvider#organization_id}.
	OrganizationId *string `field:"optional" json:"organizationId" yaml:"organizationId"`
	// Disable use of TLS. Default value is `false`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex#plaintext YandexProvider#plaintext}
	Plaintext interface{} `field:"optional" json:"plaintext" yaml:"plaintext"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex#region_id YandexProvider#region_id}.
	RegionId *string `field:"optional" json:"regionId" yaml:"regionId"`
	// Either the path to or the contents of a Service Account key file in JSON format.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex#service_account_key_file YandexProvider#service_account_key_file}
	ServiceAccountKeyFile *string `field:"optional" json:"serviceAccountKeyFile" yaml:"serviceAccountKeyFile"`
	// Yandex.Cloud storage service access key.  Used when a storage data/resource doesn't have an access key explicitly specified.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex#storage_access_key YandexProvider#storage_access_key}
	StorageAccessKey *string `field:"optional" json:"storageAccessKey" yaml:"storageAccessKey"`
	// Yandex.Cloud storage service endpoint. Default is  storage.yandexcloud.net.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex#storage_endpoint YandexProvider#storage_endpoint}
	StorageEndpoint *string `field:"optional" json:"storageEndpoint" yaml:"storageEndpoint"`
	// Yandex.Cloud storage service secret key.  Used when a storage data/resource doesn't have a secret key explicitly specified.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex#storage_secret_key YandexProvider#storage_secret_key}
	StorageSecretKey *string `field:"optional" json:"storageSecretKey" yaml:"storageSecretKey"`
	// The access token for API operations.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex#token YandexProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// Yandex.Cloud Message Queue service access key.  Used when a message queue resource doesn't have an access key explicitly specified.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex#ymq_access_key YandexProvider#ymq_access_key}
	YmqAccessKey *string `field:"optional" json:"ymqAccessKey" yaml:"ymqAccessKey"`
	// Yandex.Cloud Message Queue service endpoint. Default is  message-queue.api.cloud.yandex.net.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex#ymq_endpoint YandexProvider#ymq_endpoint}
	YmqEndpoint *string `field:"optional" json:"ymqEndpoint" yaml:"ymqEndpoint"`
	// Yandex.Cloud Message Queue service secret key.  Used when a message queue resource doesn't have a secret key explicitly specified.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex#ymq_secret_key YandexProvider#ymq_secret_key}
	YmqSecretKey *string `field:"optional" json:"ymqSecretKey" yaml:"ymqSecretKey"`
	// The zone where operations will take place. Examples are ru-central1-a, ru-central2-c, etc.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex#zone YandexProvider#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

