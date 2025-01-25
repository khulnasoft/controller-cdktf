//go:build no_runtime_type_checking

package inlinehook

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InlineHook) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (i *jsiiProxy_InlineHook) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (i *jsiiProxy_InlineHook) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_InlineHook) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_InlineHook) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_InlineHook) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_InlineHook) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_InlineHook) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_InlineHook) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_InlineHook) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_InlineHook) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_InlineHook) validateImportFromParameters(id *string) error {
	return nil
}

func (i *jsiiProxy_InlineHook) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_InlineHook) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (i *jsiiProxy_InlineHook) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (i *jsiiProxy_InlineHook) validateMoveToIdParameters(id *string) error {
	return nil
}

func (i *jsiiProxy_InlineHook) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (i *jsiiProxy_InlineHook) validatePutHeadersParameters(value interface{}) error {
	return nil
}

func validateInlineHook_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateInlineHook_IsConstructParameters(x interface{}) error {
	return nil
}

func validateInlineHook_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateInlineHook_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_InlineHook) validateSetAuthParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_InlineHook) validateSetChannelParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_InlineHook) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_InlineHook) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_InlineHook) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_InlineHook) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_InlineHook) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_InlineHook) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_InlineHook) validateSetStatusParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_InlineHook) validateSetTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_InlineHook) validateSetVersionParameters(val *string) error {
	return nil
}

func validateNewInlineHookParameters(scope constructs.Construct, id *string, config *InlineHookConfig) error {
	return nil
}

