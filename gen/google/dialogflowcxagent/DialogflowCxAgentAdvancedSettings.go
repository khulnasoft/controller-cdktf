package dialogflowcxagent


type DialogflowCxAgentAdvancedSettings struct {
	// audio_export_gcs_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/dialogflow_cx_agent#audio_export_gcs_destination DialogflowCxAgent#audio_export_gcs_destination}
	AudioExportGcsDestination *DialogflowCxAgentAdvancedSettingsAudioExportGcsDestination `field:"optional" json:"audioExportGcsDestination" yaml:"audioExportGcsDestination"`
	// dtmf_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/dialogflow_cx_agent#dtmf_settings DialogflowCxAgent#dtmf_settings}
	DtmfSettings *DialogflowCxAgentAdvancedSettingsDtmfSettings `field:"optional" json:"dtmfSettings" yaml:"dtmfSettings"`
}

