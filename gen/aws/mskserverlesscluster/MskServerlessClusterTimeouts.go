package mskserverlesscluster


type MskServerlessClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/msk_serverless_cluster#create MskServerlessCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/msk_serverless_cluster#delete MskServerlessCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

