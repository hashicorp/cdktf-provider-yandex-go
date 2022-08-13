// Prebuilt yandex Provider for Terraform CDK (cdktf)
package yandex


type FunctionContent struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/yandex/r/function#zip_filename Function#zip_filename}.
	ZipFilename *string `field:"required" json:"zipFilename" yaml:"zipFilename"`
}

