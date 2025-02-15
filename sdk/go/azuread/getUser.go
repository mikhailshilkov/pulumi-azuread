// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Gets information about an Azure Active Directory user.
// 
// > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Read directory data` within the `Windows Azure Active Directory` API.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azuread/blob/master/website/docs/d/user.html.markdown.
func LookupUser(ctx *pulumi.Context, args *GetUserArgs) (*GetUserResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["objectId"] = args.ObjectId
		inputs["userPrincipalName"] = args.UserPrincipalName
	}
	outputs, err := ctx.Invoke("azuread:index/getUser:getUser", inputs)
	if err != nil {
		return nil, err
	}
	return &GetUserResult{
		AccountEnabled: outputs["accountEnabled"],
		DisplayName: outputs["displayName"],
		Mail: outputs["mail"],
		MailNickname: outputs["mailNickname"],
		ObjectId: outputs["objectId"],
		UserPrincipalName: outputs["userPrincipalName"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getUser.
type GetUserArgs struct {
	// Specifies the Object ID of the Application within Azure Active Directory.
	ObjectId interface{}
	// The User Principal Name of the Azure AD User.
	UserPrincipalName interface{}
}

// A collection of values returned by getUser.
type GetUserResult struct {
	// `True` if the account is enabled; otherwise `False`.
	AccountEnabled interface{}
	// The Display Name of the Azure AD User.
	DisplayName interface{}
	// The primary email address of the Azure AD User.
	Mail interface{}
	// The email alias of the Azure AD User.
	MailNickname interface{}
	ObjectId interface{}
	// The User Principal Name of the Azure AD User.
	UserPrincipalName interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
