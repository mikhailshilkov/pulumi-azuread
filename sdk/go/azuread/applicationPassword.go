// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Password associated with an Application within Azure Active Directory.
// 
// > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.
type ApplicationPassword struct {
	s *pulumi.ResourceState
}

// NewApplicationPassword registers a new resource with the given unique name, arguments, and options.
func NewApplicationPassword(ctx *pulumi.Context,
	name string, args *ApplicationPasswordArgs, opts ...pulumi.ResourceOpt) (*ApplicationPassword, error) {
	if args == nil || args.ApplicationId == nil {
		return nil, errors.New("missing required argument 'ApplicationId'")
	}
	if args == nil || args.Value == nil {
		return nil, errors.New("missing required argument 'Value'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["applicationId"] = nil
		inputs["endDate"] = nil
		inputs["endDateRelative"] = nil
		inputs["keyId"] = nil
		inputs["startDate"] = nil
		inputs["value"] = nil
	} else {
		inputs["applicationId"] = args.ApplicationId
		inputs["endDate"] = args.EndDate
		inputs["endDateRelative"] = args.EndDateRelative
		inputs["keyId"] = args.KeyId
		inputs["startDate"] = args.StartDate
		inputs["value"] = args.Value
	}
	s, err := ctx.RegisterResource("azuread:index/applicationPassword:ApplicationPassword", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ApplicationPassword{s: s}, nil
}

// GetApplicationPassword gets an existing ApplicationPassword resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationPassword(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ApplicationPasswordState, opts ...pulumi.ResourceOpt) (*ApplicationPassword, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["applicationId"] = state.ApplicationId
		inputs["endDate"] = state.EndDate
		inputs["endDateRelative"] = state.EndDateRelative
		inputs["keyId"] = state.KeyId
		inputs["startDate"] = state.StartDate
		inputs["value"] = state.Value
	}
	s, err := ctx.ReadResource("azuread:index/applicationPassword:ApplicationPassword", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ApplicationPassword{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ApplicationPassword) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ApplicationPassword) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *ApplicationPassword) ApplicationId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["applicationId"])
}

// The End Date which the Password is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
func (r *ApplicationPassword) EndDate() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["endDate"])
}

// A relative duration for which the Password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
func (r *ApplicationPassword) EndDateRelative() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["endDateRelative"])
}

// A GUID used to uniquely identify this Password. If not specified a GUID will be created. Changing this field forces a new resource to be created.
func (r *ApplicationPassword) KeyId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["keyId"])
}

// The Start Date which the Password is valid from, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
func (r *ApplicationPassword) StartDate() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["startDate"])
}

// The Password for this Application .
func (r *ApplicationPassword) Value() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["value"])
}

// Input properties used for looking up and filtering ApplicationPassword resources.
type ApplicationPasswordState struct {
	ApplicationId interface{}
	// The End Date which the Password is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
	EndDate interface{}
	// A relative duration for which the Password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
	EndDateRelative interface{}
	// A GUID used to uniquely identify this Password. If not specified a GUID will be created. Changing this field forces a new resource to be created.
	KeyId interface{}
	// The Start Date which the Password is valid from, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
	StartDate interface{}
	// The Password for this Application .
	Value interface{}
}

// The set of arguments for constructing a ApplicationPassword resource.
type ApplicationPasswordArgs struct {
	ApplicationId interface{}
	// The End Date which the Password is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
	EndDate interface{}
	// A relative duration for which the Password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
	EndDateRelative interface{}
	// A GUID used to uniquely identify this Password. If not specified a GUID will be created. Changing this field forces a new resource to be created.
	KeyId interface{}
	// The Start Date which the Password is valid from, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
	StartDate interface{}
	// The Password for this Application .
	Value interface{}
}