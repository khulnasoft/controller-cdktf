package appuser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppUserConfig struct {
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
	// App to associate user with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user#app_id AppUser#app_id}
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// User associated with the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user#user_id AppUser#user_id}
	UserId *string `field:"required" json:"userId" yaml:"userId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user#id AppUser#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The password to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user#password AppUser#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The JSON profile of the App User.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user#profile AppUser#profile}
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Retain the user assignment on destroy.
	//
	// If set to true, the resource will be removed from state but not from the Okta app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user#retain_assignment AppUser#retain_assignment}
	RetainAssignment interface{} `field:"optional" json:"retainAssignment" yaml:"retainAssignment"`
	// The username to use for the app user.
	//
	// In case the user is assigned to the app with `SHARED_USERNAME_AND_PASSWORD` credentials scheme, this field will be computed and should not be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user#username AppUser#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

