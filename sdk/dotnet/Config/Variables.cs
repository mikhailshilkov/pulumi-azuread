// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;

namespace Pulumi.AzureAD.Config
{
    public static class Config
    {
        private static readonly Pulumi.Config __config = new Pulumi.Config("azuread");

        public static string? ClientCertificatePassword { get; set; } = __config.Get("clientCertificatePassword") ?? Utilities.GetEnv("ARM_CLIENT_CERTIFICATE_PASSWORD") ?? "";

        public static string? ClientCertificatePath { get; set; } = __config.Get("clientCertificatePath") ?? Utilities.GetEnv("ARM_CLIENT_CERTIFICATE_PATH") ?? "";

        public static string? ClientId { get; set; } = __config.Get("clientId") ?? Utilities.GetEnv("ARM_CLIENT_ID") ?? "";

        public static string? ClientSecret { get; set; } = __config.Get("clientSecret") ?? Utilities.GetEnv("ARM_CLIENT_SECRET") ?? "";

        public static string? Environment { get; set; } = __config.Get("environment") ?? Utilities.GetEnv("ARM_ENVIRONMENT") ?? "public";

        public static string? MsiEndpoint { get; set; } = __config.Get("msiEndpoint") ?? Utilities.GetEnv("ARM_MSI_ENDPOINT") ?? "";

        public static string? SubscriptionId { get; set; } = __config.Get("subscriptionId") ?? Utilities.GetEnv("ARM_SUBSCRIPTION_ID") ?? "";

        public static string? TenantId { get; set; } = __config.Get("tenantId") ?? Utilities.GetEnv("ARM_TENANT_ID") ?? "";

        public static bool? UseMsi { get; set; } = __config.GetBoolean("useMsi") ?? Utilities.GetEnvBoolean("ARM_USE_MSI") ?? false;

    }
    namespace ConfigTypes
    {
    }
}
