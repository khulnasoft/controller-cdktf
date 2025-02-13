package lambdafunction


type LambdaFunctionEnvironment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/lambda_function#variables LambdaFunction#variables}.
	Variables *map[string]*string `field:"optional" json:"variables" yaml:"variables"`
}

