package policyrulemfa

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PolicyRuleMfaConfig struct {
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
	// Policy Rule Name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_mfa#name PolicyRuleMfa#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// app_exclude block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_mfa#app_exclude PolicyRuleMfa#app_exclude}
	AppExclude interface{} `field:"optional" json:"appExclude" yaml:"appExclude"`
	// app_include block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_mfa#app_include PolicyRuleMfa#app_include}
	AppInclude interface{} `field:"optional" json:"appInclude" yaml:"appInclude"`
	// When a user should be prompted for MFA. It can be `CHALLENGE`, `LOGIN`, or `NEVER`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_mfa#enroll PolicyRuleMfa#enroll}
	Enroll *string `field:"optional" json:"enroll" yaml:"enroll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_mfa#id PolicyRuleMfa#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Network selection mode: `ANYWHERE`, `ZONE`, `ON_NETWORK`, or `OFF_NETWORK`. Default: `ANYWHERE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_mfa#network_connection PolicyRuleMfa#network_connection}
	NetworkConnection *string `field:"optional" json:"networkConnection" yaml:"networkConnection"`
	// Required if `network_connection` = `ZONE`. Indicates the network zones to exclude.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_mfa#network_excludes PolicyRuleMfa#network_excludes}
	NetworkExcludes *[]*string `field:"optional" json:"networkExcludes" yaml:"networkExcludes"`
	// Required if `network_connection` = `ZONE`. Indicates the network zones to include.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_mfa#network_includes PolicyRuleMfa#network_includes}
	NetworkIncludes *[]*string `field:"optional" json:"networkIncludes" yaml:"networkIncludes"`
	// Policy ID of the Rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_mfa#policy_id PolicyRuleMfa#policy_id}
	PolicyId *string `field:"optional" json:"policyId" yaml:"policyId"`
	// Rule priority.
	//
	// This attribute can be set to a valid priority. To avoid an endless diff situation an error is thrown if an invalid property is provided. The Okta API defaults to the last (lowest) if not provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_mfa#priority PolicyRuleMfa#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Policy Rule Status: `ACTIVE` or `INACTIVE`. Default: `ACTIVE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_mfa#status PolicyRuleMfa#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Set of User IDs to Exclude.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_mfa#users_excluded PolicyRuleMfa#users_excluded}
	UsersExcluded *[]*string `field:"optional" json:"usersExcluded" yaml:"usersExcluded"`
}

