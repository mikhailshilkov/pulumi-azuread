// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about an existing Application within Azure Active Directory.
// 
// > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.
func LookupApplication(ctx *pulumi.Context, args *GetApplicationArgs) (*GetApplicationResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
		inputs["objectId"] = args.ObjectId
	}
	outputs, err := ctx.Invoke("azuread:index/getApplication:getApplication", inputs)
	if err != nil {
		return nil, err
	}
	return &GetApplicationResult{
		ApplicationId: outputs["applicationId"],
		AvailableToOtherTenants: outputs["availableToOtherTenants"],
		Homepage: outputs["homepage"],
		IdentifierUris: outputs["identifierUris"],
		Name: outputs["name"],
		Oauth2AllowImplicitFlow: outputs["oauth2AllowImplicitFlow"],
		ObjectId: outputs["objectId"],
		ReplyUrls: outputs["replyUrls"],
		RequiredResourceAccesses: outputs["requiredResourceAccesses"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getApplication.
type GetApplicationArgs struct {
	// Specifies the name of the Application within Azure Active Directory.
	Name interface{}
	// Specifies the Object ID of the Application within Azure Active Directory.
	ObjectId interface{}
}

// A collection of values returned by getApplication.
type GetApplicationResult struct {
	// the Application ID of the Azure Active Directory Application.
	ApplicationId interface{}
	// Is this Azure AD Application available to other tenants?
	AvailableToOtherTenants interface{}
	Homepage interface{}
	// A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
	IdentifierUris interface{}
	Name interface{}
	// Does this Azure AD Application allow OAuth2.0 implicit flow tokens?
	Oauth2AllowImplicitFlow interface{}
	// the Object ID of the Azure Active Directory Application.
	ObjectId interface{}
	// A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
	ReplyUrls interface{}
	// A collection of `required_resource_access` blocks as documented below.
	RequiredResourceAccesses interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
