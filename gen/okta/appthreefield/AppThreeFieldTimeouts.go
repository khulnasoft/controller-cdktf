package appthreefield


type AppThreeFieldTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#create AppThreeField#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#read AppThreeField#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#update AppThreeField#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

