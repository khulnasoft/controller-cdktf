package dnsrecordset


type DnsRecordSetRoutingPolicyGeoHealthCheckedTargets struct {
	// internal_load_balancers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/dns_record_set#internal_load_balancers DnsRecordSet#internal_load_balancers}
	InternalLoadBalancers interface{} `field:"required" json:"internalLoadBalancers" yaml:"internalLoadBalancers"`
}

