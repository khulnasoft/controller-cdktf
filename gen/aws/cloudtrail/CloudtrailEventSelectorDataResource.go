package cloudtrail


type CloudtrailEventSelectorDataResource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/cloudtrail#type Cloudtrail#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/cloudtrail#values Cloudtrail#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

