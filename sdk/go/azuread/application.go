// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an Application within Azure Active Directory.
// 
// > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write owned by applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azuread/blob/master/website/docs/r/application.html.markdown.
type Application struct {
	s *pulumi.ResourceState
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOpt) (*Application, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["appRoles"] = nil
		inputs["availableToOtherTenants"] = nil
		inputs["groupMembershipClaims"] = nil
		inputs["homepage"] = nil
		inputs["identifierUris"] = nil
		inputs["name"] = nil
		inputs["oauth2AllowImplicitFlow"] = nil
		inputs["oauth2Permissions"] = nil
		inputs["publicClient"] = nil
		inputs["replyUrls"] = nil
		inputs["requiredResourceAccesses"] = nil
		inputs["type"] = nil
	} else {
		inputs["appRoles"] = args.AppRoles
		inputs["availableToOtherTenants"] = args.AvailableToOtherTenants
		inputs["groupMembershipClaims"] = args.GroupMembershipClaims
		inputs["homepage"] = args.Homepage
		inputs["identifierUris"] = args.IdentifierUris
		inputs["name"] = args.Name
		inputs["oauth2AllowImplicitFlow"] = args.Oauth2AllowImplicitFlow
		inputs["oauth2Permissions"] = args.Oauth2Permissions
		inputs["publicClient"] = args.PublicClient
		inputs["replyUrls"] = args.ReplyUrls
		inputs["requiredResourceAccesses"] = args.RequiredResourceAccesses
		inputs["type"] = args.Type
	}
	inputs["applicationId"] = nil
	inputs["objectId"] = nil
	s, err := ctx.RegisterResource("azuread:index/application:Application", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Application{s: s}, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ApplicationState, opts ...pulumi.ResourceOpt) (*Application, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["appRoles"] = state.AppRoles
		inputs["applicationId"] = state.ApplicationId
		inputs["availableToOtherTenants"] = state.AvailableToOtherTenants
		inputs["groupMembershipClaims"] = state.GroupMembershipClaims
		inputs["homepage"] = state.Homepage
		inputs["identifierUris"] = state.IdentifierUris
		inputs["name"] = state.Name
		inputs["oauth2AllowImplicitFlow"] = state.Oauth2AllowImplicitFlow
		inputs["oauth2Permissions"] = state.Oauth2Permissions
		inputs["objectId"] = state.ObjectId
		inputs["publicClient"] = state.PublicClient
		inputs["replyUrls"] = state.ReplyUrls
		inputs["requiredResourceAccesses"] = state.RequiredResourceAccesses
		inputs["type"] = state.Type
	}
	s, err := ctx.ReadResource("azuread:index/application:Application", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Application{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Application) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Application) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// A collection of `appRole` blocks as documented below. For more information https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles
func (r *Application) AppRoles() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["appRoles"])
}

// The Application ID.
func (r *Application) ApplicationId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["applicationId"])
}

// Is this Azure AD Application available to other tenants? Defaults to `false`.
func (r *Application) AvailableToOtherTenants() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["availableToOtherTenants"])
}

// Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Defaults to `SecurityGroup`. Possible values are `None`, `SecurityGroup` or `All`.
func (r *Application) GroupMembershipClaims() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["groupMembershipClaims"])
}

// The URL to the application's home page. If no homepage is specified this defaults to `https://{name}`.
func (r *Application) Homepage() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["homepage"])
}

// A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
func (r *Application) IdentifierUris() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["identifierUris"])
}

// The display name for the application.
func (r *Application) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to `false`.
func (r *Application) Oauth2AllowImplicitFlow() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["oauth2AllowImplicitFlow"])
}

// A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by a `oauth2Permission` block as documented below.
func (r *Application) Oauth2Permissions() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["oauth2Permissions"])
}

// The Application's Object ID.
func (r *Application) ObjectId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["objectId"])
}

// Is this Azure AD Application a public client? Defaults to `false`.
func (r *Application) PublicClient() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["publicClient"])
}

// A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
func (r *Application) ReplyUrls() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["replyUrls"])
}

// A collection of `requiredResourceAccess` blocks as documented below.
func (r *Application) RequiredResourceAccesses() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["requiredResourceAccesses"])
}

// Specifies whether the id property references an `OAuth2Permission` or an `AppRole`. Possible values are `Scope` or `Role`.
func (r *Application) Type() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["type"])
}

// Input properties used for looking up and filtering Application resources.
type ApplicationState struct {
	// A collection of `appRole` blocks as documented below. For more information https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles
	AppRoles interface{}
	// The Application ID.
	ApplicationId interface{}
	// Is this Azure AD Application available to other tenants? Defaults to `false`.
	AvailableToOtherTenants interface{}
	// Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Defaults to `SecurityGroup`. Possible values are `None`, `SecurityGroup` or `All`.
	GroupMembershipClaims interface{}
	// The URL to the application's home page. If no homepage is specified this defaults to `https://{name}`.
	Homepage interface{}
	// A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
	IdentifierUris interface{}
	// The display name for the application.
	Name interface{}
	// Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to `false`.
	Oauth2AllowImplicitFlow interface{}
	// A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by a `oauth2Permission` block as documented below.
	Oauth2Permissions interface{}
	// The Application's Object ID.
	ObjectId interface{}
	// Is this Azure AD Application a public client? Defaults to `false`.
	PublicClient interface{}
	// A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
	ReplyUrls interface{}
	// A collection of `requiredResourceAccess` blocks as documented below.
	RequiredResourceAccesses interface{}
	// Specifies whether the id property references an `OAuth2Permission` or an `AppRole`. Possible values are `Scope` or `Role`.
	Type interface{}
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// A collection of `appRole` blocks as documented below. For more information https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles
	AppRoles interface{}
	// Is this Azure AD Application available to other tenants? Defaults to `false`.
	AvailableToOtherTenants interface{}
	// Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Defaults to `SecurityGroup`. Possible values are `None`, `SecurityGroup` or `All`.
	GroupMembershipClaims interface{}
	// The URL to the application's home page. If no homepage is specified this defaults to `https://{name}`.
	Homepage interface{}
	// A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
	IdentifierUris interface{}
	// The display name for the application.
	Name interface{}
	// Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to `false`.
	Oauth2AllowImplicitFlow interface{}
	// A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by a `oauth2Permission` block as documented below.
	Oauth2Permissions interface{}
	// Is this Azure AD Application a public client? Defaults to `false`.
	PublicClient interface{}
	// A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
	ReplyUrls interface{}
	// A collection of `requiredResourceAccess` blocks as documented below.
	RequiredResourceAccesses interface{}
	// Specifies whether the id property references an `OAuth2Permission` or an `AppRole`. Possible values are `Scope` or `Role`.
	Type interface{}
}
