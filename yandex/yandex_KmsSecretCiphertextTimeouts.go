// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type KmsSecretCiphertextTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kms_secret_ciphertext#create KmsSecretCiphertext#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kms_secret_ciphertext#delete KmsSecretCiphertext#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/kms_secret_ciphertext#read KmsSecretCiphertext#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

