package efsfilesystem


type EfsFileSystemLifecyclePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/efs_file_system#transition_to_ia EfsFileSystem#transition_to_ia}.
	TransitionToIa *string `field:"optional" json:"transitionToIa" yaml:"transitionToIa"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/efs_file_system#transition_to_primary_storage_class EfsFileSystem#transition_to_primary_storage_class}.
	TransitionToPrimaryStorageClass *string `field:"optional" json:"transitionToPrimaryStorageClass" yaml:"transitionToPrimaryStorageClass"`
}

