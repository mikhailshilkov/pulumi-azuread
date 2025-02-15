// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Service Principal associated with an Application within Azure Active Directory.
// 
// > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API. Please see The Granting a Service Principal permission to manage AAD for the required steps.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azuread/blob/master/website/docs/r/service_principal.html.markdown.
type ServicePrincipal struct {
	s *pulumi.ResourceState
}

// NewServicePrincipal registers a new resource with the given unique name, arguments, and options.
func NewServicePrincipal(ctx *pulumi.Context,
	name string, args *ServicePrincipalArgs, opts ...pulumi.ResourceOpt) (*ServicePrincipal, error) {
	if args == nil || args.ApplicationId == nil {
		return nil, errors.New("missing required argument 'ApplicationId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["appRoleAssignmentRequired"] = nil
		inputs["applicationId"] = nil
		inputs["oauth2Permissions"] = nil
		inputs["tags"] = nil
	} else {
		inputs["appRoleAssignmentRequired"] = args.AppRoleAssignmentRequired
		inputs["applicationId"] = args.ApplicationId
		inputs["oauth2Permissions"] = args.Oauth2Permissions
		inputs["tags"] = args.Tags
	}
	inputs["displayName"] = nil
	inputs["objectId"] = nil
	s, err := ctx.RegisterResource("azuread:index/servicePrincipal:ServicePrincipal", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ServicePrincipal{s: s}, nil
}

// GetServicePrincipal gets an existing ServicePrincipal resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServicePrincipal(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ServicePrincipalState, opts ...pulumi.ResourceOpt) (*ServicePrincipal, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["appRoleAssignmentRequired"] = state.AppRoleAssignmentRequired
		inputs["applicationId"] = state.ApplicationId
		inputs["displayName"] = state.DisplayName
		inputs["oauth2Permissions"] = state.Oauth2Permissions
		inputs["objectId"] = state.ObjectId
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("azuread:index/servicePrincipal:ServicePrincipal", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ServicePrincipal{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ServicePrincipal) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ServicePrincipal) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Does this Service Principal require an AppRoleAssignment to a user or group before Azure AD will issue a user or access token to the application? Defaults to `false`.
func (r *ServicePrincipal) AppRoleAssignmentRequired() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["appRoleAssignmentRequired"])
}

// The ID of the Azure AD Application for which to create a Service Principal.
func (r *ServicePrincipal) ApplicationId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["applicationId"])
}

// The Display Name of the Azure Active Directory Application associated with this Service Principal.
func (r *ServicePrincipal) DisplayName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["displayName"])
}

// A collection of OAuth 2.0 permissions exposed by the associated application. Each permission is covered by a `oauth2Permission` block as documented below.
func (r *ServicePrincipal) Oauth2Permissions() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["oauth2Permissions"])
}

// The Service Principal's Object ID.
func (r *ServicePrincipal) ObjectId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["objectId"])
}

// A list of tags to apply to the Service Principal.
func (r *ServicePrincipal) Tags() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering ServicePrincipal resources.
type ServicePrincipalState struct {
	// Does this Service Principal require an AppRoleAssignment to a user or group before Azure AD will issue a user or access token to the application? Defaults to `false`.
	AppRoleAssignmentRequired interface{}
	// The ID of the Azure AD Application for which to create a Service Principal.
	ApplicationId interface{}
	// The Display Name of the Azure Active Directory Application associated with this Service Principal.
	DisplayName interface{}
	// A collection of OAuth 2.0 permissions exposed by the associated application. Each permission is covered by a `oauth2Permission` block as documented below.
	Oauth2Permissions interface{}
	// The Service Principal's Object ID.
	ObjectId interface{}
	// A list of tags to apply to the Service Principal.
	Tags interface{}
}

// The set of arguments for constructing a ServicePrincipal resource.
type ServicePrincipalArgs struct {
	// Does this Service Principal require an AppRoleAssignment to a user or group before Azure AD will issue a user or access token to the application? Defaults to `false`.
	AppRoleAssignmentRequired interface{}
	// The ID of the Azure AD Application for which to create a Service Principal.
	ApplicationId interface{}
	// A collection of OAuth 2.0 permissions exposed by the associated application. Each permission is covered by a `oauth2Permission` block as documented below.
	Oauth2Permissions interface{}
	// A list of tags to apply to the Service Principal.
	Tags interface{}
}
