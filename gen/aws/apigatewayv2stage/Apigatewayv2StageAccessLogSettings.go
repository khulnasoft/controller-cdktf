package apigatewayv2stage


type Apigatewayv2StageAccessLogSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/apigatewayv2_stage#destination_arn Apigatewayv2Stage#destination_arn}.
	DestinationArn *string `field:"required" json:"destinationArn" yaml:"destinationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/apigatewayv2_stage#format Apigatewayv2Stage#format}.
	Format *string `field:"required" json:"format" yaml:"format"`
}

