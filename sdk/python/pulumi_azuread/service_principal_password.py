# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class ServicePrincipalPassword(pulumi.CustomResource):
    end_date: pulumi.Output[str]
    """
    The End Date which the Password is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
    """
    end_date_relative: pulumi.Output[str]
    """
    A relative duration for which the Password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
    """
    key_id: pulumi.Output[str]
    """
    A GUID used to uniquely identify this Key. If not specified a GUID will be created. Changing this field forces a new resource to be created.
    """
    service_principal_id: pulumi.Output[str]
    """
    The ID of the Service Principal for which this password should be created. Changing this field forces a new resource to be created.
    """
    start_date: pulumi.Output[str]
    """
    The Start Date which the Password is valid from, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
    """
    value: pulumi.Output[str]
    """
    The Password for this Service Principal.
    """
    def __init__(__self__, resource_name, opts=None, end_date=None, end_date_relative=None, key_id=None, service_principal_id=None, start_date=None, value=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Password associated with a Service Principal within Azure Active Directory.
        
        > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] end_date: The End Date which the Password is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        :param pulumi.Input[str] end_date_relative: A relative duration for which the Password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] key_id: A GUID used to uniquely identify this Key. If not specified a GUID will be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] service_principal_id: The ID of the Service Principal for which this password should be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] start_date: The Start Date which the Password is valid from, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
        :param pulumi.Input[str] value: The Password for this Service Principal.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azuread/blob/master/website/docs/r/service_principal_password.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['end_date'] = end_date
            __props__['end_date_relative'] = end_date_relative
            __props__['key_id'] = key_id
            if service_principal_id is None:
                raise TypeError("Missing required property 'service_principal_id'")
            __props__['service_principal_id'] = service_principal_id
            __props__['start_date'] = start_date
            if value is None:
                raise TypeError("Missing required property 'value'")
            __props__['value'] = value
        super(ServicePrincipalPassword, __self__).__init__(
            'azuread:index/servicePrincipalPassword:ServicePrincipalPassword',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, end_date=None, end_date_relative=None, key_id=None, service_principal_id=None, start_date=None, value=None):
        """
        Get an existing ServicePrincipalPassword resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] end_date: The End Date which the Password is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        :param pulumi.Input[str] end_date_relative: A relative duration for which the Password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] key_id: A GUID used to uniquely identify this Key. If not specified a GUID will be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] service_principal_id: The ID of the Service Principal for which this password should be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] start_date: The Start Date which the Password is valid from, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
        :param pulumi.Input[str] value: The Password for this Service Principal.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azuread/blob/master/website/docs/r/service_principal_password.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["end_date"] = end_date
        __props__["end_date_relative"] = end_date_relative
        __props__["key_id"] = key_id
        __props__["service_principal_id"] = service_principal_id
        __props__["start_date"] = start_date
        __props__["value"] = value
        return ServicePrincipalPassword(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

