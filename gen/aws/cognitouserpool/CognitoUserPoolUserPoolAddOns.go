package cognitouserpool


type CognitoUserPoolUserPoolAddOns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/cognito_user_pool#advanced_security_mode CognitoUserPool#advanced_security_mode}.
	AdvancedSecurityMode *string `field:"required" json:"advancedSecurityMode" yaml:"advancedSecurityMode"`
}

