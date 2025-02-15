// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a single Group Membership within Azure Active Directory.
// 
// > **NOTE:** Do not use this resource at the same time as `azuread_group.members`.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azuread/blob/master/website/docs/r/group_member.html.markdown.
type GroupMember struct {
	s *pulumi.ResourceState
}

// NewGroupMember registers a new resource with the given unique name, arguments, and options.
func NewGroupMember(ctx *pulumi.Context,
	name string, args *GroupMemberArgs, opts ...pulumi.ResourceOpt) (*GroupMember, error) {
	if args == nil || args.GroupObjectId == nil {
		return nil, errors.New("missing required argument 'GroupObjectId'")
	}
	if args == nil || args.MemberObjectId == nil {
		return nil, errors.New("missing required argument 'MemberObjectId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["groupObjectId"] = nil
		inputs["memberObjectId"] = nil
	} else {
		inputs["groupObjectId"] = args.GroupObjectId
		inputs["memberObjectId"] = args.MemberObjectId
	}
	s, err := ctx.RegisterResource("azuread:index/groupMember:GroupMember", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &GroupMember{s: s}, nil
}

// GetGroupMember gets an existing GroupMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupMember(ctx *pulumi.Context,
	name string, id pulumi.ID, state *GroupMemberState, opts ...pulumi.ResourceOpt) (*GroupMember, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["groupObjectId"] = state.GroupObjectId
		inputs["memberObjectId"] = state.MemberObjectId
	}
	s, err := ctx.ReadResource("azuread:index/groupMember:GroupMember", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &GroupMember{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *GroupMember) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *GroupMember) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The Object ID of the Azure AD Group you want to add the Member to.  Changing this forces a new resource to be created.
func (r *GroupMember) GroupObjectId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["groupObjectId"])
}

// The Object ID of the Azure AD Object you want to add as a Member to the Group. Supported Object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
func (r *GroupMember) MemberObjectId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["memberObjectId"])
}

// Input properties used for looking up and filtering GroupMember resources.
type GroupMemberState struct {
	// The Object ID of the Azure AD Group you want to add the Member to.  Changing this forces a new resource to be created.
	GroupObjectId interface{}
	// The Object ID of the Azure AD Object you want to add as a Member to the Group. Supported Object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
	MemberObjectId interface{}
}

// The set of arguments for constructing a GroupMember resource.
type GroupMemberArgs struct {
	// The Object ID of the Azure AD Group you want to add the Member to.  Changing this forces a new resource to be created.
	GroupObjectId interface{}
	// The Object ID of the Azure AD Object you want to add as a Member to the Group. Supported Object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
	MemberObjectId interface{}
}
