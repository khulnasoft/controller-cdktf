package idpoidc

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/khulnasoft/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/khulnasoft/controller-cdktf/gen/okta/idpoidc/internal"
)

// Represents a {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc okta_idp_oidc}.
type IdpOidc interface {
	cdktf.TerraformResource
	AccountLinkAction() *string
	SetAccountLinkAction(val *string)
	AccountLinkActionInput() *string
	AccountLinkGroupInclude() *[]*string
	SetAccountLinkGroupInclude(val *[]*string)
	AccountLinkGroupIncludeInput() *[]*string
	AuthorizationBinding() *string
	SetAuthorizationBinding(val *string)
	AuthorizationBindingInput() *string
	AuthorizationUrl() *string
	SetAuthorizationUrl(val *string)
	AuthorizationUrlInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeprovisionedAction() *string
	SetDeprovisionedAction(val *string)
	DeprovisionedActionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GroupsAction() *string
	SetGroupsAction(val *string)
	GroupsActionInput() *string
	GroupsAssignment() *[]*string
	SetGroupsAssignment(val *[]*string)
	GroupsAssignmentInput() *[]*string
	GroupsAttribute() *string
	SetGroupsAttribute(val *string)
	GroupsAttributeInput() *string
	GroupsFilter() *[]*string
	SetGroupsFilter(val *[]*string)
	GroupsFilterInput() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IssuerMode() *string
	SetIssuerMode(val *string)
	IssuerModeInput() *string
	IssuerUrl() *string
	SetIssuerUrl(val *string)
	IssuerUrlInput() *string
	JwksBinding() *string
	SetJwksBinding(val *string)
	JwksBindingInput() *string
	JwksUrl() *string
	SetJwksUrl(val *string)
	JwksUrlInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxClockSkew() *float64
	SetMaxClockSkew(val *float64)
	MaxClockSkewInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PkceRequired() interface{}
	SetPkceRequired(val interface{})
	PkceRequiredInput() interface{}
	ProfileMaster() interface{}
	SetProfileMaster(val interface{})
	ProfileMasterInput() interface{}
	ProtocolType() *string
	SetProtocolType(val *string)
	ProtocolTypeInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	ProvisioningAction() *string
	SetProvisioningAction(val *string)
	ProvisioningActionInput() *string
	// Experimental.
	RawOverrides() interface{}
	RequestSignatureAlgorithm() *string
	SetRequestSignatureAlgorithm(val *string)
	RequestSignatureAlgorithmInput() *string
	RequestSignatureScope() *string
	SetRequestSignatureScope(val *string)
	RequestSignatureScopeInput() *string
	Scopes() *[]*string
	SetScopes(val *[]*string)
	ScopesInput() *[]*string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	SubjectMatchAttribute() *string
	SetSubjectMatchAttribute(val *string)
	SubjectMatchAttributeInput() *string
	SubjectMatchType() *string
	SetSubjectMatchType(val *string)
	SubjectMatchTypeInput() *string
	SuspendedAction() *string
	SetSuspendedAction(val *string)
	SuspendedActionInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TokenBinding() *string
	SetTokenBinding(val *string)
	TokenBindingInput() *string
	TokenUrl() *string
	SetTokenUrl(val *string)
	TokenUrlInput() *string
	Type() *string
	UserInfoBinding() *string
	SetUserInfoBinding(val *string)
	UserInfoBindingInput() *string
	UserInfoUrl() *string
	SetUserInfoUrl(val *string)
	UserInfoUrlInput() *string
	UsernameTemplate() *string
	SetUsernameTemplate(val *string)
	UsernameTemplateInput() *string
	UserTypeId() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAccountLinkAction()
	ResetAccountLinkGroupInclude()
	ResetDeprovisionedAction()
	ResetGroupsAction()
	ResetGroupsAssignment()
	ResetGroupsAttribute()
	ResetGroupsFilter()
	ResetId()
	ResetIssuerMode()
	ResetMaxClockSkew()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPkceRequired()
	ResetProfileMaster()
	ResetProtocolType()
	ResetProvisioningAction()
	ResetRequestSignatureAlgorithm()
	ResetRequestSignatureScope()
	ResetStatus()
	ResetSubjectMatchAttribute()
	ResetSubjectMatchType()
	ResetSuspendedAction()
	ResetUserInfoBinding()
	ResetUserInfoUrl()
	ResetUsernameTemplate()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for IdpOidc
type jsiiProxy_IdpOidc struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IdpOidc) AccountLinkAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountLinkAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) AccountLinkActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountLinkActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) AccountLinkGroupInclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountLinkGroupInclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) AccountLinkGroupIncludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountLinkGroupIncludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) AuthorizationBinding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) AuthorizationBindingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) AuthorizationUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) AuthorizationUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) DeprovisionedAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deprovisionedAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) DeprovisionedActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deprovisionedActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) GroupsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) GroupsActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) GroupsAssignment() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsAssignment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) GroupsAssignmentInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsAssignmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) GroupsAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) GroupsAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) GroupsFilter() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) GroupsFilterInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) IssuerMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) IssuerModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) IssuerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) IssuerUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) JwksBinding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) JwksBindingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) JwksUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) JwksUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) MaxClockSkew() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxClockSkew",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) MaxClockSkewInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxClockSkewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) PkceRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pkceRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) PkceRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pkceRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) ProfileMaster() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"profileMaster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) ProfileMasterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"profileMasterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) ProtocolType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) ProtocolTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) ProvisioningAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) ProvisioningActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) RequestSignatureAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestSignatureAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) RequestSignatureAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestSignatureAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) RequestSignatureScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestSignatureScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) RequestSignatureScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestSignatureScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) ScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) SubjectMatchAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectMatchAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) SubjectMatchAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectMatchAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) SubjectMatchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectMatchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) SubjectMatchTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectMatchTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) SuspendedAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suspendedAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) SuspendedActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suspendedActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) TokenBinding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) TokenBindingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) TokenUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) TokenUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) UserInfoBinding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInfoBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) UserInfoBindingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInfoBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) UserInfoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInfoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) UserInfoUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInfoUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) UsernameTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) UsernameTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpOidc) UserTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userTypeId",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc okta_idp_oidc} Resource.
func NewIdpOidc(scope constructs.Construct, id *string, config *IdpOidcConfig) IdpOidc {
	_init_.Initialize()

	if err := validateNewIdpOidcParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IdpOidc{}

	_jsii_.Create(
		"@cdktf/provider-okta.idpOidc.IdpOidc",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc okta_idp_oidc} Resource.
func NewIdpOidc_Override(i IdpOidc, scope constructs.Construct, id *string, config *IdpOidcConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.idpOidc.IdpOidc",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IdpOidc)SetAccountLinkAction(val *string) {
	if err := j.validateSetAccountLinkActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountLinkAction",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetAccountLinkGroupInclude(val *[]*string) {
	if err := j.validateSetAccountLinkGroupIncludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountLinkGroupInclude",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetAuthorizationBinding(val *string) {
	if err := j.validateSetAuthorizationBindingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationBinding",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetAuthorizationUrl(val *string) {
	if err := j.validateSetAuthorizationUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationUrl",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetDeprovisionedAction(val *string) {
	if err := j.validateSetDeprovisionedActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deprovisionedAction",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetGroupsAction(val *string) {
	if err := j.validateSetGroupsActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsAction",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetGroupsAssignment(val *[]*string) {
	if err := j.validateSetGroupsAssignmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsAssignment",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetGroupsAttribute(val *string) {
	if err := j.validateSetGroupsAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsAttribute",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetGroupsFilter(val *[]*string) {
	if err := j.validateSetGroupsFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsFilter",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetIssuerMode(val *string) {
	if err := j.validateSetIssuerModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuerMode",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetIssuerUrl(val *string) {
	if err := j.validateSetIssuerUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuerUrl",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetJwksBinding(val *string) {
	if err := j.validateSetJwksBindingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwksBinding",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetJwksUrl(val *string) {
	if err := j.validateSetJwksUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwksUrl",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetMaxClockSkew(val *float64) {
	if err := j.validateSetMaxClockSkewParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxClockSkew",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetPkceRequired(val interface{}) {
	if err := j.validateSetPkceRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pkceRequired",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetProfileMaster(val interface{}) {
	if err := j.validateSetProfileMasterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profileMaster",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetProtocolType(val *string) {
	if err := j.validateSetProtocolTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocolType",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetProvisioningAction(val *string) {
	if err := j.validateSetProvisioningActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioningAction",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetRequestSignatureAlgorithm(val *string) {
	if err := j.validateSetRequestSignatureAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestSignatureAlgorithm",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetRequestSignatureScope(val *string) {
	if err := j.validateSetRequestSignatureScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestSignatureScope",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetScopes(val *[]*string) {
	if err := j.validateSetScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetSubjectMatchAttribute(val *string) {
	if err := j.validateSetSubjectMatchAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectMatchAttribute",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetSubjectMatchType(val *string) {
	if err := j.validateSetSubjectMatchTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectMatchType",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetSuspendedAction(val *string) {
	if err := j.validateSetSuspendedActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspendedAction",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetTokenBinding(val *string) {
	if err := j.validateSetTokenBindingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenBinding",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetTokenUrl(val *string) {
	if err := j.validateSetTokenUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenUrl",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetUserInfoBinding(val *string) {
	if err := j.validateSetUserInfoBindingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userInfoBinding",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetUserInfoUrl(val *string) {
	if err := j.validateSetUserInfoUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userInfoUrl",
		val,
	)
}

func (j *jsiiProxy_IdpOidc)SetUsernameTemplate(val *string) {
	if err := j.validateSetUsernameTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernameTemplate",
		val,
	)
}

