package dataoktaappsaml

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOktaAppSamlConfig struct {
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
	// Search only ACTIVE applications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/app_saml#active_only DataOktaAppSaml#active_only}
	ActiveOnly interface{} `field:"optional" json:"activeOnly" yaml:"activeOnly"`
	// Id of application to retrieve, conflicts with label and label_prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/app_saml#id DataOktaAppSaml#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The label of the app to retrieve, conflicts with label_prefix and id.
	//
	// Label
	// 				uses the ?q=<label> query parameter exposed by Okta's API. It should be noted that at this time
	// 				this searches both name and label. This is used to avoid paginating through all applications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/app_saml#label DataOktaAppSaml#label}
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Label prefix of the app to retrieve, conflicts with label and id.
	//
	// This will tell the
	// 				provider to do a starts with query as opposed to an equals query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/app_saml#label_prefix DataOktaAppSaml#label_prefix}
	LabelPrefix *string `field:"optional" json:"labelPrefix" yaml:"labelPrefix"`
	// Denotes whether the request is compressed or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/app_saml#request_compressed DataOktaAppSaml#request_compressed}
	RequestCompressed interface{} `field:"optional" json:"requestCompressed" yaml:"requestCompressed"`
	// Ignore groups sync. This is a temporary solution until 'groups' field is supported in all the app-like resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/app_saml#skip_groups DataOktaAppSaml#skip_groups}
	SkipGroups interface{} `field:"optional" json:"skipGroups" yaml:"skipGroups"`
	// Ignore users sync. This is a temporary solution until 'users' field is supported in all the app-like resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/app_saml#skip_users DataOktaAppSaml#skip_users}
	SkipUsers interface{} `field:"optional" json:"skipUsers" yaml:"skipUsers"`
}

