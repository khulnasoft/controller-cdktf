package codedeploydeploymentgroup


type CodedeployDeploymentGroupAutoRollbackConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/codedeploy_deployment_group#enabled CodedeployDeploymentGroup#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/codedeploy_deployment_group#events CodedeployDeploymentGroup#events}.
	Events *[]*string `field:"optional" json:"events" yaml:"events"`
}

