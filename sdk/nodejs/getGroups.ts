// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Gets Object IDs or Display Names for multiple Azure Active Directory groups.
 * 
 * > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Read directory data` within the `Windows Azure Active Directory` API.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 * 
 * const groups = azuread.getGroups({
 *     names: [
 *         "group-a",
 *         "group-b",
 *     ],
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azuread/blob/master/website/docs/d/groups.html.markdown.
 */
export function getGroups(args?: GetGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupsResult> & GetGroupsResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetGroupsResult> = pulumi.runtime.invoke("azuread:index/getGroups:getGroups", {
        "names": args.names,
        "objectIds": args.objectIds,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getGroups.
 */
export interface GetGroupsArgs {
    /**
     * The Display Names of the Azure AD Groups.
     */
    readonly names?: string[];
    /**
     * The Object IDs of the Azure AD Groups.
     */
    readonly objectIds?: string[];
}

/**
 * A collection of values returned by getGroups.
 */
export interface GetGroupsResult {
    /**
     * The Display Names of the Azure AD Groups.
     */
    readonly names: string[];
    /**
     * The Object IDs of the Azure AD Groups.
     */
    readonly objectIds: string[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
