package sagemakerdomain


type SagemakerDomainDefaultUserSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/sagemaker_domain#execution_role SagemakerDomain#execution_role}.
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// canvas_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/sagemaker_domain#canvas_app_settings SagemakerDomain#canvas_app_settings}
	CanvasAppSettings *SagemakerDomainDefaultUserSettingsCanvasAppSettings `field:"optional" json:"canvasAppSettings" yaml:"canvasAppSettings"`
	// jupyter_server_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/sagemaker_domain#jupyter_server_app_settings SagemakerDomain#jupyter_server_app_settings}
	JupyterServerAppSettings *SagemakerDomainDefaultUserSettingsJupyterServerAppSettings `field:"optional" json:"jupyterServerAppSettings" yaml:"jupyterServerAppSettings"`
	// kernel_gateway_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/sagemaker_domain#kernel_gateway_app_settings SagemakerDomain#kernel_gateway_app_settings}
	KernelGatewayAppSettings *SagemakerDomainDefaultUserSettingsKernelGatewayAppSettings `field:"optional" json:"kernelGatewayAppSettings" yaml:"kernelGatewayAppSettings"`
	// r_session_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/sagemaker_domain#r_session_app_settings SagemakerDomain#r_session_app_settings}
	RSessionAppSettings *SagemakerDomainDefaultUserSettingsRSessionAppSettings `field:"optional" json:"rSessionAppSettings" yaml:"rSessionAppSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/sagemaker_domain#security_groups SagemakerDomain#security_groups}.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// sharing_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/sagemaker_domain#sharing_settings SagemakerDomain#sharing_settings}
	SharingSettings *SagemakerDomainDefaultUserSettingsSharingSettings `field:"optional" json:"sharingSettings" yaml:"sharingSettings"`
	// tensor_board_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/sagemaker_domain#tensor_board_app_settings SagemakerDomain#tensor_board_app_settings}
	TensorBoardAppSettings *SagemakerDomainDefaultUserSettingsTensorBoardAppSettings `field:"optional" json:"tensorBoardAppSettings" yaml:"tensorBoardAppSettings"`
}