// Generates CDKTF code for importing a IdpOidc resource upon running "cdktf plan <stack-name>".
func IdpOidc_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateIdpOidc_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.idpOidc.IdpOidc",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func IdpOidc_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdpOidc_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.idpOidc.IdpOidc",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IdpOidc_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdpOidc_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.idpOidc.IdpOidc",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IdpOidc_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdpOidc_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.idpOidc.IdpOidc",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IdpOidc_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-okta.idpOidc.IdpOidc",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IdpOidc) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IdpOidc) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IdpOidc) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpOidc) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpOidc) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpOidc) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpOidc) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpOidc) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpOidc) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpOidc) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpOidc) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpOidc) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpOidc) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IdpOidc) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpOidc) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IdpOidc) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IdpOidc) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IdpOidc) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IdpOidc) ResetAccountLinkAction() {
	_jsii_.InvokeVoid(
		i,
		"resetAccountLinkAction",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpOidc) ResetAccountLinkGroupInclude() {
	_jsii_.InvokeVoid(
		i,
		"resetAccountLinkGroupInclude",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpOidc) ResetDeprovisionedAction() {
	_jsii_.InvokeVoid(
		i,
		"resetDeprovisionedAction",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpOidc) ResetGroupsAction() {
	_jsii_.InvokeVoid(
		i,
		"resetGroupsAction",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpOidc) ResetGroupsAssignment() {
	_jsii_.InvokeVoid(
		i,
		"resetGroupsAssignment",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpOidc) ResetGroupsAttribute() {
	_jsii_.InvokeVoid(
		i,
		"resetGroupsAttribute",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpOidc) ResetGroupsFilter() {
	_jsii_.InvokeVoid(
		i,
		"resetGroupsFilter",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpOidc) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpOidc) ResetIssuerMode() {
	_jsii_.InvokeVoid(
		i,
		"resetIssuerMode",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpOidc) ResetMaxClockSkew() {
	_jsii_.InvokeVoid(
		i,
		"resetMaxClockSkew",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpOidc) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpOidc) ResetPkceRequired() {
	_jsii_.InvokeVoid(
		i,
		"resetPkceRequired",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpOidc) ResetProfileMaster() {
	_jsii_.InvokeVoid(
		i,
		"resetProfileMaster",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpOidc) ResetProtocolType() {
	_jsii_.InvokeVoid(
		i,
		"resetProtocolType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpOidc) ResetProvisioningAction() {
	_jsii_.InvokeVoid(
		i,
		"resetProvisioningAction",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpOidc) ResetRequestSignatureAlgorithm() {
	_jsii_.InvokeVoid(
		i,
		"resetRequestSignatureAlgorithm",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpOidc) ResetRequestSignatureScope() {
	_jsii_.InvokeVoid(
		i,
		"resetRequestSignatureScope",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpOidc) ResetStatus() {
	_jsii_.InvokeVoid(
		i,
		"resetStatus",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpOidc) ResetSubjectMatchAttribute() {
	_jsii_.InvokeVoid(
		i,
		"resetSubjectMatchAttribute",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpOidc) ResetSubjectMatchType() {
	_jsii_.InvokeVoid(
		i,
		"resetSubjectMatchType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpOidc) ResetSuspendedAction() {
	_jsii_.InvokeVoid(
		i,
		"resetSuspendedAction",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpOidc) ResetUserInfoBinding() {
	_jsii_.InvokeVoid(
		i,
		"resetUserInfoBinding",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpOidc) ResetUserInfoUrl() {
	_jsii_.InvokeVoid(
		i,
		"resetUserInfoUrl",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpOidc) ResetUsernameTemplate() {
	_jsii_.InvokeVoid(
		i,
		"resetUsernameTemplate",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpOidc) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpOidc) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpOidc) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpOidc) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpOidc) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpOidc) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

