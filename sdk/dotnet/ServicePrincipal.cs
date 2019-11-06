// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD
{
    /// <summary>
    /// Manages a Service Principal associated with an Application within Azure Active Directory.
    /// 
    /// &gt; **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API. Please see The Granting a Service Principal permission to manage AAD for the required steps.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azuread/blob/master/website/docs/r/service_principal.html.markdown.
    /// </summary>
    public partial class ServicePrincipal : Pulumi.CustomResource
    {
        /// <summary>
        /// Does this Service Principal require an AppRoleAssignment to a user or group before Azure AD will issue a user or access token to the application? Defaults to `false`.
        /// </summary>
        [Output("appRoleAssignmentRequired")]
        public Output<bool?> AppRoleAssignmentRequired { get; private set; } = null!;

        /// <summary>
        /// The ID of the Azure AD Application for which to create a Service Principal.
        /// </summary>
        [Output("applicationId")]
        public Output<string> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// The Display Name of the Azure Active Directory Application associated with this Service Principal.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// A collection of OAuth 2.0 permissions exposed by the associated application. Each permission is covered by a `oauth2_permission` block as documented below.
        /// </summary>
        [Output("oauth2Permissions")]
        public Output<ImmutableArray<Outputs.ServicePrincipalOauth2Permissions>> Oauth2Permissions { get; private set; } = null!;

        /// <summary>
        /// The Service Principal's Object ID.
        /// </summary>
        [Output("objectId")]
        public Output<string> ObjectId { get; private set; } = null!;

        /// <summary>
        /// A list of tags to apply to the Service Principal.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ServicePrincipal resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServicePrincipal(string name, ServicePrincipalArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/servicePrincipal:ServicePrincipal", name, args, MakeResourceOptions(options, ""))
        {
        }

        private ServicePrincipal(string name, Input<string> id, ServicePrincipalState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/servicePrincipal:ServicePrincipal", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServicePrincipal resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServicePrincipal Get(string name, Input<string> id, ServicePrincipalState? state = null, CustomResourceOptions? options = null)
        {
            return new ServicePrincipal(name, id, state, options);
        }
    }

    public sealed class ServicePrincipalArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Does this Service Principal require an AppRoleAssignment to a user or group before Azure AD will issue a user or access token to the application? Defaults to `false`.
        /// </summary>
        [Input("appRoleAssignmentRequired")]
        public Input<bool>? AppRoleAssignmentRequired { get; set; }

        /// <summary>
        /// The ID of the Azure AD Application for which to create a Service Principal.
        /// </summary>
        [Input("applicationId", required: true)]
        public Input<string> ApplicationId { get; set; } = null!;

        [Input("oauth2Permissions")]
        private InputList<Inputs.ServicePrincipalOauth2PermissionsArgs>? _oauth2Permissions;

        /// <summary>
        /// A collection of OAuth 2.0 permissions exposed by the associated application. Each permission is covered by a `oauth2_permission` block as documented below.
        /// </summary>
        public InputList<Inputs.ServicePrincipalOauth2PermissionsArgs> Oauth2Permissions
        {
            get => _oauth2Permissions ?? (_oauth2Permissions = new InputList<Inputs.ServicePrincipalOauth2PermissionsArgs>());
            set => _oauth2Permissions = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of tags to apply to the Service Principal.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public ServicePrincipalArgs()
        {
        }
    }

    public sealed class ServicePrincipalState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Does this Service Principal require an AppRoleAssignment to a user or group before Azure AD will issue a user or access token to the application? Defaults to `false`.
        /// </summary>
        [Input("appRoleAssignmentRequired")]
        public Input<bool>? AppRoleAssignmentRequired { get; set; }

        /// <summary>
        /// The ID of the Azure AD Application for which to create a Service Principal.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// The Display Name of the Azure Active Directory Application associated with this Service Principal.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("oauth2Permissions")]
        private InputList<Inputs.ServicePrincipalOauth2PermissionsGetArgs>? _oauth2Permissions;

        /// <summary>
        /// A collection of OAuth 2.0 permissions exposed by the associated application. Each permission is covered by a `oauth2_permission` block as documented below.
        /// </summary>
        public InputList<Inputs.ServicePrincipalOauth2PermissionsGetArgs> Oauth2Permissions
        {
            get => _oauth2Permissions ?? (_oauth2Permissions = new InputList<Inputs.ServicePrincipalOauth2PermissionsGetArgs>());
            set => _oauth2Permissions = value;
        }

        /// <summary>
        /// The Service Principal's Object ID.
        /// </summary>
        [Input("objectId")]
        public Input<string>? ObjectId { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of tags to apply to the Service Principal.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public ServicePrincipalState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ServicePrincipalOauth2PermissionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the admin consent.
        /// </summary>
        [Input("adminConsentDescription")]
        public Input<string>? AdminConsentDescription { get; set; }

        /// <summary>
        /// The display name of the admin consent.
        /// </summary>
        [Input("adminConsentDisplayName")]
        public Input<string>? AdminConsentDisplayName { get; set; }

        /// <summary>
        /// The unique identifier for one of the `OAuth2Permission`.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Is this permission enabled?
        /// </summary>
        [Input("isEnabled")]
        public Input<bool>? IsEnabled { get; set; }

        /// <summary>
        /// The type of the permission.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The description of the user consent.
        /// </summary>
        [Input("userConsentDescription")]
        public Input<string>? UserConsentDescription { get; set; }

        /// <summary>
        /// The display name of the user consent.
        /// </summary>
        [Input("userConsentDisplayName")]
        public Input<string>? UserConsentDisplayName { get; set; }

        /// <summary>
        /// The name of this permission.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public ServicePrincipalOauth2PermissionsArgs()
        {
        }
    }

    public sealed class ServicePrincipalOauth2PermissionsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the admin consent.
        /// </summary>
        [Input("adminConsentDescription")]
        public Input<string>? AdminConsentDescription { get; set; }

        /// <summary>
        /// The display name of the admin consent.
        /// </summary>
        [Input("adminConsentDisplayName")]
        public Input<string>? AdminConsentDisplayName { get; set; }

        /// <summary>
        /// The unique identifier for one of the `OAuth2Permission`.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Is this permission enabled?
        /// </summary>
        [Input("isEnabled")]
        public Input<bool>? IsEnabled { get; set; }

        /// <summary>
        /// The type of the permission.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The description of the user consent.
        /// </summary>
        [Input("userConsentDescription")]
        public Input<string>? UserConsentDescription { get; set; }

        /// <summary>
        /// The display name of the user consent.
        /// </summary>
        [Input("userConsentDisplayName")]
        public Input<string>? UserConsentDisplayName { get; set; }

        /// <summary>
        /// The name of this permission.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public ServicePrincipalOauth2PermissionsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ServicePrincipalOauth2Permissions
    {
        /// <summary>
        /// The description of the admin consent.
        /// </summary>
        public readonly string AdminConsentDescription;
        /// <summary>
        /// The display name of the admin consent.
        /// </summary>
        public readonly string AdminConsentDisplayName;
        /// <summary>
        /// The unique identifier for one of the `OAuth2Permission`.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Is this permission enabled?
        /// </summary>
        public readonly bool IsEnabled;
        /// <summary>
        /// The type of the permission.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The description of the user consent.
        /// </summary>
        public readonly string UserConsentDescription;
        /// <summary>
        /// The display name of the user consent.
        /// </summary>
        public readonly string UserConsentDisplayName;
        /// <summary>
        /// The name of this permission.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private ServicePrincipalOauth2Permissions(
            string adminConsentDescription,
            string adminConsentDisplayName,
            string id,
            bool isEnabled,
            string type,
            string userConsentDescription,
            string userConsentDisplayName,
            string value)
        {
            AdminConsentDescription = adminConsentDescription;
            AdminConsentDisplayName = adminConsentDisplayName;
            Id = id;
            IsEnabled = isEnabled;
            Type = type;
            UserConsentDescription = userConsentDescription;
            UserConsentDisplayName = userConsentDisplayName;
            Value = value;
        }
    }
    }
}