package mskconnectconnector


type MskconnectConnectorTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/mskconnect_connector#create MskconnectConnector#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/mskconnect_connector#delete MskconnectConnector#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/mskconnect_connector#update MskconnectConnector#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

