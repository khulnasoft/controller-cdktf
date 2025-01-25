package emailsenderverification

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EmailSenderVerificationConfig struct {
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
	// Email sender ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/email_sender_verification#sender_id EmailSenderVerification#sender_id}
	SenderId *string `field:"required" json:"senderId" yaml:"senderId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/email_sender_verification#id EmailSenderVerification#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

