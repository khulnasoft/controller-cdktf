package iottopicrule


type IotTopicRuleLambda struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/iot_topic_rule#function_arn IotTopicRule#function_arn}.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
}

