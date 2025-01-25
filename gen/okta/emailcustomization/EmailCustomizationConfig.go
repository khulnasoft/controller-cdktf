package emailcustomization

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EmailCustomizationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Brand ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/email_customization#brand_id EmailCustomization#brand_id}
	BrandId *string `field:"required" json:"brandId" yaml:"brandId"`
	// Template Name - Example values: `AccountLockout`,`ADForgotPassword`,`ADForgotPasswordDenied`,`ADSelfServiceUnlock`,`ADUserActivation`,`AuthenticatorEnrolled`,`AuthenticatorReset`,`ChangeEmailConfirmation`,`EmailChallenge`,`EmailChangeConfirmation`,`EmailFactorVerification`,`ForgotPassword`,`ForgotPasswordDenied`,`IGAReviewerEndNotification`,`IGAReviewerNotification`,`IGAReviewerPendingNotification`,`IGAReviewerReassigned`,`LDAPForgotPassword`,`LDAPForgotPasswordDenied`,`LDAPSelfServiceUnlock`,`LDAPUserActivation`,`MyAccountChangeConfirmation`,`NewSignOnNotification`,`OktaVerifyActivation`,`PasswordChanged`,`PasswordResetByAdmin`,`PendingEmailChange`,`RegistrationActivation`,`RegistrationEmailVerification`,`SelfServiceUnlock`,`SelfServiceUnlockOnUnlockedAccount`,`UserActivation`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/email_customization#template_name EmailCustomization#template_name}
	TemplateName *string `field:"required" json:"templateName" yaml:"templateName"`
	// The body of the customization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/email_customization#body EmailCustomization#body}
	Body *string `field:"optional" json:"body" yaml:"body"`
	// Force is_default on the create and delete by deleting all email customizations.
	//
	// Comma separated string with values of 'create' or 'destroy' or both `create,destroy'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/email_customization#force_is_default EmailCustomization#force_is_default}
	ForceIsDefault *string `field:"optional" json:"forceIsDefault" yaml:"forceIsDefault"`
	// Whether the customization is the default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/email_customization#is_default EmailCustomization#is_default}
	IsDefault interface{} `field:"optional" json:"isDefault" yaml:"isDefault"`
	// The language supported by the customization - Example values from [supported languages](https://developer.okta.com/docs/reference/api/brands/#supported-languages).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/email_customization#language EmailCustomization#language}
	Language *string `field:"optional" json:"language" yaml:"language"`
	// The subject of the customization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/email_customization#subject EmailCustomization#subject}
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
}

