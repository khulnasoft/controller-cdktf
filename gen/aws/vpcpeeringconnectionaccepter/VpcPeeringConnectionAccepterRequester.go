package vpcpeeringconnectionaccepter


type VpcPeeringConnectionAccepterRequester struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/vpc_peering_connection_accepter#allow_classic_link_to_remote_vpc VpcPeeringConnectionAccepterA#allow_classic_link_to_remote_vpc}.
	AllowClassicLinkToRemoteVpc interface{} `field:"optional" json:"allowClassicLinkToRemoteVpc" yaml:"allowClassicLinkToRemoteVpc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/vpc_peering_connection_accepter#allow_remote_vpc_dns_resolution VpcPeeringConnectionAccepterA#allow_remote_vpc_dns_resolution}.
	AllowRemoteVpcDnsResolution interface{} `field:"optional" json:"allowRemoteVpcDnsResolution" yaml:"allowRemoteVpcDnsResolution"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/vpc_peering_connection_accepter#allow_vpc_to_remote_classic_link VpcPeeringConnectionAccepterA#allow_vpc_to_remote_classic_link}.
	AllowVpcToRemoteClassicLink interface{} `field:"optional" json:"allowVpcToRemoteClassicLink" yaml:"allowVpcToRemoteClassicLink"`
}

