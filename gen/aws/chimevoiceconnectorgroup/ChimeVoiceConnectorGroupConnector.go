package chimevoiceconnectorgroup


type ChimeVoiceConnectorGroupConnector struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/chime_voice_connector_group#priority ChimeVoiceConnectorGroup#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/chime_voice_connector_group#voice_connector_id ChimeVoiceConnectorGroup#voice_connector_id}.
	VoiceConnectorId *string `field:"required" json:"voiceConnectorId" yaml:"voiceConnectorId"`
}

