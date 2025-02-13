package s3bucketreplicationconfiguration


type S3BucketReplicationConfigurationRuleFilterAnd struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/s3_bucket_replication_configuration#prefix S3BucketReplicationConfigurationA#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/s3_bucket_replication_configuration#tags S3BucketReplicationConfigurationA#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

