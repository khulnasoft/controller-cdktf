package appuserschemaproperty


type AppUserSchemaPropertyOneOf struct {
	// Enum value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user_schema_property#const AppUserSchemaProperty#const}
	Const *string `field:"required" json:"const" yaml:"const"`
	// Enum title.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user_schema_property#title AppUserSchemaProperty#title}
	Title *string `field:"required" json:"title" yaml:"title"`
}

